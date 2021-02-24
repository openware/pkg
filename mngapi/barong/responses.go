package barong

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
