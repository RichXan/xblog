package designpattern

import (
	"fmt"
	"time"
)

// 函数选项模式
// 简化了初始化结构体字段过多的问题，解决了创建结构体时自带默认值。
type Client struct {
	Name      string
	Timeout   time.Duration
	WriteTime time.Duration
	ReadTime  time.Duration
}

type Option func(*Client)

func WithName(name string) Option {
	return func(c *Client) {
		c.Name = name
	}
}

func WithTimeout(time time.Duration) Option {
	return func(c *Client) {
		c.Timeout = time
	}
}
func WithWriteTime(time time.Duration) Option {
	return func(c *Client) {
		c.WriteTime = time
	}
}
func WithReadTime(time time.Duration) Option {
	return func(c *Client) {
		c.ReadTime = time
	}
}

func NewClient(opts ...Option) *Client {
	client := &Client{
		Name:      "localhost",
		Timeout:   time.Duration(10),
		WriteTime: time.Duration(10),
		ReadTime:  time.Duration(10),
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func CreateClient() {
	clt := NewClient(WithName("xan"))
	fmt.Println("clent: ", clt)
}
