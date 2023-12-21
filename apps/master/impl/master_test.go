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

func TestDownloadK8sImage(t *testing.T) {
	err := svc.DownloadK8sImage(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInitialK8s(t *testing.T) {
	err := svc.InitialK8s(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetJoinK8sToken(t *testing.T) {
	tl, err := svc.GetJoinK8sToken(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tl)
}
