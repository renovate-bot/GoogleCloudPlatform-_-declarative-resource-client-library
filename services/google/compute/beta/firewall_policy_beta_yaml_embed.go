// Copyright 2021 Google LLC. All Rights Reserved.
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
// gen_go_data -package beta -var YAML_firewall_policy blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/firewall_policy.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/firewall_policy.yaml
var YAML_firewall_policy = []byte("info:\n  title: Compute/FirewallPolicy\n  description: DCL Specification for the Compute FirewallPolicy resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a FirewallPolicy\n    parameters:\n    - name: FirewallPolicy\n      required: true\n      description: A full instance of a FirewallPolicy\n  apply:\n    description: The function used to apply information about a FirewallPolicy\n    parameters:\n    - name: FirewallPolicy\n      required: true\n      description: A full instance of a FirewallPolicy\n  delete:\n    description: The function used to delete a FirewallPolicy\n    parameters:\n    - name: FirewallPolicy\n      required: true\n      description: A full instance of a FirewallPolicy\n  deleteAll:\n    description: The function used to delete all FirewallPolicy\n    parameters:\n    - name: parent\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many FirewallPolicy\n    parameters:\n    - name: parent\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    FirewallPolicy:\n      title: FirewallPolicy\n      x-dcl-id: locations/global/firewallPolicies/{{name}}\n      x-dcl-locations:\n      - global\n      type: object\n      required:\n      - shortName\n      - parent\n      properties:\n        creationTimestamp:\n          type: string\n          x-dcl-go-name: CreationTimestamp\n          readOnly: true\n          description: Creation timestamp in RFC3339 text format.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description of this resource. Provide this property\n            when you create the resource.\n        fingerprint:\n          type: string\n          x-dcl-go-name: Fingerprint\n          readOnly: true\n          description: Fingerprint of the resource. This field is used internally\n            during updates of this resource.\n          x-kubernetes-immutable: true\n        id:\n          type: string\n          x-dcl-go-name: Id\n          readOnly: true\n          description: The unique identifier for the resource. This identifier is\n            defined by the server.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Name of the resource. It is a numeric ID allocated by GCP which\n            uniquely identifies the Firewall Policy.\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        parent:\n          type: string\n          x-dcl-go-name: Parent\n          description: The parent of the firewall policy.\n          x-kubernetes-immutable: true\n          x-dcl-forward-slash-allowed: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Folder\n            field: name\n            parent: true\n          - resource: Cloudresourcemanager/Organization\n            field: name\n            parent: true\n        ruleTupleCount:\n          type: integer\n          format: int64\n          x-dcl-go-name: RuleTupleCount\n          readOnly: true\n          description: Total count of all firewall policy rule tuples. A firewall\n            policy can not exceed a set number of tuples.\n          x-kubernetes-immutable: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: Server-defined URL for the resource.\n          x-kubernetes-immutable: true\n        selfLinkWithId:\n          type: string\n          x-dcl-go-name: SelfLinkWithId\n          readOnly: true\n          description: Server-defined URL for this resource with the resource id.\n          x-kubernetes-immutable: true\n        shortName:\n          type: string\n          x-dcl-go-name: ShortName\n          description: User-provided name of the Organization firewall policy. The\n            name should be unique in the organization in which the firewall policy\n            is created. The name must be 1-63 characters long, and comply with RFC1035.\n            Specifically, the name must be 1-63 characters long and match the regular\n            expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first character\n            must be a lowercase letter, and all following characters must be a dash,\n            lowercase letter, or digit, except the last character, which cannot be\n            a dash.\n          x-kubernetes-immutable: true\n")

// 4506 bytes
// MD5: 7174827c18d30678176160b585f17183