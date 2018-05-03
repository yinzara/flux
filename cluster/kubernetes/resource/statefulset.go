package resource

import (
	"github.com/weaveworks/flux/resource"
)

type StatefulSet struct {
	BaseObject
	Spec StatefulSetSpec
}

type StatefulSetSpec struct {
	Replicas int
	Template PodTemplate
}

func (ss StatefulSet) Containers() []resource.Container {
	return ss.Spec.Template.Containers()
}
