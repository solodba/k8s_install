package impl_test

import "testing"

func TestUploadDockerCompressInstallFile(t *testing.T) {
	err := svc.UploadDockerCompressInstallFile(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUncompressDockerInstallFIle(t *testing.T) {
	err := svc.UncompressDockerInstallFIle(ctx)
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

func TestUploadContainerdCompressInstallFile(t *testing.T) {
	err := svc.UploadContainerdCompressInstallFile(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUncompressContainerdInstallFIle(t *testing.T) {
	err := svc.UncompressContainerdInstallFIle(ctx)
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
