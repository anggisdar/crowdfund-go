package user

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Ocuppation string `json:"ocuppation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

//ini struct user input untuk mapping
//harus diawali huruf besar agar bisa diakses publik

// bagian login
type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `~json:"password" binding:"required"`
}
