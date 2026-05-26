package shipping

// CreateOrderResponse nhận về sau khi tạo đơn giao hàng thành công
type CreateOrderResponse struct {
	OrderCode   string  `json:"order_code"`
	Status      string  `json:"status"`       // created, picking, delivering, delivered
	TrackingURL string  `json:"tracking_url"`
	Fee         float64 `json:"fee"`
	EstimatedAt string  `json:"estimated_at"`
}

// TrackingResponse nhận về khi theo dõi trạng thái đơn hàng
type TrackingResponse struct {
	OrderCode string         `json:"order_code"`
	Status    string         `json:"status"`
	Logs      []TrackingLog  `json:"logs"`
}

type TrackingLog struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

// CancelOrderResponse nhận về sau khi hủy đơn
type CancelOrderResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// apiResponse wrapper chung nếu Project B trả về dạng { success, data, message }
type apiResponse[T any] struct {
	Success bool   `json:"success"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}
