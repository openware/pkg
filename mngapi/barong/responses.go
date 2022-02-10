package barong

// User represents a User creation response
type User struct {
	Email       string `json:"email"`
	UID         string `json:"uid"`
	Role        string `json:"role"`
	Level       uint64 `json:"level"`
	OTP         bool   `json:"otp"`
	State       string `json:"state"`
	ReferralUID string `json:"referral_uid"`
	Data        string `json:"data"`
}

// ServiceAccount represents a Service Account creation response
type ServiceAccount struct {
	Email     string `json:"email"`
	UID       string `json:"uid"`
	Role      string `json:"role"`
	Level     uint64 `json:"level"`
	State     string `json:"state"`
	User      User   `json:"user"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// APIKey represents an API Key creation response
type APIKey struct {
	KID       string   `json:"kid"`
	Algorithm string   `json:"algorithm"`
	Scope     []string `json:"scope"`
	State     string   `json:"state"`
	Secret    string   `json:"secret"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type Attachment struct {
	ID  uint64 `json:"id"`
	UID string `json:"user_uid"`
}
