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
// gen_go_data -package cloudkms -var YAML_ekm_connection blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudkms/ekm_connection.yaml

package cloudkms

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudkms/ekm_connection.yaml
var YAML_ekm_connection = []byte("info:\n  title: Cloudkms/EkmConnection\n  description: The Cloudkms EkmConnection resource\n  x-dcl-struct-name: EkmConnection\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a EkmConnection\n    parameters:\n    - name: ekmConnection\n      required: true\n      description: A full instance of a EkmConnection\n  apply:\n    description: The function used to apply information about a EkmConnection\n    parameters:\n    - name: ekmConnection\n      required: true\n      description: A full instance of a EkmConnection\n  list:\n    description: The function used to list information about many EkmConnection\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    EkmConnection:\n      title: EkmConnection\n      x-dcl-id: projects/{{project}}/locations/{{location}}/ekmConnections/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: true\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - serviceResolvers\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time at which the EkmConnection was created.\n          x-kubernetes-immutable: true\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Optional. This checksum is computed by the server based on\n            the value of other fields, and may be sent on update requests to ensure\n            the client has an up-to-date value before proceeding.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n          x-dcl-parameter: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The resource name for the EkmConnection in the format `projects/*/locations/*/ekmConnections/*`.\n          x-dcl-has-long-form: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-parameter: true\n        serviceResolvers:\n          type: array\n          x-dcl-go-name: ServiceResolvers\n          description: A list of ServiceResolvers where the EKM can be reached. There\n            should be one ServiceResolver per EKM replica. Currently, only a single\n            ServiceResolver is supported.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: EkmConnectionServiceResolvers\n            required:\n            - serviceDirectoryService\n            - hostname\n            - serverCertificates\n            properties:\n              endpointFilter:\n                type: string\n                x-dcl-go-name: EndpointFilter\n                description: Optional. The filter applied to the endpoints of the\n                  resolved service. If no filter is specified, all endpoints will\n                  be considered. An endpoint will be chosen arbitrarily from the filtered\n                  list for each request. For endpoint filter syntax and examples,\n                  see https://cloud.google.com/service-directory/docs/reference/rpc/google.cloud.servicedirectory.v1#resolveservicerequest.\n              hostname:\n                type: string\n                x-dcl-go-name: Hostname\n                description: Required. The hostname of the EKM replica used at TLS\n                  and HTTP layers.\n              serverCertificates:\n                type: array\n                x-dcl-go-name: ServerCertificates\n                description: Required. A list of leaf server certificates used to\n                  authenticate HTTPS connections to the EKM replica. Currently, a\n                  maximum of 10 Certificate is supported.\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: object\n                  x-dcl-go-type: EkmConnectionServiceResolversServerCertificates\n                  required:\n                  - rawDer\n                  properties:\n                    issuer:\n                      type: string\n                      x-dcl-go-name: Issuer\n                      readOnly: true\n                      description: Output only. The issuer distinguished name in RFC\n                        2253 format. Only present if parsed is true.\n                    notAfterTime:\n                      type: string\n                      format: date-time\n                      x-dcl-go-name: NotAfterTime\n                      readOnly: true\n                      description: Output only. The certificate is not valid after\n                        this time. Only present if parsed is true.\n                    notBeforeTime:\n                      type: string\n                      format: date-time\n                      x-dcl-go-name: NotBeforeTime\n                      readOnly: true\n                      description: Output only. The certificate is not valid before\n                        this time. Only present if parsed is true.\n                    parsed:\n                      type: boolean\n                      x-dcl-go-name: Parsed\n                      readOnly: true\n                      description: Output only. True if the certificate was parsed\n                        successfully.\n                    rawDer:\n                      type: string\n                      x-dcl-go-name: RawDer\n                      description: Required. The raw certificate bytes in DER format.\n                    serialNumber:\n                      type: string\n                      x-dcl-go-name: SerialNumber\n                      readOnly: true\n                      description: Output only. The certificate serial number as a\n                        hex string. Only present if parsed is true.\n                    sha256Fingerprint:\n                      type: string\n                      x-dcl-go-name: Sha256Fingerprint\n                      readOnly: true\n                      description: Output only. The SHA-256 certificate fingerprint\n                        as a hex string. Only present if parsed is true.\n                    subject:\n                      type: string\n                      x-dcl-go-name: Subject\n                      readOnly: true\n                      description: Output only. The subject distinguished name in\n                        RFC 2253 format. Only present if parsed is true.\n                    subjectAlternativeDnsNames:\n                      type: array\n                      x-dcl-go-name: SubjectAlternativeDnsNames\n                      readOnly: true\n                      description: Output only. The subject Alternative DNS names.\n                        Only present if parsed is true.\n                      x-dcl-list-type: list\n                      items:\n                        type: string\n                        x-dcl-go-type: string\n              serviceDirectoryService:\n                type: string\n                x-dcl-go-name: ServiceDirectoryService\n                description: Required. The resource name of the Service Directory\n                  service pointing to an EKM replica, in the format `projects/*/locations/*/namespaces/*/services/*`.\n                x-dcl-references:\n                - resource: Servicedirectory/Service\n                  field: name\n")

// 7966 bytes
// MD5: 07add9cfce9c1c7b0c699d1ea86d37ac
