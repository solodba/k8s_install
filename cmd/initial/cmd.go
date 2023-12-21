package initial

import "github.com/spf13/cobra"

// 项目初始化配置子命令
var Cmd = &cobra.Command{
	Use:     "init",
	Short:   "k8s_install init service",
	Long:    "k8s_install init service",
	Example: "k8s_install init -f xxx",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
