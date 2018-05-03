package resource

import (
	"github.com/weaveworks/flux/resource"
)

type DaemonSet struct {
	BaseObject
	Spec DaemonSetSpec
}

type DaemonSetSpec struct {
	Template PodTemplate
}

func (ds DaemonSet) Containers() []resource.Container {
	return ds.Spec.Template.Containers()
}
