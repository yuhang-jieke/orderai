package pkg

import "context"

// WebService Web服务接口 - 抽象产品
type WebService interface {
	// Start 启动服务
	Start() error

	// Stop 停止服务
	Stop(ctx context.Context) error

	// Name 获取服务名称
	Name() string

	// Addr 获取服务地址
	Addr() string
}

// ServiceConfig 服务配置
type ServiceConfig struct {
	Name    string // 服务名称
	Host    string // 主机地址
	Port    int    // 端口
	Timeout int    // 超时时间(秒)
}

// ServiceType 服务类型枚举
type ServiceType string

const (
	ServiceTypeHTTP ServiceType = "http" // HTTP服务 (Gin)
	ServiceTypeGRPC ServiceType = "grpc" // gRPC服务
)
