package shipping

// Address dùng chung cho sender và receiver
type Address struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Street   string `json:"street"`
	District string `json:"district"`
	City     string `json:"city"`
}

// CreateOrderRequest gửi sang Project B để tạo đơn giao hàng
type CreateOrderRequest struct {
	ReferenceID string  `json:"reference_id"` // ID nội bộ để đối soát
	Sender      Address `json:"sender"`
	Receiver    Address `json:"receiver"`
	WeightGram  int     `json:"weight_gram"`
	CODAmount   float64 `json:"cod_amount"`   // 0 nếu không thu hộ
	Note        string  `json:"note"`
}

// CancelOrderRequest gửi sang Project B để hủy đơn
type CancelOrderRequest struct {
	OrderCode string `json:"order_code"`
	Reason    string `json:"reason"`
}
