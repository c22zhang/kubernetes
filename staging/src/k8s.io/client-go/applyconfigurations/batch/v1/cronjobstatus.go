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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// CronJobStatusApplyConfiguration represents a declarative configuration of the CronJobStatus type for use
// with apply.
type CronJobStatusApplyConfiguration struct {
	Active             []corev1.ObjectReferenceApplyConfiguration `json:"active,omitempty"`
	LastScheduleTime   *metav1.Time                               `json:"lastScheduleTime,omitempty"`
	LastSuccessfulTime *metav1.Time                               `json:"lastSuccessfulTime,omitempty"`
}

// CronJobStatusApplyConfiguration constructs a declarative configuration of the CronJobStatus type for use with
// apply.
func CronJobStatus() *CronJobStatusApplyConfiguration {
	return &CronJobStatusApplyConfiguration{}
}
func (b CronJobStatusApplyConfiguration) IsApplyConfiguration() {}

// WithActive adds the given value to the Active field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Active field.
func (b *CronJobStatusApplyConfiguration) WithActive(values ...*corev1.ObjectReferenceApplyConfiguration) *CronJobStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithActive")
		}
		b.Active = append(b.Active, *values[i])
	}
	return b
}

// WithLastScheduleTime sets the LastScheduleTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastScheduleTime field is set to the value of the last call.
func (b *CronJobStatusApplyConfiguration) WithLastScheduleTime(value metav1.Time) *CronJobStatusApplyConfiguration {
	b.LastScheduleTime = &value
	return b
}

// WithLastSuccessfulTime sets the LastSuccessfulTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastSuccessfulTime field is set to the value of the last call.
func (b *CronJobStatusApplyConfiguration) WithLastSuccessfulTime(value metav1.Time) *CronJobStatusApplyConfiguration {
	b.LastSuccessfulTime = &value
	return b
}
