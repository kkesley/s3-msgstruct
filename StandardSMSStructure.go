package msgstruct

//StandardSMSStructure json supported structure to make a sms request
type StandardSMSStructure struct {
	ReferenceID    string            `json:"reference_id"` //THE FIRST SECTION IS THE SERVICE WHO CALLED e.g. newsletter/arn::itea::1::platform::client::1/SPECIFIC_TOPIC or platform/arn::itea::1::platform::client::1/SPECIFIC_TOPIC
	Method         string            `json:"method"`       //should be sms
	Principal      string            `json:"principal"`
	Sender         *string           `json:"sender"`
	Receiver       *string           `json:"receiver"`
	TextBody       *string           `json:"text"`
	CallbackURL    string            `json:"callback_url"`
	CallbackBody   map[string]string `json:"callback_body"`
	CallbackHeader map[string]string `json:"callback_header"`
}
