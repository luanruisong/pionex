package market

import (
	"net/http"

	"github.com/luanruisong/pionex/api"
	"github.com/valyala/fasthttp"
)

type (
	SymbolsReq struct {
		Symbols string `form:"symbols"` //Concatenate multiple symbols with ','
		Type    string `form:"type"`    //Type, if the symbol is specified, the type is irrelevant. If the symbol is not specified, the default is SPOT, with the possible values being SPOT / PERP.
	}
	SymbolInfo struct {
		Symbol             string `json:"symbol"`
		Name               string `json:"name"`            // Name, only for PERP.
		Type               string `json:"type"`            // SPOT / PERP.
		BaseCurrency       string `json:"baseCurrency"`    // Base coin.
		QuoteCurrency      string `json:"quoteCurrency"`   // Quote coin.
		BasePrecision      int    `json:"basePrecision"`   // Precision digits of base currency price.
		QuotePrecision     int    `json:"quotePrecision"`  // Precision digits of quote currency price.
		AmountPrecision    int    `json:"amountPrecision"` // Precision digits of the amount of market price buying order.
		MinAmount          string `json:"minAmount"`       // Minimum amount of the order, only for SPOT
		MinTradeSize       string `json:"minTradeSize"`    // Minimum limit order quantity.
		MaxTradeSize       string `json:"maxTradeSize"`    // Maximum limit order quantity.
		MinTradeDumping    string `json:"minTradeDumping"` // Minimum quantity of market price selling order.
		MaxTradeDumping    string `json:"maxTradeDumping"` // Maximum quantity of market price selling order.
		Enable             bool   `json:"enable"`          // Enable trading.
		BuyCeiling         string `json:"buyCeiling"`      // Maximum ratio of buying price, cannot be higher than a multiple of the latest highest buying price.
		SellFloor          string `json:"sellFloor"`       // Minimum ratio of selling price, cannot be lower than a multiple of the latest lowest selling price.
		MinNotional        string `json:"minNotional"`     // Only for PERP.
		BaseStep           string `json:"baseStep"`
		QuoteStep          string `json:"quoteStep"`
		MinSizeLimit       string `json:"minSizeLimit"`
		MaxSizeLimit       string `json:"maxSizeLimit"`
		MaxImpactLimit     string `json:"maxImpactLimit"`
		MinSizeMarket      string `json:"minSizeMarket"`
		MaxSizeMarket      string `json:"maxSizeMarket"`
		MaxImpactMarket    string `json:"maxImpactMarket"` // Max impact price of market order, only for PERP.
		MaxOrderNum        int    `json:"maxOrderNum"`
		Status             string `json:"status"`
		LiquidationFeeRate string `json:"liquidationFeeRate"` // Liquidation fee rate, only for PERP.
	}
	SymbolsRes struct {
		Symbols []SymbolInfo `json:"symbols"`
	}

	SymbolReq struct {
		Symbol string `form:"symbol"` //Concatenate multiple symbols with ','
		Limit  int64  `form:"limit"`  //Type, if the symbol is specified, the type is irrelevant. If the symbol is not specified, the default is SPOT, with the possible values being SPOT / PERP.
	}

	Trade struct {
		Symbol    string `json:"symbol"`
		TradeId   string `json:"tradeId"`
		Price     string `json:"price"`
		Size      string `json:"size"`
		Side      string `json:"side"`
		Timestamp int64  `json:"timestamp"`
	}

	GetTradesRes struct {
		Trades []Trade `json:"trades"`
	}

	GetDepthRes struct {
		Bids       [][]string `json:"bids"`
		Asks       [][]string `json:"asks"`
		UpdateTime int64      `json:"updateTime"`
	}

	TickerReq struct {
		Symbol string `form:"symbol"` //Concatenate multiple symbols with ','
		Type   string `form:"type"`   //Type, if the symbol is specified, the type is irrelevant. If the symbol is not specified, the default is SPOT, with the possible values being SPOT / PERP.
	}

	Ticker struct {
		Symbol string `json:"symbol"`
		Time   int64  `json:"time"`
		Open   string `json:"open"`
		Close  string `json:"close"`
		High   string `json:"high"`
		Low    string `json:"low"`
		Volume string `json:"volume"`
		Amount string `json:"amount"`
		Count  int    `json:"count"`
	}

	Get24hrTickerRes struct {
		Tickers []Ticker `json:"tickers"`
	}

	BookTicker struct {
		Symbol    string `json:"symbol"`
		BidPrice  string `json:"bidPrice"`
		BidSize   string `json:"bidSize"`
		AskPrice  string `json:"askPrice"`
		AskSize   string `json:"askSize"`
		Timestamp int64  `json:"timestamp"`
	}

	GetBookTickerRes struct {
		Tickers []BookTicker `json:"tickers"`
	}

	GetKlineReq struct {
		Symbol   string `form:"symbol"`
		Interval string `form:"interval"`
		EndTime  int64  `form:"endTime"`
		Limit    int64  `form:"limit"`
	}
	Kline struct {
		Time   int64  `json:"time"`
		Open   string `json:"open"`
		Close  string `json:"close"`
		High   string `json:"high"`
		Low    string `json:"low"`
		Volume string `json:"volume"`
	}

	GetKlineRes struct {
		Klines []Kline `json:"klines"`
	}

	Market struct {
		symbols       *api.Api[*SymbolsReq, *SymbolsRes]
		getDepth      *api.Api[*SymbolReq, *GetDepthRes]
		getTrades     *api.Api[*SymbolReq, *GetTradesRes]
		get24hrTicker *api.Api[*TickerReq, *Get24hrTickerRes]
		getBookTicker *api.Api[*TickerReq, *GetBookTickerRes]
		getKline      *api.Api[*GetKlineReq, *GetKlineRes]
	}
)

