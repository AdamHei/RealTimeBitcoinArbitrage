package tickermodels

type ITicker interface {
	GetExchangeData() LimitedJson
}

type LimitedJson map[string]map[string]string

const BID = "Bid"
const ASK = "Ask"
