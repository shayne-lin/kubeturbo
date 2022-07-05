/*
Copyright 2022.

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

package v1alpha1

// PolicyReference represents a declarative configuration of the reference to a Policy
type PolicyReference struct {
	// The kind of the policy referent
	// +kubebuilder:validation:Enum=SLOHorizontalScale
	Kind string `json:"kind"`

	// The name of the policy referent
	Name string `json:"name"`
}
