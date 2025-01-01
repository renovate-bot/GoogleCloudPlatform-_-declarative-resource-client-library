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
// gen_go_data -package beta -var YAML_release blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/firebaserules/beta/release.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/firebaserules/beta/release.yaml
var YAML_release = []byte("info:\n  title: Firebaserules/Release\n  description: \"\"\n  x-dcl-struct-name: Release\n  x-dcl-has-iam: false\n  x-dcl-ref:\n    text: Firebase Rules API Documentation\n    url: https://firebase.google.com/docs/reference/rules/rest#rest-resource:-v1.projects.releases\n  x-dcl-guides:\n  - text: Get started with Firebase Security Rules\n    url: https://firebase.google.com/docs/rules/get-started\npaths:\n  get:\n    description: The function used to get information about a Release\n    parameters:\n    - name: release\n      required: true\n      description: A full instance of a Release\n  apply:\n    description: The function used to apply information about a Release\n    parameters:\n    - name: release\n      required: true\n      description: A full instance of a Release\n  delete:\n    description: The function used to delete a Release\n    parameters:\n    - name: release\n      required: true\n      description: A full instance of a Release\n  deleteAll:\n    description: The function used to delete all Release\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Release\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Release:\n      title: Release\n      x-dcl-id: projects/{{project}}/releases/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - rulesetName\n      - project\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Time the release was created.\n          x-kubernetes-immutable: true\n        disabled:\n          type: boolean\n          x-dcl-go-name: Disabled\n          readOnly: true\n          description: Disable the release to keep it from being served. The response\n            code of NOT_FOUND will be given for executables generated from this Release.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Format: `projects/{project_id}/releases/{release_id}`\\Firestore\n            Rules Releases will **always** have the name ''cloud.firestore'''\n          x-kubernetes-immutable: true\n          x-dcl-forward-slash-allowed: true\n          x-dcl-has-long-form: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-parameter: true\n        rulesetName:\n          type: string\n          x-dcl-go-name: RulesetName\n          description: Name of the `Ruleset` referred to by this `Release`. The `Ruleset`\n            must exist for the `Release` to be created.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Firebaserules/Ruleset\n            field: name\n          x-dcl-has-long-form: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Time the release was updated.\n          x-kubernetes-immutable: true\n")

// 3528 bytes
// MD5: 2bf22d0aeeb46bf2d464c57e11e431df
