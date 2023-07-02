package lawg

type Tags map[string]interface{}

type UpdateEvent struct {
	ID          *string        `json:"id,omitempty"`
	Title       *string        `json:"title,omitempty"`
	Description *string        `json:"description,omitempty"`
	Emoji       *string        `json:"emoji,omitempty"`
	Timestamp   *string        `json:"timestamp,omitempty"`
	Metadata    *EventMetadata `json:"metadata,omitempty"`
	Notify      *bool          `json:"notify,omitempty"`
}

type CreateEvent struct {
	Title       string         `json:"title"`
	Description *string        `json:"description,omitempty"`
	Emoji       *string        `json:"emoji,omitempty"`
	Timestamp   *string        `json:"timestamp,omitempty"`
	Metadata    *EventMetadata `json:"metadata,omitempty"`
	Notify      *bool          `json:"notify,omitempty"`
}

type EventMetadata struct {
	UA   *string `json:"ua,omitempty"`
	Tags Tags    `json:"tags,omitempty"`
}
