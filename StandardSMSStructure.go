package msgstruct

type StandardSMSStructure struct {
	Sender   *string `json:"sender"`
	Receiver *string `json:"receiver"`
	TextBody *string `json:"text"`
}
