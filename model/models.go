package model

type (
	User struct {
		UserId    string `db:"user_id"`
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		Email     string `db:"email"`
		Phone     string `db:"phone"`
		Password  string `db:"password"`
	}

	Token struct {
		UserID string `db:"user_id"`
		Email  string `db:"email"`
	}
)
