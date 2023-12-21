package impl_test

import "testing"

func TestUploadDockerInstallFile(t *testing.T) {
	err := svc.UploadDockerInstallFile(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUbuntuDockerInstall(t *testing.T) {
	err := svc.UbuntuDockerInstall(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUploadContainerdInstallFile(t *testing.T) {
	err := svc.UploadContainerdInstallFile(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUbuntuContainerdInstall(t *testing.T) {
	err := svc.UbuntuContainerdInstall(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUbuntuKubeadmKubectlKuheletInstall(t *testing.T) {
	err := svc.UbuntuKubeadmKubectlKuheletInstall(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestJoinK8s(t *testing.T) {
	err := svc.JoinK8s(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
