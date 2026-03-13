package pkg

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

// GRPCService gRPC服务实现 - 具体产品
type GRPCService struct {
	config        *ServiceConfig
	server        *grpc.Server
	listener      net.Listener
	isRunning     bool
	registerFuncs []func(*grpc.Server) // 服务注册函数列表
}

// NewGRPCService 创建gRPC服务实例
func NewGRPCService(cfg *ServiceConfig) *GRPCService {
	if cfg == nil {
		cfg = &ServiceConfig{
			Name:    "grpc-service",
			Host:    "",
			Port:    50051,
			Timeout: 5,
		}
	}
	if cfg.Port == 0 {
		cfg.Port = 50051
	}
	if cfg.Timeout == 0 {
		cfg.Timeout = 5
	}

	return &GRPCService{
		config:        cfg,
		registerFuncs: make([]func(*grpc.Server), 0),
	}
}

// RegisterService 注册gRPC服务
func (s *GRPCService) RegisterService(registerFunc func(*grpc.Server)) {
	s.registerFuncs = append(s.registerFuncs, registerFunc)
}

// Start 启动gRPC服务
func (s *GRPCService) Start() error {
	// 创建监听器
	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("gRPC监听失败: %w", err)
	}
	s.listener = listener

	// 创建gRPC服务器
	s.server = grpc.NewServer()

	// 注册所有服务
	for _, registerFunc := range s.registerFuncs {
		registerFunc(s.server)
	}

	// 启动信号监听
	go s.listenSignals()

	// 启动服务
	log.Printf("[GRPC] 服务启动中... 地址: %s", addr)
	s.isRunning = true
	if err := s.server.Serve(listener); err != nil {
		return fmt.Errorf("gRPC服务启动失败: %w", err)
	}

	return nil
}

// Stop 停止gRPC服务
func (s *GRPCService) Stop(ctx context.Context) error {
	if s.server == nil {
		return nil
	}

	log.Println("[GRPC] 服务正在关闭...")

	// 优雅停止
	stopped := make(chan struct{})
	go func() {
		s.server.GracefulStop()
		close(stopped)
	}()

	select {
	case <-ctx.Done():
		// 超时，强制停止
		s.server.Stop()
		log.Println("[GRPC] 服务强制关闭")
	case <-stopped:
		log.Println("[GRPC] 服务优雅关闭")
	}

	s.isRunning = false
	return nil
}

// Name 获取服务名称
func (s *GRPCService) Name() string {
	return s.config.Name
}

// Addr 获取服务地址
func (s *GRPCService) Addr() string {
	return fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
}

// listenSignals 监听系统信号
func (s *GRPCService) listenSignals() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[GRPC] 接收到关闭信号...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.config.Timeout)*time.Second)
	defer cancel()

	if err := s.Stop(ctx); err != nil {
		log.Printf("[GRPC] 关闭错误: %v", err)
	}
}

// GetServer 获取gRPC服务器实例
func (s *GRPCService) GetServer() *grpc.Server {
	return s.server
}

// IsRunning 检查服务是否运行中
func (s *GRPCService) IsRunning() bool {
	return s.isRunning
}
