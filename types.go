package monobullet

type PushType string

const (
	NoteType PushType = "note"
	LinkType          = "link"
	FileType          = "file"
)

type User struct {
	Created         float32 `json:"created"`
	Email           string  `json:"email"`
	EmailNormalized string  `json:"email_normalized"`
	Iden            string  `json:"iden"`
	ImageURL        string  `json:"image_url"`
	MaxUploadSize   float32 `json:"max_upload_size"`
	Modified        float32 `json:"modified"`
	Name            string  `json:"name"`
	PlanID          string  `json:"plan_id"`
	Active          bool    `json:"active"`
	Pro             bool    `json:"pro"`
}

type Push struct {
	Type       PushType `json:"type"`
	DeviceIden string   `json:"device_iden"`
	Email      string   `json:"email"`
	ChannelTag string   `json:"channel_tag"`
	ClientIden string   `json:"client_iden"`
}

type Note struct {
	Push
	Title string `json:"title"`
	Body  string `json:"body"`
}
type Link struct {
	Push
	Title string `json:"title"`
	Body  string `json:"body"`
	URL   string `json:"url"`
}
type File struct {
	Push
	Body     string `json:"body"`
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	FileURL  string `json:"file_url"`
}
