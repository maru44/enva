package domain

import "time"

type (
	Subscription struct {
		ID     string `json:"id"`
		UserID UserID `json:"user_id,omitempty"`
		OrgID  OrgID  `json:"org_id,omitempty"`

		StripeSubscriptionID     string `json:"stripe_subscription_id"`
		StripeCustomerID         string `json:"stripe_customer_id"`
		StripeProductID          string `json:"stripe_product_id"`
		StripeSubscriptionStatus string `json:"stripe_subscription_status"`

		IsValid   bool       `json:"is_valid"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	// @TODO this repo should be embed stripe sdk
	ISubscriptionInteractor interface {
		CreateSubscription()
		UpdateSubscription()

		CreateCustomer()
	}
)
