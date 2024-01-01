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
// gen_go_data -package compute -var YAML_subnetwork blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/subnetwork.yaml

package compute

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/subnetwork.yaml
var YAML_subnetwork = []byte("info:\n  title: Compute/Subnetwork\n  description: The Compute Subnetwork resource\n  x-dcl-struct-name: Subnetwork\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Subnetwork\n    parameters:\n    - name: subnetwork\n      required: true\n      description: A full instance of a Subnetwork\n  apply:\n    description: The function used to apply information about a Subnetwork\n    parameters:\n    - name: subnetwork\n      required: true\n      description: A full instance of a Subnetwork\n  delete:\n    description: The function used to delete a Subnetwork\n    parameters:\n    - name: subnetwork\n      required: true\n      description: A full instance of a Subnetwork\n  deleteAll:\n    description: The function used to delete all Subnetwork\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: region\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Subnetwork\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: region\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Subnetwork:\n      title: Subnetwork\n      x-dcl-id: projects/{{project}}/regions/{{region}}/subnetworks/{{name}}\n      x-dcl-locations:\n      - region\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - ipCidrRange\n      - name\n      - network\n      - region\n      - project\n      properties:\n        creationTimestamp:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreationTimestamp\n          readOnly: true\n          description: Creation timestamp in RFC3339 text format.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: 'An optional description of this resource. Provide this property\n            when you create the resource. This field can be set only at resource creation\n            time. '\n          x-kubernetes-immutable: true\n        enableFlowLogs:\n          type: boolean\n          x-dcl-go-name: EnableFlowLogs\n          description: Whether to enable flow logging for this subnetwork. If this\n            field is not explicitly set, it will not appear in `get` listings. If\n            not set the default behavior is to disable flow logging. This field isn't\n            supported with the `purpose` field set to `INTERNAL_HTTPS_LOAD_BALANCER`.\n          x-kubernetes-immutable: true\n          x-dcl-server-default: true\n        fingerprint:\n          type: string\n          x-dcl-go-name: Fingerprint\n          readOnly: true\n          description: 'Fingerprint of this resource. This field is used internally\n            during updates of this resource. '\n          x-kubernetes-immutable: true\n        gatewayAddress:\n          type: string\n          x-dcl-go-name: GatewayAddress\n          readOnly: true\n          description: 'The gateway address for default routes to reach destination\n            addresses outside this subnetwork. '\n          x-kubernetes-immutable: true\n        ipCidrRange:\n          type: string\n          x-dcl-go-name: IPCidrRange\n          description: 'The range of internal addresses that are owned by this subnetwork.\n            Provide this property when you create the subnetwork. For example, 10.0.0.0/8\n            or 192.168.0.0/16. Ranges must be unique and non-overlapping within a\n            network. Only IPv4 is supported. '\n        logConfig:\n          type: object\n          x-dcl-go-name: LogConfig\n          x-dcl-go-type: SubnetworkLogConfig\n          description: 'Denotes the logging options for the subnetwork flow logs.\n            If logging is enabled logs will be exported to Cloud Logging. '\n          properties:\n            aggregationInterval:\n              type: string\n              x-dcl-go-name: AggregationInterval\n              x-dcl-go-type: SubnetworkLogConfigAggregationIntervalEnum\n              description: 'Can only be specified if VPC flow logging for this subnetwork\n                is enabled. Toggles the aggregation interval for collecting flow logs.\n                Increasing the interval time will reduce the amount of generated flow\n                logs for long lasting connections. Default is an interval of 5 seconds\n                per connection. Possible values are INTERVAL_5_SEC, INTERVAL_30_SEC,\n                INTERVAL_1_MIN, INTERVAL_5_MIN, INTERVAL_10_MIN, INTERVAL_15_MIN '\n              default: INTERVAL_5_SEC\n              enum:\n              - INTERVAL_5_SEC\n              - INTERVAL_30_SEC\n              - INTERVAL_1_MIN\n              - INTERVAL_5_MIN\n              - INTERVAL_10_MIN\n              - INTERVAL_15_MIN\n            flowSampling:\n              type: number\n              format: double\n              x-dcl-go-name: FlowSampling\n              description: 'Can only be specified if VPC flow logging for this subnetwork\n                is enabled. The value of the field must be in [0, 1]. Set the sampling\n                rate of VPC flow logs within the subnetwork where 1.0 means all collected\n                logs are reported and 0.0 means no logs are reported. Default is 0.5\n                which means half of all collected logs are reported. '\n              default: 0.5\n            metadata:\n              type: string\n              x-dcl-go-name: Metadata\n              x-dcl-go-type: SubnetworkLogConfigMetadataEnum\n              description: 'Can only be specified if VPC flow logging for this subnetwork\n                is enabled. Configures whether metadata fields should be added to\n                the reported VPC flow logs. Default is `INCLUDE_ALL_METADATA`.  Possible\n                values: EXCLUDE_ALL_METADATA, INCLUDE_ALL_METADATA'\n              default: INCLUDE_ALL_METADATA\n              enum:\n              - EXCLUDE_ALL_METADATA\n              - INCLUDE_ALL_METADATA\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'The name of the resource, provided by the client when initially\n            creating the resource. The name must be 1-63 characters long, and comply\n            with RFC1035. Specifically, the name must be 1-63 characters long and\n            match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means\n            the first character must be a lowercase letter, and all following characters\n            must be a dash, lowercase letter, or digit, except the last character,\n            which cannot be a dash. '\n          x-kubernetes-immutable: true\n        network:\n          type: string\n          x-dcl-go-name: Network\n          description: 'The network this subnet belongs to. Only networks that are\n            in the distributed mode can have subnetworks. '\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Compute/Network\n            field: selfLink\n        privateIPGoogleAccess:\n          type: boolean\n          x-dcl-go-name: PrivateIPGoogleAccess\n          description: 'When enabled, VMs in this subnetwork without external IP addresses\n            can access Google APIs and services by using Private Google Access. '\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project id of the resource.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        purpose:\n          type: string\n          x-dcl-go-name: Purpose\n          x-dcl-go-type: SubnetworkPurposeEnum\n          description: 'The purpose of the resource. This field can be either PRIVATE\n            or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER\n            is a user-created subnetwork that is reserved for Internal HTTP(S) Load\n            Balancing. If unspecified, the purpose defaults to PRIVATE.  If set to\n            INTERNAL_HTTPS_LOAD_BALANCER you must also set the role. '\n          x-kubernetes-immutable: true\n          x-dcl-server-default: true\n          enum:\n          - INTERNAL_HTTPS_LOAD_BALANCER\n          - PRIVATE\n          - AGGREGATE\n          - PRIVATE_SERVICE_CONNECT\n          - CLOUD_EXTENSION\n          - PRIVATE_NAT\n        region:\n          type: string\n          x-dcl-go-name: Region\n          description: 'The GCP region for this subnetwork. '\n          x-kubernetes-immutable: true\n        role:\n          type: string\n          x-dcl-go-name: Role\n          x-dcl-go-type: SubnetworkRoleEnum\n          description: 'The role of subnetwork. Currenly, this field is only used\n            when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE\n            or BACKUP. An ACTIVE subnetwork is one that is currently being used for\n            Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready\n            to be promoted to ACTIVE or is currently draining. '\n          enum:\n          - ACTIVE\n          - BACKUP\n        secondaryIPRanges:\n          type: array\n          x-dcl-go-name: SecondaryIPRanges\n          description: 'An array of configurations for secondary IP ranges for VM\n            instances contained in this subnetwork. The primary IP of such VM must\n            belong to the primary ipCidrRange of the subnetwork. The alias IPs may\n            belong to either primary or secondary ranges. This field uses attr-as-block\n            mode to avoid breaking users during the 0.12 upgrade. See [the Attr-as-Block\n            page](https://www.terraform.io/docs/configuration/attr-as-blocks.html)\n            for more details. '\n          x-dcl-server-default: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: SubnetworkSecondaryIPRanges\n            required:\n            - rangeName\n            - ipCidrRange\n            properties:\n              ipCidrRange:\n                type: string\n                x-dcl-go-name: IPCidrRange\n                description: 'The range of IP addresses belonging to this subnetwork\n                  secondary range. Provide this property when you create the subnetwork.\n                  Ranges must be unique and non-overlapping with all primary and secondary\n                  IP ranges within a network. Only IPv4 is supported. '\n              rangeName:\n                type: string\n                x-dcl-go-name: RangeName\n                description: 'The name associated with this subnetwork secondary range,\n                  used when adding an alias IP range to a VM instance. The name must\n                  be 1-63 characters long, and comply with RFC1035. The name must\n                  be unique within the subnetwork. '\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: '[Output Only] Server-defined URL for the resource.'\n          x-kubernetes-immutable: true\n")

// 11259 bytes
// MD5: b0e87e9d881036b91acaf94ee532acef
