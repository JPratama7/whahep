package whahep

type Button struct {
	ButtonId    string `json:"button_id,omitempty"`
	DisplayText string `json:"display_text,omitempty"`
}

type ButtonsMessage struct {
	HeaderText  string `json:"header_text,omitempty"`
	ContentText string `json:"content_text,omitempty"`
	FooterText  string `json:"footer_text,omitempty"`
}

type ButtonsMessages struct {
	Message ButtonsMessage `json:"message,omitempty"`
	Buttons []Button       `json:"buttons,omitempty"`
}

type ListMessage struct {
	Title       string
	Description string
	FooterText  string
	ButtonText  string
	Sections    []ListSection
}

type ListSection struct {
	Title string
	Rows  []ListRow
}

type ListRow struct {
	Title       string
	Description string
	RowId       string
}
