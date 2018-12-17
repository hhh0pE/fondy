package fondy

type FinalResponse struct {
	OrderID             string `json:"order_id"`
	MerchantID          int    `json:"merchant_id"`
	Amount              int    `json:"amount"`
	Currency            string `json:"currency"`
	OrderStatus         string `json:"order_status"`
	ResponseStatus      string `json:"response_status"`
	Signature           string `json:"signature"`
	MaskedCard          string `json:"masked_card"`
	CardBin             int    `json:"card_bin"`
	CardType            string `json:"card_type"`
	ResponseCode        int    `json:"response_code"`
	ResponseDescription string `json:"response_description"`
	Eci                 int    `json:"eci"`
	PaymentID           int    `json:"payment_id"`
	Rectoken            string `json:"rectoken"`
}
