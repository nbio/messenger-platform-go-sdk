package messenger

type AttachmentType string

const (
	AttachmentTypeTemplate AttachmentType = "template"
	AttachmentTypeImage    AttachmentType = "image"
	AttachmentTypeVideo    AttachmentType = "video"
)

type Attachment struct {
	Type    AttachmentType `json:"type"`
	Payload interface{}    `json:"payload,omitempty"`
}

// func (a *Attachment) MarshalJSON() ([]byte, error) {
// 	if a.Type == "" {
// 		switch a.Payload.(type) {
// 		case template.Payload:
// 			a.Type = AttachmentTypeTemplate
// 		case Resource:
// 			a.Type = AttachmentTypeImage //best guess
// 		default:
// 			return []byte{}, errors.New("Invalid payload")
// 		}
// 	}
// 	return json.NewEncoder()
// }

type Resource struct {
	URL string `json:"url"`
}