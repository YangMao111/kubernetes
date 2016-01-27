/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

// DO NOT EDIT. THIS FILE IS AUTO-GENERATED BY $KUBEROOT/hack/update-generated-conversions.sh

package v1

import (
	reflect "reflect"

	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
	extensions "k8s.io/kubernetes/pkg/apis/extensions"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func autoConvert_api_ObjectMeta_To_v1_ObjectMeta(in *api.ObjectMeta, out *v1.ObjectMeta, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*api.ObjectMeta))(in)
	}
	out.Name = in.Name
	out.GenerateName = in.GenerateName
	out.Namespace = in.Namespace
	out.SelfLink = in.SelfLink
	out.UID = in.UID
	out.ResourceVersion = in.ResourceVersion
	out.Generation = in.Generation
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.CreationTimestamp, &out.CreationTimestamp, s); err != nil {
		return err
	}
	// unable to generate simple pointer conversion for unversioned.Time -> unversioned.Time
	if in.DeletionTimestamp != nil {
		out.DeletionTimestamp = new(unversioned.Time)
		if err := api.Convert_unversioned_Time_To_unversioned_Time(in.DeletionTimestamp, out.DeletionTimestamp, s); err != nil {
			return err
		}
	} else {
		out.DeletionTimestamp = nil
	}
	if in.DeletionGracePeriodSeconds != nil {
		out.DeletionGracePeriodSeconds = new(int64)
		*out.DeletionGracePeriodSeconds = *in.DeletionGracePeriodSeconds
	} else {
		out.DeletionGracePeriodSeconds = nil
	}
	if in.Labels != nil {
		out.Labels = make(map[string]string)
		for key, val := range in.Labels {
			out.Labels[key] = val
		}
	} else {
		out.Labels = nil
	}
	if in.Annotations != nil {
		out.Annotations = make(map[string]string)
		for key, val := range in.Annotations {
			out.Annotations[key] = val
		}
	} else {
		out.Annotations = nil
	}
	return nil
}

func Convert_api_ObjectMeta_To_v1_ObjectMeta(in *api.ObjectMeta, out *v1.ObjectMeta, s conversion.Scope) error {
	return autoConvert_api_ObjectMeta_To_v1_ObjectMeta(in, out, s)
}

func autoConvert_v1_ObjectMeta_To_api_ObjectMeta(in *v1.ObjectMeta, out *api.ObjectMeta, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*v1.ObjectMeta))(in)
	}
	out.Name = in.Name
	out.GenerateName = in.GenerateName
	out.Namespace = in.Namespace
	out.SelfLink = in.SelfLink
	out.UID = in.UID
	out.ResourceVersion = in.ResourceVersion
	out.Generation = in.Generation
	if err := api.Convert_unversioned_Time_To_unversioned_Time(&in.CreationTimestamp, &out.CreationTimestamp, s); err != nil {
		return err
	}
	// unable to generate simple pointer conversion for unversioned.Time -> unversioned.Time
	if in.DeletionTimestamp != nil {
		out.DeletionTimestamp = new(unversioned.Time)
		if err := api.Convert_unversioned_Time_To_unversioned_Time(in.DeletionTimestamp, out.DeletionTimestamp, s); err != nil {
			return err
		}
	} else {
		out.DeletionTimestamp = nil
	}
	if in.DeletionGracePeriodSeconds != nil {
		out.DeletionGracePeriodSeconds = new(int64)
		*out.DeletionGracePeriodSeconds = *in.DeletionGracePeriodSeconds
	} else {
		out.DeletionGracePeriodSeconds = nil
	}
	if in.Labels != nil {
		out.Labels = make(map[string]string)
		for key, val := range in.Labels {
			out.Labels[key] = val
		}
	} else {
		out.Labels = nil
	}
	if in.Annotations != nil {
		out.Annotations = make(map[string]string)
		for key, val := range in.Annotations {
			out.Annotations[key] = val
		}
	} else {
		out.Annotations = nil
	}
	return nil
}

func Convert_v1_ObjectMeta_To_api_ObjectMeta(in *v1.ObjectMeta, out *api.ObjectMeta, s conversion.Scope) error {
	return autoConvert_v1_ObjectMeta_To_api_ObjectMeta(in, out, s)
}

func autoConvert_autoscaling_Scale_To_v1_Scale(in *autoscaling.Scale, out *Scale, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*autoscaling.Scale))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_api_ObjectMeta_To_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_autoscaling_ScaleSpec_To_v1_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_autoscaling_ScaleStatus_To_v1_ScaleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_autoscaling_Scale_To_v1_Scale(in *autoscaling.Scale, out *Scale, s conversion.Scope) error {
	return autoConvert_autoscaling_Scale_To_v1_Scale(in, out, s)
}

