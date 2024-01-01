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
// gen_go_data -package beta -var YAML_topic blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/pubsub/beta/topic.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/pubsub/beta/topic.yaml
var YAML_topic = []byte("info:\n  title: Pubsub/Topic\n  description: The Pubsub Topic resource\n  x-dcl-struct-name: Topic\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Topic\n    parameters:\n    - name: topic\n      required: true\n      description: A full instance of a Topic\n  apply:\n    description: The function used to apply information about a Topic\n    parameters:\n    - name: topic\n      required: true\n      description: A full instance of a Topic\n  delete:\n    description: The function used to delete a Topic\n    parameters:\n    - name: topic\n      required: true\n      description: A full instance of a Topic\n  deleteAll:\n    description: The function used to delete all Topic\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Topic\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Topic:\n      title: Topic\n      x-dcl-id: projects/{{project}}/topics/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - project\n      properties:\n        kmsKeyName:\n          type: string\n          x-dcl-go-name: KmsKeyName\n          description: 'The resource name of the Cloud KMS CryptoKey to be used to\n            protect access to messages published on this topic. Your project''s Pub/Sub\n            service account (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`)\n            must have `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.  The\n            expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*` '\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: 'A set of key/value label pairs to assign to this Topic. '\n        messageStoragePolicy:\n          type: object\n          x-dcl-go-name: MessageStoragePolicy\n          x-dcl-go-type: TopicMessageStoragePolicy\n          description: 'Policy constraining the set of Google Cloud Platform regions\n            where messages published to the topic may be stored. If not present, then\n            no constraints are in effect. '\n          x-dcl-server-default: true\n          properties:\n            allowedPersistenceRegions:\n              type: array\n              x-dcl-go-name: AllowedPersistenceRegions\n              description: 'A list of IDs of GCP regions where messages that are published\n                to the topic may be persisted in storage. Messages published by publishers\n                running in non-allowed GCP regions (or running outside of GCP altogether)\n                will be routed for storage in one of the allowed regions. An empty\n                list means that no regions are allowed, and is not a valid configuration. '\n              x-dcl-send-empty: true\n              x-dcl-list-type: set\n              items:\n                type: string\n                x-dcl-go-type: string\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Name of the topic.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project id of the resource.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n")

// 3700 bytes
// MD5: 84e85a406f64e0c43cdb983f2c1596a5
