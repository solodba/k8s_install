package impl

import (
	"github.com/solodba/k8s_install/apps/master"
	"github.com/solodba/k8s_install/apps/node1"
	"github.com/solodba/k8s_install/conf"
	"github.com/solodba/mcube/apps"
)

var (
	svc = &impl{}
)

// 业务实现类
type impl struct {
	c         *conf.Config
	masterSvc master.Service
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return node1.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	i.c = conf.C()
	i.masterSvc = apps.GetInternalApp(master.AppName).(master.Service)
	return nil
}

// 注册实例类
func init() {
	apps.RegistryInternalApp(svc)
}
