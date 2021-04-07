package barong

// CreateServiceAccountParams contain all the allowed params for Service Account creation
type CreateServiceAccountParams struct {
	OwnerUID string `json:"owner_uid"`
	Role     string `json:"service_account_role"`
	UID      string `json:"service_account_uid,omitempty"`
	Email    string `json:"service_account_email,omitempty"`
	State    string `json:"service_account_state,omitempty"`
	Level    int    `json:"service_account_level,omitempty"`
}

// CreateAPIKeyParams contain all the allowed params for API Key creation
type CreateAPIKeyParams struct {
	Algorithm string `json:"algorithm"`
	UID       string `json:"uid,omitempty"`
	Scopes    string `json:"scopes,omitempty"`
}

// CreateAttachmentParams contain all the allowed params for API Key creation
type CreateAttachmentParams struct {
	UID      string `json:"uid,omitempty"`
	FileName string `json:"filename,omitempty"`
	FileExt  string `json:"file_ext,omitempty"`
	Upload   string `json:"upload,omitempty"`
}
