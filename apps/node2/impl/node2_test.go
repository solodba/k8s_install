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

func TestJoinK8s(t *testing.T) {
	err := svc.JoinK8s(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
