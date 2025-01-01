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
// gen_go_data -package compute -var YAML_network_firewall_policy blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/network_firewall_policy.yaml

package compute

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/network_firewall_policy.yaml
var YAML_network_firewall_policy = []byte("info:\n  title: Compute/NetworkFirewallPolicy\n  description: The Compute NetworkFirewallPolicy resource\n  x-dcl-struct-name: NetworkFirewallPolicy\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a NetworkFirewallPolicy\n    parameters:\n    - name: networkFirewallPolicy\n      required: true\n      description: A full instance of a NetworkFirewallPolicy\n  apply:\n    description: The function used to apply information about a NetworkFirewallPolicy\n    parameters:\n    - name: networkFirewallPolicy\n      required: true\n      description: A full instance of a NetworkFirewallPolicy\n  delete:\n    description: The function used to delete a NetworkFirewallPolicy\n    parameters:\n    - name: networkFirewallPolicy\n      required: true\n      description: A full instance of a NetworkFirewallPolicy\n  deleteAll:\n    description: The function used to delete all NetworkFirewallPolicy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many NetworkFirewallPolicy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    NetworkFirewallPolicy:\n      title: NetworkFirewallPolicy\n      x-dcl-id: projects/{{project}}/global/firewallPolicies/{{name}}\n      x-dcl-locations:\n      - region\n      - global\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - project\n      properties:\n        creationTimestamp:\n          type: string\n          x-dcl-go-name: CreationTimestamp\n          readOnly: true\n          description: Creation timestamp in RFC3339 text format.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description of this resource. Provide this property\n            when you create the resource.\n        fingerprint:\n          type: string\n          x-dcl-go-name: Fingerprint\n          readOnly: true\n          description: Fingerprint of the resource. This field is used internally\n            during updates of this resource.\n          x-kubernetes-immutable: true\n        id:\n          type: string\n          x-dcl-go-name: Id\n          readOnly: true\n          description: The unique identifier for the resource. This identifier is\n            defined by the server.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location of this resource.\n          x-kubernetes-immutable: true\n          x-dcl-parameter: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: User-provided name of the Network firewall policy. The name\n            should be unique in the project in which the firewall policy is created.\n            The name must be 1-63 characters long, and comply with RFC1035. Specifically,\n            the name must be 1-63 characters long and match the regular expression\n            [a-z]([-a-z0-9]*[a-z0-9])? which means the first character must be a lowercase\n            letter, and all following characters must be a dash, lowercase letter,\n            or digit, except the last character, which cannot be a dash.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-parameter: true\n        region:\n          type: string\n          x-dcl-go-name: Region\n          readOnly: true\n          description: '[Output Only] URL of the region where the regional firewall\n            policy resides. This field is not applicable to global firewall policies.\n            You must specify this field as part of the HTTP request URL. It is not\n            settable as a field in the request body.'\n          x-kubernetes-immutable: true\n        ruleTupleCount:\n          type: integer\n          format: int64\n          x-dcl-go-name: RuleTupleCount\n          readOnly: true\n          description: Total count of all firewall policy rule tuples. A firewall\n            policy can not exceed a set number of tuples.\n          x-kubernetes-immutable: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: Server-defined URL for the resource.\n          x-kubernetes-immutable: true\n        selfLinkWithId:\n          type: string\n          x-dcl-go-name: SelfLinkWithId\n          readOnly: true\n          description: Server-defined URL for this resource with the resource id.\n          x-kubernetes-immutable: true\n")

// 5186 bytes
// MD5: 8d915e95d8b74cedd1d78026811ff342
