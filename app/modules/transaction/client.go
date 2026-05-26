package transaction

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
			cfg.GetString("TRANSACTION_SERVICE_URL"),
			httpclient.WithAPIKey(cfg.GetString("TRANSACTION_SERVICE_API_KEY")),
		),
	}
}

// Create tạo giao dịch mới
func (c *Client) Create(ctx context.Context, req CreateTransactionRequest) (*TransactionResponse, error) {
	var wrapper apiResponse[TransactionResponse]
	if err := c.http.Post(ctx, "/api/transactions", req, &wrapper); err != nil {
		return nil, fmt.Errorf("transaction.Create: %w", err)
	}
	return &wrapper.Data, nil
}

// GetByID lấy chi tiết giao dịch
func (c *Client) GetByID(ctx context.Context, transactionID string) (*TransactionDetailResponse, error) {
	var wrapper apiResponse[TransactionDetailResponse]
	if err := c.http.Get(ctx, "/api/transactions/"+transactionID, &wrapper); err != nil {
		return nil, fmt.Errorf("transaction.GetByID: %w", err)
	}
	return &wrapper.Data, nil
}

// Refund hoàn tiền giao dịch
func (c *Client) Refund(ctx context.Context, req RefundTransactionRequest) (*RefundResponse, error) {
	var wrapper apiResponse[RefundResponse]
	if err := c.http.Post(ctx, "/api/transactions/refund", req, &wrapper); err != nil {
		return nil, fmt.Errorf("transaction.Refund: %w", err)
	}
	return &wrapper.Data, nil
}
