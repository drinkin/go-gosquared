package gosquared

import "time"

type Person struct {
	ID          string                 `json:"id"`
	Email       string                 `json:"email,omitempty"`
	Name        string                 `json:"name,omitempty"`
	FirstName   string                 `json:"first_name,omitempty"`
	LastName    string                 `json:"last_name,omitempty"`
	Username    string                 `json:"username,omitempty"`
	Description string                 `json:"description,omitempty"`
	Avatar      string                 `json:"avatar,omitempty"`
	Phone       string                 `json:"phone,omitempty"`
	CreatedAt   *time.Time             `json:"created_at,omitempty"`
	Custom      map[string]interface{} `json:"custom,omitempty"`
}
