/*
Copyright 2019 Banzai Cloud.

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

package pilot

import (
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/banzaicloud/istio-operator/pkg/resources/templates"
)

func (r *Reconciler) clusterRoleGalley() runtime.Object {
	return &rbacv1.ClusterRole{
		ObjectMeta: templates.ObjectMetaClusterScope(clusterRoleNameGalley, galleyLabels, r.Config),
		Rules: []rbacv1.PolicyRule{
			{
				// For reading Istio resources
				APIGroups: []string{"authentication.istio.io", "config.istio.io", "networking.istio.io", "rbac.istio.io", "security.istio.io"},
				Resources: []string{"*"},
				Verbs:     []string{"get", "list", "watch"},
			},
			{
				// For updating Istio resource statuses
				APIGroups: []string{"authentication.istio.io", "config.istio.io", "networking.istio.io", "rbac.istio.io", "security.istio.io"},
				Resources: []string{"*/status"},
				Verbs:     []string{"update"},
			},
			{
				APIGroups: []string{"extensions", "apps"},
				Resources: []string{"deployments"},
				Verbs:     []string{"get", "list", "watch"},
			},
			{
				APIGroups: []string{""},
				Resources: []string{"pods", "nodes", "services", "endpoints", "namespaces"},
				Verbs:     []string{"get", "list", "watch"},
			},
			{
				APIGroups: []string{"extensions"},
				Resources: []string{"ingresses"},
				Verbs:     []string{"get", "list", "watch"},
			},
			{
				APIGroups:     []string{"extensions"},
				Resources:     []string{"deployments/finalizers"},
				ResourceNames: []string{"istio-galley"},
				Verbs:         []string{"update"},
			},
			{
				APIGroups: []string{"apiextensions.k8s.io"},
				Resources: []string{"customresourcedefinitions"},
				Verbs:     []string{"get", "list", "watch"},
			},
			{
				APIGroups: []string{"rbac.authorization.k8s.io"},
				Resources: []string{"clusterroles"},
				Verbs:     []string{"get", "list", "watch"},
			},
		},
	}
}