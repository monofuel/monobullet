package monobullet

type User struct {
	Created         float32 `json:"created"`
	Email           string  `json:"email"`
	EmailNormalized string  `json:"email_normalized"`
	Iden            string  `json:"iden"`
	ImageUrl        string  `json:"image_url"`
	MaxUploadSize   float32 `json:"max_upload_size"`
	Modified        float32 `json:"modified"`
	Name            string  `json:"Name"`
}
