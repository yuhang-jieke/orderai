package pkg

import (
	"fmt"
	"log"
	"sync"
)

// WebServiceFactory Web服务工厂 - 抽象工厂
type WebServiceFactory interface {
	CreateService(cfg *ServiceConfig) (WebService, error)
}

// GinHTTPFactory Gin HTTP服务工厂 - 具体工厂
type GinHTTPFactory struct{}

// NewGinHTTPFactory 创建Gin HTTP工厂
func NewGinHTTPFactory() *GinHTTPFactory {
	return &GinHTTPFactory{}
}

// CreateService 创建Gin HTTP服务
func (f *GinHTTPFactory) CreateService(cfg *ServiceConfig) (WebService, error) {
	return NewGinHTTPService(cfg), nil
}

// GRPCFactory gRPC服务工厂 - 具体工厂
type GRPCFactory struct{}

// NewGRPCFactory 创建gRPC工厂
func NewGRPCFactory() *GRPCFactory {
	return &GRPCFactory{}
}

// CreateService 创建gRPC服务
func (f *GRPCFactory) CreateService(cfg *ServiceConfig) (WebService, error) {
	return NewGRPCService(cfg), nil
}

// =====================================================
// 简单工厂模式 (Simple Factory)
// =====================================================

// SimpleFactory 简单工厂 - 根据类型创建服务
type SimpleFactory struct{}

// NewSimpleFactory 创建简单工厂
func NewSimpleFactory() *SimpleFactory {
	return &SimpleFactory{}
}

// CreateService 根据服务类型创建服务
func (f *SimpleFactory) CreateService(serviceType ServiceType, cfg *ServiceConfig) (WebService, error) {
	switch serviceType {
	case ServiceTypeHTTP:
		return NewGinHTTPService(cfg), nil
	case ServiceTypeGRPC:
		return NewGRPCService(cfg), nil
	default:
		return nil, fmt.Errorf("不支持的服务类型: %s", serviceType)
	}
}

// =====================================================
// 服务管理器 (Service Manager) - 管理多个服务
// =====================================================

// ServiceManager 服务管理器
type ServiceManager struct {
	services map[string]WebService
	mu       sync.RWMutex
}

// NewServiceManager 创建服务管理器
func NewServiceManager() *ServiceManager {
	return &ServiceManager{
		services: make(map[string]WebService),
	}
}

// AddService 添加服务
func (m *ServiceManager) AddService(name string, service WebService) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.services[name] = service
}

// GetService 获取服务
func (m *ServiceManager) GetService(name string) (WebService, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	service, ok := m.services[name]
	return service, ok
}

// StartAll 启动所有服务
func (m *ServiceManager) StartAll() error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for name, service := range m.services {
		go func(n string, s WebService) {
			log.Printf("[ServiceManager] 启动服务: %s", n)
			if err := s.Start(); err != nil {
				log.Printf("[ServiceManager] 服务 %s 启动失败: %v", n, err)
			}
		}(name, service)
	}
	return nil
}

// StopAll 停止所有服务
func (m *ServiceManager) StopAll() {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for name, service := range m.services {
		log.Printf("[ServiceManager] 停止服务: %s", name)
		// 这里简化处理，实际应该使用context
		_ = service.Stop(nil)
	}
}

// ListServices 列出所有服务
func (m *ServiceManager) ListServices() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	names := make([]string, 0, len(m.services))
	for name := range m.services {
		names = append(names, name)
	}
	return names
}

// RemoveService 移除服务
func (m *ServiceManager) RemoveService(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.services, name)
}
