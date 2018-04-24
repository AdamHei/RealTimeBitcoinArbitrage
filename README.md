# honorsproject
API for providing real-time Bitcoin arbitrage data across several exchanges

## Installation
`go get github.com/adamhei/honorsproject`

## Usage
`go run driver.go`

## Endpoints
- `/live-spread/widest` - Best opportunity at the given moment
- `/live-spread/all` - all exchange bid/ask data

### Currently supported exchanges:
- Binance
- Bitfinex
- Bitstamp
- ExMo
- GDAX
- Gemini
- Kraken
- Poloniex
