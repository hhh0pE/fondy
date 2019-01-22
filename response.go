package fondy

// FinalResponse contains fields which are available in fondy final payment response
type FinalResponse struct {
	OrderID        string      `json:"order_id"`
	MerchantID     int         `json:"merchant_id"`
	Amount         string      `json:"amount"`
	Currency       string      `json:"currency"`
	OrderStatus    string      `json:"order_status"`
	ResponseStatus string      `json:"response_status"`
	Signature      string      `json:"signature"`
	MaskedCard     string      `json:"masked_card"`
	CardBin        interface{} `json:"card_bin"`
	CardType       string      `json:"card_type"`
	// ResponseCode sometimes is int and sometimes string
	ResponseCode        interface{} `json:"response_code,omitempty"`
	ResponseDescription string      `json:"response_description"`
	Eci                 string      `json:"eci"`
	PaymentID           int         `json:"payment_id"`
	Rectoken            string      `json:"rectoken"`
}
