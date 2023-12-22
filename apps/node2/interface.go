package node2

import "context"

// 模块名称
const (
	AppName = "node2"
)

// k8s安装程序接口
type Service interface {
	// 上传docker压缩安装文件
	UploadDockerCompressInstallFile(context.Context) error
	// 解压docker压缩安装文件
	UncompressDockerInstallFIle(context.Context) error
	// docker安装
	UbuntuDockerInstall(context.Context) error
	// 上传containerd压缩文件文件
	UploadContainerdCompressInstallFile(context.Context) error
	// 解压Containerd压缩安装文件
	UncompressContainerdInstallFIle(context.Context) error
	// containerd安装
	UbuntuContainerdInstall(context.Context) error
	// 安装kubeadm、kubectl、kubelet
	UbuntuKubeadmKubectlKuheletInstall(context.Context) error
	// 加入k8s集群
	JoinK8s(context.Context) error
}
