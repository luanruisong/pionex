package order

import (
	"github.com/luanruisong/pionex/api"
	"net/http"
)

type (
	// NewOrder
	NewOrderReq struct {
		Symbol        string `json:"symbol"`
		Side          string `json:"side"`
		Type          string `json:"type"`
		ClientOrderId string `json:"clientOrderId,omitempty"`
		Size          string `json:"size,omitempty"`
		Price         string `json:"price,omitempty"`
		Amount        string `json:"amount,omitempty"`
		IOC           bool   `json:"IOC,omitempty"`
	}

	NewOrderRes struct {
		OrderId       int64  `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	}

	CreateOrder struct {
		Side          string `json:"side"`
		Type          string `json:"type"`
		ClientOrderId string `json:"clientOrderId,omitempty"`
		Size          string `json:"size,omitempty"`
		Price         string `json:"price,omitempty"`
	}
	// NewMultipleOrder
	NewMultipleOrderReq struct {
		Symbol string        `json:"symbol"`
		Orders []CreateOrder `json:"orders"`
	}

	NewMultipleOrderRes struct {
		OrderIds []NewOrderRes `json:"orderIds"`
	}

	// CancelOrder
	CancelOrderReq struct {
		Symbol  string `json:"symbol"`
		OrderId int64  `json:"orderId"`
	}

	// GetOrder
	GetOrderReq struct {
		OrderId int64 `form:"orderId"` // Order id.
	}

	GetOrderRes struct {
		OrderId       int64  `json:"orderId"`       // Order id.
		Symbol        string `json:"symbol"`        // Symbol.
		Type          string `json:"type"`          // LIMIT / MARKET.
		Side          string `json:"side"`          // BUY / SELL.
		Price         string `json:"price"`         // Price.
		Size          string `json:"size"`          // Order quantity.
		Amount        string `json:"amount"`        // The amount of market buy order.
		FilledSize    string `json:"filledSize"`    // Filled quantity of order.
		FilledAmount  string `json:"filledAmount"`  // Filled amount of order.
		Fee           string `json:"fee"`           // Transaction fee.
		FeeCoin       string `json:"feeCoin"`       // Currency of transaction fee.
		Status        string `json:"status"`        // OPEN / CLOSED.
		IOC           bool   `json:"IOC"`           // IOC
		ClientOrderId string `json:"clientOrderId"` // Client id.
		Source        string `json:"source"`        // Source of order, MANUAL / API
		CreateTime    int64  `json:"createTime"`
		UpdateTime    int64  `json:"updateTime"`
	}

	// GetOrderByClientId
	GetOrderByClientIdReq struct {
		ClientOrderId int64 `form:"clientOrderId"` // Order id.
	}

	// GetOpenOrders
	GetOpenOrdersReq struct {
		Symbol string `form:"symbol"`
	}

	OrdersRes struct {
		Orders []GetOrderRes `json:"orders"`
	}

	// GetAllOrders
	GetAllOrdersReq struct {
		Symbol    string `form:"symbol"`
		StartTime int64  `form:"startTime"`
		EndTime   int64  `form:"endTime"`
		Limit     int64  `form:"limit"`
	}

	// GetFills
	GetFillsReq struct {
		Symbol    string `form:"symbol"`
		StartTime int64  `form:"startTime"`
		EndTime   int64  `form:"endTime"`
	}

	FillInfo struct {
		Id        int64  `json:"id"`
		OrderId   int    `json:"orderId"`
		Symbol    string `json:"symbol"`
		Side      string `json:"side"`
		Role      string `json:"role"`
		Price     string `json:"price"`
		Size      string `json:"size"`
		Fee       string `json:"fee"`
		FeeCoin   string `json:"feeCoin"`
		Timestamp int64  `json:"timestamp"`
	}

	GetFillsRes struct {
		Fills []FillInfo `json:"fills"`
	}

	// GetFills
	GetFillsByOrderIdeq struct {
		OrderId int64 `form:"orderId"`
		FromId  int64 `form:"fromId"`
	}

	// CancelAllOrders
	CancelAllOrdersReq struct {
		Symbol string `json:"symbol"`
	}

	Trans struct {
		s *api.Singer
		c *http.Client
	}
)

var (
	// apis
	newOrder = &api.Api[*NewOrderReq, *NewOrderRes]{
		Method: http.MethodPost,
		Path:   "/api/v1/trade/order",
	}

	newMultipleOrder = &api.Api[*NewMultipleOrderReq, *NewMultipleOrderRes]{
		Method: http.MethodPost,
		Path:   "/api/v1/trade/massOrder",
	}

	cancelOrder = &api.Api[*CancelOrderReq, struct{}]{
		Method: http.MethodDelete,
		Path:   "/api/v1/trade/order",
	}

	getOrder = &api.Api[*GetOrderReq, *GetOrderRes]{
		Method: http.MethodGet,
		Path:   "/api/v1/trade/order",
	}

	getOrderByClientId = &api.Api[*GetOrderByClientIdReq, *GetOrderRes]{
		Method: http.MethodGet,
		Path:   "/api/v1/trade/orderByClientOrderId",
	}

	getOpenOrders = &api.Api[*GetOpenOrdersReq, *OrdersRes]{
		Method: http.MethodGet,
		Path:   "/api/v1/trade/openOrders",
	}

	getAllOrders = &api.Api[*GetAllOrdersReq, *OrdersRes]{
		Method: http.MethodGet,
		Path:   "/api/v1/trade/allOrders",
	}

	getFills = &api.Api[*GetFillsReq, *GetFillsRes]{
		Method: http.MethodGet,
		Path:   "/api/v1/trade/fills",
	}

	getFillsByOrderId = &api.Api[*GetFillsByOrderIdeq, *GetFillsRes]{
		Method: http.MethodGet,
		Path:   "/api/v1/trade/fillsByOrderId",
	}

	cancelAllOrders = &api.Api[*CancelAllOrdersReq, struct{}]{
		Method: http.MethodDelete,
		Path:   "/api/v1/trade/allOrders",
	}
)

// NewOrder https://pionex-doc.gitbook.io/apidocs/restful/orders/new-order
func (t *Trans) NewOrder(req *NewOrderReq) (*api.Ret[*NewOrderRes], error) {
	return newOrder.Do(req, t.s, t.c)
}

// NewMultipleOrder https://pionex-doc.gitbook.io/apidocs/restful/orders/new-multiple-order
func (t *Trans) NewMultipleOrder(req *NewMultipleOrderReq) (*api.Ret[*NewMultipleOrderRes], error) {
	return newMultipleOrder.Do(req, t.s, t.c)
}

// CancelOrder https://pionex-doc.gitbook.io/apidocs/restful/orders/cancel-order
func (t *Trans) CancelOrder(req *CancelOrderReq) (*api.Ret[struct{}], error) {
	return cancelOrder.Do(req, t.s, t.c)
}

// GetOrder https://pionex-doc.gitbook.io/apidocs/restful/orders/get-order
func (t *Trans) GetOrder(req *GetOrderReq) (*api.Ret[*GetOrderRes], error) {
	return getOrder.Do(req, t.s, t.c)
}

// GetOrderByClientId https://pionex-doc.gitbook.io/apidocs/restful/orders/get-order-by-client-order-id
func (t *Trans) GetOrderByClientId(req *GetOrderByClientIdReq) (*api.Ret[*GetOrderRes], error) {
	return getOrderByClientId.Do(req, t.s, t.c)
}

// GetOpenOrders https://pionex-doc.gitbook.io/apidocs/restful/orders/get-open-orders
func (t *Trans) GetOpenOrders(req *GetOpenOrdersReq) (*api.Ret[*OrdersRes], error) {
	return getOpenOrders.Do(req, t.s, t.c)
}

// GetAllOrders https://pionex-doc.gitbook.io/apidocs/restful/orders/get-all-orders
func (t *Trans) GetAllOrders(req *GetAllOrdersReq) (*api.Ret[*OrdersRes], error) {
	return getAllOrders.Do(req, t.s, t.c)
}

// GetFills https://pionex-doc.gitbook.io/apidocs/restful/orders/get-fills
func (t *Trans) GetFills(req *GetFillsReq) (*api.Ret[*GetFillsRes], error) {
	return getFills.Do(req, t.s, t.c)
}

// GetFillsByOrderId https://pionex-doc.gitbook.io/apidocs/restful/orders/get-fills-by-order-id
func (t *Trans) GetFillsByOrderId(req *GetFillsByOrderIdeq) (*api.Ret[*GetFillsRes], error) {
	return getFillsByOrderId.Do(req, t.s, t.c)
}

// CancelAllOrders https://pionex-doc.gitbook.io/apidocs/restful/orders/cancel-all-orders
func (t *Trans) CancelAllOrders(req *CancelAllOrdersReq) (*api.Ret[struct{}], error) {
	return cancelAllOrders.Do(req, t.s, t.c)
}

func NewTrans(s *api.Singer, c *http.Client) *Trans {
	return &Trans{
		s: s,
		c: c,
	}
}
