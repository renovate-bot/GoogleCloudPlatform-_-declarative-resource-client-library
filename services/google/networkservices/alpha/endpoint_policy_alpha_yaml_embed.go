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
// gen_go_data -package alpha -var YAML_endpoint_policy blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/alpha/endpoint_policy.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/alpha/endpoint_policy.yaml
var YAML_endpoint_policy = []byte("info:\n  title: NetworkServices/EndpointPolicy\n  description: The NetworkServices EndpointPolicy resource\n  x-dcl-struct-name: EndpointPolicy\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a EndpointPolicy\n    parameters:\n    - name: endpointPolicy\n      required: true\n      description: A full instance of a EndpointPolicy\n  apply:\n    description: The function used to apply information about a EndpointPolicy\n    parameters:\n    - name: endpointPolicy\n      required: true\n      description: A full instance of a EndpointPolicy\n  delete:\n    description: The function used to delete a EndpointPolicy\n    parameters:\n    - name: endpointPolicy\n      required: true\n      description: A full instance of a EndpointPolicy\n  deleteAll:\n    description: The function used to delete all EndpointPolicy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many EndpointPolicy\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    EndpointPolicy:\n      title: EndpointPolicy\n      x-dcl-id: projects/{{project}}/locations/{{location}}/endpointPolicies/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - type\n      - endpointMatcher\n      - project\n      - location\n      properties:\n        authorizationPolicy:\n          type: string\n          x-dcl-go-name: AuthorizationPolicy\n          description: Optional. This field specifies the URL of AuthorizationPolicy\n            resource that applies authorization policies to the inbound traffic at\n            the matched endpoints. Refer to Authorization. If this field is not specified,\n            authorization is disabled(no authz checks) for this endpoint.\n          x-dcl-references:\n          - resource: Networksecurity/AuthorizationPolicy\n            field: name\n        clientTlsPolicy:\n          type: string\n          x-dcl-go-name: ClientTlsPolicy\n          description: Optional. A URL referring to a ClientTlsPolicy resource. ClientTlsPolicy\n            can be set to specify the authentication for traffic from the proxy to\n            the actual endpoints. More specifically, it is applied to the outgoing\n            traffic from the proxy to the endpoint. This is typically used for sidecar\n            model where the proxy identifies itself as endpoint to the control plane,\n            with the connection between sidecar and endpoint requiring authentication.\n            If this field is not set, authentication is disabled(open). Applicable\n            only when EndpointPolicyType is SIDECAR_PROXY.\n          x-dcl-references:\n          - resource: Networksecurity/ClientTlsPolicy\n            field: name\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A free-text description of the resource. Max length\n            1024 characters.\n        endpointMatcher:\n          type: object\n          x-dcl-go-name: EndpointMatcher\n          x-dcl-go-type: EndpointPolicyEndpointMatcher\n          description: Required. A matcher that selects endpoints to which the policies\n            should be applied.\n          properties:\n            metadataLabelMatcher:\n              type: object\n              x-dcl-go-name: MetadataLabelMatcher\n              x-dcl-go-type: EndpointPolicyEndpointMatcherMetadataLabelMatcher\n              description: The matcher is based on node metadata presented by xDS\n                clients.\n              properties:\n                metadataLabelMatchCriteria:\n                  type: string\n                  x-dcl-go-name: MetadataLabelMatchCriteria\n                  x-dcl-go-type: EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum\n                  description: 'Specifies how matching should be done. Supported values\n                    are: MATCH_ANY: At least one of the Labels specified in the matcher\n                    should match the metadata presented by xDS client. MATCH_ALL:\n                    The metadata presented by the xDS client should contain all of\n                    the labels specified here. The selection is determined based on\n                    the best match. For example, suppose there are three EndpointPolicy\n                    resources P1, P2 and P3 and if P1 has a the matcher as MATCH_ANY\n                    , P2 has MATCH_ALL , and P3 has MATCH_ALL . If a client with label\n                    connects, the config from P1 will be selected. If a client with\n                    label connects, the config from P2 will be selected. If a client\n                    with label connects, the config from P3 will be selected. If there\n                    is more than one best match, (for example, if a config P4 with\n                    selector exists and if a client with label connects), an error\n                    will be thrown. Possible values: METADATA_LABEL_MATCH_CRITERIA_UNSPECIFIED,\n                    MATCH_ANY, MATCH_ALL'\n                  enum:\n                  - METADATA_LABEL_MATCH_CRITERIA_UNSPECIFIED\n                  - MATCH_ANY\n                  - MATCH_ALL\n                metadataLabels:\n                  type: array\n                  x-dcl-go-name: MetadataLabels\n                  description: The list of label value pairs that must match labels\n                    in the provided metadata based on filterMatchCriteria This list\n                    can have at most 64 entries. The list can be empty if the match\n                    criteria is MATCH_ANY, to specify a wildcard match (i.e this matches\n                    any client).\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: object\n                    x-dcl-go-type: EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels\n                    required:\n                    - labelName\n                    - labelValue\n                    properties:\n                      labelName:\n                        type: string\n                        x-dcl-go-name: LabelName\n                        description: Required. Label name presented as key in xDS\n                          Node Metadata.\n                      labelValue:\n                        type: string\n                        x-dcl-go-name: LabelValue\n                        description: Required. Label value presented as value corresponding\n                          to the above key, in xDS Node Metadata.\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. Set of label tags associated with the EndpointPolicy\n            resource.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n          x-dcl-parameter: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. Name of the EndpointPolicy resource.\n          x-kubernetes-immutable: true\n          x-dcl-has-long-form: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-parameter: true\n        serverTlsPolicy:\n          type: string\n          x-dcl-go-name: ServerTlsPolicy\n          description: Optional. A URL referring to ServerTlsPolicy resource. ServerTlsPolicy\n            is used to determine the authentication policy to be applied to terminate\n            the inbound traffic at the identified backends. If this field is not set,\n            authentication is disabled(open) for this endpoint.\n          x-dcl-references:\n          - resource: Networksecurity/ServerTlsPolicy\n            field: name\n        trafficPortSelector:\n          type: object\n          x-dcl-go-name: TrafficPortSelector\n          x-dcl-go-type: EndpointPolicyTrafficPortSelector\n          description: Optional. Port selector for the (matched) endpoints. If no\n            port selector is provided, the matched config is applied to all ports.\n          properties:\n            ports:\n              type: array\n              x-dcl-go-name: Ports\n              description: Optional. A list of ports. Can be port numbers or port\n                range (example, specifies all ports from 80 to 90, including 80 and\n                90) or named ports or * to specify all ports. If the list is empty,\n                all ports are selected.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n        type:\n          type: string\n          x-dcl-go-name: Type\n          x-dcl-go-type: EndpointPolicyTypeEnum\n          description: 'Required. The type of endpoint config. This is primarily used\n            to validate the configuration. Possible values: ENDPOINT_CONFIG_SELECTOR_TYPE_UNSPECIFIED,\n            SIDECAR_PROXY, GRPC_SERVER'\n          enum:\n          - ENDPOINT_CONFIG_SELECTOR_TYPE_UNSPECIFIED\n          - SIDECAR_PROXY\n          - GRPC_SERVER\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was updated.\n          x-kubernetes-immutable: true\n")

// 10326 bytes
// MD5: 58a3e2b7d939af840810d9d34afc0ec2
