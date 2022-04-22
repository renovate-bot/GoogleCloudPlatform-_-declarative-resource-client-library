// Copyright 2022 Google LLC. All Rights Reserved.
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
// gen_go_data -package beta -var YAML_dataset blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/bigquery/beta/dataset.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/bigquery/beta/dataset.yaml
var YAML_dataset = []byte("info:\n  title: Bigquery/Dataset\n  description: The Bigquery Dataset resource\n  x-dcl-struct-name: Dataset\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Dataset\n    parameters:\n    - name: Dataset\n      required: true\n      description: A full instance of a Dataset\n  apply:\n    description: The function used to apply information about a Dataset\n    parameters:\n    - name: Dataset\n      required: true\n      description: A full instance of a Dataset\n  delete:\n    description: The function used to delete a Dataset\n    parameters:\n    - name: Dataset\n      required: true\n      description: A full instance of a Dataset\n  deleteAll:\n    description: The function used to delete all Dataset\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Dataset\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Dataset:\n      title: Dataset\n      x-dcl-id: projects/{{project}}/datasets/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - project\n      properties:\n        access:\n          type: array\n          x-dcl-go-name: Access\n          description: 'Optional. An array of objects that define dataset access for\n            one or more entities. You can set this property when inserting or updating\n            a dataset in order to control who is allowed to access the data. If unspecified\n            at dataset creation time, BigQuery adds default dataset access for the\n            following entities: access.specialGroup: projectReaders; access.role:\n            READER; access.specialGroup: projectWriters; access.role: WRITER; access.specialGroup:\n            projectOwners; access.role: OWNER; access.userByEmail: ; access.role:\n            OWNER;'\n          x-dcl-send-empty: true\n          x-dcl-list-type: set\n          items:\n            type: object\n            x-dcl-go-type: DatasetAccess\n            required:\n            - role\n            properties:\n              domain:\n                type: string\n                x-dcl-go-name: Domain\n                description: 'A domain to grant access to. Any users signed in with\n                  the domain specified will be granted the specified access. Example:\n                  \"example.com\". Maps to IAM policy member \"domain:DOMAIN\".'\n                x-dcl-conflicts:\n                - userByEmail\n                - groupByEmail\n                - specialGroup\n                - iamMember\n                - view\n                - routine\n              groupByEmail:\n                type: string\n                x-dcl-go-name: GroupByEmail\n                description: An email address of a Google Group to grant access to.\n                  Maps to IAM policy member \"group:GROUP\".\n                x-dcl-conflicts:\n                - userByEmail\n                - domain\n                - specialGroup\n                - iamMember\n                - view\n                - routine\n              iamMember:\n                type: string\n                x-dcl-go-name: IamMember\n                description: Some other type of member that appears in the IAM Policy\n                  but isn't a user, group, domain, or special group.\n                x-dcl-conflicts:\n                - userByEmail\n                - groupByEmail\n                - domain\n                - specialGroup\n                - view\n                - routine\n              role:\n                type: string\n                x-dcl-go-name: Role\n                description: 'Required. An IAM role ID that should be granted to the\n                  user, group, or domain specified in this access entry. The following\n                  legacy mappings will be applied: OWNER <=> roles/bigquery.dataOwner\n                  WRITER <=> roles/bigquery.dataEditor READER <=> roles/bigquery.dataViewer\n                  This field will accept any of the above formats, but will return\n                  only the legacy format. For example, if you set this field to \"roles/bigquery.dataOwner\",\n                  it will be returned back as \"OWNER\".'\n              routine:\n                type: object\n                x-dcl-go-name: Routine\n                x-dcl-go-type: DatasetAccessRoutine\n                description: A routine from a different dataset to grant access to.\n                  Queries executed against that routine will have read access to views/tables/routines\n                  in this dataset. Only UDF is supported for now. The role field is\n                  not required when this field is set. If that routine is updated\n                  by any user, access to the routine needs to be granted again via\n                  an update operation.\n                x-dcl-conflicts:\n                - userByEmail\n                - groupByEmail\n                - domain\n                - specialGroup\n                - iamMember\n                - view\n                required:\n                - projectId\n                - datasetId\n                - routineId\n                properties:\n                  datasetId:\n                    type: string\n                    x-dcl-go-name: DatasetId\n                    description: Required. The ID of the dataset containing this routine.\n                    x-dcl-references:\n                    - resource: Bigquery/Dataset\n                      field: name\n                  projectId:\n                    type: string\n                    x-dcl-go-name: ProjectId\n                    description: Required. The ID of the project containing this routine.\n                    x-dcl-references:\n                    - resource: Cloudresourcemanager/Project\n                      field: name\n                  routineId:\n                    type: string\n                    x-dcl-go-name: RoutineId\n                    description: Required. The ID of the routine. The ID must contain\n                      only letters (a-z, A-Z), numbers (0-9), or underscores (_).\n                      The maximum length is 256 characters.\n                    x-dcl-references:\n                    - resource: Bigquery/Routine\n                      field: name\n              specialGroup:\n                type: string\n                x-dcl-go-name: SpecialGroup\n                description: 'A special group to grant access to. Possible values\n                  include: projectOwners: Owners of the enclosing project. projectReaders:\n                  Readers of the enclosing project. projectWriters: Writers of the\n                  enclosing project. allAuthenticatedUsers: All authenticated BigQuery\n                  users. Maps to similarly-named IAM members.'\n                x-dcl-conflicts:\n                - userByEmail\n                - groupByEmail\n                - domain\n                - iamMember\n                - view\n                - routine\n              userByEmail:\n                type: string\n                x-dcl-go-name: UserByEmail\n                description: 'An email address of a user to grant access to. For example:\n                  fred@example.com. Maps to IAM policy member \"user:EMAIL\" or \"serviceAccount:EMAIL\".'\n                x-dcl-conflicts:\n                - groupByEmail\n                - domain\n                - specialGroup\n                - iamMember\n                - view\n                - routine\n              view:\n                type: object\n                x-dcl-go-name: View\n                x-dcl-go-type: DatasetAccessView\n                description: A view from a different dataset to grant access to. Queries\n                  executed against that view will have read access to views/tables/routines\n                  in this dataset. The role field is not required when this field\n                  is set. If that view is updated by any user, access to the view\n                  needs to be granted again via an update operation.\n                x-dcl-conflicts:\n                - userByEmail\n                - groupByEmail\n                - domain\n                - specialGroup\n                - iamMember\n                - routine\n                required:\n                - projectId\n                - datasetId\n                - tableId\n                properties:\n                  datasetId:\n                    type: string\n                    x-dcl-go-name: DatasetId\n                    description: Required. The ID of the dataset containing this table.\n                    x-dcl-references:\n                    - resource: Bigquery/Dataset\n                      field: name\n                  projectId:\n                    type: string\n                    x-dcl-go-name: ProjectId\n                    description: Required. The ID of the project containing this table.\n                    x-dcl-references:\n                    - resource: Cloudresourcemanager/Project\n                      field: name\n                  tableId:\n                    type: string\n                    x-dcl-go-name: TableId\n                    description: Required. The ID of the table. The ID must contain\n                      only letters (a-z, A-Z), numbers (0-9), or underscores (_).\n                      The maximum length is 1,024 characters. Certain operations allow\n                      suffixing of the table ID with a partition decorator, such as\n                      `sample_table$20190123`.\n                    x-dcl-references:\n                    - resource: Bigquery/Table\n                      field: name\n        creationTime:\n          type: integer\n          format: int64\n          x-dcl-go-name: CreationTime\n          readOnly: true\n          description: Output only. The time when this dataset was created, in milliseconds\n            since the epoch.\n          x-kubernetes-immutable: true\n        defaultEncryptionConfiguration:\n          type: object\n          x-dcl-go-name: DefaultEncryptionConfiguration\n          x-dcl-go-type: DatasetDefaultEncryptionConfiguration\n          description: The default encryption key for all tables in the dataset. Once\n            this property is set, all newly-created partitioned tables in the dataset\n            will have encryption key set to this value, unless table creation request\n            (or query) overrides the key.\n          x-kubernetes-immutable: true\n          properties:\n            kmsKeyName:\n              type: string\n              x-dcl-go-name: KmsKeyName\n              description: Optional. Describes the Cloud KMS encryption key that will\n                be used to protect destination BigQuery table. The BigQuery Service\n                Account associated with your project requires access to this encryption\n                key.\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Cloudkms/CryptoKey\n                field: name\n        defaultPartitionExpirationMs:\n          type: string\n          x-dcl-go-name: DefaultPartitionExpirationMs\n          description: This default partition expiration, expressed in milliseconds.\n            When new time-partitioned tables are created in a dataset where this property\n            is set, the table will inherit this value, propagated as the `TimePartitioning.expirationMs`\n            property on the new table. If you set `TimePartitioning.expirationMs`\n            explicitly when creating a table, the `defaultPartitionExpirationMs` of\n            the containing dataset is ignored. When creating a partitioned table,\n            if `defaultPartitionExpirationMs` is set, the `defaultTableExpirationMs`\n            value is ignored and the table will not inherit a table expiration deadline.\n        defaultTableExpirationMs:\n          type: string\n          x-dcl-go-name: DefaultTableExpirationMs\n          description: Optional. The default lifetime of all tables in the dataset,\n            in milliseconds. The minimum lifetime value is 3600000 milliseconds (one\n            hour). To clear an existing default expiration with a PATCH request, set\n            to 0. Once this property is set, all newly-created tables in the dataset\n            will have an expirationTime property set to the creation time plus the\n            value in this property, and changing the value will only affect new tables,\n            not existing ones. When the expirationTime for a given table is reached,\n            that table will be deleted automatically. If a table's expirationTime\n            is modified or removed before the table expires, or if you provide an\n            explicit expirationTime when creating a table, that value takes precedence\n            over the default expiration time indicated by this property.\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A user-friendly description of the dataset.\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Output only. A hash of the resource.\n          x-kubernetes-immutable: true\n        friendlyName:\n          type: string\n          x-dcl-go-name: FriendlyName\n          description: Optional. A descriptive name for the dataset.\n        id:\n          type: string\n          x-dcl-go-name: Id\n          readOnly: true\n          description: Output only. The fully-qualified unique name of the dataset\n            in the format projectId:datasetId. The dataset name without the project\n            name is given in the datasetId field. When creating a new dataset, leave\n            this field blank, and instead specify the datasetId field.\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: The labels associated with this dataset. You can use these\n            to organize and group your datasets. You can set this property when inserting\n            or updating a dataset. See (/bigquery/docs/creating-managing-labels#creating_and_updating_dataset_labels)\n            for more information.\n        lastModifiedTime:\n          type: integer\n          format: int64\n          x-dcl-go-name: LastModifiedTime\n          readOnly: true\n          description: Output only. The date when this dataset was last modified,\n            in milliseconds since the epoch.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The geographic location where the dataset should reside. See\n            https://cloud.google.com/bigquery/docs/locations for supported locations.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. A unique ID for this dataset, without the project\n            name. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores\n            (_). The maximum length is 1,024 characters.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The ID of the project containing this dataset.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        published:\n          type: boolean\n          x-dcl-go-name: Published\n          description: Whether this dataset is visible to all users in public search.\n            This field can only be set to true if READER access is given to allAuthenticatedUsers\n            in the access list. The default value is false.\n          x-kubernetes-immutable: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: Output only. A URL that can be used to access the resource\n            again. You can use this URL in Get or Update requests to the resource.\n          x-kubernetes-immutable: true\n")

// 16213 bytes
// MD5: c63b2185f10218c780ba731dd0bee202