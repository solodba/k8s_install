package master

import (
	"context"
)

// 模块名称
const (
	AppName = "master"
)

// k8s安装程序接口
type Service interface {
	// 上传docker安装文件
	UploadDockerInstallFile(context.Context) error
	// docker安装
	UbuntuDockerInstall(context.Context) error
	// 上传containerd安装文件
	UploadContainerdInstallFile(context.Context) error
	// containerd安装
	UbuntuContainerdInstall(context.Context) error
	// 安装kubeadm、kubectl、kubelet
	UbuntuKubeadmKubectlKuheletInstall(context.Context) error
	// 下载安装k8s镜像
	DownloadK8sImage(context.Context) error
	// 初始化k8s集群
	InitialK8s(context.Context) error
	// 获取token
	GetJoinK8sToken(context.Context) (*TokenList, error)
}
