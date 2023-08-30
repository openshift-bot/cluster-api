//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package builder

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/cluster-api/api/v1beta1"
	apiv1beta1 "sigs.k8s.io/cluster-api/exp/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapTemplateBuilder) DeepCopyInto(out *BootstrapTemplateBuilder) {
	*out = *in
	if in.obj != nil {
		in, out := &in.obj, &out.obj
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapTemplateBuilder.
func (in *BootstrapTemplateBuilder) DeepCopy() *BootstrapTemplateBuilder {
	if in == nil {
		return nil
	}
	out := new(BootstrapTemplateBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBuilder) DeepCopyInto(out *ClusterBuilder) {
	*out = *in
	if in.labels != nil {
		in, out := &in.labels, &out.labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.annotations != nil {
		in, out := &in.annotations, &out.annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.topology != nil {
		in, out := &in.topology, &out.topology
		*out = new(v1beta1.Topology)
		(*in).DeepCopyInto(*out)
	}
	if in.infrastructureCluster != nil {
		in, out := &in.infrastructureCluster, &out.infrastructureCluster
		*out = (*in).DeepCopy()
	}
	if in.controlPlane != nil {
		in, out := &in.controlPlane, &out.controlPlane
		*out = (*in).DeepCopy()
	}
	if in.network != nil {
		in, out := &in.network, &out.network
		*out = new(v1beta1.ClusterNetwork)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBuilder.
func (in *ClusterBuilder) DeepCopy() *ClusterBuilder {
	if in == nil {
		return nil
	}
	out := new(ClusterBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterClassBuilder) DeepCopyInto(out *ClusterClassBuilder) {
	*out = *in
	if in.infrastructureClusterTemplate != nil {
		in, out := &in.infrastructureClusterTemplate, &out.infrastructureClusterTemplate
		*out = (*in).DeepCopy()
	}
	if in.controlPlaneMetadata != nil {
		in, out := &in.controlPlaneMetadata, &out.controlPlaneMetadata
		*out = new(v1beta1.ObjectMeta)
		(*in).DeepCopyInto(*out)
	}
	if in.controlPlaneTemplate != nil {
		in, out := &in.controlPlaneTemplate, &out.controlPlaneTemplate
		*out = (*in).DeepCopy()
	}
	if in.controlPlaneInfrastructureMachineTemplate != nil {
		in, out := &in.controlPlaneInfrastructureMachineTemplate, &out.controlPlaneInfrastructureMachineTemplate
		*out = (*in).DeepCopy()
	}
	if in.controlPlaneMHC != nil {
		in, out := &in.controlPlaneMHC, &out.controlPlaneMHC
		*out = new(v1beta1.MachineHealthCheckClass)
		(*in).DeepCopyInto(*out)
	}
	if in.controlPlaneNodeDrainTimeout != nil {
		in, out := &in.controlPlaneNodeDrainTimeout, &out.controlPlaneNodeDrainTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.controlPlaneNodeVolumeDetachTimeout != nil {
		in, out := &in.controlPlaneNodeVolumeDetachTimeout, &out.controlPlaneNodeVolumeDetachTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.controlPlaneNodeDeletionTimeout != nil {
		in, out := &in.controlPlaneNodeDeletionTimeout, &out.controlPlaneNodeDeletionTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.controlPlaneNamingStrategy != nil {
		in, out := &in.controlPlaneNamingStrategy, &out.controlPlaneNamingStrategy
		*out = new(v1beta1.ControlPlaneClassNamingStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.machineDeploymentClasses != nil {
		in, out := &in.machineDeploymentClasses, &out.machineDeploymentClasses
		*out = make([]v1beta1.MachineDeploymentClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.variables != nil {
		in, out := &in.variables, &out.variables
		*out = make([]v1beta1.ClusterClassVariable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.statusVariables != nil {
		in, out := &in.statusVariables, &out.statusVariables
		*out = make([]v1beta1.ClusterClassStatusVariable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.patches != nil {
		in, out := &in.patches, &out.patches
		*out = make([]v1beta1.ClusterClassPatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterClassBuilder.
func (in *ClusterClassBuilder) DeepCopy() *ClusterClassBuilder {
	if in == nil {
		return nil
	}
	out := new(ClusterClassBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTopologyBuilder) DeepCopyInto(out *ClusterTopologyBuilder) {
	*out = *in
	if in.workers != nil {
		in, out := &in.workers, &out.workers
		*out = new(v1beta1.WorkersTopology)
		(*in).DeepCopyInto(*out)
	}
	if in.controlPlaneMHC != nil {
		in, out := &in.controlPlaneMHC, &out.controlPlaneMHC
		*out = new(v1beta1.MachineHealthCheckTopology)
		(*in).DeepCopyInto(*out)
	}
	if in.variables != nil {
		in, out := &in.variables, &out.variables
		*out = make([]v1beta1.ClusterVariable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTopologyBuilder.
func (in *ClusterTopologyBuilder) DeepCopy() *ClusterTopologyBuilder {
	if in == nil {
		return nil
	}
	out := new(ClusterTopologyBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneBuilder) DeepCopyInto(out *ControlPlaneBuilder) {
	*out = *in
	if in.obj != nil {
		in, out := &in.obj, &out.obj
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneBuilder.
func (in *ControlPlaneBuilder) DeepCopy() *ControlPlaneBuilder {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneTemplateBuilder) DeepCopyInto(out *ControlPlaneTemplateBuilder) {
	*out = *in
	if in.obj != nil {
		in, out := &in.obj, &out.obj
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneTemplateBuilder.
func (in *ControlPlaneTemplateBuilder) DeepCopy() *ControlPlaneTemplateBuilder {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneTemplateBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureClusterBuilder) DeepCopyInto(out *InfrastructureClusterBuilder) {
	*out = *in
	if in.obj != nil {
		in, out := &in.obj, &out.obj
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureClusterBuilder.
func (in *InfrastructureClusterBuilder) DeepCopy() *InfrastructureClusterBuilder {
	if in == nil {
		return nil
	}
	out := new(InfrastructureClusterBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureClusterTemplateBuilder) DeepCopyInto(out *InfrastructureClusterTemplateBuilder) {
	*out = *in
	if in.obj != nil {
		in, out := &in.obj, &out.obj
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureClusterTemplateBuilder.
func (in *InfrastructureClusterTemplateBuilder) DeepCopy() *InfrastructureClusterTemplateBuilder {
	if in == nil {
		return nil
	}
	out := new(InfrastructureClusterTemplateBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureMachineTemplateBuilder) DeepCopyInto(out *InfrastructureMachineTemplateBuilder) {
	*out = *in
	if in.obj != nil {
		in, out := &in.obj, &out.obj
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureMachineTemplateBuilder.
func (in *InfrastructureMachineTemplateBuilder) DeepCopy() *InfrastructureMachineTemplateBuilder {
	if in == nil {
		return nil
	}
	out := new(InfrastructureMachineTemplateBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineBuilder) DeepCopyInto(out *MachineBuilder) {
	*out = *in
	if in.version != nil {
		in, out := &in.version, &out.version
		*out = new(string)
		**out = **in
	}
	if in.bootstrap != nil {
		in, out := &in.bootstrap, &out.bootstrap
		*out = (*in).DeepCopy()
	}
	if in.labels != nil {
		in, out := &in.labels, &out.labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineBuilder.
func (in *MachineBuilder) DeepCopy() *MachineBuilder {
	if in == nil {
		return nil
	}
	out := new(MachineBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeploymentBuilder) DeepCopyInto(out *MachineDeploymentBuilder) {
	*out = *in
	if in.bootstrapTemplate != nil {
		in, out := &in.bootstrapTemplate, &out.bootstrapTemplate
		*out = (*in).DeepCopy()
	}
	if in.infrastructureTemplate != nil {
		in, out := &in.infrastructureTemplate, &out.infrastructureTemplate
		*out = (*in).DeepCopy()
	}
	if in.selector != nil {
		in, out := &in.selector, &out.selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.version != nil {
		in, out := &in.version, &out.version
		*out = new(string)
		**out = **in
	}
	if in.replicas != nil {
		in, out := &in.replicas, &out.replicas
		*out = new(int32)
		**out = **in
	}
	if in.generation != nil {
		in, out := &in.generation, &out.generation
		*out = new(int64)
		**out = **in
	}
	if in.labels != nil {
		in, out := &in.labels, &out.labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.status != nil {
		in, out := &in.status, &out.status
		*out = new(v1beta1.MachineDeploymentStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeploymentBuilder.
func (in *MachineDeploymentBuilder) DeepCopy() *MachineDeploymentBuilder {
	if in == nil {
		return nil
	}
	out := new(MachineDeploymentBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeploymentClassBuilder) DeepCopyInto(out *MachineDeploymentClassBuilder) {
	*out = *in
	if in.infrastructureMachineTemplate != nil {
		in, out := &in.infrastructureMachineTemplate, &out.infrastructureMachineTemplate
		*out = (*in).DeepCopy()
	}
	if in.bootstrapTemplate != nil {
		in, out := &in.bootstrapTemplate, &out.bootstrapTemplate
		*out = (*in).DeepCopy()
	}
	if in.labels != nil {
		in, out := &in.labels, &out.labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.annotations != nil {
		in, out := &in.annotations, &out.annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.machineHealthCheckClass != nil {
		in, out := &in.machineHealthCheckClass, &out.machineHealthCheckClass
		*out = new(v1beta1.MachineHealthCheckClass)
		(*in).DeepCopyInto(*out)
	}
	if in.failureDomain != nil {
		in, out := &in.failureDomain, &out.failureDomain
		*out = new(string)
		**out = **in
	}
	if in.nodeDrainTimeout != nil {
		in, out := &in.nodeDrainTimeout, &out.nodeDrainTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.nodeVolumeDetachTimeout != nil {
		in, out := &in.nodeVolumeDetachTimeout, &out.nodeVolumeDetachTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.nodeDeletionTimeout != nil {
		in, out := &in.nodeDeletionTimeout, &out.nodeDeletionTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.minReadySeconds != nil {
		in, out := &in.minReadySeconds, &out.minReadySeconds
		*out = new(int32)
		**out = **in
	}
	if in.strategy != nil {
		in, out := &in.strategy, &out.strategy
		*out = new(v1beta1.MachineDeploymentStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.namingStrategy != nil {
		in, out := &in.namingStrategy, &out.namingStrategy
		*out = new(v1beta1.MachineDeploymentClassNamingStrategy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeploymentClassBuilder.
func (in *MachineDeploymentClassBuilder) DeepCopy() *MachineDeploymentClassBuilder {
	if in == nil {
		return nil
	}
	out := new(MachineDeploymentClassBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDeploymentTopologyBuilder) DeepCopyInto(out *MachineDeploymentTopologyBuilder) {
	*out = *in
	if in.replicas != nil {
		in, out := &in.replicas, &out.replicas
		*out = new(int32)
		**out = **in
	}
	if in.mhc != nil {
		in, out := &in.mhc, &out.mhc
		*out = new(v1beta1.MachineHealthCheckTopology)
		(*in).DeepCopyInto(*out)
	}
	if in.variables != nil {
		in, out := &in.variables, &out.variables
		*out = make([]v1beta1.ClusterVariable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDeploymentTopologyBuilder.
func (in *MachineDeploymentTopologyBuilder) DeepCopy() *MachineDeploymentTopologyBuilder {
	if in == nil {
		return nil
	}
	out := new(MachineDeploymentTopologyBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineHealthCheckBuilder) DeepCopyInto(out *MachineHealthCheckBuilder) {
	*out = *in
	if in.ownerRefs != nil {
		in, out := &in.ownerRefs, &out.ownerRefs
		*out = make([]v1.OwnerReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.selector.DeepCopyInto(&out.selector)
	if in.conditions != nil {
		in, out := &in.conditions, &out.conditions
		*out = make([]v1beta1.UnhealthyCondition, len(*in))
		copy(*out, *in)
	}
	if in.maxUnhealthy != nil {
		in, out := &in.maxUnhealthy, &out.maxUnhealthy
		*out = new(intstr.IntOrString)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineHealthCheckBuilder.
func (in *MachineHealthCheckBuilder) DeepCopy() *MachineHealthCheckBuilder {
	if in == nil {
		return nil
	}
	out := new(MachineHealthCheckBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachinePoolBuilder) DeepCopyInto(out *MachinePoolBuilder) {
	*out = *in
	if in.bootstrapTemplate != nil {
		in, out := &in.bootstrapTemplate, &out.bootstrapTemplate
		*out = (*in).DeepCopy()
	}
	if in.infrastructureTemplate != nil {
		in, out := &in.infrastructureTemplate, &out.infrastructureTemplate
		*out = (*in).DeepCopy()
	}
	if in.version != nil {
		in, out := &in.version, &out.version
		*out = new(string)
		**out = **in
	}
	if in.replicas != nil {
		in, out := &in.replicas, &out.replicas
		*out = new(int32)
		**out = **in
	}
	if in.labels != nil {
		in, out := &in.labels, &out.labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.status != nil {
		in, out := &in.status, &out.status
		*out = new(apiv1beta1.MachinePoolStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachinePoolBuilder.
func (in *MachinePoolBuilder) DeepCopy() *MachinePoolBuilder {
	if in == nil {
		return nil
	}
	out := new(MachinePoolBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSetBuilder) DeepCopyInto(out *MachineSetBuilder) {
	*out = *in
	if in.bootstrapTemplate != nil {
		in, out := &in.bootstrapTemplate, &out.bootstrapTemplate
		*out = (*in).DeepCopy()
	}
	if in.infrastructureTemplate != nil {
		in, out := &in.infrastructureTemplate, &out.infrastructureTemplate
		*out = (*in).DeepCopy()
	}
	if in.replicas != nil {
		in, out := &in.replicas, &out.replicas
		*out = new(int32)
		**out = **in
	}
	if in.labels != nil {
		in, out := &in.labels, &out.labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ownerRefs != nil {
		in, out := &in.ownerRefs, &out.ownerRefs
		*out = make([]v1.OwnerReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSetBuilder.
func (in *MachineSetBuilder) DeepCopy() *MachineSetBuilder {
	if in == nil {
		return nil
	}
	out := new(MachineSetBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestBootstrapTemplateBuilder) DeepCopyInto(out *TestBootstrapTemplateBuilder) {
	*out = *in
	if in.obj != nil {
		in, out := &in.obj, &out.obj
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestBootstrapTemplateBuilder.
func (in *TestBootstrapTemplateBuilder) DeepCopy() *TestBootstrapTemplateBuilder {
	if in == nil {
		return nil
	}
	out := new(TestBootstrapTemplateBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestControlPlaneBuilder) DeepCopyInto(out *TestControlPlaneBuilder) {
	*out = *in
	if in.obj != nil {
		in, out := &in.obj, &out.obj
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestControlPlaneBuilder.
func (in *TestControlPlaneBuilder) DeepCopy() *TestControlPlaneBuilder {
	if in == nil {
		return nil
	}
	out := new(TestControlPlaneBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestControlPlaneTemplateBuilder) DeepCopyInto(out *TestControlPlaneTemplateBuilder) {
	*out = *in
	if in.obj != nil {
		in, out := &in.obj, &out.obj
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestControlPlaneTemplateBuilder.
func (in *TestControlPlaneTemplateBuilder) DeepCopy() *TestControlPlaneTemplateBuilder {
	if in == nil {
		return nil
	}
	out := new(TestControlPlaneTemplateBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestInfrastructureClusterBuilder) DeepCopyInto(out *TestInfrastructureClusterBuilder) {
	*out = *in
	if in.obj != nil {
		in, out := &in.obj, &out.obj
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestInfrastructureClusterBuilder.
func (in *TestInfrastructureClusterBuilder) DeepCopy() *TestInfrastructureClusterBuilder {
	if in == nil {
		return nil
	}
	out := new(TestInfrastructureClusterBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestInfrastructureClusterTemplateBuilder) DeepCopyInto(out *TestInfrastructureClusterTemplateBuilder) {
	*out = *in
	if in.obj != nil {
		in, out := &in.obj, &out.obj
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestInfrastructureClusterTemplateBuilder.
func (in *TestInfrastructureClusterTemplateBuilder) DeepCopy() *TestInfrastructureClusterTemplateBuilder {
	if in == nil {
		return nil
	}
	out := new(TestInfrastructureClusterTemplateBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestInfrastructureMachineTemplateBuilder) DeepCopyInto(out *TestInfrastructureMachineTemplateBuilder) {
	*out = *in
	if in.obj != nil {
		in, out := &in.obj, &out.obj
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestInfrastructureMachineTemplateBuilder.
func (in *TestInfrastructureMachineTemplateBuilder) DeepCopy() *TestInfrastructureMachineTemplateBuilder {
	if in == nil {
		return nil
	}
	out := new(TestInfrastructureMachineTemplateBuilder)
	in.DeepCopyInto(out)
	return out
}
