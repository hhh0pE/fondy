package fondy

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	MerchantID int
	Password   string
}

// Request is an wrapper for PaymentRequestParameters, needed for payment creation request
type Request struct {
	PaymentRequestParameters `json:"request"`
}

// PaymentRequestParameters contains mandatory data for payment creation
type PaymentRequestParameters struct {
	Amount            int     `json:"amount"`
	Currency          string  `json:"currency"`
	MerchantID        int     `json:"merchant_id"`
	OrderDesc         string  `json:"order_desc"`
	OrderID           string  `json:"order_id"`
	Rectoken          *string `json:"rectoken,omitempty"`
	ResponseURL       *string `json:"response_url,omitempty"`
	RequiredRectoken  *string `json:"required_rectoken,omitempty"`
	ServerCallbackURL *string `json:"server_callback_url,omitempty"`
	Signature         string  `json:"signature"`
}

// GenerateSignature generates payment signature needed for authenticating the payment.
// The user is supposed to generate signature for payment request
// TODO tests
func (c Client) GenerateSignature(p PaymentRequestParameters) string {
	sig := sha1.New()

	sigString := fmt.Sprintf("%s|%d|%s|%d|%s|%s", c.Password, p.Amount, p.Currency,
		p.MerchantID, p.OrderDesc, p.OrderID)

	// TODO upgrade it somehow
	if p.Rectoken != nil {
		sigString += "|" + *p.Rectoken
	}
	if p.ResponseURL != nil {
		sigString += "|" + *p.ResponseURL
	}
	if p.RequiredRectoken != nil {
		sigString += "|" + *p.RequiredRectoken
	}
	if p.ServerCallbackURL != nil {
		sigString += "|" + *p.ServerCallbackURL
	}

	fmt.Fprint(sig, sigString)
	return hex.EncodeToString(sig.Sum(nil))
}

type Response struct {
	InterimResponseParameters `json:"response"`
}

// InterimResponseParameters is the interim response, which returns URL where
// user needs to finish his payment
type InterimResponseParameters struct {
	ResponseStatus string `json:"response_status"`

	// Fields below are filled when ResponseStatus is "success"
	CheckoutURL string `json:"checkout_url"`
	PaymentID   string `json:"payment_id"`

	// Fields below are filled when ResponseStatus is "failure"
	ErrorCode    *int    `json:"error_code"`
	ErrorMessage *string `json:"error_message"`
}

const requestPaymentURL = "https://api.fondy.eu/api/checkout/url"

func (c Client) RequestPayment(r Request) (Response, error) {
	payload, err := json.Marshal(r)
	if err != nil {
		return Response{}, err
	}
	buf := bytes.NewBuffer(payload)
	request, err := http.NewRequest(http.MethodPost, requestPaymentURL, buf)
	if err != nil {
		return Response{}, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	var responseParams Response
	err = json.NewDecoder(resp.Body).Decode(&responseParams)
	return responseParams, err
}
