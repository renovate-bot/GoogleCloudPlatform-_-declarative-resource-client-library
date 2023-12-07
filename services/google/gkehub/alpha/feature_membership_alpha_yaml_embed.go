// Copyright 2023 Google LLC. All Rights Reserved.
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
// gen_go_data -package alpha -var YAML_feature_membership blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/gkehub/alpha/feature_membership.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/gkehub/alpha/feature_membership.yaml
var YAML_feature_membership = []byte("info:\n  title: GkeHub/FeatureMembership\n  description: The GkeHub FeatureMembership resource\n  x-dcl-struct-name: FeatureMembership\n  x-dcl-has-iam: false\n  x-dcl-mutex: '{{project}}/{{location}}/{{feature}}'\npaths:\n  get:\n    description: The function used to get information about a FeatureMembership\n    parameters:\n    - name: featureMembership\n      required: true\n      description: A full instance of a FeatureMembership\n  apply:\n    description: The function used to apply information about a FeatureMembership\n    parameters:\n    - name: featureMembership\n      required: true\n      description: A full instance of a FeatureMembership\n  delete:\n    description: The function used to delete a FeatureMembership\n    parameters:\n    - name: featureMembership\n      required: true\n      description: A full instance of a FeatureMembership\n  deleteAll:\n    description: The function used to delete all FeatureMembership\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: feature\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many FeatureMembership\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: feature\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    FeatureMembership:\n      title: FeatureMembership\n      x-dcl-id: projects/{{project}}/locations/{{location}}/features/{{feature}}/memberships/{{membership}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - project\n      - location\n      - feature\n      - membership\n      properties:\n        configmanagement:\n          type: object\n          x-dcl-go-name: Configmanagement\n          x-dcl-go-type: FeatureMembershipConfigmanagement\n          description: Config Management-specific spec.\n          properties:\n            binauthz:\n              type: object\n              x-dcl-go-name: Binauthz\n              x-dcl-go-type: FeatureMembershipConfigmanagementBinauthz\n              description: '**DEPRECATED** Binauthz configuration for the cluster.\n                This field will be ignored and should not be set.'\n              x-dcl-server-default: true\n              properties:\n                enabled:\n                  type: boolean\n                  x-dcl-go-name: Enabled\n                  description: Whether binauthz is enabled in this cluster.\n            configSync:\n              type: object\n              x-dcl-go-name: ConfigSync\n              x-dcl-go-type: FeatureMembershipConfigmanagementConfigSync\n              description: Config Sync configuration for the cluster.\n              x-dcl-send-empty: true\n              properties:\n                git:\n                  type: object\n                  x-dcl-go-name: Git\n                  x-dcl-go-type: FeatureMembershipConfigmanagementConfigSyncGit\n                  properties:\n                    gcpServiceAccountEmail:\n                      type: string\n                      x-dcl-go-name: GcpServiceAccountEmail\n                      description: The GCP Service Account Email used for auth when\n                        secretType is gcpServiceAccount.\n                      x-dcl-references:\n                      - resource: Iam/ServiceAccount\n                        field: email\n                    httpsProxy:\n                      type: string\n                      x-dcl-go-name: HttpsProxy\n                      description: URL for the HTTPS proxy to be used when communicating\n                        with the Git repo.\n                    policyDir:\n                      type: string\n                      x-dcl-go-name: PolicyDir\n                      description: 'The path within the Git repository that represents\n                        the top level of the repo to sync. Default: the root directory\n                        of the repository.'\n                    secretType:\n                      type: string\n                      x-dcl-go-name: SecretType\n                      description: Type of secret configured for access to the Git\n                        repo. Must be one of ssh, cookiefile, gcenode, token, gcpserviceaccount\n                        or none. The validation of this is case-sensitive.\n                    syncBranch:\n                      type: string\n                      x-dcl-go-name: SyncBranch\n                      description: 'The branch of the repository to sync from. Default:\n                        master.'\n                    syncRepo:\n                      type: string\n                      x-dcl-go-name: SyncRepo\n                      description: The URL of the Git repository to use as the source\n                        of truth.\n                    syncRev:\n                      type: string\n                      x-dcl-go-name: SyncRev\n                      description: Git revision (tag or hash) to check out. Default\n                        HEAD.\n                    syncWaitSecs:\n                      type: string\n                      x-dcl-go-name: SyncWaitSecs\n                      description: 'Period in seconds between consecutive syncs. Default:\n                        15.'\n                metricsGcpServiceAccountEmail:\n                  type: string\n                  x-dcl-go-name: MetricsGcpServiceAccountEmail\n                  description: The Email of the Google Cloud Service Account (GSA)\n                    used for exporting Config Sync metrics to Cloud Monitoring. The\n                    GSA should have the Monitoring Metric Writer(roles/monitoring.metricWriter)\n                    IAM role. The Kubernetes ServiceAccount `default` in the namespace\n                    `config-management-monitoring` should be bound to the GSA.\n                  x-dcl-references:\n                  - resource: Iam/ServiceAccount\n                    field: email\n                oci:\n                  type: object\n                  x-dcl-go-name: Oci\n                  x-dcl-go-type: FeatureMembershipConfigmanagementConfigSyncOci\n                  properties:\n                    gcpServiceAccountEmail:\n                      type: string\n                      x-dcl-go-name: GcpServiceAccountEmail\n                      description: 'The GCP Service Account Email used for auth when\n                        secret_type is gcpserviceaccount. '\n                      x-dcl-references:\n                      - resource: Iam/ServiceAccount\n                        field: email\n                    policyDir:\n                      type: string\n                      x-dcl-go-name: PolicyDir\n                      description: 'The absolute path of the directory that contains\n                        the local resources. Default: the root directory of the image.'\n                    secretType:\n                      type: string\n                      x-dcl-go-name: SecretType\n                      description: Type of secret configured for access to the OCI\n                        Image. Must be one of gcenode, gcpserviceaccount or none.\n                        The validation of this is case-sensitive.\n                    syncRepo:\n                      type: string\n                      x-dcl-go-name: SyncRepo\n                      description: The OCI image repository URL for the package to\n                        sync from. e.g. LOCATION-docker.pkg.dev/PROJECT_ID/REPOSITORY_NAME/PACKAGE_NAME.\n                    syncWaitSecs:\n                      type: string\n                      x-dcl-go-name: SyncWaitSecs\n                      description: 'Period in seconds(int64 format) between consecutive\n                        syncs. Default: 15.'\n                preventDrift:\n                  type: boolean\n                  x-dcl-go-name: PreventDrift\n                  description: Set to true to enable the Config Sync admission webhook\n                    to prevent drifts. If set to `false`, disables the Config Sync\n                    admission webhook and does not prevent drifts.\n                  x-dcl-server-default: true\n                sourceFormat:\n                  type: string\n                  x-dcl-go-name: SourceFormat\n                  description: Specifies whether the Config Sync Repo is in \"hierarchical\"\n                    or \"unstructured\" mode.\n            hierarchyController:\n              type: object\n              x-dcl-go-name: HierarchyController\n              x-dcl-go-type: FeatureMembershipConfigmanagementHierarchyController\n              description: Hierarchy Controller configuration for the cluster.\n              x-dcl-send-empty: true\n              properties:\n                enableHierarchicalResourceQuota:\n                  type: boolean\n                  x-dcl-go-name: EnableHierarchicalResourceQuota\n                  description: Whether hierarchical resource quota is enabled in this\n                    cluster.\n                  x-dcl-send-empty: true\n                enablePodTreeLabels:\n                  type: boolean\n                  x-dcl-go-name: EnablePodTreeLabels\n                  description: Whether pod tree labels are enabled in this cluster.\n                  x-dcl-send-empty: true\n                enabled:\n                  type: boolean\n                  x-dcl-go-name: Enabled\n                  description: Whether Hierarchy Controller is enabled in this cluster.\n                  x-dcl-send-empty: true\n            policyController:\n              type: object\n              x-dcl-go-name: PolicyController\n              x-dcl-go-type: FeatureMembershipConfigmanagementPolicyController\n              description: Policy Controller configuration for the cluster.\n              properties:\n                auditIntervalSeconds:\n                  type: string\n                  x-dcl-go-name: AuditIntervalSeconds\n                  description: Sets the interval for Policy Controller Audit Scans\n                    (in seconds). When set to 0, this disables audit functionality\n                    altogether.\n                enabled:\n                  type: boolean\n                  x-dcl-go-name: Enabled\n                  description: Enables the installation of Policy Controller. If false,\n                    the rest of PolicyController fields take no effect.\n                exemptableNamespaces:\n                  type: array\n                  x-dcl-go-name: ExemptableNamespaces\n                  description: The set of namespaces that are excluded from Policy\n                    Controller checks. Namespaces do not need to currently exist on\n                    the cluster.\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: string\n                    x-dcl-go-type: string\n                logDeniesEnabled:\n                  type: boolean\n                  x-dcl-go-name: LogDeniesEnabled\n                  description: Logs all denies and dry run failures.\n                monitoring:\n                  type: object\n                  x-dcl-go-name: Monitoring\n                  x-dcl-go-type: FeatureMembershipConfigmanagementPolicyControllerMonitoring\n                  description: 'Specifies the backends Policy Controller should export\n                    metrics to. For example, to specify metrics should be exported\n                    to Cloud Monitoring and Prometheus, specify backends: [\"cloudmonitoring\",\n                    \"prometheus\"]. Default: [\"cloudmonitoring\", \"prometheus\"]'\n                  x-dcl-server-default: true\n                  properties:\n                    backends:\n                      type: array\n                      x-dcl-go-name: Backends\n                      description: ' Specifies the list of backends Policy Controller\n                        will export to. Specifying an empty value `[]` disables metrics\n                        export.'\n                      x-dcl-server-default: true\n                      x-dcl-send-empty: true\n                      x-dcl-list-type: list\n                      items:\n                        type: string\n                        x-dcl-go-type: FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum\n                        enum:\n                        - MONITORING_BACKEND_UNSPECIFIED\n                        - PROMETHEUS\n                        - CLOUD_MONITORING\n                mutationEnabled:\n                  type: boolean\n                  x-dcl-go-name: MutationEnabled\n                  description: Enable or disable mutation in policy controller. If\n                    true, mutation CRDs, webhook and controller deployment will be\n                    deployed to the cluster.\n                referentialRulesEnabled:\n                  type: boolean\n                  x-dcl-go-name: ReferentialRulesEnabled\n                  description: Enables the ability to use Constraint Templates that\n                    reference to objects other than the object currently being evaluated.\n                templateLibraryInstalled:\n                  type: boolean\n                  x-dcl-go-name: TemplateLibraryInstalled\n                  description: Installs the default template library along with Policy\n                    Controller.\n            version:\n              type: string\n              x-dcl-go-name: Version\n              description: Optional. Version of ACM to install. Defaults to the latest\n                version.\n              x-dcl-server-default: true\n        feature:\n          type: string\n          x-dcl-go-name: Feature\n          description: The name of the feature\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Gkehub/Feature\n            field: name\n            parent: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location of the feature\n          x-kubernetes-immutable: true\n        membership:\n          type: string\n          x-dcl-go-name: Membership\n          description: The name of the membership\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Gkehub/Membership\n            field: name\n        membershipLocation:\n          type: string\n          x-dcl-go-name: MembershipLocation\n          description: The location of the membership\n          x-kubernetes-immutable: true\n        mesh:\n          type: object\n          x-dcl-go-name: Mesh\n          x-dcl-go-type: FeatureMembershipMesh\n          description: Manage Mesh Features\n          properties:\n            controlPlane:\n              type: string\n              x-dcl-go-name: ControlPlane\n              x-dcl-go-type: FeatureMembershipMeshControlPlaneEnum\n              description: '**DEPRECATED** Whether to automatically manage Service\n                Mesh control planes. Possible values: CONTROL_PLANE_MANAGEMENT_UNSPECIFIED,\n                AUTOMATIC, MANUAL'\n              enum:\n              - CONTROL_PLANE_MANAGEMENT_UNSPECIFIED\n              - AUTOMATIC\n              - MANUAL\n            management:\n              type: string\n              x-dcl-go-name: Management\n              x-dcl-go-type: FeatureMembershipMeshManagementEnum\n              description: 'Whether to automatically manage Service Mesh. Possible\n                values: MANAGEMENT_UNSPECIFIED, MANAGEMENT_AUTOMATIC, MANAGEMENT_MANUAL'\n              enum:\n              - MANAGEMENT_UNSPECIFIED\n              - MANAGEMENT_AUTOMATIC\n              - MANAGEMENT_MANUAL\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project of the feature\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n")

// 16166 bytes
// MD5: 9bb422f2e116935c665d320385a8cb2e
