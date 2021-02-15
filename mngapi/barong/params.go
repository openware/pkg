package barong

type CreateServiceAccountParams struct {
	OwnerUID string `json:"owner_uid"`
	Role     string `json:"service_account_role"`
	UID      string `json:"service_account_uid,omitempty"`
	Email    string `json:"service_account_email,omitempty"`
}
