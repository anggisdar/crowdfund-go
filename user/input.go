package user

type RegisterUserInput struct {
	Name       string
	Ocuppation string
	Email      string
	Password   string
}

//ini struct user input untuk mapping
//harus diawali huruf besar agar bisa diakses publik
