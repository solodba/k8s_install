package impl

import (
	"context"
	"fmt"
	"strings"

	"github.com/solodba/k8s_install/apps/master"
	"github.com/solodba/mcube/logger"
)

// 上传docker安装文件
func (i *impl) UploadDockerInstallFile(context.Context) error {
	_, err := i.c.Master.UploadFile(i.c.Docker.LimitsFile, i.c.Docker.LimitsFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Docker.LimitsFile)
	_, err = i.c.Master.UploadFile(i.c.Docker.SysctlFile, i.c.Docker.SysctlFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Docker.SysctlFile)
	_, err = i.c.Master.UploadFile(i.c.Docker.DockerBinaryFile, i.c.Docker.DockerBinaryFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Docker.DockerBinaryFile)
	_, err = i.c.Master.UploadFile(i.c.Docker.DockerDaemonFile, i.c.Docker.DockerDaemonFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Docker.DockerDaemonFile)
	_, err = i.c.Master.UploadFile(i.c.Docker.ContainerdServiceFile, i.c.Docker.ContainerdServiceFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Docker.ContainerdServiceFile)
	_, err = i.c.Master.UploadFile(i.c.Docker.DockerServiceFile, i.c.Docker.DockerServiceFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Docker.DockerServiceFile)
	_, err = i.c.Master.UploadFile(i.c.Docker.DockerSocketFile, i.c.Docker.DockerSocketFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Docker.DockerSocketFile)
	_, err = i.c.Master.UploadFile(i.c.Docker.DockerComposeFile, i.c.Docker.DockerComposeFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Docker.DockerComposeFile)
	return nil
}

