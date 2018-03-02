package orders

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	amazonpay "github.com/qor/amazon-pay-sdk-go"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/models/users"
	"github.com/qor/qor-example/utils"
	"github.com/qor/transition"
)

var (
	// OrderState order's state machine
	OrderState = transition.New(&Order{})

	// ItemState order item's state machine
	ItemState = transition.New(&OrderItem{})
)

var (
	// DraftState draft state
	DraftState = "draft"
)

func init() {
	// Define Order's States
	OrderState.Initial("draft")

	OrderState.State("pending").Enter(func(value interface{}, tx *gorm.DB) (err error) {
		order := value.(*Order)
		tx.Model(order).Association("OrderItems").Find(&order.OrderItems)
		if order.AmazonOrderReferenceID != "" {
			_, err = config.AmazonPay.SetOrderReferenceDetails(order.AmazonOrderReferenceID, amazonpay.OrderReferenceAttributes{
				OrderTotal: amazonpay.OrderTotal{CurrencyCode: "JPY", Amount: utils.FormatPrice(order.Amount())},
			})

			if err == nil {
				err = config.AmazonPay.ConfirmOrderReference(order.AmazonOrderReferenceID)
			}

			var orderDetail amazonpay.GetOrderReferenceDetailsResponse
			if err == nil {
				orderDetail, err = config.AmazonPay.GetOrderReferenceDetails(order.AmazonOrderReferenceID, order.AmazonAddressAccessToken)
			}

			if err == nil {
				address := orderDetail.GetOrderReferenceDetailsResult.OrderReferenceDetails.Destination.PhysicalDestination
				amazonAddress := users.Address{}
				amazonAddress.ContactName = address.Name
				amazonAddress.Phone = address.Phone
				amazonAddress.Address1 = address.District + " " + address.AddressLine1
				amazonAddress.Address2 = address.AddressLine2 + " " + address.AddressLine3
				amazonAddress.City = address.City
				order.ShippingAddress = amazonAddress
				order.BillingAddress = amazonAddress

				result, _ := json.Marshal(orderDetail)
				order.PaymentLog += "\n\nSetOrderReferenceDetails\n" + string(result)
				order.PaymentMethod = AmazonPay
			}
		} else {
			order.PaymentMethod = COD
		}

		if err != nil {
			order.PaymentLog += "\n" + err.Error()
		} else {
			for idx, orderItem := range order.OrderItems {
				order.OrderItems[idx].Price = orderItem.SellingPrice()
			}
			order.PaymentAmount = order.Amount()
			order.PaymentTotal = order.Total()
		}
		return err
	})

	OrderState.State("processing").Enter(func(value interface{}, tx *gorm.DB) (err error) {
		order := value.(*Order)

		switch order.PaymentMethod {
		case AmazonPay:
			var result amazonpay.AuthorizeResponse
			result, err = config.AmazonPay.Authorize(order.AmazonOrderReferenceID, order.UniqueExternalID(),
				amazonpay.Price{
					Amount:       utils.FormatPrice(order.PaymentTotal),
					CurrencyCode: config.Config.AmazonPay.CurrencyCode,
				},
				amazonpay.AuthorizeInput{},
			)

			if err == nil {
				order.AmazonAuthorizationID = result.AuthorizeResult.AuthorizationDetails.AmazonAuthorizationID
			}

			log, _ := json.Marshal(result)
			order.PaymentLog += "\n\nAuthorizeResponse\n" + string(log)
		case COD:
		default:
			err = errors.New("unsupported pay method")
		}

		if err != nil {
			order.PaymentLog += "\n" + err.Error()
		}

		return
	})

	OrderState.State("cancelled").Enter(func(value interface{}, tx *gorm.DB) (err error) {
		order := value.(*Order)
		method := ""

		switch order.PaymentMethod {
		case AmazonPay:
			if order.AmazonAuthorizationID != "" {
				method = "CloseAuthorization"
				err = config.AmazonPay.CloseAuthorization(order.AmazonAuthorizationID, "cancel order")
			} else if order.AmazonOrderReferenceID != "" {
				method = "CancelOrderReference"
				err = config.AmazonPay.CancelOrderReference(order.AmazonOrderReferenceID, "cancel order")
			}
		case COD:
		default:
			err = errors.New("unsupported pay method")
		}

		order.PaymentLog += "\n\n" + method + "\n" + fmt.Sprintf("Order cancelled at %#v", time.Now().String())

		if err != nil {
			order.PaymentLog += fmt.Sprintf("with error %v", err.Error())
		} else {
			now := time.Now()
			order.CancelledAt = &now
		}
		return
	})

	OrderState.State("shipped").Enter(func(value interface{}, tx *gorm.DB) (err error) {
		order := value.(*Order)

		switch order.PaymentMethod {
		case AmazonPay:
			if order.AmazonAuthorizationID != "" {
				var result amazonpay.CaptureResponse
				result, err = config.AmazonPay.Capture(order.AmazonAuthorizationID, order.UniqueExternalID(),
					amazonpay.Price{
						Amount:       utils.FormatPrice(order.PaymentTotal),
						CurrencyCode: config.Config.AmazonPay.CurrencyCode,
					},
					amazonpay.CaptureInput{},
				)

				if err == nil {
					order.AmazonCaptureID = result.CaptureResult.CaptureDetails.AmazonCaptureID
				}
				log, _ := json.Marshal(result)
				order.PaymentLog += "\n\nCapture\n" + string(log)
			}
		case COD:
		default:
			err = errors.New("unsupported pay method")
		}

		if err != nil {
			order.PaymentLog += "\n" + err.Error()
		} else {
			now := time.Now()
			order.ShippedAt = &now
		}
		return
	})

	OrderState.State("paid_cancelled").Enter(func(value interface{}, tx *gorm.DB) (err error) {
		order := value.(*Order)

		switch order.PaymentMethod {
		case AmazonPay:
			var result amazonpay.RefundResponse
			result, err = config.AmazonPay.Refund(order.AmazonCaptureID, order.UniqueExternalID(), amazonpay.Price{
				Amount:       utils.FormatPrice(order.PaymentTotal),
				CurrencyCode: config.Config.AmazonPay.CurrencyCode,
			}, amazonpay.RefundInput{})

			if err == nil {
				order.AmazonRefundID = result.RefundResult.RefundDetails.AmazonRefundID
			}

			log, _ := json.Marshal(result)
			order.PaymentLog += "\n\n" + string(log)
		case COD:
		default:
			err = errors.New("unsupported pay method")
		}

		order.PaymentLog += "\n\nRefund\n" + fmt.Sprintf("Order paid cancelled at %#v", time.Now().String())

		if err != nil {
			order.PaymentLog += fmt.Sprintf("with error %v", err.Error())
		} else {
			now := time.Now()
			order.CancelledAt = &now
		}
		return
	})

	OrderState.State("returned").Enter(func(value interface{}, tx *gorm.DB) error {
		order := value.(*Order)

		// check returned or not
		now := time.Now()
		order.ReturnedAt = &now
		return nil
	})

	OrderState.Event("checkout").To("pending").From("draft").After(func(value interface{}, tx *gorm.DB) (err error) {
		order := value.(*Order)
		for _, item := range order.OrderItems {
			ItemState.Trigger("checkout", &item, tx)
		}
		return nil
	})

	OrderState.Event("process").To("processing").From("pending").After(func(value interface{}, tx *gorm.DB) (err error) {
		order := value.(*Order)
		tx.Model(order).Association("OrderItems").Find(&order.OrderItems)

		for _, item := range order.OrderItems {
			ItemState.Trigger("process", &item, tx)
		}
		return nil
	})

	cancelEvent := OrderState.Event("cancel")
	cancelEvent.To("cancelled").From("draft", "pending").After(func(value interface{}, tx *gorm.DB) (err error) {
		order := value.(*Order)
		tx.Model(order).Association("OrderItems").Find(&order.OrderItems)

		for _, item := range order.OrderItems {
			ItemState.Trigger("cancel", &item, tx)
		}
		return nil
	})

	cancelEvent.To("paid_cancelled").From("processing", "shipped").After(func(value interface{}, tx *gorm.DB) (err error) {
		order := value.(*Order)
		tx.Model(order).Association("OrderItems").Find(&order.OrderItems)

		for _, item := range order.OrderItems {
			ItemState.Trigger("cancel", &item, tx)
		}
		return nil
	})

	OrderState.Event("ship").To("shipped").From("processing").After(func(value interface{}, tx *gorm.DB) (err error) {
		order := value.(*Order)
		tx.Model(order).Association("OrderItems").Find(&order.OrderItems)

		for _, item := range order.OrderItems {
			ItemState.Trigger("ship", &item, tx)
		}
		return nil
	})

	OrderState.Event("return").To("returned").From("shipped").After(func(value interface{}, tx *gorm.DB) (err error) {
		order := value.(*Order)
		tx.Model(order).Association("OrderItems").Find(&order.OrderItems)

		for _, item := range order.OrderItems {
			ItemState.Trigger("return", &item, tx)
		}
		return nil
	})

	// Define ItemItem's States
	ItemState.Initial("draft")
	ItemState.State("pending").Enter(func(value interface{}, tx *gorm.DB) error {
		// freeze stock, update order state
		return nil
	})
	ItemState.State("cancelled").Enter(func(value interface{}, tx *gorm.DB) error {
		// release stock, upate order state
		return nil
	})
	ItemState.State("processing")
	ItemState.State("shipped")
	ItemState.State("paid_cancelled").Enter(func(value interface{}, tx *gorm.DB) error {
		// do refund, release stock, update order state
		return nil
	})
	ItemState.State("returned")

	ItemState.Event("checkout").To("pending").From("draft")
	ItemState.Event("process").To("processing").From("checkout")
	cancelItemEvent := ItemState.Event("cancel")
	cancelItemEvent.To("cancelled").From("checkout")
	cancelItemEvent.To("paid_cancelled").From("paid")
	ItemState.Event("process").To("processing").From("paid")
	ItemState.Event("ship").To("shipped").From("processing")
	ItemState.Event("return").To("returned").From("shipped")
}
