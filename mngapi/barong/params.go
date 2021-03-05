package barong

// CreateServiceAccountParams contain all the allowed params for Service Account creation
type CreateServiceAccountParams struct {
	OwnerUID string `json:"owner_uid"`
	Role     string `json:"service_account_role"`
	UID      string `json:"service_account_uid,omitempty"`
	Email    string `json:"service_account_email,omitempty"`
}

// CreateAPIKeyParams contain all the allowed params for API Key creation
type CreateAPIKeyParams struct {
	Algorithm string   `json:"algorithm"`
	UID       string   `json:"uid,omitempty"`
	Scopes    string `json:"scopes,omitempty"`
}
