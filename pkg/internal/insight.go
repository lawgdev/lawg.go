package lawg

type CreateInsight struct {
	Title string       `json:"title"`
	Value InsightValue `json:"value"`
	Emoji string       `json:"emoji,omitempty"`
}

type InsightValue struct {
	Set       *int `json:"set,omitempty"`
	Increment *int `json:"increment,omitempty"`
}

type UpdateInsight struct {
	ID        string `json:"id,omitempty"`
	Emoji     string `json:"emoji,omitempty"`
	Set       *int   `json:"set,omitempty"`
	Increment *int   `json:"increment,omitempty"`
}
