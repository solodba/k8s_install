package node2

import "context"

// 模块名称
const (
	AppName = "node2"
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
	// 加入k8s集群
	JoinK8s(context.Context) error
}
