package tickermodels

type ITicker interface {
	GetExchangeData() Ticker
}

type Ticker struct {
	Exchange string `json:"exchange"`
	Bid      string `json:"bid"`
	Ask      string `json:"ask"`
}
