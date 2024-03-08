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
// gen_go_data -package beta -var YAML_target blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/clouddeploy/beta/target.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/clouddeploy/beta/target.yaml
var YAML_target = []byte("info:\n  title: Clouddeploy/Target\n  description: The Cloud Deploy `Target` resource\n  x-dcl-struct-name: Target\n  x-dcl-has-iam: false\n  x-dcl-ref:\n    text: REST API\n    url: https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.targets\npaths:\n  get:\n    description: The function used to get information about a Target\n    parameters:\n    - name: target\n      required: true\n      description: A full instance of a Target\n  apply:\n    description: The function used to apply information about a Target\n    parameters:\n    - name: target\n      required: true\n      description: A full instance of a Target\n  delete:\n    description: The function used to delete a Target\n    parameters:\n    - name: target\n      required: true\n      description: A full instance of a Target\n  deleteAll:\n    description: The function used to delete all Target\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Target\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Target:\n      title: Target\n      x-dcl-id: projects/{{project}}/locations/{{location}}/targets/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        annotations:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Annotations\n          description: Optional. User annotations. These attributes can only be set\n            and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations\n            for more details such as format and size limitations.\n        anthosCluster:\n          type: object\n          x-dcl-go-name: AnthosCluster\n          x-dcl-go-type: TargetAnthosCluster\n          description: Information specifying an Anthos Cluster.\n          x-dcl-conflicts:\n          - gke\n          - run\n          - multiTarget\n          properties:\n            membership:\n              type: string\n              x-dcl-go-name: Membership\n              description: Membership of the GKE Hub-registered cluster to which to\n                apply the Skaffold configuration. Format is `projects/{project}/locations/{location}/memberships/{membership_name}`.\n              x-dcl-references:\n              - resource: Gkehub/Membership\n                field: selfLink\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. Time at which the `Target` was created.\n          x-kubernetes-immutable: true\n        deployParameters:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: DeployParameters\n          description: Optional. The deploy parameters to use for this target.\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. Description of the `Target`. Max length is 255 characters.\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Optional. This checksum is computed by the server based on\n            the value of other fields, and may be sent on update and delete requests\n            to ensure the client has an up-to-date value before proceeding.\n          x-kubernetes-immutable: true\n        executionConfigs:\n          type: array\n          x-dcl-go-name: ExecutionConfigs\n          description: Configurations for all execution that relates to this `Target`.\n            Each `ExecutionEnvironmentUsage` value may only be used in a single configuration;\n            using the same value multiple times is an error. When one or more configurations\n            are specified, they must include the `RENDER` and `DEPLOY` `ExecutionEnvironmentUsage`\n            values. When no configurations are specified, execution will use the default\n            specified in `DefaultPool`.\n          x-dcl-server-default: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: TargetExecutionConfigs\n            required:\n            - usages\n            properties:\n              artifactStorage:\n                type: string\n                x-dcl-go-name: ArtifactStorage\n                description: Optional. Cloud Storage location in which to store execution\n                  outputs. This can either be a bucket (\"gs://my-bucket\") or a path\n                  within a bucket (\"gs://my-bucket/my-dir\"). If unspecified, a default\n                  bucket located in the same region will be used.\n                x-dcl-server-default: true\n              executionTimeout:\n                type: string\n                x-dcl-go-name: ExecutionTimeout\n                description: Optional. Execution timeout for a Cloud Build Execution.\n                  This must be between 10m and 24h in seconds format. If unspecified,\n                  a default timeout of 1h is used.\n                x-dcl-server-default: true\n              serviceAccount:\n                type: string\n                x-dcl-go-name: ServiceAccount\n                description: Optional. Google service account to use for execution.\n                  If unspecified, the project execution service account (-compute@developer.gserviceaccount.com)\n                  is used.\n                x-dcl-server-default: true\n              usages:\n                type: array\n                x-dcl-go-name: Usages\n                description: Required. Usages when this configuration should be applied.\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: string\n                  x-dcl-go-type: TargetExecutionConfigsUsagesEnum\n                  enum:\n                  - EXECUTION_ENVIRONMENT_USAGE_UNSPECIFIED\n                  - RENDER\n                  - DEPLOY\n              workerPool:\n                type: string\n                x-dcl-go-name: WorkerPool\n                description: Optional. The resource name of the `WorkerPool`, with\n                  the format `projects/{project}/locations/{location}/workerPools/{worker_pool}`.\n                  If this optional field is unspecified, the default Cloud Build pool\n                  will be used.\n                x-dcl-references:\n                - resource: Cloudbuild/WorkerPool\n                  field: selfLink\n        gke:\n          type: object\n          x-dcl-go-name: Gke\n          x-dcl-go-type: TargetGke\n          description: Information specifying a GKE Cluster.\n          x-dcl-conflicts:\n          - anthosCluster\n          - run\n          - multiTarget\n          properties:\n            cluster:\n              type: string\n              x-dcl-go-name: Cluster\n              description: Information specifying a GKE Cluster. Format is `projects/{project_id}/locations/{location_id}/clusters/{cluster_id}.\n              x-dcl-references:\n              - resource: Container/Cluster\n                field: selfLink\n            internalIP:\n              type: boolean\n              x-dcl-go-name: InternalIP\n              description: Optional. If true, `cluster` is accessed using the private\n                IP address of the control plane endpoint. Otherwise, the default IP\n                address of the control plane endpoint is used. The default IP address\n                is the private IP address for clusters with private control-plane\n                endpoints and the public IP address otherwise. Only specify this option\n                when `cluster` is a [private GKE cluster](https://cloud.google.com/kubernetes-engine/docs/concepts/private-cluster-concept).\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: 'Optional. Labels are attributes that can be set and used by\n            both the user and by Google Cloud Deploy. Labels must meet the following\n            constraints: * Keys and values can contain only lowercase letters, numeric\n            characters, underscores, and dashes. * All characters must use UTF-8 encoding,\n            and international characters are allowed. * Keys must start with a lowercase\n            letter or international character. * Each resource is limited to a maximum\n            of 64 labels. Both keys and values are additionally constrained to be\n            <= 128 bytes.'\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n          x-dcl-parameter: true\n        multiTarget:\n          type: object\n          x-dcl-go-name: MultiTarget\n          x-dcl-go-type: TargetMultiTarget\n          description: Information specifying a multiTarget.\n          x-dcl-conflicts:\n          - gke\n          - anthosCluster\n          - run\n          required:\n          - targetIds\n          properties:\n            targetIds:\n              type: array\n              x-dcl-go-name: TargetIds\n              description: Required. The target_ids of this multiTarget.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Name of the `Target`. Format is [a-z][a-z0-9\\-]{0,62}.\n          x-kubernetes-immutable: true\n          x-dcl-parameter: true\n          x-dcl-has-long-form: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-parameter: true\n        requireApproval:\n          type: boolean\n          x-dcl-go-name: RequireApproval\n          description: Optional. Whether or not the `Target` requires approval.\n        run:\n          type: object\n          x-dcl-go-name: Run\n          x-dcl-go-type: TargetRun\n          description: Information specifying a Cloud Run deployment target.\n          x-dcl-conflicts:\n          - gke\n          - anthosCluster\n          - multiTarget\n          required:\n          - location\n          properties:\n            location:\n              type: string\n              x-dcl-go-name: Location\n              description: Required. The location where the Cloud Run Service should\n                be located. Format is `projects/{project}/locations/{location}`.\n        targetId:\n          type: string\n          x-dcl-go-name: TargetId\n          readOnly: true\n          description: Output only. Resource id of the `Target`.\n          x-kubernetes-immutable: true\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. Unique identifier of the `Target`.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. Most recent time at which the `Target` was updated.\n          x-kubernetes-immutable: true\n")

// 11819 bytes
// MD5: 1a1bfcd4668c3c13fc74f43377cdc63b
