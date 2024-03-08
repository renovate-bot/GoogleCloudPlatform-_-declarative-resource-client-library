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
// gen_go_data -package containeraws -var YAML_cluster blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/containeraws/cluster.yaml

package containeraws

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/containeraws/cluster.yaml
var YAML_cluster = []byte("info:\n  title: ContainerAws/Cluster\n  description: An Anthos cluster running on AWS.\n  x-dcl-struct-name: Cluster\n  x-dcl-has-iam: false\n  x-dcl-ref:\n    text: API reference\n    url: https://cloud.google.com/anthos/clusters/docs/multi-cloud/reference/rest/v1/projects.locations.awsClusters\n  x-dcl-guides:\n  - text: Multicloud overview\n    url: https://cloud.google.com/anthos/clusters/docs/multi-cloud\npaths:\n  get:\n    description: The function used to get information about a Cluster\n    parameters:\n    - name: cluster\n      required: true\n      description: A full instance of a Cluster\n  apply:\n    description: The function used to apply information about a Cluster\n    parameters:\n    - name: cluster\n      required: true\n      description: A full instance of a Cluster\n  delete:\n    description: The function used to delete a Cluster\n    parameters:\n    - name: cluster\n      required: true\n      description: A full instance of a Cluster\n  deleteAll:\n    description: The function used to delete all Cluster\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Cluster\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Cluster:\n      title: Cluster\n      x-dcl-id: projects/{{project}}/locations/{{location}}/awsClusters/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - networking\n      - awsRegion\n      - controlPlane\n      - authorization\n      - project\n      - location\n      - fleet\n      properties:\n        annotations:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Annotations\n          description: 'Optional. Annotations on the cluster. This field has the same\n            restrictions as Kubernetes annotations. The total size of all keys and\n            values combined is limited to 256k. Key can have 2 segments: prefix (optional)\n            and name (required), separated by a slash (/). Prefix must be a DNS subdomain.\n            Name must be 63 characters or less, begin and end with alphanumerics,\n            with dashes (-), underscores (_), dots (.), and alphanumerics between.'\n        authorization:\n          type: object\n          x-dcl-go-name: Authorization\n          x-dcl-go-type: ClusterAuthorization\n          description: Configuration related to the cluster RBAC settings.\n          required:\n          - adminUsers\n          properties:\n            adminGroups:\n              type: array\n              x-dcl-go-name: AdminGroups\n              description: Groups of users that can perform operations as a cluster\n                admin. A managed ClusterRoleBinding will be created to grant the `cluster-admin`\n                ClusterRole to the groups. Up to ten admin groups can be provided.\n                For more info on RBAC, see https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: ClusterAuthorizationAdminGroups\n                required:\n                - group\n                properties:\n                  group:\n                    type: string\n                    x-dcl-go-name: Group\n                    description: The name of the group, e.g. `my-group@domain.com`.\n            adminUsers:\n              type: array\n              x-dcl-go-name: AdminUsers\n              description: Users to perform operations as a cluster admin. A managed\n                ClusterRoleBinding will be created to grant the `cluster-admin` ClusterRole\n                to the users. Up to ten admin users can be provided. For more info\n                on RBAC, see https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: ClusterAuthorizationAdminUsers\n                required:\n                - username\n                properties:\n                  username:\n                    type: string\n                    x-dcl-go-name: Username\n                    description: The name of the user, e.g. `my-gcp-id@gmail.com`.\n        awsRegion:\n          type: string\n          x-dcl-go-name: AwsRegion\n          description: The AWS region where the cluster runs. Each Google Cloud region\n            supports a subset of nearby AWS regions. You can call to list all supported\n            AWS regions within a given Google Cloud region.\n          x-kubernetes-immutable: true\n        binaryAuthorization:\n          type: object\n          x-dcl-go-name: BinaryAuthorization\n          x-dcl-go-type: ClusterBinaryAuthorization\n          description: Configuration options for the Binary Authorization feature.\n          x-dcl-server-default: true\n          properties:\n            evaluationMode:\n              type: string\n              x-dcl-go-name: EvaluationMode\n              x-dcl-go-type: ClusterBinaryAuthorizationEvaluationModeEnum\n              description: 'Mode of operation for Binary Authorization policy evaluation.\n                Possible values: DISABLED, PROJECT_SINGLETON_POLICY_ENFORCE'\n              x-dcl-server-default: true\n              enum:\n              - DISABLED\n              - PROJECT_SINGLETON_POLICY_ENFORCE\n        controlPlane:\n          type: object\n          x-dcl-go-name: ControlPlane\n          x-dcl-go-type: ClusterControlPlane\n          description: Configuration related to the cluster control plane.\n          required:\n          - version\n          - subnetIds\n          - configEncryption\n          - iamInstanceProfile\n          - databaseEncryption\n          - awsServicesAuthentication\n          properties:\n            awsServicesAuthentication:\n              type: object\n              x-dcl-go-name: AwsServicesAuthentication\n              x-dcl-go-type: ClusterControlPlaneAwsServicesAuthentication\n              description: Authentication configuration for management of AWS resources.\n              required:\n              - roleArn\n              properties:\n                roleArn:\n                  type: string\n                  x-dcl-go-name: RoleArn\n                  description: The Amazon Resource Name (ARN) of the role that the\n                    Anthos Multi-Cloud API will assume when managing AWS resources\n                    on your account.\n                roleSessionName:\n                  type: string\n                  x-dcl-go-name: RoleSessionName\n                  description: Optional. An identifier for the assumed role session.\n                    When unspecified, it defaults to `multicloud-service-agent`.\n                  x-dcl-server-default: true\n            configEncryption:\n              type: object\n              x-dcl-go-name: ConfigEncryption\n              x-dcl-go-type: ClusterControlPlaneConfigEncryption\n              description: The ARN of the AWS KMS key used to encrypt cluster configuration.\n              required:\n              - kmsKeyArn\n              properties:\n                kmsKeyArn:\n                  type: string\n                  x-dcl-go-name: KmsKeyArn\n                  description: The ARN of the AWS KMS key used to encrypt cluster\n                    configuration.\n            databaseEncryption:\n              type: object\n              x-dcl-go-name: DatabaseEncryption\n              x-dcl-go-type: ClusterControlPlaneDatabaseEncryption\n              description: The ARN of the AWS KMS key used to encrypt cluster secrets.\n              x-kubernetes-immutable: true\n              required:\n              - kmsKeyArn\n              properties:\n                kmsKeyArn:\n                  type: string\n                  x-dcl-go-name: KmsKeyArn\n                  description: The ARN of the AWS KMS key used to encrypt cluster\n                    secrets.\n                  x-kubernetes-immutable: true\n            iamInstanceProfile:\n              type: string\n              x-dcl-go-name: IamInstanceProfile\n              description: The name of the AWS IAM instance pofile to assign to each\n                control plane replica.\n            instanceType:\n              type: string\n              x-dcl-go-name: InstanceType\n              description: Optional. The AWS instance type. When unspecified, it defaults\n                to `m5.large`.\n              x-dcl-server-default: true\n            mainVolume:\n              type: object\n              x-dcl-go-name: MainVolume\n              x-dcl-go-type: ClusterControlPlaneMainVolume\n              description: Optional. Configuration related to the main volume provisioned\n                for each control plane replica. The main volume is in charge of storing\n                all of the cluster's etcd state. Volumes will be provisioned in the\n                availability zone associated with the corresponding subnet. When unspecified,\n                it defaults to 8 GiB with the GP2 volume type.\n              x-kubernetes-immutable: true\n              x-dcl-server-default: true\n              properties:\n                iops:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: Iops\n                  description: Optional. The number of I/O operations per second (IOPS)\n                    to provision for GP3 volume.\n                  x-kubernetes-immutable: true\n                  x-dcl-server-default: true\n                kmsKeyArn:\n                  type: string\n                  x-dcl-go-name: KmsKeyArn\n                  description: Optional. The Amazon Resource Name (ARN) of the Customer\n                    Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified,\n                    the default Amazon managed key associated to the AWS region where\n                    this cluster runs will be used.\n                  x-kubernetes-immutable: true\n                sizeGib:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: SizeGib\n                  description: Optional. The size of the volume, in GiBs. When unspecified,\n                    a default value is provided. See the specific reference in the\n                    parent resource.\n                  x-kubernetes-immutable: true\n                  x-dcl-server-default: true\n                throughput:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: Throughput\n                  description: Optional. The throughput to provision for the volume,\n                    in MiB/s. Only valid if the volume type is GP3. If volume type\n                    is gp3 and throughput is not specified, the throughput will defaults\n                    to 125.\n                  x-kubernetes-immutable: true\n                  x-dcl-server-default: true\n                volumeType:\n                  type: string\n                  x-dcl-go-name: VolumeType\n                  x-dcl-go-type: ClusterControlPlaneMainVolumeVolumeTypeEnum\n                  description: 'Optional. Type of the EBS volume. When unspecified,\n                    it defaults to GP2 volume. Possible values: VOLUME_TYPE_UNSPECIFIED,\n                    GP2, GP3'\n                  x-kubernetes-immutable: true\n                  x-dcl-server-default: true\n                  enum:\n                  - VOLUME_TYPE_UNSPECIFIED\n                  - GP2\n                  - GP3\n            proxyConfig:\n              type: object\n              x-dcl-go-name: ProxyConfig\n              x-dcl-go-type: ClusterControlPlaneProxyConfig\n              description: Proxy configuration for outbound HTTP(S) traffic.\n              required:\n              - secretArn\n              - secretVersion\n              properties:\n                secretArn:\n                  type: string\n                  x-dcl-go-name: SecretArn\n                  description: The ARN of the AWS Secret Manager secret that contains\n                    the HTTP(S) proxy configuration.\n                secretVersion:\n                  type: string\n                  x-dcl-go-name: SecretVersion\n                  description: The version string of the AWS Secret Manager secret\n                    that contains the HTTP(S) proxy configuration.\n            rootVolume:\n              type: object\n              x-dcl-go-name: RootVolume\n              x-dcl-go-type: ClusterControlPlaneRootVolume\n              description: Optional. Configuration related to the root volume provisioned\n                for each control plane replica. Volumes will be provisioned in the\n                availability zone associated with the corresponding subnet. When unspecified,\n                it defaults to 32 GiB with the GP2 volume type.\n              x-dcl-server-default: true\n              properties:\n                iops:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: Iops\n                  description: Optional. The number of I/O operations per second (IOPS)\n                    to provision for GP3 volume.\n                  x-dcl-server-default: true\n                kmsKeyArn:\n                  type: string\n                  x-dcl-go-name: KmsKeyArn\n                  description: Optional. The Amazon Resource Name (ARN) of the Customer\n                    Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified,\n                    the default Amazon managed key associated to the AWS region where\n                    this cluster runs will be used.\n                sizeGib:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: SizeGib\n                  description: Optional. The size of the volume, in GiBs. When unspecified,\n                    a default value is provided. See the specific reference in the\n                    parent resource.\n                  x-dcl-server-default: true\n                throughput:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: Throughput\n                  description: Optional. The throughput to provision for the volume,\n                    in MiB/s. Only valid if the volume type is GP3. If volume type\n                    is gp3 and throughput is not specified, the throughput will defaults\n                    to 125.\n                  x-dcl-server-default: true\n                volumeType:\n                  type: string\n                  x-dcl-go-name: VolumeType\n                  x-dcl-go-type: ClusterControlPlaneRootVolumeVolumeTypeEnum\n                  description: 'Optional. Type of the EBS volume. When unspecified,\n                    it defaults to GP2 volume. Possible values: VOLUME_TYPE_UNSPECIFIED,\n                    GP2, GP3'\n                  x-dcl-server-default: true\n                  enum:\n                  - VOLUME_TYPE_UNSPECIFIED\n                  - GP2\n                  - GP3\n            securityGroupIds:\n              type: array\n              x-dcl-go-name: SecurityGroupIds\n              description: Optional. The IDs of additional security groups to add\n                to control plane replicas. The Anthos Multi-Cloud API will automatically\n                create and manage security groups with the minimum rules needed for\n                a functioning cluster.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            sshConfig:\n              type: object\n              x-dcl-go-name: SshConfig\n              x-dcl-go-type: ClusterControlPlaneSshConfig\n              description: Optional. SSH configuration for how to access the underlying\n                control plane machines.\n              required:\n              - ec2KeyPair\n              properties:\n                ec2KeyPair:\n                  type: string\n                  x-dcl-go-name: Ec2KeyPair\n                  description: The name of the EC2 key pair used to login into cluster\n                    machines.\n            subnetIds:\n              type: array\n              x-dcl-go-name: SubnetIds\n              description: The list of subnets where control plane replicas will run.\n                A replica will be provisioned on each subnet and up to three values\n                can be provided. Each subnet must be in a different AWS Availability\n                Zone (AZ).\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            tags:\n              type: object\n              additionalProperties:\n                type: string\n              x-dcl-go-name: Tags\n              description: Optional. A set of AWS resource tags to propagate to all\n                underlying managed AWS resources. Specify at most 50 pairs containing\n                alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127\n                Unicode characters. Values can be up to 255 Unicode characters.\n            version:\n              type: string\n              x-dcl-go-name: Version\n              description: The Kubernetes version to run on control plane replicas\n                (e.g. `1.19.10-gke.1000`). You can list all supported versions on\n                a given Google Cloud region by calling .\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time at which this cluster was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A human readable description of this cluster. Cannot\n            be longer than 255 UTF-8 encoded bytes.\n        endpoint:\n          type: string\n          x-dcl-go-name: Endpoint\n          readOnly: true\n          description: Output only. The endpoint of the cluster's API server.\n          x-kubernetes-immutable: true\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Allows clients to perform consistent read-modify-writes through\n            optimistic concurrency control. May be sent on update and delete requests\n            to ensure the client has an up-to-date value before proceeding.\n          x-kubernetes-immutable: true\n        fleet:\n          type: object\n          x-dcl-go-name: Fleet\n          x-dcl-go-type: ClusterFleet\n          description: Fleet configuration.\n          x-kubernetes-immutable: true\n          required:\n          - project\n          properties:\n            membership:\n              type: string\n              x-dcl-go-name: Membership\n              readOnly: true\n              description: The name of the managed Hub Membership resource associated\n                to this cluster. Membership names are formatted as projects/<project-number>/locations/global/membership/<cluster-id>.\n              x-kubernetes-immutable: true\n            project:\n              type: string\n              x-dcl-go-name: Project\n              description: The number of the Fleet host project where this cluster\n                will be registered.\n              x-kubernetes-immutable: true\n              x-dcl-references:\n              - resource: Cloudresourcemanager/Project\n                field: name\n                parent: true\n              x-dcl-has-long-form: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n          x-dcl-parameter: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of this resource.\n          x-kubernetes-immutable: true\n          x-dcl-has-long-form: true\n        networking:\n          type: object\n          x-dcl-go-name: Networking\n          x-dcl-go-type: ClusterNetworking\n          description: Cluster-wide networking configuration.\n          required:\n          - vpcId\n          - podAddressCidrBlocks\n          - serviceAddressCidrBlocks\n          properties:\n            perNodePoolSgRulesDisabled:\n              type: boolean\n              x-dcl-go-name: PerNodePoolSgRulesDisabled\n              description: Disable the per node pool subnet security group rules on\n                the control plane security group. When set to true, you must also\n                provide one or more security groups that ensure node pools are able\n                to send requests to the control plane on TCP/443 and TCP/8132. Failure\n                to do so may result in unavailable node pools.\n            podAddressCidrBlocks:\n              type: array\n              x-dcl-go-name: PodAddressCidrBlocks\n              description: All pods in the cluster are assigned an RFC1918 IPv4 address\n                from these ranges. Only a single range is supported. This field cannot\n                be changed after creation.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            serviceAddressCidrBlocks:\n              type: array\n              x-dcl-go-name: ServiceAddressCidrBlocks\n              description: All services in the cluster are assigned an RFC1918 IPv4\n                address from these ranges. Only a single range is supported. This\n                field cannot be changed after creation.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            vpcId:\n              type: string\n              x-dcl-go-name: VPCId\n              description: The VPC associated with the cluster. All component clusters\n                (i.e. control plane and node pools) run on a single VPC. This field\n                cannot be changed after creation.\n              x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-parameter: true\n        reconciling:\n          type: boolean\n          x-dcl-go-name: Reconciling\n          readOnly: true\n          description: Output only. If set, there are currently changes in flight\n            to the cluster.\n          x-kubernetes-immutable: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: ClusterStateEnum\n          readOnly: true\n          description: 'Output only. The current state of the cluster. Possible values:\n            STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR,\n            DEGRADED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - PROVISIONING\n          - RUNNING\n          - RECONCILING\n          - STOPPING\n          - ERROR\n          - DEGRADED\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. A globally unique identifier for the cluster.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time at which this cluster was last updated.\n          x-kubernetes-immutable: true\n        workloadIdentityConfig:\n          type: object\n          x-dcl-go-name: WorkloadIdentityConfig\n          x-dcl-go-type: ClusterWorkloadIdentityConfig\n          readOnly: true\n          description: Output only. Workload Identity settings.\n          x-kubernetes-immutable: true\n          properties:\n            identityProvider:\n              type: string\n              x-dcl-go-name: IdentityProvider\n              description: The ID of the OIDC Identity Provider (IdP) associated to\n                the Workload Identity Pool.\n              x-kubernetes-immutable: true\n            issuerUri:\n              type: string\n              x-dcl-go-name: IssuerUri\n              description: The OIDC issuer URL for this cluster.\n              x-kubernetes-immutable: true\n            workloadPool:\n              type: string\n              x-dcl-go-name: WorkloadPool\n              description: The Workload Identity Pool associated to the cluster.\n              x-kubernetes-immutable: true\n")

// 25309 bytes
// MD5: ff832b576e7405f22f533af7cba5000b
