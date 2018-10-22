package msgstruct

//StandardEmailStructure json supported structure to make an email request
type StandardEmailStructure struct {
	ReferenceID    string            `json:"reference_id"` //THE FIRST SECTION IS THE SERVICE WHO CALLED e.g. newsletter/arn::itea::1::platform::client::1/SPECIFIC_TOPIC or platform/arn::itea::1::platform::client::1/SPECIFIC_TOPIC
	Method         string            `json:"method"`       //should be email
	Principal      string            `json:"principal"`
	Sender         *string           `json:"sender"`
	Receivers      []*string         `json:"receivers"`
	CCs            []*string         `json:"ccs"`
	BCCs           []*string         `json:"bccs"`
	HTMLBody       *string           `json:"html"`
	TextBody       *string           `json:"text"`
	Subject        *string           `json:"subject"`
	ReplyTo        []*string         `json:"reply_to"`
	CallbackURL    string            `json:"callback_url"`
	CallbackBody   map[string]string `json:"callback_body"`
	CallbackHeader map[string]string `json:"callback_header"`
}
