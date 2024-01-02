package photo

import "rest-api/models"

// PhotoResponse adalah struktur untuk respon foto dengan informasi pengguna
type PhotoResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	User     UserResponse `json:"user,omitempty"`
}

// PhotoRegularResponse adalah struktur untuk respon foto tanpa informasi pengguna
type PhotoRegularResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}

// UserResponse adalah struktur untuk respon informasi pengguna pada foto
type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// FormatPhoto mengonversi model Photo menjadi bentuk respon yang diinginkan
func FormatPhoto(photo *models.Photo, typeRes string) interface{} {
	if typeRes == "regular" {
		return PhotoRegularResponse{
			ID:       photo.ID,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoURL: photo.PhotoURL,
		}
	}

	return PhotoResponse{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		User: UserResponse{
			ID:       photo.User.ID,
			Username: photo.User.Username,
			Email:    photo.User.Email,
		},
	}
}
