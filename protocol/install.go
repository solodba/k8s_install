package protocol

import (
	"context"
)

var (
	ctx = context.Background()
)

// k8s安装服务结构体
type K8sInstallSvc struct {
}

// k8s安装服务结构体构造函数
func Newk8sInstallSvc() *K8sInstallSvc {
	return &K8sInstallSvc{}
}

// K8s安装服务
func (m *K8sInstallSvc) K8sInstall(ctx context.Context) error {
	return nil
}
