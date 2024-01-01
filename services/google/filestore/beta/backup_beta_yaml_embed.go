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
// gen_go_data -package beta -var YAML_backup blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/filestore/beta/backup.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/filestore/beta/backup.yaml
var YAML_backup = []byte("info:\n  title: Filestore/Backup\n  description: The Filestore Backup resource\n  x-dcl-struct-name: Backup\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Backup\n    parameters:\n    - name: backup\n      required: true\n      description: A full instance of a Backup\n  apply:\n    description: The function used to apply information about a Backup\n    parameters:\n    - name: backup\n      required: true\n      description: A full instance of a Backup\n  delete:\n    description: The function used to delete a Backup\n    parameters:\n    - name: backup\n      required: true\n      description: A full instance of a Backup\n  deleteAll:\n    description: The function used to delete all Backup\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Backup\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Backup:\n      title: Backup\n      x-dcl-id: projects/{{project}}/locations/{{location}}/backups/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - sourceInstance\n      - sourceFileShare\n      - project\n      - location\n      properties:\n        capacityGb:\n          type: integer\n          format: int64\n          x-dcl-go-name: CapacityGb\n          readOnly: true\n          description: Output only. Capacity of the source file share when the backup\n            was created.\n          x-kubernetes-immutable: true\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time when the backup was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: A description of the backup with 2048 characters or less. Requests\n            with longer descriptions will be rejected.\n        downloadBytes:\n          type: integer\n          format: int64\n          x-dcl-go-name: DownloadBytes\n          readOnly: true\n          description: Output only. Amount of bytes that will be downloaded if the\n            backup is restored. This may be different than storage bytes, since sequential\n            backups of the same disk will share storage.\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Resource labels to represent user provided metadata.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The resource name of the backup.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        sourceFileShare:\n          type: string\n          x-dcl-go-name: SourceFileShare\n          description: Name of the file share in the source Cloud Filestore instance\n            that the backup is created from.\n          x-kubernetes-immutable: true\n        sourceInstance:\n          type: string\n          x-dcl-go-name: SourceInstance\n          description: The resource name of the source Cloud Filestore instance, in\n            the format projects/{project_number}/locations/{location_id}/instances/{instance_id},\n            used to create this backup.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Filestore/Instance\n            field: name\n        sourceInstanceTier:\n          type: string\n          x-dcl-go-name: SourceInstanceTier\n          x-dcl-go-type: BackupSourceInstanceTierEnum\n          readOnly: true\n          description: 'Output only. The service tier of the source Cloud Filestore\n            instance that this backup is created from. Possible values: TIER_UNSPECIFIED,\n            STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD'\n          x-kubernetes-immutable: true\n          enum:\n          - TIER_UNSPECIFIED\n          - STANDARD\n          - PREMIUM\n          - BASIC_HDD\n          - BASIC_SSD\n          - HIGH_SCALE_SSD\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: BackupStateEnum\n          readOnly: true\n          description: 'Output only. The backup state. Possible values: STATE_UNSPECIFIED,\n            CREATING, READY, REPAIRING, DELETING, ERROR, RESTORING'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - CREATING\n          - READY\n          - REPAIRING\n          - DELETING\n          - ERROR\n          - RESTORING\n        storageBytes:\n          type: integer\n          format: int64\n          x-dcl-go-name: StorageBytes\n          readOnly: true\n          description: Output only. The size of the storage used by the backup. As\n            backups share storage, this number is expected to change with backup creation/deletion.\n          x-kubernetes-immutable: true\n")

// 5786 bytes
// MD5: 20cf6134f3cc93bde459fd31b7115f76
