package protocol

import (
	"context"

	"github.com/solodba/k8s_install/apps/master"
	"github.com/solodba/k8s_install/apps/node1"
	"github.com/solodba/k8s_install/apps/node2"
	"github.com/solodba/mcube/apps"
	"github.com/solodba/mcube/logger"
)

var (
	ctx = context.Background()
)

// k8s安装服务结构体
type K8sInstallSvc struct {
	masterSvc master.Service
	node1Svc  node1.Service
	node2Svc  node2.Service
}

// k8s安装服务结构体构造函数
func Newk8sInstallSvc() *K8sInstallSvc {
	return &K8sInstallSvc{
		masterSvc: apps.GetInternalApp(master.AppName).(master.Service),
		node1Svc:  apps.GetInternalApp(node1.AppName).(node1.Service),
		node2Svc:  apps.GetInternalApp(node2.AppName).(node2.Service),
	}
}

// K8s安装服务
func (m *K8sInstallSvc) K8sInstall(ctx context.Context) error {
	// master节点安装containerd
	logger.L().Info().Msgf("======主节点开始安装Containerd======")
	err := m.masterSvc.UploadContainerdCompressInstallFile(ctx)
	if err != nil {
		return err
	}
	err = m.masterSvc.UncompressContainerdInstallFIle(ctx)
	if err != nil {
		return err
	}
	err = m.masterSvc.UbuntuContainerdInstall(ctx)
	if err != nil {
		return err
	}
	err = m.masterSvc.UbuntuKubeadmKubectlKuheletInstall(ctx)
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("======主节点安装Containerd成功======")
	// node1节点安装containerd
	logger.L().Info().Msgf("======Node1节点开始安装Containerd======")
	err = m.node1Svc.UploadContainerdCompressInstallFile(ctx)
	if err != nil {
		return err
	}
	err = m.node1Svc.UncompressContainerdInstallFIle(ctx)
	if err != nil {
		return err
	}
	err = m.node1Svc.UbuntuContainerdInstall(ctx)
	if err != nil {
		return err
	}
	err = m.node1Svc.UbuntuKubeadmKubectlKuheletInstall(ctx)
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("======Node1节点安装Containerd成功======")
	// node2节点安装containerd
	logger.L().Info().Msgf("======Node2节点开始安装Containerd======")
	err = m.node2Svc.UploadContainerdCompressInstallFile(ctx)
	if err != nil {
		return err
	}
	err = m.node2Svc.UncompressContainerdInstallFIle(ctx)
	if err != nil {
		return err
	}
	err = m.node2Svc.UbuntuContainerdInstall(ctx)
	if err != nil {
		return err
	}
	err = m.node2Svc.UbuntuKubeadmKubectlKuheletInstall(ctx)
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("======Node2节点安装Containerd成功======")
	// master节点下载安装k8s镜像
	logger.L().Info().Msgf("======master节点开始下载k8s镜像======")
	err = m.masterSvc.DownloadK8sImage(ctx)
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("======master节点下载k8s镜像成功======")
	// master节点初始化k8s集群
	logger.L().Info().Msgf("======master节点开始初始化k8s集群======")
	err = m.masterSvc.InitialK8s(ctx)
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("======master节点初始化k8s集群成功======")
	// node1节点加入k8s集群
	logger.L().Info().Msgf("======node1节点开始加入k8s集群======")
	err = m.node1Svc.JoinK8s(ctx)
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("======node1节点加入k8s集群成功======")
	// node2节点加入k8s集群
	logger.L().Info().Msgf("======node2节点开始加入k8s集群======")
	err = m.node2Svc.JoinK8s(ctx)
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("======node2节点加入k8s集群成功======")
	return nil
}
