package schemas

type AddUserInput struct {
	Username string `json:"username"  binding:"required"`
	Email    string `json:"email"  binding:"required"`
}
