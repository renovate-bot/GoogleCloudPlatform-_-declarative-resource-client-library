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
// gen_go_data -package alpha -var YAML_firebase_project blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/firebase/alpha/firebase_project.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/firebase/alpha/firebase_project.yaml
var YAML_firebase_project = []byte("info:\n  title: Firebase/FirebaseProject\n  description: \"\"\n  x-dcl-struct-name: FirebaseProject\n  x-dcl-has-iam: false\n  x-dcl-ref:\n    text: Firebase Project API Documentation\n    url: https://firebase.google.com/docs/projects/api/reference/rest#rest-resource:-v1beta1.projects\n  x-dcl-guides:\n  - text: Get started with Firebase Projects and Apps\n    url: https://firebase.google.com/docs/projects/api/workflow_set-up-and-manage-project\npaths:\n  get:\n    description: The function used to get information about a FirebaseProject\n    parameters:\n    - name: firebaseProject\n      required: true\n      description: A full instance of a FirebaseProject\n  apply:\n    description: The function used to apply information about a FirebaseProject\n    parameters:\n    - name: firebaseProject\n      required: true\n      description: A full instance of a FirebaseProject\n  list:\n    description: The function used to list information about many FirebaseProject\n    parameters: []\ncomponents:\n  schemas:\n    FirebaseProject:\n      title: FirebaseProject\n      x-dcl-id: projects/{{project}}\n      x-dcl-parent-container: project\n      x-dcl-labels: annotations\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - project\n      properties:\n        annotations:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Annotations\n          description: '  // Set of user-defined annotations for the [FirebaseProject][]\n            as per [AIP-128](https://google.aip.dev/128#annotations).  These annotations\n            are intended solely for developers and client-side tools Firebase services\n            will not mutate this annotation set.  This field may only be assigned\n            on Update'\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: The user-assigned display name of the Project.  This field\n            may only be assigned on Update\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource.  FirebaseProjects are generally\n            referneced by the GCP Project they augment.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-parameter: true\n          x-dcl-has-long-form: true\n        projectId:\n          type: string\n          x-dcl-go-name: ProjectId\n          readOnly: true\n          description: Immutable. A user-assigned unique identifier for the Project.\n            This identifier may appear in URLs or names for some Firebase resources\n            associated with the Project, but it should generally be treated as a convenience\n            alias to reference the Project.\n          x-kubernetes-immutable: true\n        projectNumber:\n          type: integer\n          format: int64\n          x-dcl-go-name: ProjectNumber\n          readOnly: true\n          description: Immutable. The globally unique, Google-assigned canonical identifier\n            for the Project. Use this identifier when configuring integrations and/or\n            making API calls to Firebase or third-party services.\n          x-kubernetes-immutable: true\n        resources:\n          type: object\n          x-dcl-go-name: Resources\n          x-dcl-go-type: FirebaseProjectResources\n          readOnly: true\n          description: The default Firebase resources associated with the Project.\n          x-kubernetes-immutable: true\n          properties:\n            hostingSite:\n              type: string\n              x-dcl-go-name: HostingSite\n              readOnly: true\n              description: 'The default Firebase Hosting site name, in the format:\n                `PROJECT_ID` Though rare, your `projectId` might already be used as\n                the name for an existing Hosting site in another project (learn more\n                about creating non-default, [additional sites](https://firebase.google.com/docs/hosting/multisites)).\n                In these cases, your `projectId` is appended with a hyphen then five\n                alphanumeric characters to create your default Hosting site name.\n                For example, if your `projectId` is `myproject123`, your default Hosting\n                site name might be: `myproject123-a5c16`'\n              x-kubernetes-immutable: true\n            locationId:\n              type: string\n              x-dcl-go-name: LocationId\n              readOnly: true\n              description: The ID of the Project's default GCP resource location.\n                The location is one of the available [GCP resource locations](https://firebase.google.com/docs/projects/locations).\n                This field is omitted if the default GCP resource location has not\n                been finalized yet. To set a Project's default GCP resource location,\n                call [`FinalizeDefaultLocation`](../projects.defaultLocation/finalize)\n                after you add Firebase resources to the Project.\n              x-kubernetes-immutable: true\n            realtimeDatabaseInstance:\n              type: string\n              x-dcl-go-name: RealtimeDatabaseInstance\n              readOnly: true\n              description: 'The default Firebase Realtime Database instance name,\n                in the format: `PROJECT_ID` Though rare, your `projectId` might already\n                be used as the name for an existing Realtime Database instance in\n                another project (learn more about [database sharding](https://firebase.google.com/docs/database/usage/sharding)).\n                In these cases, your `projectId` is appended with a hyphen then five\n                alphanumeric characters to create your default Realtime Database instance\n                name. For example, if your `projectId` is `myproject123`, your default\n                database instance name might be: `myproject123-a5c16`'\n              x-kubernetes-immutable: true\n            storageBucket:\n              type: string\n              x-dcl-go-name: StorageBucket\n              readOnly: true\n              description: 'The default Cloud Storage for Firebase storage bucket,\n                in the format: `PROJECT_ID.appspot.com`'\n              x-kubernetes-immutable: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: FirebaseProjectStateEnum\n          readOnly: true\n          description: 'Output only. The lifecycle state of the Project. Updates to\n            the state must be performed via com.google.cloudresourcemanager.v1.Projects.DeleteProject\n            and com.google.cloudresourcemanager.v1.Projects.UndeleteProject Possible\n            values: STATE_UNSPECIFIED, ACTIVE, DELETED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - ACTIVE\n          - DELETED\n")

// 7043 bytes
// MD5: b1d7bebdbfd6660f994cbf951b4f198d
