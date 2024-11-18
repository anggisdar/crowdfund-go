package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Ocuppation string `json:"ocuppation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Ocuppation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	return formatter
}

// mengubah format pada struct user yang ada di db dari handler agar ditampilkan sesuai dengan keinginan (tidak seluruh data)