func autoConvert_autoscaling_ScaleSpec_To_v1_ScaleSpec(in *autoscaling.ScaleSpec, out *ScaleSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*autoscaling.ScaleSpec))(in)
	}
	out.Replicas = int32(in.Replicas)
	return nil
}

func Convert_autoscaling_ScaleSpec_To_v1_ScaleSpec(in *autoscaling.ScaleSpec, out *ScaleSpec, s conversion.Scope) error {
	return autoConvert_autoscaling_ScaleSpec_To_v1_ScaleSpec(in, out, s)
}

func autoConvert_autoscaling_ScaleStatus_To_v1_ScaleStatus(in *autoscaling.ScaleStatus, out *ScaleStatus, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*autoscaling.ScaleStatus))(in)
	}
	out.Replicas = int32(in.Replicas)
	out.Selector = in.Selector
	return nil
}

func Convert_autoscaling_ScaleStatus_To_v1_ScaleStatus(in *autoscaling.ScaleStatus, out *ScaleStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_ScaleStatus_To_v1_ScaleStatus(in, out, s)
}

func autoConvert_v1_HorizontalPodAutoscalerSpec_To_extensions_HorizontalPodAutoscalerSpec(in *HorizontalPodAutoscalerSpec, out *extensions.HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*HorizontalPodAutoscalerSpec))(in)
	}
	// in.ScaleTargetRef has no peer in out
	if in.MinReplicas != nil {
		out.MinReplicas = new(int)
		*out.MinReplicas = int(*in.MinReplicas)
	} else {
		out.MinReplicas = nil
	}
	out.MaxReplicas = int(in.MaxReplicas)
	// in.TargetCPUUtilizationPercentage has no peer in out
	return nil
}

func autoConvert_v1_Scale_To_autoscaling_Scale(in *Scale, out *autoscaling.Scale, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*Scale))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := Convert_v1_ObjectMeta_To_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	if err := Convert_v1_ScaleSpec_To_autoscaling_ScaleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ScaleStatus_To_autoscaling_ScaleStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_Scale_To_autoscaling_Scale(in *Scale, out *autoscaling.Scale, s conversion.Scope) error {
	return autoConvert_v1_Scale_To_autoscaling_Scale(in, out, s)
}

func autoConvert_v1_ScaleSpec_To_autoscaling_ScaleSpec(in *ScaleSpec, out *autoscaling.ScaleSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*ScaleSpec))(in)
	}
	out.Replicas = int(in.Replicas)
	return nil
}

func Convert_v1_ScaleSpec_To_autoscaling_ScaleSpec(in *ScaleSpec, out *autoscaling.ScaleSpec, s conversion.Scope) error {
	return autoConvert_v1_ScaleSpec_To_autoscaling_ScaleSpec(in, out, s)
}

func autoConvert_v1_ScaleStatus_To_autoscaling_ScaleStatus(in *ScaleStatus, out *autoscaling.ScaleStatus, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*ScaleStatus))(in)
	}
	out.Replicas = int(in.Replicas)
	out.Selector = in.Selector
	return nil
}

func Convert_v1_ScaleStatus_To_autoscaling_ScaleStatus(in *ScaleStatus, out *autoscaling.ScaleStatus, s conversion.Scope) error {
	return autoConvert_v1_ScaleStatus_To_autoscaling_ScaleStatus(in, out, s)
}

func autoConvert_extensions_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec(in *extensions.HorizontalPodAutoscalerSpec, out *HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*extensions.HorizontalPodAutoscalerSpec))(in)
	}
	// in.ScaleRef has no peer in out
	if in.MinReplicas != nil {
		out.MinReplicas = new(int32)
		*out.MinReplicas = int32(*in.MinReplicas)
	} else {
		out.MinReplicas = nil
	}
	out.MaxReplicas = int32(in.MaxReplicas)
	// in.CPUUtilization has no peer in out
	return nil
}

func init() {
	err := api.Scheme.AddGeneratedConversionFuncs(
		autoConvert_api_ObjectMeta_To_v1_ObjectMeta,
		autoConvert_autoscaling_ScaleSpec_To_v1_ScaleSpec,
		autoConvert_autoscaling_ScaleStatus_To_v1_ScaleStatus,
		autoConvert_autoscaling_Scale_To_v1_Scale,
		autoConvert_extensions_HorizontalPodAutoscalerSpec_To_v1_HorizontalPodAutoscalerSpec,
		autoConvert_v1_HorizontalPodAutoscalerSpec_To_extensions_HorizontalPodAutoscalerSpec,
		autoConvert_v1_ObjectMeta_To_api_ObjectMeta,
		autoConvert_v1_ScaleSpec_To_autoscaling_ScaleSpec,
		autoConvert_v1_ScaleStatus_To_autoscaling_ScaleStatus,
		autoConvert_v1_Scale_To_autoscaling_Scale,
	)
	if err != nil {
		// If one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}
}