// MarketData https://pionex-doc.gitbook.io/apidocs/restful/common/market-data
func (m *Market) GetSymbols(req *SymbolsReq) (*api.Ret[*SymbolsRes], error) {
	return m.symbols.Do(req)
}

// GetDepth https://pionex-doc.gitbook.io/apidocs/restful/markets/get-depth
func (m *Market) GetDepth(req *SymbolReq) (*api.Ret[*GetDepthRes], error) {
	return m.getDepth.Do(req)
}

// GetTrades https://pionex-doc.gitbook.io/apidocs/restful/markets/get-trades
func (m *Market) GetTrades(req *SymbolReq) (*api.Ret[*GetTradesRes], error) {
	return m.getTrades.Do(req)
}

// Get24hrTicker https://pionex-doc.gitbook.io/apidocs/restful/markets/get-24hr-ticker
func (m *Market) Get24hrTicker(req *TickerReq) (*api.Ret[*Get24hrTickerRes], error) {
	return m.get24hrTicker.Do(req)
}

// GetBookTicker https://pionex-doc.gitbook.io/apidocs/restful/markets/get-book-ticker
func (m *Market) GetBookTicker(req *TickerReq) (*api.Ret[*GetBookTickerRes], error) {
	return m.getBookTicker.Do(req)
}

// GetKline https://pionex-doc.gitbook.io/apidocs/restful/markets/get-klines
func (m *Market) GetKline(req *GetKlineReq) (*api.Ret[*GetKlineRes], error) {
	return m.getKline.Do(req)
}

func NewMarket(c *fasthttp.Client) *Market {
	return &Market{
		symbols:       api.NewApi[*SymbolsReq, *SymbolsRes](http.MethodGet, "/api/v1/common/symbols", api.WithClient(c)),
		getDepth:      api.NewApi[*SymbolReq, *GetDepthRes](http.MethodGet, "/api/v1/market/depth", api.WithClient(c)),
		getTrades:     api.NewApi[*SymbolReq, *GetTradesRes](http.MethodGet, "/api/v1/market/trades", api.WithClient(c)),
		get24hrTicker: api.NewApi[*TickerReq, *Get24hrTickerRes](http.MethodGet, "/api/v1/market/tickers", api.WithClient(c)),
		getBookTicker: api.NewApi[*TickerReq, *GetBookTickerRes](http.MethodGet, "/api/v1/market/bookTickers", api.WithClient(c)),
		getKline:      api.NewApi[*GetKlineReq, *GetKlineRes](http.MethodGet, "/api/v1/market/klines", api.WithClient(c)),
	}
}
