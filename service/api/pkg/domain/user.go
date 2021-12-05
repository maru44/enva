package domain

type (
	User struct {
		ID              string    `json:"id"`
		UserName        string    `json:"username"`
		Email           string    `json:"email"`
		Explanation     string    `json:"explanation"`
		ImageURL        *string   `json:"image_url"`
		IsActive        bool      `json:"is_active"`
		IsEmailVerified bool      `json:"is_email_verified"`
		UserType        *UserType `json:"user_type"`
	}

	UserType string
)

const (
	UserTypeAdmin  = UserType("admin")
	UserTypeNormal = UserType("normal")
	UserTypeGuest  = UserType("guest")
)

func (u *User) IsAdmin() bool {
	if u.UserType == nil {
		return false
	}
	return *u.UserType == UserTypeAdmin
}
