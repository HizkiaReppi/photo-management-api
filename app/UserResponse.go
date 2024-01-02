package user

import "rest-api/models"

// UserResponse adalah struktur untuk respon informasi pengguna
type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UserResponseWithToken adalah struktur untuk respon informasi pengguna dengan token
type UserResponseWithToken struct {
	UserResponse
	Token string `json:"token"`
}

// FormatUserResponse mengonversi model User menjadi bentuk respon yang diinginkan
func FormatUserResponse(user models.User, token string) interface{} {
	if token == "" {
		return UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}
	}

	return UserResponseWithToken{
		UserResponse: UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
		Token: token,
	}
}
