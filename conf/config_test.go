package conf_test

import (
	"testing"

	"github.com/solodba/k8s_install/conf"
)

func TestLoadConfigFromToml(t *testing.T) {
	err := conf.LoadConfigFromToml("test/test.toml")
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(conf.C().Docker)
	// t.Log(conf.C().Master)
	// t.Log(conf.C().Node1)
	// t.Log(conf.C().Node2)
	// t.Log(conf.C().Docker.DockerServiceFilePath())
	// t.Log(conf.C().Docker.DockerSocketFilePath())
	// t.Log(conf.C().Docker.DockerComposeFilePath())
	// t.Log(conf.C().Docker.DockerDaemonFilePath())
	// t.Log(conf.C().Docker.DockerBinaryFilePath())
	// t.Log(conf.C().Docker.ContainerdServiceFilePath())
	// t.Log(conf.C().Containerd)
	// t.Log(conf.C().K8s.DownloadImageScriptPath())
	// t.Log(conf.C().Containerd.ContainerdFilePath())
	// t.Log(conf.C().Containerd.ContainerdCompressInstallFilePath())
	t.Log(conf.C().Docker.CompressInstallFile)
	t.Log(conf.C().Docker.DockerCompressInstallFilePath())
}

func TestLoadConfigFromEnv(t *testing.T) {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(conf.C().Docker)
	// t.Log(conf.C().Master)
	// t.Log(conf.C().Node1)
	// t.Log(conf.C().Node2)
	// t.Log(conf.C().Docker.DockerServiceFilePath())
	// t.Log(conf.C().Docker.DockerSocketFilePath())
	// t.Log(conf.C().Docker.DockerComposeFilePath())
	// t.Log(conf.C().Docker.DockerDaemonFilePath())
	// t.Log(conf.C().Docker.DockerBinaryFilePath())
	// t.Log(conf.C().Docker.ContainerdServiceFilePath())
	// t.Log(conf.C().Containerd)
	// t.Log(conf.C().K8s.DownloadImageScriptPath())
	// t.Log(conf.C().Containerd.ContainerdFilePath())
	// t.Log(conf.C().Containerd.CompressInstallFile)
	// t.Log(conf.C().Containerd.ContainerdCompressInstallFilePath())
	t.Log(conf.C().Docker.CompressInstallFile)
	t.Log(conf.C().Docker.DockerCompressInstallFilePath())
}

func TestUploadFile(t *testing.T) {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	srcFile := "mysql-8.0.25-linux-glibc2.12-x86_64.tar.xz"
	dstFile := "/root/mysql-8.0.25-linux-glibc2.12-x86_64.tar.xz"
	result, err := conf.C().Master.UploadFile(srcFile, dstFile)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestRunShell(t *testing.T) {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	cmd := `/bin/bash -c "ls -l"`
	result, err := conf.C().Master.RunShell(cmd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
