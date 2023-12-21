package impl_test

import (
	"context"

	"github.com/solodba/k8s_install/apps/node2"
	"github.com/solodba/k8s_install/test/tools"
	"github.com/solodba/mcube/apps"
)

var (
	svc node2.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(node2.AppName).(node2.Service)
}
