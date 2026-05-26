package queue

import (
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/spf13/viper"
)

// Client dùng để enqueue job từ handler/service
type Client struct {
	asynq *asynq.Client
}

func NewClient(cfg *viper.Viper) *Client {
	opt := asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", cfg.GetString("REDIS_HOST"), cfg.GetString("REDIS_PORT")),
		Password: cfg.GetString("REDIS_PASSWORD"),
		DB:       cfg.GetInt("REDIS_DB"),
	}
	return &Client{asynq: asynq.NewClient(opt)}
}

func (c *Client) Enqueue(task *asynq.Task, opts ...asynq.Option) error {
	_, err := c.asynq.Enqueue(task, opts...)
	return err
}

func (c *Client) Close() {
	c.asynq.Close()
}
