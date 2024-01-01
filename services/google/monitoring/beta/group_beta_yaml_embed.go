// Copyright 2024 Google LLC. All Rights Reserved.
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
// gen_go_data -package beta -var YAML_group blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/beta/group.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/monitoring/beta/group.yaml
var YAML_group = []byte("info:\n  title: Monitoring/Group\n  description: The Monitoring Group resource\n  x-dcl-struct-name: Group\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Group\n    parameters:\n    - name: group\n      required: true\n      description: A full instance of a Group\n  apply:\n    description: The function used to apply information about a Group\n    parameters:\n    - name: group\n      required: true\n      description: A full instance of a Group\n  delete:\n    description: The function used to delete a Group\n    parameters:\n    - name: group\n      required: true\n      description: A full instance of a Group\n  deleteAll:\n    description: The function used to delete all Group\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Group\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Group:\n      title: Group\n      x-dcl-id: projects/{{project}}/groups/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - displayName\n      - filter\n      - project\n      properties:\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: A user-assigned name for this group, used only for display\n            purposes.\n        filter:\n          type: string\n          x-dcl-go-name: Filter\n          description: The filter used to determine which monitored resources belong\n            to this group.\n        isCluster:\n          type: boolean\n          x-dcl-go-name: IsCluster\n          description: If true, the members of this group are considered to be a cluster.\n            The system can perform additional analysis on groups that are clusters.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Output only. The name of this group. The format is: `projects/{{project}}/groups/{{name}}`,\n            which is generated automatically.'\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        parentName:\n          type: string\n          x-dcl-go-name: ParentName\n          description: 'The name of the group''s parent, if it has one. The format\n            is: projects/ For groups with no parent, `parent_name` is the empty string,\n            ``.'\n          x-dcl-send-empty: true\n          x-dcl-references:\n          - resource: Monitoring/Group\n            field: name\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project of the group\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n")

// 2994 bytes
// MD5: 5b0f18ecb40670caa9a101b3177f4844
