// Copyright 2025 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// GENERATED BY gen_go_data.go
// gen_go_data -package beta -var YAML_identity_aware_proxy_client blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iap/beta/identity_aware_proxy_client.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iap/beta/identity_aware_proxy_client.yaml
var YAML_identity_aware_proxy_client = []byte("info:\n  title: Iap/IdentityAwareProxyClient\n  description: The Iap IdentityAwareProxyClient resource\n  x-dcl-struct-name: IdentityAwareProxyClient\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a IdentityAwareProxyClient\n    parameters:\n    - name: identityAwareProxyClient\n      required: true\n      description: A full instance of a IdentityAwareProxyClient\n  apply:\n    description: The function used to apply information about a IdentityAwareProxyClient\n    parameters:\n    - name: identityAwareProxyClient\n      required: true\n      description: A full instance of a IdentityAwareProxyClient\n  delete:\n    description: The function used to delete a IdentityAwareProxyClient\n    parameters:\n    - name: identityAwareProxyClient\n      required: true\n      description: A full instance of a IdentityAwareProxyClient\n  deleteAll:\n    description: The function used to delete all IdentityAwareProxyClient\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: brand\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many IdentityAwareProxyClient\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: brand\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    IdentityAwareProxyClient:\n      title: IdentityAwareProxyClient\n      x-dcl-id: projects/{{project}}/brands/{{brand}}/identityAwareProxyClients/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - project\n      - brand\n      properties:\n        brand:\n          type: string\n          x-dcl-go-name: Brand\n          description: The brand for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Iap/Brand\n            field: name\n            parent: true\n          x-dcl-parameter: true\n          x-dcl-has-long-form: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Human-friendly name given to the OAuth client.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. Unique identifier of the OAuth client.\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n          x-dcl-has-long-form: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-parameter: true\n          x-dcl-has-long-form: true\n        secret:\n          type: string\n          x-dcl-go-name: Secret\n          readOnly: true\n          description: Output only. Client secret of the OAuth client.\n          x-kubernetes-immutable: true\n")

// 3187 bytes
// MD5: 7006bd02c1b199e602dc7b2e3f63e61f
