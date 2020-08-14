package requestpayload

var (
	//MaxLength request Body length
	MaxLength = int64(1000)

	//YellowRabbitURL url for yellowrabbit service
	YellowRabbitURL = "http://localhost:8095/"
)

//PayloadCollection payload
type PayloadCollection struct {
	Payloads []Payload `json:"data"`
}

//Payload struct
type Payload struct {
	MessageID string `json:"messageId"`
	Content   string `json:"content"`
}
