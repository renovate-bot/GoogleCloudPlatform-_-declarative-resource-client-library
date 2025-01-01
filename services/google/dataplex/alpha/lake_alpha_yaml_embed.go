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
// gen_go_data -package alpha -var YAML_lake blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/dataplex/alpha/lake.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/dataplex/alpha/lake.yaml
var YAML_lake = []byte("info:\n  title: Dataplex/Lake\n  description: The Dataplex Lake resource\n  x-dcl-struct-name: Lake\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Lake\n    parameters:\n    - name: lake\n      required: true\n      description: A full instance of a Lake\n  apply:\n    description: The function used to apply information about a Lake\n    parameters:\n    - name: lake\n      required: true\n      description: A full instance of a Lake\n  delete:\n    description: The function used to delete a Lake\n    parameters:\n    - name: lake\n      required: true\n      description: A full instance of a Lake\n  deleteAll:\n    description: The function used to delete all Lake\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Lake\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Lake:\n      title: Lake\n      x-dcl-id: projects/{{project}}/locations/{{location}}/lakes/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        assetStatus:\n          type: object\n          x-dcl-go-name: AssetStatus\n          x-dcl-go-type: LakeAssetStatus\n          readOnly: true\n          description: Output only. Aggregated status of the underlying assets of\n            the lake.\n          properties:\n            activeAssets:\n              type: integer\n              format: int64\n              x-dcl-go-name: ActiveAssets\n              description: Number of active assets.\n            securityPolicyApplyingAssets:\n              type: integer\n              format: int64\n              x-dcl-go-name: SecurityPolicyApplyingAssets\n              description: Number of assets that are in process of updating the security\n                policy on attached resources.\n            updateTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: UpdateTime\n              description: Last update time of the status.\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time when the lake was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. Description of the lake.\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Optional. User friendly display name.\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. User-defined labels for the lake.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n          x-dcl-parameter: true\n        metastore:\n          type: object\n          x-dcl-go-name: Metastore\n          x-dcl-go-type: LakeMetastore\n          description: Optional. Settings to manage lake and Dataproc Metastore service\n            instance association.\n          properties:\n            service:\n              type: string\n              x-dcl-go-name: Service\n              description: 'Optional. A relative reference to the Dataproc Metastore\n                (https://cloud.google.com/dataproc-metastore/docs) service associated\n                with the lake: `projects/{project_id}/locations/{location_id}/services/{service_id}`'\n        metastoreStatus:\n          type: object\n          x-dcl-go-name: MetastoreStatus\n          x-dcl-go-type: LakeMetastoreStatus\n          readOnly: true\n          description: Output only. Metastore status of the lake.\n          properties:\n            endpoint:\n              type: string\n              x-dcl-go-name: Endpoint\n              description: The URI of the endpoint used to access the Metastore service.\n            message:\n              type: string\n              x-dcl-go-name: Message\n              description: Additional information about the current status.\n            state:\n              type: string\n              x-dcl-go-name: State\n              x-dcl-go-type: LakeMetastoreStatusStateEnum\n              description: 'Current state of association. Possible values: STATE_UNSPECIFIED,\n                NONE, READY, UPDATING, ERROR'\n              enum:\n              - STATE_UNSPECIFIED\n              - NONE\n              - READY\n              - UPDATING\n              - ERROR\n            updateTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: UpdateTime\n              description: Last update time of the metastore status of the lake.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of the lake.\n          x-dcl-references:\n          - resource: Dataplex/Lake\n            field: selfLink\n            parent: true\n          x-dcl-has-long-form: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-parameter: true\n        serviceAccount:\n          type: string\n          x-dcl-go-name: ServiceAccount\n          readOnly: true\n          description: Output only. Service account associated with this lake. This\n            service account must be authorized to access or operate on resources managed\n            by the lake.\n          x-kubernetes-immutable: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: LakeStateEnum\n          readOnly: true\n          description: 'Output only. Current state of the lake. Possible values: STATE_UNSPECIFIED,\n            ACTIVE, CREATING, DELETING, ACTION_REQUIRED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - ACTIVE\n          - CREATING\n          - DELETING\n          - ACTION_REQUIRED\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. System generated globally unique ID for the lake.\n            This ID will be different if the lake is deleted and re-created with the\n            same name.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time when the lake was last updated.\n          x-kubernetes-immutable: true\n")

// 7170 bytes
// MD5: f09e63b35b7ba0f6b3deb51885021302
