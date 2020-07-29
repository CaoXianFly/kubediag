// +build !ignore_autogenerated

/*
Copyright 2020 The Kube Diagnoser Authors.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Abnormal) DeepCopyInto(out *Abnormal) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Abnormal.
func (in *Abnormal) DeepCopy() *Abnormal {
	if in == nil {
		return nil
	}
	out := new(Abnormal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Abnormal) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AbnormalCondition) DeepCopyInto(out *AbnormalCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AbnormalCondition.
func (in *AbnormalCondition) DeepCopy() *AbnormalCondition {
	if in == nil {
		return nil
	}
	out := new(AbnormalCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AbnormalList) DeepCopyInto(out *AbnormalList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Abnormal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AbnormalList.
func (in *AbnormalList) DeepCopy() *AbnormalList {
	if in == nil {
		return nil
	}
	out := new(AbnormalList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AbnormalList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AbnormalSpec) DeepCopyInto(out *AbnormalSpec) {
	*out = *in
	if in.KubernetesEvent != nil {
		in, out := &in.KubernetesEvent, &out.KubernetesEvent
		*out = new(corev1.Event)
		(*in).DeepCopyInto(*out)
	}
	if in.AssignedInformationCollectors != nil {
		in, out := &in.AssignedInformationCollectors, &out.AssignedInformationCollectors
		*out = make([]NamespacedName, len(*in))
		copy(*out, *in)
	}
	if in.AssignedDiagnosers != nil {
		in, out := &in.AssignedDiagnosers, &out.AssignedDiagnosers
		*out = make([]NamespacedName, len(*in))
		copy(*out, *in)
	}
	if in.AssignedRecoverers != nil {
		in, out := &in.AssignedRecoverers, &out.AssignedRecoverers
		*out = make([]NamespacedName, len(*in))
		copy(*out, *in)
	}
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AbnormalSpec.
func (in *AbnormalSpec) DeepCopy() *AbnormalSpec {
	if in == nil {
		return nil
	}
	out := new(AbnormalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AbnormalStatus) DeepCopyInto(out *AbnormalStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AbnormalCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.StartTime.DeepCopyInto(&out.StartTime)
	out.Diagnoser = in.Diagnoser
	out.Recoverer = in.Recoverer
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AbnormalStatus.
func (in *AbnormalStatus) DeepCopy() *AbnormalStatus {
	if in == nil {
		return nil
	}
	out := new(AbnormalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Diagnoser) DeepCopyInto(out *Diagnoser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Diagnoser.
func (in *Diagnoser) DeepCopy() *Diagnoser {
	if in == nil {
		return nil
	}
	out := new(Diagnoser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Diagnoser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiagnoserList) DeepCopyInto(out *DiagnoserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Diagnoser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiagnoserList.
func (in *DiagnoserList) DeepCopy() *DiagnoserList {
	if in == nil {
		return nil
	}
	out := new(DiagnoserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiagnoserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiagnoserSpec) DeepCopyInto(out *DiagnoserSpec) {
	*out = *in
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiagnoserSpec.
func (in *DiagnoserSpec) DeepCopy() *DiagnoserSpec {
	if in == nil {
		return nil
	}
	out := new(DiagnoserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiagnoserStatus) DeepCopyInto(out *DiagnoserStatus) {
	*out = *in
	in.LastDiagnosis.DeepCopyInto(&out.LastDiagnosis)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiagnoserStatus.
func (in *DiagnoserStatus) DeepCopy() *DiagnoserStatus {
	if in == nil {
		return nil
	}
	out := new(DiagnoserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Diagnosis) DeepCopyInto(out *Diagnosis) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	in.Abnormal.DeepCopyInto(&out.Abnormal)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Diagnosis.
func (in *Diagnosis) DeepCopy() *Diagnosis {
	if in == nil {
		return nil
	}
	out := new(Diagnosis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InformationCollector) DeepCopyInto(out *InformationCollector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InformationCollector.
func (in *InformationCollector) DeepCopy() *InformationCollector {
	if in == nil {
		return nil
	}
	out := new(InformationCollector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InformationCollector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InformationCollectorList) DeepCopyInto(out *InformationCollectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InformationCollector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InformationCollectorList.
func (in *InformationCollectorList) DeepCopy() *InformationCollectorList {
	if in == nil {
		return nil
	}
	out := new(InformationCollectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InformationCollectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InformationCollectorSpec) DeepCopyInto(out *InformationCollectorSpec) {
	*out = *in
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InformationCollectorSpec.
func (in *InformationCollectorSpec) DeepCopy() *InformationCollectorSpec {
	if in == nil {
		return nil
	}
	out := new(InformationCollectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InformationCollectorStatus) DeepCopyInto(out *InformationCollectorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InformationCollectorStatus.
func (in *InformationCollectorStatus) DeepCopy() *InformationCollectorStatus {
	if in == nil {
		return nil
	}
	out := new(InformationCollectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Label) DeepCopyInto(out *Label) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Label.
func (in *Label) DeepCopy() *Label {
	if in == nil {
		return nil
	}
	out := new(Label)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedName) DeepCopyInto(out *NamespacedName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedName.
func (in *NamespacedName) DeepCopy() *NamespacedName {
	if in == nil {
		return nil
	}
	out := new(NamespacedName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Recoverer) DeepCopyInto(out *Recoverer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Recoverer.
func (in *Recoverer) DeepCopy() *Recoverer {
	if in == nil {
		return nil
	}
	out := new(Recoverer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Recoverer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecovererList) DeepCopyInto(out *RecovererList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Recoverer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecovererList.
func (in *RecovererList) DeepCopy() *RecovererList {
	if in == nil {
		return nil
	}
	out := new(RecovererList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecovererList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecovererSpec) DeepCopyInto(out *RecovererSpec) {
	*out = *in
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecovererSpec.
func (in *RecovererSpec) DeepCopy() *RecovererSpec {
	if in == nil {
		return nil
	}
	out := new(RecovererSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecovererStatus) DeepCopyInto(out *RecovererStatus) {
	*out = *in
	in.LastRecovery.DeepCopyInto(&out.LastRecovery)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecovererStatus.
func (in *RecovererStatus) DeepCopy() *RecovererStatus {
	if in == nil {
		return nil
	}
	out := new(RecovererStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Recovery) DeepCopyInto(out *Recovery) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	in.Abnormal.DeepCopyInto(&out.Abnormal)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Recovery.
func (in *Recovery) DeepCopy() *Recovery {
	if in == nil {
		return nil
	}
	out := new(Recovery)
	in.DeepCopyInto(out)
	return out
}
