/*
/*
@Time : 2019/12/26 2:07 下午
@Author : tianpeng.du
@File : register.go
@Software: GoLand
*/
package pkg

import (
	"github.com/du2016/scheduler-framework-test/pkg/test"
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func Register() *cobra.Command {
	return app.NewSchedulerCommand(
		app.WithPlugin(test.Name, test.New))
}
