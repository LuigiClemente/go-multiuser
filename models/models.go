package models

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

// SubscriptionPlan represents a subscription plan
type SubscriptionPlan struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	MemberLimit int    `json:"memberLimit"`
}

// UserSubscription represents a user's subscription
type UserSubscription struct {
	ID                 int    `json:"id"`
	UserID             int    `json:"userId"`
	SubscriptionPlanID int    `json:"subscriptionPlanId"`
	Role               string `json:"role"`
}
