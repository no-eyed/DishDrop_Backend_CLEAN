package requests

type UserRequest struct {
	Data struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		UserName  string `json:"username"`
		Email     []struct {
			EmailAddress string `json:"email_address"`
		} `json:"email_addresses"`

		UUID string `json:"id"`

		CreatedAt int64 `json:"created_at"`
		UpdatedAt int64 `json:"updated_at"`
	} `json:"data"`

	Type string `json:"type"`
}
