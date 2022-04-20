package parser

import "time"

type TypeFormData struct {
	EventID      string        `json:"event_id"`
	EventType    string        `json:"event_type"`
	FormResponse *FormResponse `json:"form_response"`
}

type FormResponse struct {
	FormID      string            `json:"form_id"`
	Token       string            `json:"token"`
	LandedAt    *time.Time        `json:"landed_at"`
	SubmittedAt *time.Time        `json:"submitted_at"`
	Hidden      map[string]string `json:"hidden"`
	Definition  *Definition       `json:"definition"`
	Answers     []Answers         `json:"answers"`
}

type Definition struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Fields []Fields `json:"fields"`
}

type Fields struct {
	ID                      string         `json:"id"`
	Ref                     string         `json:"ref"`
	Type                    string         `json:"type"`
	Title                   string         `json:"title"`
	Properties              Properties     `json:"properties"`
	AllowMultipleSelections bool           `json:"allow_multiple_selections,omitempty"`
	Choices                 []FieldChoices `json:"choices,omitempty"`
}

type Properties struct {
}

type FieldChoices struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type Answers struct {
	Type        string   `json:"type"`
	PhoneNumber string   `json:"phone_number,omitempty"`
	Field       *Field   `json:"field"`
	Email       string   `json:"email,omitempty"`
	URL         string   `json:"url,omitempty"`
	Text        string   `json:"text,omitempty"`
	Number      int      `json:"number,omitempty"`
	Date        string   `json:"date,omitempty"`
	FileURL     string   `json:"file_url,omitempty"`
	Choices     *Choices `json:"choices,omitempty"`
	Choice      *Choice  `json:"choice,omitempty"`
	Boolean     bool     `json:"boolean,omitempty"`
}

type Field struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Ref  string `json:"ref"`
}

type Choices struct {
	Labels []string `json:"labels"`
}

type Choice struct {
	Label string `json:"label"`
}
