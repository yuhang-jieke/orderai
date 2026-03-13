package pkg

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuhang-jieke/orderai/srv/api-getaway/router"
)

// GinHTTPService Gin HTTP服务实现 - 具体产品
type GinHTTPService struct {
	config    *ServiceConfig
	server    *http.Server
	engine    *gin.Engine
	isRunning bool
}

// NewGinHTTPService 创建Gin HTTP服务实例
func NewGinHTTPService(cfg *ServiceConfig) *GinHTTPService {
	if cfg == nil {
		cfg = &ServiceConfig{
			Name:    "gin-http-service",
			Host:    "",
			Port:    8080,
			Timeout: 5,
		}
	}
	if cfg.Port == 0 {
		cfg.Port = 8080
	}
	if cfg.Timeout == 0 {
		cfg.Timeout = 5
	}

	return &GinHTTPService{
		config: cfg,
	}
}

// Start 启动Gin HTTP服务
func (s *GinHTTPService) Start() error {
	// 获取路由引擎
	s.engine = router.Router()

	// 设置服务模式
	gin.SetMode(gin.ReleaseMode)

	// 创建HTTP服务器
	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	s.server = &http.Server{
		Addr:         addr,
		Handler:      s.engine,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// 启动信号监听
	go s.listenSignals()

	// 启动服务
	log.Printf("[GinHTTP] 服务启动中... 地址: %s", addr)
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("Gin HTTP服务启动失败: %w", err)
	}

	return nil
}

// Stop 停止Gin HTTP服务
func (s *GinHTTPService) Stop(ctx context.Context) error {
	if s.server == nil {
		return nil
	}

	log.Println("[GinHTTP] 服务正在关闭...")

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("Gin HTTP服务关闭失败: %w", err)
	}

	s.isRunning = false
	log.Println("[GinHTTP] 服务已关闭")
	return nil
}

// Name 获取服务名称
func (s *GinHTTPService) Name() string {
	return s.config.Name
}

// Addr 获取服务地址
func (s *GinHTTPService) Addr() string {
	return fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
}

// listenSignals 监听系统信号
func (s *GinHTTPService) listenSignals() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[GinHTTP] 接收到关闭信号...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.config.Timeout)*time.Second)
	defer cancel()

	if err := s.Stop(ctx); err != nil {
		log.Printf("[GinHTTP] 关闭错误: %v", err)
	}
}

// SetRouter 设置自定义路由
func (s *GinHTTPService) SetRouter(engine *gin.Engine) {
	s.engine = engine
}

// GetEngine 获取Gin引擎
func (s *GinHTTPService) GetEngine() *gin.Engine {
	return s.engine
}
