package shipping

import (
	"context"
	"fmt"
	"myapp/pkg/httpclient"

	"github.com/spf13/viper"
)

type Client struct {
	http *httpclient.Client
}

func NewClient(cfg *viper.Viper) *Client {
	return &Client{
		http: httpclient.New(
			cfg.GetString("SHIPPING_SERVICE_URL"),
			httpclient.WithAPIKey(cfg.GetString("SHIPPING_SERVICE_API_KEY")),
		),
	}
}

// CreateOrder tạo đơn giao hàng mới
func (c *Client) CreateOrder(ctx context.Context, req CreateOrderRequest) (*CreateOrderResponse, error) {
	var wrapper apiResponse[CreateOrderResponse]
	if err := c.http.Post(ctx, "/api/orders", req, &wrapper); err != nil {
		return nil, fmt.Errorf("shipping.CreateOrder: %w", err)
	}
	return &wrapper.Data, nil
}

// Track theo dõi trạng thái đơn hàng
func (c *Client) Track(ctx context.Context, orderCode string) (*TrackingResponse, error) {
	var wrapper apiResponse[TrackingResponse]
	if err := c.http.Get(ctx, "/api/orders/"+orderCode+"/tracking", &wrapper); err != nil {
		return nil, fmt.Errorf("shipping.Track: %w", err)
	}
	return &wrapper.Data, nil
}

// Cancel hủy đơn giao hàng
func (c *Client) Cancel(ctx context.Context, req CancelOrderRequest) (*CancelOrderResponse, error) {
	var wrapper apiResponse[CancelOrderResponse]
	if err := c.http.Post(ctx, "/api/orders/cancel", req, &wrapper); err != nil {
		return nil, fmt.Errorf("shipping.Cancel: %w", err)
	}
	return &wrapper.Data, nil
}
