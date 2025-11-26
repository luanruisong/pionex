package order

import (
	"net/http"

	"github.com/luanruisong/pionex/api"
	"github.com/valyala/fasthttp"
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
		newOrder           *api.Api[*NewOrderReq, *NewOrderRes]
		newMultipleOrder   *api.Api[*NewMultipleOrderReq, *NewMultipleOrderRes]
		cancelOrder        *api.Api[*CancelOrderReq, struct{}]
		getOrder           *api.Api[*GetOrderReq, *GetOrderRes]
		getOrderByClientId *api.Api[*GetOrderByClientIdReq, *GetOrderRes]
		getOpenOrders      *api.Api[*GetOpenOrdersReq, *OrdersRes]
		getAllOrders       *api.Api[*GetAllOrdersReq, *OrdersRes]
		getFills           *api.Api[*GetFillsReq, *GetFillsRes]
		getFillsByOrderId  *api.Api[*GetFillsByOrderIdeq, *GetFillsRes]
		cancelAllOrders    *api.Api[*CancelAllOrdersReq, struct{}]
	}
)

var ()

// NewOrder https://pionex-doc.gitbook.io/apidocs/restful/orders/new-order
func (t *Trans) NewOrder(req *NewOrderReq) (*api.Ret[*NewOrderRes], error) {
	return t.newOrder.Do(req)
}

// NewMultipleOrder https://pionex-doc.gitbook.io/apidocs/restful/orders/new-multiple-order
func (t *Trans) NewMultipleOrder(req *NewMultipleOrderReq) (*api.Ret[*NewMultipleOrderRes], error) {
	return t.newMultipleOrder.Do(req)
}

// CancelOrder https://pionex-doc.gitbook.io/apidocs/restful/orders/cancel-order
func (t *Trans) CancelOrder(req *CancelOrderReq) (*api.Ret[struct{}], error) {
	return t.cancelOrder.Do(req)
}

// GetOrder https://pionex-doc.gitbook.io/apidocs/restful/orders/get-order
func (t *Trans) GetOrder(req *GetOrderReq) (*api.Ret[*GetOrderRes], error) {
	return t.getOrder.Do(req)
}

// GetOrderByClientId https://pionex-doc.gitbook.io/apidocs/restful/orders/get-order-by-client-order-id
func (t *Trans) GetOrderByClientId(req *GetOrderByClientIdReq) (*api.Ret[*GetOrderRes], error) {
	return t.getOrderByClientId.Do(req)
}

// GetOpenOrders https://pionex-doc.gitbook.io/apidocs/restful/orders/get-open-orders
func (t *Trans) GetOpenOrders(req *GetOpenOrdersReq) (*api.Ret[*OrdersRes], error) {
	return t.getOpenOrders.Do(req)
}

// GetAllOrders https://pionex-doc.gitbook.io/apidocs/restful/orders/get-all-orders
func (t *Trans) GetAllOrders(req *GetAllOrdersReq) (*api.Ret[*OrdersRes], error) {
	return t.getAllOrders.Do(req)
}

// GetFills https://pionex-doc.gitbook.io/apidocs/restful/orders/get-fills
func (t *Trans) GetFills(req *GetFillsReq) (*api.Ret[*GetFillsRes], error) {
	return t.getFills.Do(req)
}

// GetFillsByOrderId https://pionex-doc.gitbook.io/apidocs/restful/orders/get-fills-by-order-id
func (t *Trans) GetFillsByOrderId(req *GetFillsByOrderIdeq) (*api.Ret[*GetFillsRes], error) {
	return t.getFillsByOrderId.Do(req)
}

// CancelAllOrders https://pionex-doc.gitbook.io/apidocs/restful/orders/cancel-all-orders
func (t *Trans) CancelAllOrders(req *CancelAllOrdersReq) (*api.Ret[struct{}], error) {
	return t.cancelAllOrders.Do(req)
}

func NewTrans(s *api.Singer, c *fasthttp.Client) *Trans {
	return &Trans{
		newOrder: api.NewApi[*NewOrderReq, *NewOrderRes](
			http.MethodPost,
			"/api/v1/trade/order",
			api.WithClient(c),
			api.WithSigner(s),
		),
		newMultipleOrder: api.NewApi[*NewMultipleOrderReq, *NewMultipleOrderRes](
			http.MethodPost,
			"/api/v1/trade/massOrder",
			api.WithClient(c),
			api.WithSigner(s),
		),
		cancelOrder: api.NewApi[*CancelOrderReq, struct{}](
			http.MethodDelete,
			"/api/v1/trade/order",
			api.WithClient(c),
			api.WithSigner(s),
		),
		getOrder: api.NewApi[*GetOrderReq, *GetOrderRes](
			http.MethodGet,
			"/api/v1/trade/order",
			api.WithClient(c),
			api.WithSigner(s),
		),
		getOrderByClientId: api.NewApi[*GetOrderByClientIdReq, *GetOrderRes](
			http.MethodGet,
			"/api/v1/trade/orderByClientOrderId",
			api.WithClient(c),
			api.WithSigner(s),
		),
		getOpenOrders: api.NewApi[*GetOpenOrdersReq, *OrdersRes](
			http.MethodGet,
			"/api/v1/trade/openOrders",
			api.WithClient(c),
			api.WithSigner(s),
		),
		getAllOrders: api.NewApi[*GetAllOrdersReq, *OrdersRes](
			http.MethodGet,
			"/api/v1/trade/allOrders",
			api.WithClient(c),
			api.WithSigner(s),
		),
		getFills: api.NewApi[*GetFillsReq, *GetFillsRes](
			http.MethodGet,
			"/api/v1/trade/fills",
			api.WithClient(c),
			api.WithSigner(s),
		),
		getFillsByOrderId: api.NewApi[*GetFillsByOrderIdeq, *GetFillsRes](
			http.MethodGet,
			"/api/v1/trade/fillsByOrderId",
			api.WithClient(c),
			api.WithSigner(s),
		),
		cancelAllOrders: api.NewApi[*CancelAllOrdersReq, struct{}](
			http.MethodDelete,
			"/api/v1/trade/allOrders",
			api.WithClient(c),
			api.WithSigner(s),
		),
	}
}
