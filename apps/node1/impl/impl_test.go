package impl_test

import (
	"context"

	"github.com/solodba/k8s_install/apps/node1"
	"github.com/solodba/k8s_install/test/tools"
	"github.com/solodba/mcube/apps"
)

var (
	svc node1.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(node1.AppName).(node1.Service)
}
