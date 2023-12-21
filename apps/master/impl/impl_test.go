package impl_test

import (
	"context"

	"github.com/solodba/k8s_install/apps/master"
	"github.com/solodba/k8s_install/test/tools"
	"github.com/solodba/mcube/apps"
)

var (
	svc master.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(master.AppName).(master.Service)
}
