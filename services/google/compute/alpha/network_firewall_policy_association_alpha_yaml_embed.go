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
// gen_go_data -package alpha -var YAML_network_firewall_policy_association blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/alpha/network_firewall_policy_association.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/alpha/network_firewall_policy_association.yaml
var YAML_network_firewall_policy_association = []byte("info:\n  title: Compute/NetworkFirewallPolicyAssociation\n  description: The Compute NetworkFirewallPolicyAssociation resource\n  x-dcl-struct-name: NetworkFirewallPolicyAssociation\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a NetworkFirewallPolicyAssociation\n    parameters:\n    - name: networkFirewallPolicyAssociation\n      required: true\n      description: A full instance of a NetworkFirewallPolicyAssociation\n  apply:\n    description: The function used to apply information about a NetworkFirewallPolicyAssociation\n    parameters:\n    - name: networkFirewallPolicyAssociation\n      required: true\n      description: A full instance of a NetworkFirewallPolicyAssociation\n  delete:\n    description: The function used to delete a NetworkFirewallPolicyAssociation\n    parameters:\n    - name: networkFirewallPolicyAssociation\n      required: true\n      description: A full instance of a NetworkFirewallPolicyAssociation\n  deleteAll:\n    description: The function used to delete all NetworkFirewallPolicyAssociation\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: firewallPolicy\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many NetworkFirewallPolicyAssociation\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: firewallPolicy\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    NetworkFirewallPolicyAssociation:\n      title: NetworkFirewallPolicyAssociation\n      x-dcl-id: projects/{{project}}/global/firewallPolicies/{{firewall_policy}}/getAssociation?name={{name}}\n      x-dcl-locations:\n      - global\n      - region\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - attachmentTarget\n      - firewallPolicy\n      - project\n      properties:\n        attachmentTarget:\n          type: string\n          x-dcl-go-name: AttachmentTarget\n          description: The target that the firewall policy is attached to.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Compute/Network\n            field: name\n        firewallPolicy:\n          type: string\n          x-dcl-go-name: FirewallPolicy\n          description: The firewall policy ID of the association.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Compute/FirewallPolicy\n            field: name\n            parent: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location of this resource.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name for an association.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        shortName:\n          type: string\n          x-dcl-go-name: ShortName\n          readOnly: true\n          description: The short name of the firewall policy of the association.\n          x-kubernetes-immutable: true\n")

// 3708 bytes
// MD5: 9a99fc07bea3e88db875d4291b7214cb
