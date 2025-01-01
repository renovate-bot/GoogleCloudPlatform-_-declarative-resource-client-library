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
// gen_go_data -package alpha -var YAML_log_exclusion blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/logging/alpha/log_exclusion.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/logging/alpha/log_exclusion.yaml
var YAML_log_exclusion = []byte("info:\n  title: Logging/LogExclusion\n  description: The Logging LogExclusion resource\n  x-dcl-struct-name: LogExclusion\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a LogExclusion\n    parameters:\n    - name: logExclusion\n      required: true\n      description: A full instance of a LogExclusion\n  apply:\n    description: The function used to apply information about a LogExclusion\n    parameters:\n    - name: logExclusion\n      required: true\n      description: A full instance of a LogExclusion\n  delete:\n    description: The function used to delete a LogExclusion\n    parameters:\n    - name: logExclusion\n      required: true\n      description: A full instance of a LogExclusion\n  deleteAll:\n    description: The function used to delete all LogExclusion\n    parameters:\n    - name: parent\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many LogExclusion\n    parameters:\n    - name: parent\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    LogExclusion:\n      title: LogExclusion\n      x-dcl-id: '{{parent}}/exclusions/{{name}}'\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - filter\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The creation timestamp of the exclusion. This\n            field may not be present for older exclusions.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A description of this exclusion.\n        disabled:\n          type: boolean\n          x-dcl-go-name: Disabled\n          description: Optional. If set to True, then this exclusion is disabled and\n            it does not exclude any log entries. You can update an exclusion to change\n            the value of this field.\n        filter:\n          type: string\n          x-dcl-go-name: Filter\n          description: 'Required. An (https://cloud.google.com/logging/docs/view/advanced-queries#sample),\n            you can exclude less than 100% of the matching log entries. For example,\n            the following query matches 99% of low-severity log entries from Google\n            Cloud Storage buckets: `\"resource.type=gcs_bucket severity'\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. A client-assigned identifier, such as `\"load-balancer-exclusion\"`.\n            Identifiers are limited to 100 characters and can include only letters,\n            digits, underscores, hyphens, and periods. First character has to be alphanumeric.\n          x-kubernetes-immutable: true\n        parent:\n          type: string\n          x-dcl-go-name: Parent\n          description: 'The parent resource in which to create the exclusion: \"projects/[PROJECT_ID]\"\n            \"organizations/[ORGANIZATION_ID]\" \"billingAccounts/[BILLING_ACCOUNT_ID]\"\n            \"folders/[FOLDER_ID]\" Examples: \"projects/my-logging-project\", \"organizations/123456789\".\n            Authorization requires the following IAM permission on the specified resource\n            parent: logging.exclusions.create'\n          x-kubernetes-immutable: true\n          x-dcl-forward-slash-allowed: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          - resource: Cloudresourcemanager/Folder\n            field: name\n            parent: true\n          - resource: Cloudresourcemanager/Organization\n            field: name\n            parent: true\n          - resource: Cloudresourcemanager/BillingAccount\n            field: name\n            parent: true\n          x-dcl-parameter: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The last update timestamp of the exclusion. This\n            field may not be present for older exclusions.\n          x-kubernetes-immutable: true\n")

// 4317 bytes
// MD5: 84c7d2661d1f0fad087c1a5c0f093cde
