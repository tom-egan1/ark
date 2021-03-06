/*
Copyright 2018 the Heptio Ark contributors.

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

package install

import (
	"fmt"

	arkv1 "github.com/heptio/ark/pkg/apis/ark/v1"

	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CRDs returns a list of the CRD types for all of the required Ark CRDs
func CRDs() []*apiextv1beta1.CustomResourceDefinition {
	return []*apiextv1beta1.CustomResourceDefinition{
		crd("Backup", "backups"),
		crd("Schedule", "schedules"),
		crd("Restore", "restores"),
		crd("Config", "configs"),
		crd("DownloadRequest", "downloadrequests"),
		crd("DeleteBackupRequest", "deletebackuprequests"),
		crd("PodVolumeBackup", "podvolumebackups"),
		crd("PodVolumeRestore", "podvolumerestores"),
		crd("ResticRepository", "resticrepositories"),
	}
}

func crd(kind, plural string) *apiextv1beta1.CustomResourceDefinition {
	return &apiextv1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s.%s", plural, arkv1.GroupName),
		},
		Spec: apiextv1beta1.CustomResourceDefinitionSpec{
			Group:   arkv1.GroupName,
			Version: arkv1.SchemeGroupVersion.Version,
			Scope:   apiextv1beta1.NamespaceScoped,
			Names: apiextv1beta1.CustomResourceDefinitionNames{
				Plural: plural,
				Kind:   kind,
			},
		},
	}
}
