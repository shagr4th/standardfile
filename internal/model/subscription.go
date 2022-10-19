package model

type Subscription struct {
	ID               string `json:"uuid"`
	Name             string `json:"planName"`
	EndsAt           int64  `json:"endsAt"`
	CreatedAt        int64  `json:"createdAt"`
	UpdatedAt        int64  `json:"updatedAt"`
	Cancelled        int64  `json:"cancelled"`
	SubscriptionId   int64  `json:"subscriptionId"`
	SubscriptionType string `json:"subscriptionType"`
}

type SubscriptionRole struct {
	ID   string `json:"uuid"`
	Name string `json:"name"`
}
