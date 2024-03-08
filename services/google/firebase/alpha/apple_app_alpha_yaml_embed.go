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
// gen_go_data -package alpha -var YAML_apple_app blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/firebase/alpha/apple_app.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/firebase/alpha/apple_app.yaml
var YAML_apple_app = []byte("info:\n  title: Firebase/AppleApp\n  description: \"\"\n  x-dcl-struct-name: AppleApp\n  x-dcl-has-iam: false\n  x-dcl-ref:\n    text: Firebase AppleApp API Documentation\n    url: https://firebase.google.com/docs/projects/api/reference/rest#rest-resource:-v1beta1.projects.iosapps\n  x-dcl-guides:\n  - text: Get started with Firebase Projects and Apps\n    url: https://firebase.google.com/docs/projects/api/workflow_set-up-and-manage-project\npaths:\n  get:\n    description: The function used to get information about a AppleApp\n    parameters:\n    - name: appleApp\n      required: true\n      description: A full instance of a AppleApp\n  apply:\n    description: The function used to apply information about a AppleApp\n    parameters:\n    - name: appleApp\n      required: true\n      description: A full instance of a AppleApp\n  delete:\n    description: The function used to delete a AppleApp\n    parameters:\n    - name: appleApp\n      required: true\n      description: A full instance of a AppleApp\n  deleteAll:\n    description: The function used to delete all AppleApp\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many AppleApp\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    AppleApp:\n      title: AppleApp\n      x-dcl-id: projects/{{project}}/iosApps/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - bundleId\n      - project\n      properties:\n        apiKeyId:\n          type: string\n          x-dcl-go-name: ApiKeyId\n          description: The key_id of the GCP ApiKey associated with this App. If set\n            must have no restrictions, or only have restrictions that are valid for\n            the associated Firebase App. Cannot be set in create requests, instead\n            an existing valid API Key will be chosen, or if no valid API Keys exist,\n            one will be provisioned for you. Cannot be set to an empty value in update\n            requests.\n          x-dcl-references:\n          - resource: Apikeys/Key\n            field: name\n        appId:\n          type: string\n          x-dcl-go-name: AppId\n          readOnly: true\n          description: Output only. Immutable. The globally unique, Firebase-assigned\n            identifier for the `AppleApp`. This identifier should be treated as an\n            opaque token, as the data format is not specified.\n          x-kubernetes-immutable: true\n        appStoreId:\n          type: string\n          x-dcl-go-name: AppStoreId\n          description: The automatically generated Apple ID assigned to the Apple\n            app by Apple in the Apple App Store.\n        bundleId:\n          type: string\n          x-dcl-go-name: BundleId\n          description: Immutable. The canonical bundle ID of the Apple app as it would\n            appear in the Apple AppStore.\n          x-kubernetes-immutable: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: The user-assigned display name for the `AppleApp`.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'The resource name of the AppleApp, in the format: `projects/PROJECT_IDENTIFIER/iosApps/APP_ID`\n            * PROJECT_IDENTIFIER: the parent Project''s [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number)\n            ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id).\n            Learn more about using project identifiers in Google''s [AIP 2510 standard](https://google.aip.dev/cloud/2510).\n            Note that the value for PROJECT_IDENTIFIER in any response body will be\n            the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier\n            for the App (see [`appId`](../projects.iosApps#IosApp.FIELDS.app_id)).'\n          x-dcl-server-generated-parameter: true\n          x-dcl-has-long-form: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-parameter: true\n          x-dcl-has-long-form: true\n        projectId:\n          type: string\n          x-dcl-go-name: ProjectId\n          readOnly: true\n          description: Output only. Immutable. A user-assigned unique identifier of\n            the parent FirebaseProject for the `AppleApp`.\n          x-kubernetes-immutable: true\n        teamId:\n          type: string\n          x-dcl-go-name: TeamId\n          description: The Apple Developer Team ID associated with the App in the\n            App Store.\n")

// 4983 bytes
// MD5: 559841e14e641fac23db6ac316284ace
