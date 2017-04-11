// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package innsmouth

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	builders "k8s.io/apiserver-builder/pkg/builders"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_innsmouth_DeepOne, InType: reflect.TypeOf(&DeepOne{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_innsmouth_DeepOneList, InType: reflect.TypeOf(&DeepOneList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_innsmouth_DeepOneSpec, InType: reflect.TypeOf(&DeepOneSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_innsmouth_DeepOneStatus, InType: reflect.TypeOf(&DeepOneStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_innsmouth_DeepOneStatusStrategy, InType: reflect.TypeOf(&DeepOneStatusStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_innsmouth_DeepOneStrategy, InType: reflect.TypeOf(&DeepOneStrategy{})},
	)
}

func DeepCopy_innsmouth_DeepOne(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeepOne)
		out := out.(*DeepOne)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		return nil
	}
}

func DeepCopy_innsmouth_DeepOneList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeepOneList)
		out := out.(*DeepOneList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]DeepOne, len(*in))
			for i := range *in {
				if err := DeepCopy_innsmouth_DeepOne(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_innsmouth_DeepOneSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeepOneSpec)
		out := out.(*DeepOneSpec)
		*out = *in
		return nil
	}
}

func DeepCopy_innsmouth_DeepOneStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeepOneStatus)
		out := out.(*DeepOneStatus)
		*out = *in
		return nil
	}
}

func DeepCopy_innsmouth_DeepOneStatusStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeepOneStatusStrategy)
		out := out.(*DeepOneStatusStrategy)
		*out = *in
		if newVal, err := c.DeepCopy(&in.DefaultStatusStorageStrategy); err != nil {
			return err
		} else {
			out.DefaultStatusStorageStrategy = *newVal.(*builders.DefaultStatusStorageStrategy)
		}
		return nil
	}
}

func DeepCopy_innsmouth_DeepOneStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeepOneStrategy)
		out := out.(*DeepOneStrategy)
		*out = *in
		if newVal, err := c.DeepCopy(&in.DefaultStorageStrategy); err != nil {
			return err
		} else {
			out.DefaultStorageStrategy = *newVal.(*builders.DefaultStorageStrategy)
		}
		return nil
	}
}