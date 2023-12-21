package impl

import (
	"github.com/solodba/k8s_install/apps/master"
	"github.com/solodba/k8s_install/conf"
	"github.com/solodba/mcube/apps"
)

var (
	svc = &impl{}
)

// 业务实现类
type impl struct {
	c *conf.Config
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return master.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	i.c = conf.C()
	return nil
}

// 注册实例类
func init() {
	apps.RegistryInternalApp(svc)
}
