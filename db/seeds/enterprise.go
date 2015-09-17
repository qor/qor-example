// +build enterprise

package main

import (
	"log"
	"time"

	"github.com/qor/qor-example/db"
	"github.com/qor/qor-example/db/seeds"
	"github.com/theplant/qor-enterprise/promotion"
)

var (
	fake           = seeds.Fake
	truncateTables = seeds.TruncateTables

	Seeds  = seeds.Seeds
	Tables = []interface{}{
		&promotion.PromotionDiscount{},
		&promotion.PromotionRule{},
		&promotion.PromotionBenefit{},
		&promotion.PromotionCoupon{},
		&promotion.BenefitRecord{},
	}
)

func main() {
	truncateTables(Tables...)
	createRecords()
}

func createRecords() {
	for _, enterpriseData := range Seeds.Enterprises {
		begins, _ := time.Parse("2006-01-02 15:04:05", enterpriseData.Begins)
		expires, _ := time.Parse("2006-01-02 15:04:05", enterpriseData.Expires)

		enterprise := promotion.PromotionDiscount{}
		enterprise.Name = enterpriseData.Name
		enterprise.Begins = &begins
		enterprise.Expires = &expires
		enterprise.RequiresCoupon = enterpriseData.RequiresCoupon
		enterprise.Unique = enterpriseData.Unique

		if err := db.DB.Create(&enterprise).Error; err != nil {
			log.Fatalf("create enterprise (%v) failure, got err %v", enterprise, err)
		}

		for _, couponData := range enterpriseData.Coupons {
			coupon := promotion.PromotionCoupon{}
			coupon.DiscountID = enterprise.ID
			coupon.Code = couponData.Code
			if err := db.DB.Create(&coupon).Error; err != nil {
				log.Fatalf("create coupon (%v) failure, got err %v", coupon, err)
			}
		}

		for _, ruleData := range enterpriseData.Rules {
			rule := promotion.PromotionRule{}
			rule.DiscountID = enterprise.ID
			rule.Kind = ruleData.Kind
			rule.Value = ruleData.Value
			if err := db.DB.Create(&rule).Error; err != nil {
				log.Fatalf("create rule (%v) failure, got err %v", rule, err)
			}
		}

		for _, benefitData := range enterpriseData.Benefits {
			benefit := promotion.PromotionBenefit{}
			benefit.DiscountID = enterprise.ID
			benefit.Kind = benefitData.Kind
			benefit.Value = benefitData.Value
			if err := db.DB.Create(&benefit).Error; err != nil {
				log.Fatalf("create benefit (%v) failure, got err %v", benefit, err)
			}
		}
	}
}
