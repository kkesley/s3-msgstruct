package msgstruct

//StandardSMSStructure json supported structure to make a sms request
type StandardSMSStructure struct {
	Sender         *string           `json:"sender"`
	Receiver       *string           `json:"receiver"`
	TextBody       *string           `json:"text"`
	CallbackURL    string            `json:"callback_url"`
	CallbackBody   map[string]string `json:"callback_body"`
	CallbackHeader map[string]string `json:"callback_header"`
}
