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

type Device struct {
	Iden     string  `json:"iden"`
	Nickname string  `json:"nickname"`
	Active   bool    `json:"active"`
	Created  float32 `json:"created"`
	Modified float32 `json:"modified"`
	Pushable bool    `json:"pushable"`
	Icon     string  `json:"icon"`
}

type Push struct {
	Iden                    string   `json:"iden"`
	Type                    PushType `json:"type"`
	DeviceIden              string   `json:"device_iden"`
	Email                   string   `json:"email"`
	ChannelTag              string   `json:"channel_tag"`
	ClientIden              string   `json:"client_iden"`
	Active                  bool     `json:"active"`
	Modified                float32  `json:"modified"`
	Dismissed               bool     `json:"dismissed"`
	SenderIden              string   `json:"sender_iden"`
	SenderName              string   `json:"sender_name"`
	SenderEmail             string   `json:"sender_email"`
	SenderEmailNormalized   string   `json:"sender_email_normalized"`
	Created                 float32  `json:"created"`
	ReceiverEmail           string   `json:"receiver_email"`
	ReceiverEmailNormalized string   `json:"receiver_email_normalized"`
	ReceiverIden            string   `json:"receiver_iden"`
	TargetDeviceIden        string   `json:"target_device_iden"`
	SourceDeviceIden        string   `json:"source_device_iden"`
	Direction               string   `json:"direction"`
	Title                   string   `json:"title"`
	Body                    string   `json:"body"`
	URL                     string   `json:"url"`
	FileName                string   `json:"file_name"`
	FileType                string   `json:"file_type"`
	FileURL                 string   `json:"file_url"`
	AwakeAppGuids           []string `json:"awake_app_guids"`
}
