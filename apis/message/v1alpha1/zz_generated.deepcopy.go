//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Queue) DeepCopyInto(out *Queue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Queue.
func (in *Queue) DeepCopy() *Queue {
	if in == nil {
		return nil
	}
	out := new(Queue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Queue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueInitParameters) DeepCopyInto(out *QueueInitParameters) {
	*out = *in
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(string)
		**out = **in
	}
	if in.AccessKeyRef != nil {
		in, out := &in.AccessKeyRef, &out.AccessKeyRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AccessKeySelector != nil {
		in, out := &in.AccessKeySelector, &out.AccessKeySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ContentBasedDeduplication != nil {
		in, out := &in.ContentBasedDeduplication, &out.ContentBasedDeduplication
		*out = new(bool)
		**out = **in
	}
	if in.DelaySeconds != nil {
		in, out := &in.DelaySeconds, &out.DelaySeconds
		*out = new(float64)
		**out = **in
	}
	if in.FifoQueue != nil {
		in, out := &in.FifoQueue, &out.FifoQueue
		*out = new(bool)
		**out = **in
	}
	if in.MaxMessageSize != nil {
		in, out := &in.MaxMessageSize, &out.MaxMessageSize
		*out = new(float64)
		**out = **in
	}
	if in.MessageRetentionSeconds != nil {
		in, out := &in.MessageRetentionSeconds, &out.MessageRetentionSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.ReceiveWaitTimeSeconds != nil {
		in, out := &in.ReceiveWaitTimeSeconds, &out.ReceiveWaitTimeSeconds
		*out = new(float64)
		**out = **in
	}
	if in.RedrivePolicy != nil {
		in, out := &in.RedrivePolicy, &out.RedrivePolicy
		*out = new(string)
		**out = **in
	}
	if in.RegionID != nil {
		in, out := &in.RegionID, &out.RegionID
		*out = new(string)
		**out = **in
	}
	if in.VisibilityTimeoutSeconds != nil {
		in, out := &in.VisibilityTimeoutSeconds, &out.VisibilityTimeoutSeconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueInitParameters.
func (in *QueueInitParameters) DeepCopy() *QueueInitParameters {
	if in == nil {
		return nil
	}
	out := new(QueueInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueList) DeepCopyInto(out *QueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Queue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueList.
func (in *QueueList) DeepCopy() *QueueList {
	if in == nil {
		return nil
	}
	out := new(QueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueObservation) DeepCopyInto(out *QueueObservation) {
	*out = *in
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ContentBasedDeduplication != nil {
		in, out := &in.ContentBasedDeduplication, &out.ContentBasedDeduplication
		*out = new(bool)
		**out = **in
	}
	if in.DelaySeconds != nil {
		in, out := &in.DelaySeconds, &out.DelaySeconds
		*out = new(float64)
		**out = **in
	}
	if in.FifoQueue != nil {
		in, out := &in.FifoQueue, &out.FifoQueue
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MaxMessageSize != nil {
		in, out := &in.MaxMessageSize, &out.MaxMessageSize
		*out = new(float64)
		**out = **in
	}
	if in.MessageRetentionSeconds != nil {
		in, out := &in.MessageRetentionSeconds, &out.MessageRetentionSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.ReceiveWaitTimeSeconds != nil {
		in, out := &in.ReceiveWaitTimeSeconds, &out.ReceiveWaitTimeSeconds
		*out = new(float64)
		**out = **in
	}
	if in.RedrivePolicy != nil {
		in, out := &in.RedrivePolicy, &out.RedrivePolicy
		*out = new(string)
		**out = **in
	}
	if in.RegionID != nil {
		in, out := &in.RegionID, &out.RegionID
		*out = new(string)
		**out = **in
	}
	if in.VisibilityTimeoutSeconds != nil {
		in, out := &in.VisibilityTimeoutSeconds, &out.VisibilityTimeoutSeconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueObservation.
func (in *QueueObservation) DeepCopy() *QueueObservation {
	if in == nil {
		return nil
	}
	out := new(QueueObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueParameters) DeepCopyInto(out *QueueParameters) {
	*out = *in
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(string)
		**out = **in
	}
	if in.AccessKeyRef != nil {
		in, out := &in.AccessKeyRef, &out.AccessKeyRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AccessKeySelector != nil {
		in, out := &in.AccessKeySelector, &out.AccessKeySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ContentBasedDeduplication != nil {
		in, out := &in.ContentBasedDeduplication, &out.ContentBasedDeduplication
		*out = new(bool)
		**out = **in
	}
	if in.DelaySeconds != nil {
		in, out := &in.DelaySeconds, &out.DelaySeconds
		*out = new(float64)
		**out = **in
	}
	if in.FifoQueue != nil {
		in, out := &in.FifoQueue, &out.FifoQueue
		*out = new(bool)
		**out = **in
	}
	if in.MaxMessageSize != nil {
		in, out := &in.MaxMessageSize, &out.MaxMessageSize
		*out = new(float64)
		**out = **in
	}
	if in.MessageRetentionSeconds != nil {
		in, out := &in.MessageRetentionSeconds, &out.MessageRetentionSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.ReceiveWaitTimeSeconds != nil {
		in, out := &in.ReceiveWaitTimeSeconds, &out.ReceiveWaitTimeSeconds
		*out = new(float64)
		**out = **in
	}
	if in.RedrivePolicy != nil {
		in, out := &in.RedrivePolicy, &out.RedrivePolicy
		*out = new(string)
		**out = **in
	}
	if in.RegionID != nil {
		in, out := &in.RegionID, &out.RegionID
		*out = new(string)
		**out = **in
	}
	if in.SecretKeySecretRef != nil {
		in, out := &in.SecretKeySecretRef, &out.SecretKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.VisibilityTimeoutSeconds != nil {
		in, out := &in.VisibilityTimeoutSeconds, &out.VisibilityTimeoutSeconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueParameters.
func (in *QueueParameters) DeepCopy() *QueueParameters {
	if in == nil {
		return nil
	}
	out := new(QueueParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueSpec) DeepCopyInto(out *QueueSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueSpec.
func (in *QueueSpec) DeepCopy() *QueueSpec {
	if in == nil {
		return nil
	}
	out := new(QueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueStatus) DeepCopyInto(out *QueueStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueStatus.
func (in *QueueStatus) DeepCopy() *QueueStatus {
	if in == nil {
		return nil
	}
	out := new(QueueStatus)
	in.DeepCopyInto(out)
	return out
}
