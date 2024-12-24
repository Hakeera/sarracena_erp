package response

type UserResponse struct {
	Email		string	`json:"email"`
	Password	string	`json:"password"`
	Name		string	`json:"name"`
	Age			int8	`json:"age"`
}