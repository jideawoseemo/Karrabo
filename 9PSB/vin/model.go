package vin

type NinePSBVirtualAuthenticationRequest struct {
	Publickey  string `json:"publickey"`
	Privatekey string `json:"privatekey"`
}

type NinePSBCreateStaticVirtualRequest struct {
	Transaction         Transaction         `json:"transaction"`
	Order               Order               `json:"order"`
	Customer            Customer            `json:"customer"`
	Beneficiarytocredit Beneficiarytocredit `json:"benefiticiarytocredit"`
}

type Transaction struct {
	Reference string `json:"reference"`
}

type Order struct {
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
	Country     string  `json:"country"`
	AmountType  string  `json:"amountType"`
}

type Customer struct {
	Account Account `json:"account"`
}

type Account struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Beneficiarytocredit struct {
	Accountnumber string  `json:"accountnumber"`
	Bankcode      string  `json:"bankcode"`
	FeeAmount     float64 `json:"feeamount"`
}

type NinePSBVirtualAuthenticationResponse struct {
	Access_token string `json:"access_token"`
	Expires_in   int64  `json:"expires_in"`
	Code         string `json:"code"`
	Message      string `json:"message"`
}

type NinePSBStaticVirtualAccountResponse struct {
	Code        string      `json:"code"`
	Message     string      `json:"message"`
	Transaction Transaction `json:"transaction"`
	Order       Order       `json:"order"`
	Customer    Customer    `json:"customer"`
}

type AccountDReq struct {
	Number string `json:"number"`
}

type CustomerDReq struct {
	AccountDReq AccountDReq `json:"account""`
}

type AccountDResp struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Expiry string `json:"expiry"`
	Number string `json:"number"`
	Bank   string `json:"bank"`
}

type CustomerDResp struct {
	AccountDResp AccountDResp `json:"account"`
}

type NinePSBDeactivateVirtualAccountRequest struct {
	Transaction   Transaction  `json:"transaction"`
	CustomerDResp CustomerDReq `json:"customer"`
}

type NinePSBDeactivateVirtualAccountResponse struct {
	Code          string        `json:"code"`
	Message       string        `json:"message"`
	Transaction   Transaction   `json:"transaction"`
	CustomerDResp CustomerDResp `json:"customer"`
}

type NinePSBReactivateVirtualAccountRequest struct {
	Transaction  Transaction  `json:"transaction"`
	CustomerDReq CustomerDReq `json:"customer"`
}

type NinePSBReactivateVirtualAccountResponse struct {
	Code          string        `json:"code"`
	Message       string        `json:"message"`
	Transaction   Transaction   `json:"transaction"`
	CustomerDResp CustomerDResp `json:"customer"`
}

type ReactivateVirtualAccountRequest struct {
	Transaction  Transaction  `json:"transaction"`
	CustomerDReq CustomerDReq `json:"customer"`
}

type ReactivateVirtualAccountResponse struct {
	Code          string        `json:"code"`
	Message       string        `json:"message"`
	Transaction   Transaction   `json:"transaction"`
	CustomerDResp CustomerDResp `json:"customer"`
}

type ExpiryR struct {
	Hours string `json:"hours"`
	Date  string `json:"date"`
}

type AccountRResp struct {
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	ExpiryR ExpiryR `json:"expiry"`
	Number  string  `json:"number"`
	Bank    string  `json:"bank"`
}

type CustomerRResp struct {
	AccountRResp AccountRResp `json:"account"`
}

type ReallocateVirtualAccountRequest struct {
	Transaction         Transaction         `json:"transaction"`
	Order               Order               `json:"order"`
	CustomerDReq        CustomerDReq        `json:"customer"`
	Beneficiarytocredit Beneficiarytocredit `json:"beneficiarytocredit"`
}

type ReallocateVirtualAccountResponse struct {
	Code          string        `json:"code"`
	Message       string        `json:"message"`
	Transaction   Transaction   `json:"transaction"`
	Order         Order         `json:"order"`
	CustomerRResp CustomerRResp `json:"customer"`
}

type ConfirmVirtualPayment struct {
	Reference     string `json:"reference"`
	Sessionid     string `json:"sessionid"`
	Amount        string `json:"amount"`
	Accountnumber string `json:"accountnumber"`
}
