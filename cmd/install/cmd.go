package install

import (
	"context"

	"github.com/solodba/k8s_install/protocol"
	"github.com/spf13/cobra"
)

var (
	ctx = context.Background()
)

// k8s服务结构体
type Server struct {
	K8sInstallSvc *protocol.K8sInstallSvc
}

// k8s服务结构体初始化函数
func NewServer() *Server {
	return &Server{
		K8sInstallSvc: protocol.Newk8sInstallSvc(),
	}
}

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "install",
	Short:   "ms_install install service",
	Long:    "ms_install install service",
	Example: "ms_install install -f etc/config.toml",
	RunE: func(cmd *cobra.Command, args []string) error {
		svc := NewServer()
		err := svc.K8sInstallSvc.K8sInstall(ctx)
		if err != nil {
			return err
		}
		return nil
	},
}
