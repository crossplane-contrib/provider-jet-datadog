//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Downtime) DeepCopyInto(out *Downtime) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Downtime.
func (in *Downtime) DeepCopy() *Downtime {
	if in == nil {
		return nil
	}
	out := new(Downtime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Downtime) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DowntimeList) DeepCopyInto(out *DowntimeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Downtime, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DowntimeList.
func (in *DowntimeList) DeepCopy() *DowntimeList {
	if in == nil {
		return nil
	}
	out := new(DowntimeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DowntimeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DowntimeObservation) DeepCopyInto(out *DowntimeObservation) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.ActiveChildID != nil {
		in, out := &in.ActiveChildID, &out.ActiveChildID
		*out = new(float64)
		**out = **in
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DowntimeObservation.
func (in *DowntimeObservation) DeepCopy() *DowntimeObservation {
	if in == nil {
		return nil
	}
	out := new(DowntimeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DowntimeParameters) DeepCopyInto(out *DowntimeParameters) {
	*out = *in
	if in.End != nil {
		in, out := &in.End, &out.End
		*out = new(float64)
		**out = **in
	}
	if in.EndDate != nil {
		in, out := &in.EndDate, &out.EndDate
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.MonitorID != nil {
		in, out := &in.MonitorID, &out.MonitorID
		*out = new(float64)
		**out = **in
	}
	if in.MonitorTags != nil {
		in, out := &in.MonitorTags, &out.MonitorTags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MuteFirstRecoveryNotification != nil {
		in, out := &in.MuteFirstRecoveryNotification, &out.MuteFirstRecoveryNotification
		*out = new(bool)
		**out = **in
	}
	if in.Recurrence != nil {
		in, out := &in.Recurrence, &out.Recurrence
		*out = make([]RecurrenceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = new(float64)
		**out = **in
	}
	if in.StartDate != nil {
		in, out := &in.StartDate, &out.StartDate
		*out = new(string)
		**out = **in
	}
	if in.Timezone != nil {
		in, out := &in.Timezone, &out.Timezone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DowntimeParameters.
func (in *DowntimeParameters) DeepCopy() *DowntimeParameters {
	if in == nil {
		return nil
	}
	out := new(DowntimeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DowntimeSpec) DeepCopyInto(out *DowntimeSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DowntimeSpec.
func (in *DowntimeSpec) DeepCopy() *DowntimeSpec {
	if in == nil {
		return nil
	}
	out := new(DowntimeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DowntimeStatus) DeepCopyInto(out *DowntimeStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DowntimeStatus.
func (in *DowntimeStatus) DeepCopy() *DowntimeStatus {
	if in == nil {
		return nil
	}
	out := new(DowntimeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecurrenceObservation) DeepCopyInto(out *RecurrenceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecurrenceObservation.
func (in *RecurrenceObservation) DeepCopy() *RecurrenceObservation {
	if in == nil {
		return nil
	}
	out := new(RecurrenceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecurrenceParameters) DeepCopyInto(out *RecurrenceParameters) {
	*out = *in
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(float64)
		**out = **in
	}
	if in.Rrule != nil {
		in, out := &in.Rrule, &out.Rrule
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UntilDate != nil {
		in, out := &in.UntilDate, &out.UntilDate
		*out = new(float64)
		**out = **in
	}
	if in.UntilOccurrences != nil {
		in, out := &in.UntilOccurrences, &out.UntilOccurrences
		*out = new(float64)
		**out = **in
	}
	if in.WeekDays != nil {
		in, out := &in.WeekDays, &out.WeekDays
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecurrenceParameters.
func (in *RecurrenceParameters) DeepCopy() *RecurrenceParameters {
	if in == nil {
		return nil
	}
	out := new(RecurrenceParameters)
	in.DeepCopyInto(out)
	return out
}
