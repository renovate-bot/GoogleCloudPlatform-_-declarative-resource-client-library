// Copyright 2022 Google LLC. All Rights Reserved.
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
// gen_go_data -package alpha -var YAML_web_app blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/firebase/alpha/web_app.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/firebase/alpha/web_app.yaml
var YAML_web_app = []byte("info:\n  title: Firebase/WebApp\n  description: \"\"\n  x-dcl-struct-name: WebApp\n  x-dcl-has-create: true\n  x-dcl-has-iam: false\n  x-dcl-ref:\n    text: Firebase WebApps API Documentation\n    url: https://firebase.google.com/docs/projects/api/reference/rest#rest-resource:-v1beta1.projects.webapps\n  x-dcl-guides:\n  - text: Get started with Firebase Projects and Apps\n    url: https://firebase.google.com/docs/projects/api/workflow_set-up-and-manage-project\npaths:\n  get:\n    description: The function used to get information about a WebApp\n    parameters:\n    - name: WebApp\n      required: true\n      description: A full instance of a WebApp\n  apply:\n    description: The function used to apply information about a WebApp\n    parameters:\n    - name: WebApp\n      required: true\n      description: A full instance of a WebApp\n  delete:\n    description: The function used to delete a WebApp\n    parameters:\n    - name: WebApp\n      required: true\n      description: A full instance of a WebApp\n  deleteAll:\n    description: The function used to delete all WebApp\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many WebApp\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    WebApp:\n      title: WebApp\n      x-dcl-id: projects/{{project}}/webApps/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - project\n      properties:\n        apiKeyId:\n          type: string\n          x-dcl-go-name: ApiKeyId\n          description: The key_id of the GCP ApiKey associated with this App. If set\n            must have no restrictions, or only have restrictions that are valid for\n            the associated Firebase App. Cannot be set in create requests, instead\n            an existing valid API Key will be chosen, or if no valid API Keys exist,\n            one will be provisioned for you. Cannot be set to an empty value in update\n            requests.\n        appId:\n          type: string\n          x-dcl-go-name: AppId\n          readOnly: true\n          description: Immutable. The globally unique, Firebase-assigned identifier\n            for the `WebApp`. This identifier should be treated as an opaque token,\n            as the data format is not specified.\n          x-kubernetes-immutable: true\n        appUrls:\n          type: array\n          x-dcl-go-name: AppUrls\n          description: The URLs where the `WebApp` is hosted.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: The user-assigned display name for the `WebApp`.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Output Only.  The resource name of the WebApp, in the format:\n            `projects/PROJECT_IDENTIFIER/webApps/APP_ID` * PROJECT_IDENTIFIER: the\n            parent Project''s [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number)\n            ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id).\n            Learn more about using project identifiers in Google''s [AIP 2510 standard](https://google.aip.dev/cloud/2510).\n            Note that the value for PROJECT_IDENTIFIER in any response body will be\n            the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier\n            for the App (see [`appId`](../projects.webApps#WebApp.FIELDS.app_id)).'\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        projectId:\n          type: string\n          x-dcl-go-name: ProjectId\n          readOnly: true\n          description: Immutable. A user-assigned unique identifier of the parent\n            FirebaseProject for the `WebApp`.\n          x-kubernetes-immutable: true\n        webId:\n          type: string\n          x-dcl-go-name: WebId\n          readOnly: true\n          description: Output only. Immutable. A unique, Firebase-assigned identifier\n            for the `WebApp`. This identifier is only used to populate the `namespace`\n            value for the `WebApp`. For most use cases, use `appId` to identify or\n            reference the App. The `webId` value is only unique within a `FirebaseProject`\n            and its associated Apps.\n          x-kubernetes-immutable: true\n")

// 4848 bytes
// MD5: d27a095a80197c2fd0989fce76a5b2d2