// docker安装
func (i *impl) UbuntuDockerInstall(ctx context.Context) error {
	cmd := `grep "Ubuntu" /etc/issue | awk -F ' ' '{print $1}'`
	res, err := i.c.Master.RunShell(cmd)
	if err != nil {
		return err
	}
	if strings.TrimRight(res, "\n") == "Ubuntu" {
		logger.L().Info().Msgf("当前系统是Ubuntu,即将开始系统初始化、配置docker-compose与安装docker")
		cmd = fmt.Sprintf(`/bin/bash -c "cp %s /etc/security/limits.conf"`, i.c.Docker.LimitsFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("limits.conf文件配置完成!")
		cmd = fmt.Sprintf(`/bin/bash -c "cp %s /etc/sysctl.conf;modprobe ip_conntrack;modprobe br_netfilter;sysctl -p"`, i.c.Docker.SysctlFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("sysctl.conf文件配置完成!")
		cmd = fmt.Sprintf(`/bin/bash -c "tar -zxf %s -C %s"`, i.c.Docker.DockerBinaryFilePath(), i.c.Docker.Dir)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("docker二进制压缩文件解压完成!")
		cmd = fmt.Sprintf(`/bin/bash -c "cp %s/docker/* /usr/local/bin"`, i.c.Docker.Dir)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("docker二进制文件拷贝到/usr/local/bin目录完成!")
		cmd = fmt.Sprintf(`/bin/bash -c "mkdir /etc/docker;cp %s /etc/docker"`, i.c.Docker.DockerDaemonFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("daemon.json文件配置完成!")
		cmd = fmt.Sprintf(`/bin/bash -c "cp %s /lib/systemd/system/containerd.service"`, i.c.Docker.ContainerdServiceFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("containerd.service文件配置完成!")
		cmd = fmt.Sprintf(`/bin/bash -c "cp %s /lib/systemd/system/docker.service"`, i.c.Docker.DockerServiceFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("docker.service文件配置完成!")
		cmd = fmt.Sprintf(`/bin/bash -c "cp %s /lib/systemd/system/docker.socket"`, i.c.Docker.DockerSocketFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("docker.socket文件配置完成!")
		cmd = fmt.Sprintf(`/bin/bash -c "cp %s /usr/bin/docker-compose;chmod +x /usr/bin/docker-compose"`, i.c.Docker.DockerComposeFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("docker-compose文件配置完成!")
		cmd = fmt.Sprintf(`/bin/bash -c "groupadd %s;useradd %s -r -m -s /sbin/nologin -g %s"`, i.c.Docker.DockerUser, i.c.Docker.DockerUser, i.c.Docker.DockerUser)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("docker用户配置完成!")
		logger.L().Info().Msgf("正在启动docker server并设置为开机自启动!")
		cmd = fmt.Sprintf(`/bin/bash -c "systemctl enable docker.service;systemctl restart docker.service"`)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		cmd = fmt.Sprintf(`/bin/bash -c "systemctl enable docker.socket && systemctl restart docker.socket"`)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("docker server安装完成,欢迎进入docker世界!")
	} else {
		return fmt.Errorf("当前系统不是Ubuntu系统,无法完成安装!")
	}
	return nil
}

// 上传containerd安装文件
func (i *impl) UploadContainerdInstallFile(context.Context) error {
	_, err := i.c.Master.UploadFile(i.c.Docker.LimitsFile, i.c.Docker.LimitsFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Docker.LimitsFile)
	_, err = i.c.Master.UploadFile(i.c.Docker.SysctlFile, i.c.Docker.SysctlFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Docker.SysctlFile)
	_, err = i.c.Master.UploadFile(i.c.Docker.ContainerdServiceFile, i.c.Docker.ContainerdServiceFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Docker.ContainerdServiceFile)
	_, err = i.c.Master.UploadFile(i.c.Containerd.ContainerdFile, i.c.Containerd.ContainerdFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Containerd.ContainerdFile)
	_, err = i.c.Master.UploadFile(i.c.Containerd.ContainerdConfigFile, i.c.Containerd.ContainerdConfigFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Containerd.ContainerdConfigFile)
	_, err = i.c.Master.UploadFile(i.c.Containerd.NerdctlFile, i.c.Containerd.NerdctlFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Containerd.NerdctlFile)
	_, err = i.c.Master.UploadFile(i.c.Containerd.NerdctlConfigFile, i.c.Containerd.NerdctlConfigFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Containerd.NerdctlConfigFile)
	_, err = i.c.Master.UploadFile(i.c.Containerd.CniFile, i.c.Containerd.CniFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Containerd.CniFile)
	_, err = i.c.Master.UploadFile(i.c.Containerd.RuncFile, i.c.Containerd.RuncFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Containerd.RuncFile)
	_, err = i.c.Master.UploadFile(i.c.Containerd.IpvsModuleFile, i.c.Containerd.IpvsModuleFilePath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.Containerd.IpvsModuleFile)
	return nil
}

// containerd安装
func (i *impl) UbuntuContainerdInstall(context.Context) error {
	cmd := `grep "Ubuntu" /etc/issue | awk -F ' ' '{print $1}'`
	res, err := i.c.Master.RunShell(cmd)
	if err != nil {
		return err
	}
	if strings.TrimRight(res, "\n") == "Ubuntu" {
		logger.L().Info().Msgf("当前系统是Ubuntu,即将开始系统初始化、安装配置containerd")
		cmd = fmt.Sprintf(`/bin/bash -c "cp %s /etc/security/limits.conf"`, i.c.Docker.LimitsFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("limits.conf文件配置完成!")
		cmd = fmt.Sprintf(`/bin/bash -c "cp %s /etc/sysctl.conf;modprobe ip_conntrack;modprobe br_netfilter;sysctl -p"`, i.c.Docker.SysctlFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("sysctl.conf文件配置完成!")
		cmd = fmt.Sprintf(`/bin/bash -c "mkdir -p /etc/containerd /etc/nerdctl"`)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("/etc/containerd和/etc/nerdctl目录创建完成!")
		cmd = fmt.Sprintf(`/bin/bash -c "tar -zxf %s -C %s"`, i.c.Containerd.ContainerdFilePath(), i.c.Containerd.Dir)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("%s文件解压完成!", i.c.Containerd.ContainerdFile)
		cmd = fmt.Sprintf(`/bin/bash -c "cp %s/bin/* /usr/local/bin/"`, i.c.Containerd.Dir)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("Containerd二进制文件拷贝完成!")
		cmd = fmt.Sprintf(`/bin/bash -c "cp %s /usr/bin/runc;chmod a+x /usr/bin/runc"`, i.c.Containerd.RuncFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("%s文件拷贝完成!", i.c.Containerd.RuncFile)
		cmd = fmt.Sprintf(`/bin/bash -c "cp %s /etc/containerd/config.toml"`, i.c.Containerd.ContainerdConfigFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("%s文件拷贝完成!", i.c.Containerd.ContainerdConfigFile)
		cmd = fmt.Sprintf(`/bin/bash -c "cp %s /lib/systemd/system/containerd.service"`, i.c.Docker.ContainerdServiceFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("%s文件拷贝完成!", i.c.Docker.ContainerdServiceFile)
		cmd = fmt.Sprintf(`/bin/bash -c "mkdir /opt/cni/bin -p"`)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("/opt/cni/bin目录创建完成!")
		cmd = fmt.Sprintf(`/bin/bash -c "tar -zxf %s -C /opt/cni/bin"`, i.c.Containerd.CniFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("%s文件解压完成!", i.c.Containerd.CniFile)
		cmd = fmt.Sprintf(`/bin/bash -c "tar -zxf %s -C /usr/local/bin"`, i.c.Containerd.NerdctlFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("%s文件解压完成!", i.c.Containerd.NerdctlFile)
		cmd = fmt.Sprintf(`/bin/bash -c "cp %s /etc/nerdctl/nerdctl.toml"`, i.c.Containerd.NerdctlConfigFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("%s文件拷贝完成!", i.c.Containerd.NerdctlConfigFile)
		cmd = fmt.Sprintf(`/bin/bash -c "cp %s /etc/modules-load.d/modules.conf"`, i.c.Containerd.IpvsModuleFilePath())
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("%s文件拷贝完成!", i.c.Containerd.IpvsModuleFile)
		logger.L().Info().Msg("正在启动containerd server并设置为开机自启动!")
		cmd = fmt.Sprintf(`/bin/bash -c "systemctl daemon-reload;systemctl restart containerd"`)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msgf("containerd服务启动完成!")
		cmd = fmt.Sprintf(`/bin/bash -c "systemctl is-active containerd"`)
		res, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		if strings.TrimRight(res, "\n") == "active" {
			logger.L().Info().Msgf("containerd服务状态正常!")
		} else {
			return fmt.Errorf("containerd服务状态异常!")
		}
		logger.L().Info().Msg("containerd server安装完成,欢迎进入containerd的容器世界!")
	} else {
		return fmt.Errorf("当前系统不是Ubuntu系统,无法完成安装!")
	}
	return nil
}

// 安装kubeadm、kubectl、kubelet
func (i *impl) UbuntuKubeadmKubectlKuheletInstall(context.Context) error {
	cmd := `grep "Ubuntu" /etc/issue | awk -F ' ' '{print $1}'`
	res, err := i.c.Master.RunShell(cmd)
	if err != nil {
		return err
	}
	if strings.TrimRight(res, "\n") == "Ubuntu" {
		cmd = fmt.Sprintf(`/bin/bash -c "apt-get update && apt-get install -y apt-transport-https -y"`)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		cmd = fmt.Sprintf(`/bin/bash -c "curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -"`)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		cmd = fmt.Sprintf(`/bin/bash -c "echo 'deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main' > /etc/apt/sources.list.d/kubernetes.list"`)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		cmd = fmt.Sprintf(`/bin/bash -c "apt-get update;apt install -y kubeadm=%s kubelet=%s kubectl=%s"`, i.c.K8s.KubeadmVersion, i.c.K8s.KubeletVersion, i.c.K8s.KubectlVersion)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return err
		}
		logger.L().Info().Msg("Ubuntu系统安装kubeadm、kubectl、kubelet完成!")
	} else {
		return fmt.Errorf("当前系统不是Ubuntu系统,无法完成安装!")
	}
	return nil
}

// 下载安装k8s镜像
func (i *impl) DownloadK8sImage(context.Context) error {
	_, err := i.c.Master.UploadFile(i.c.K8s.DownloadImageScript, i.c.K8s.DownloadImageScriptPath())
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("%s文件上传成功!", i.c.K8s.DownloadImageScript)
	cmd := fmt.Sprintf(`/bin/bash -c "chmod +x %s;bash %s"`, i.c.K8s.DownloadImageScriptPath(), i.c.K8s.DownloadImageScriptPath())
	_, err = i.c.Master.RunShell(cmd)
	if err != nil {
		return err
	}
	logger.L().Info().Msg("安装k8s所需镜像下载完成!")
	return nil
}

// 初始化k8s集群
func (i *impl) InitialK8s(context.Context) error {
	cmd := fmt.Sprintf(`/bin/bash -c "kubeadm init --apiserver-advertise-address=%s --apiserver-bind-port=%d --kubernetes-version=%s --pod-network-cidr=%s --service-cidr=%s --service-dns-domain=%s --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --ignore-preflight-errors=swap"`,
		i.c.K8s.ApiserverAddress,
		i.c.K8s.ApiserverPort,
		i.c.K8s.Version,
		i.c.K8s.PodNetwork,
		i.c.K8s.ServiceNetwork,
		i.c.K8s.ServiceDnsDomain)
	_, err := i.c.Master.RunShell(cmd)
	if err != nil {
		return err
	}
	cmd = fmt.Sprintf(`/bin/bash -c "mkdir -p /root/.kube;cp -i /etc/kubernetes/admin.conf /root/.kube/config;chown $(id -u):$(id -g) /root/.kube/config"`)
	_, err = i.c.Master.RunShell(cmd)
	if err != nil {
		return err
	}
	logger.L().Info().Msg("主节点初始化k8s集群完成!")
	return nil
}

// 获取token
func (i *impl) GetJoinK8sToken(context.Context) (*master.TokenList, error) {
	tl := master.NewTokenList()
	cmd := fmt.Sprintf(`/bin/bash -c "kubeadm token list | grep -v 'TOKEN'"`)
	res, err := i.c.Master.RunShell(cmd)
	if err != nil {
		return nil, err
	}
	resList := strings.Split(res, " ")
	token := resList[0]
	tl.Token = token
	cmd = fmt.Sprintf(`/bin/bash -c "cat /etc/kubernetes/pki/ca.crt | openssl x509 -pubkey -noout | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'"`)
	res, err = i.c.Master.RunShell(cmd)
	if err != nil {
		return nil, err
	}
	tl.TokenHash = "sha256:" + strings.TrimRight(res, "\n")
	return tl, nil
}
