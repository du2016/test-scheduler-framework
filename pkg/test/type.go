/*
/*
@Time : 2019/12/25 11:03 上午
@Author : tianpeng.du
@File : type.go
@Software: GoLand
*/
package test

import (
	"context"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	schedulernodeinfo "k8s.io/kubernetes/pkg/scheduler/nodeinfo"
	"strconv"
	"time"
)

const Name = "test"

var _ framework.PreFilterPlugin = &TestPlugin{}

type Args struct {
	KubeConfig string `json:"kubeconfig,omitempty"`
	Master     string `json:"master,omitempty"`
}

type TestPlugin struct {
	handle framework.FrameworkHandle
}

func New(rargs *runtime.Unknown, handle framework.FrameworkHandle) (framework.Plugin, error) {
	args:=Args{}
	if err := framework.DecodeInto(rargs, args); err != nil {
		return nil, err
	}
	klog.Info(args)
	return &TestPlugin{
		handle: handle,
	}, nil
}

func (self *TestPlugin) Name() string {
	return Name
}

func (self *TestPlugin) Less(p1 *v1.Pod, p2 *v1.Pod) bool {
	if p1.CreationTimestamp.Unix() > p2.CreationTimestamp.Unix() {
		return true
	}
	return false
}

func (self *TestPlugin) PreFilter(ctx context.Context, state *framework.CycleState, p *v1.Pod) *framework.Status {
	klog.Error("into controller test")
	state.Write()
	var dtime int64
	var err error
	if v,ok:=p.Annotations["delay"];ok {
		if dtime,err=strconv.ParseInt(v,10,64);err!=nil {
			return nil
		}
		if time.Now().Unix()-p.CreationTimestamp.Unix()>=dtime {
			klog.Infof("scheduler: %s/%s",p.Namespace,p.Name)
			return nil
		}
		klog.Infof("not reatch scheduler time: %s/%s",p.Namespace,p.Name)
		return framework.NewStatus(framework.Skip,"not reatch scheduler time")
	}
	return nil
}

func (self *TestPlugin) AddPod(ctx context.Context, state *framework.CycleState, podToSchedule *v1.Pod, podToAdd *v1.Pod, nodeInfo *schedulernodeinfo.NodeInfo) *framework.Status {
	return nil
}

func (self *TestPlugin) RemovePod(ctx context.Context, state *framework.CycleState, podToSchedule *v1.Pod, podToRemove *v1.Pod, nodeInfo *schedulernodeinfo.NodeInfo) *framework.Status {
	return nil
}

func (self *TestPlugin) PreFilterExtensions() framework.PreFilterExtensions {
	return self
}
