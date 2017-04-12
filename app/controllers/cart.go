package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qor/action_bar"
	// "github.com/qor/transition"

	"dukeondope.ru/mlm/sandbox/app/models"
	"dukeondope.ru/mlm/sandbox/config"
	"dukeondope.ru/mlm/sandbox/config/admin"
	"dukeondope.ru/mlm/sandbox/config/cart"
	"dukeondope.ru/mlm/sandbox/config/seo"
)

func AddToCartHandler(ctx *gin.Context) {
	var (
		curCart, _ = cart.GetCart(ctx)
		cartItem   cart.CartItem
	)

	if err := ctx.Bind(&cartItem); err == nil {
		if _, ok := curCart.Add(&cartItem); ok {
			ctx.JSON(
				http.StatusCreated,
				gin.H{
					"status":  http.StatusCreated,
					"message": "Cart item added successfully!",
					"itemID":  cartItem.SizeVariationID,
				},
			)
		} else {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"status":  http.StatusBadRequest,
					"message": "Cart item can not be added!",
				},
			)
		}
	}
}

func RemoveFromCartHandler(ctx *gin.Context) {
	var (
		curCart, _ = cart.GetCart(ctx)
		id, _      = strconv.Atoi(ctx.Param("id"))
	)

	curCart.Remove(uint(id))

	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func ShowCartHandler(ctx *gin.Context) {
	var (
		curCart, _     = cart.GetCart(ctx)
		cartItems      = curCart.GetContent()
		sizeVariations []models.SizeVariation
		// orderItems     []models.OrderItem
	)

	DB(ctx).Preload("Size").Preload("ColorVariation.Color").Preload("ColorVariation.Product").Where(curCart.GetItemsIDS()).Find(&sizeVariations)

	for _, item := range curCart.GetContent() {
		var (
			oi models.OrderItem
			sv models.SizeVariation
		)
		item.Bind(&oi)
		DB(ctx).Preload("Size").Preload("ColorVariation.Color").Preload("ColorVariation.Product").Where([]uint{oi.SizeVariationID}).First(&sv)
		oi.SizeVariation = sv
	}

	var (
		cartAmount, cartItemAmount float32
		extCartItems               []interface{}
	)
	for _, item := range sizeVariations {
		cartItemAmount = float32(uint(item.ColorVariation.Product.Price*100)*cartItems[item.ID].Quantity) / 100
		cartAmount = cartAmount + cartItemAmount
		extCartItems = append(extCartItems, struct {
			cart.CartItem
			MainImageURL string
			ProductName  string
			ColorName    string
			SizeName     string
			Price        float32
			Amount       float32
		}{
			*cartItems[item.ID],
			item.ColorVariation.Product.MainImageURL(),
			item.ColorVariation.Product.Name,
			item.ColorVariation.Color.Name,
			item.Size.Name,
			item.ColorVariation.Product.Price,
			cartItemAmount,
		})
	}

	config.View.Funcs(funcsMap(ctx)).Execute(
		"cart/cart_show",
		gin.H{
			"ActionBarTag":  admin.ActionBar.Actions(action_bar.Action{Name: "Edit SEO", Link: seo.SEOCollection.SEOSettingURL("/help")}).Render(ctx.Writer, ctx.Request),
			"showCartItems": extCartItems,
			"cartAmount":    cartAmount,
			"Categories":    CategoriesList(ctx),

			"CurrentUser":   CurrentUser(ctx),
			"CurrentLocale": CurrentLocale(ctx),
		},
		ctx.Request,
		ctx.Writer,
	)

	/* 	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Found %v cart items", len(extCartItems)),
		"data": gin.H{
			"items":  extCartItems,
			"count":  len(extCartItems),
			"amount": cartAmount,
		},
	}) */

}

func CheckoutCartHandler(ctx *gin.Context) {
	var (
		curCart, _  = cart.GetCart(ctx)
		currentUser = CurrentUser(ctx)
		addresses   []models.Address
	)

	if currentUser == nil {
		http.Redirect(ctx.Writer, ctx.Request, "/auth/login", http.StatusTemporaryRedirect)
	}
	if curCart.IsEmpty() {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	DB(ctx).Model(&currentUser).Related(&addresses)

	config.View.Funcs(funcsMap(ctx)).Execute(
		"cart/order_create",
		gin.H{
			"ActionBarTag": admin.ActionBar.Actions(action_bar.Action{Name: "Edit SEO", Link: seo.SEOCollection.SEOSettingURL("/help")}).Render(ctx.Writer, ctx.Request),
			"Addresses":    addresses,
			// "cartAmount":    cartAmount,
			"Categories": CategoriesList(ctx),

			"CurrentUser":   CurrentUser(ctx),
			"CurrentLocale": CurrentLocale(ctx),
		},
		ctx.Request,
		ctx.Writer,
	)
}

func OrderCartHandler(ctx *gin.Context) {
	var (
		order          models.Order
		orderItems     []models.OrderItem
		sizeVariations []models.SizeVariation
		curCart, _     = cart.GetCart(ctx)
		cartItems      = curCart.GetContent()
	)

	DB(ctx).Create(&order)

	ctx.Bind(&order)
	DB(ctx).Preload("ColorVariation.Product").Where(curCart.GetItemsIDS()).Find(&sizeVariations)
	for _, item := range sizeVariations {
		orderItems = append(orderItems, models.OrderItem{
			SizeVariation: item,
			Quantity:      cartItems[item.ID].Quantity,
			Price:         item.ColorVariation.Product.Price,
			DiscountRate:  0,
		})
	}
	DB(ctx).Model(&order).Update(&models.Order{
		User:       *CurrentUser(ctx),
		OrderItems: orderItems,
	})

	DB(ctx).Model(&order).Update(&models.Order{
		PaymentAmount: order.Amount(),
	})

	if err := models.OrderState.Trigger("checkout", &order, DB(ctx)); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	DB(ctx).Save(&order)
	curCart.EmptyCart()

	if err := models.OrderState.Trigger("pay", &order, DB(ctx)); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	} else {
		if rest := order.User.Balance - order.PaymentAmount; rest >= 0 {
			DB(ctx).Model(&order.User).Update("balance", rest)
		} else {
			models.OrderState.Trigger("cancel", &order, DB(ctx), "No finance")
		}
	}

	DB(ctx).Save(&order)
}
