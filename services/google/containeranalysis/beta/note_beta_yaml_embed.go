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
// gen_go_data -package beta -var YAML_note blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/containeranalysis/beta/note.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/containeranalysis/beta/note.yaml
var YAML_note = []byte("info:\n  title: ContainerAnalysis/Note\n  description: The ContainerAnalysis Note resource\n  x-dcl-struct-name: Note\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Note\n    parameters:\n    - name: note\n      required: true\n      description: A full instance of a Note\n  apply:\n    description: The function used to apply information about a Note\n    parameters:\n    - name: note\n      required: true\n      description: A full instance of a Note\n  delete:\n    description: The function used to delete a Note\n    parameters:\n    - name: note\n      required: true\n      description: A full instance of a Note\n  deleteAll:\n    description: The function used to delete all Note\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Note\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Note:\n      title: Note\n      x-dcl-id: projects/{{project}}/notes/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - project\n      properties:\n        attestation:\n          type: object\n          x-dcl-go-name: Attestation\n          x-dcl-go-type: NoteAttestation\n          description: A note describing an attestation role.\n          x-dcl-conflicts:\n          - vulnerability\n          - build\n          - image\n          - package\n          - deployment\n          - discovery\n          properties:\n            hint:\n              type: object\n              x-dcl-go-name: Hint\n              x-dcl-go-type: NoteAttestationHint\n              description: Hint hints at the purpose of the attestation authority.\n              required:\n              - humanReadableName\n              properties:\n                humanReadableName:\n                  type: string\n                  x-dcl-go-name: HumanReadableName\n                  description: Required. The human readable name of this attestation\n                    authority, for example \"qa\".\n        build:\n          type: object\n          x-dcl-go-name: Build\n          x-dcl-go-type: NoteBuild\n          description: A note describing build provenance for a verifiable build.\n          x-dcl-conflicts:\n          - vulnerability\n          - image\n          - package\n          - deployment\n          - discovery\n          - attestation\n          required:\n          - builderVersion\n          properties:\n            builderVersion:\n              type: string\n              x-dcl-go-name: BuilderVersion\n              description: Required. Immutable. Version of the builder which produced\n                this build.\n            signature:\n              type: object\n              x-dcl-go-name: Signature\n              x-dcl-go-type: NoteBuildSignature\n              description: Signature of the build in occurrences pointing to this\n                build note containing build details.\n              required:\n              - signature\n              properties:\n                keyId:\n                  type: string\n                  x-dcl-go-name: KeyId\n                  description: An ID for the key used to sign. This could be either\n                    an ID for the key stored in `public_key` (such as the ID or fingerprint\n                    for a PGP key, or the CN for a cert), or a reference to an external\n                    key (such as a reference to a key in Cloud Key Management Service).\n                keyType:\n                  type: string\n                  x-dcl-go-name: KeyType\n                  x-dcl-go-type: NoteBuildSignatureKeyTypeEnum\n                  description: 'The type of the key, either stored in `public_key`\n                    or referenced in `key_id`. Possible values: KEY_TYPE_UNSPECIFIED,\n                    PGP_ASCII_ARMORED, PKIX_PEM'\n                  enum:\n                  - KEY_TYPE_UNSPECIFIED\n                  - PGP_ASCII_ARMORED\n                  - PKIX_PEM\n                publicKey:\n                  type: string\n                  x-dcl-go-name: PublicKey\n                  description: 'Public key of the builder which can be used to verify\n                    that the related findings are valid and unchanged. If `key_type`\n                    is empty, this defaults to PEM encoded public keys. This field\n                    may be empty if `key_id` references an external key. For Cloud\n                    Build based signatures, this is a PEM encoded public key. To verify\n                    the Cloud Build signature, place the contents of this field into\n                    a file (public.pem). The signature field is base64-decoded into\n                    its binary representation in signature.bin, and the provenance\n                    bytes from `BuildDetails` are base64-decoded into a binary representation\n                    in signed.bin. OpenSSL can then verify the signature: `openssl\n                    sha256 -verify public.pem -signature signature.bin signed.bin`'\n                signature:\n                  type: string\n                  x-dcl-go-name: Signature\n                  description: Required. Signature of the related `BuildProvenance`.\n                    In JSON, this is base-64 encoded.\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time this note was created. This field can\n            be used as a filter in list requests.\n          x-kubernetes-immutable: true\n        deployment:\n          type: object\n          x-dcl-go-name: Deployment\n          x-dcl-go-type: NoteDeployment\n          description: A note describing something that can be deployed.\n          x-dcl-conflicts:\n          - vulnerability\n          - build\n          - image\n          - package\n          - discovery\n          - attestation\n          required:\n          - resourceUri\n          properties:\n            resourceUri:\n              type: array\n              x-dcl-go-name: ResourceUri\n              description: Required. Resource URI for the artifact being deployed.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n        discovery:\n          type: object\n          x-dcl-go-name: Discovery\n          x-dcl-go-type: NoteDiscovery\n          description: A note describing the initial analysis of a resource.\n          x-dcl-conflicts:\n          - vulnerability\n          - build\n          - image\n          - package\n          - deployment\n          - attestation\n          required:\n          - analysisKind\n          properties:\n            analysisKind:\n              type: string\n              x-dcl-go-name: AnalysisKind\n              x-dcl-go-type: NoteDiscoveryAnalysisKindEnum\n              description: 'The kind of analysis that is handled by this discovery.\n                Possible values: NOTE_KIND_UNSPECIFIED, VULNERABILITY, BUILD, IMAGE,\n                PACKAGE, DEPLOYMENT, DISCOVERY, ATTESTATION, UPGRADE'\n              enum:\n              - NOTE_KIND_UNSPECIFIED\n              - VULNERABILITY\n              - BUILD\n              - IMAGE\n              - PACKAGE\n              - DEPLOYMENT\n              - DISCOVERY\n              - ATTESTATION\n              - UPGRADE\n        expirationTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: ExpirationTime\n          description: Time of expiration for this note. Empty if note does not expire.\n        image:\n          type: object\n          x-dcl-go-name: Image\n          x-dcl-go-type: NoteImage\n          description: A note describing a base image.\n          x-dcl-conflicts:\n          - vulnerability\n          - build\n          - package\n          - deployment\n          - discovery\n          - attestation\n          required:\n          - resourceUrl\n          - fingerprint\n          properties:\n            fingerprint:\n              type: object\n              x-dcl-go-name: Fingerprint\n              x-dcl-go-type: NoteImageFingerprint\n              description: Required. Immutable. The fingerprint of the base image.\n              required:\n              - v1Name\n              - v2Blob\n              properties:\n                v1Name:\n                  type: string\n                  x-dcl-go-name: V1Name\n                  description: Required. The layer ID of the final layer in the Docker\n                    image's v1 representation.\n                v2Blob:\n                  type: array\n                  x-dcl-go-name: V2Blob\n                  description: Required. The ordered list of v2 blobs that represent\n                    a given image.\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: string\n                    x-dcl-go-type: string\n                v2Name:\n                  type: string\n                  x-dcl-go-name: V2Name\n                  readOnly: true\n                  description: 'Output only. The name of the image''s v2 blobs computed\n                    via: ) Only the name of the final blob is kept.'\n            resourceUrl:\n              type: string\n              x-dcl-go-name: ResourceUrl\n              description: Required. Immutable. The resource_url for the resource\n                representing the basis of associated occurrence images.\n        longDescription:\n          type: string\n          x-dcl-go-name: LongDescription\n          description: A detailed description of this note.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. Immutable. The name of the package.\n          x-kubernetes-immutable: true\n          x-dcl-has-long-form: true\n        package:\n          type: object\n          x-dcl-go-name: Package\n          x-dcl-go-type: NotePackage\n          description: Required for non-Windows OS. The package this Upgrade is for.\n          x-dcl-conflicts:\n          - vulnerability\n          - build\n          - image\n          - deployment\n          - discovery\n          - attestation\n          required:\n          - name\n          properties:\n            distribution:\n              type: array\n              x-dcl-go-name: Distribution\n              description: The various channels by which a package is distributed.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: NotePackageDistribution\n                required:\n                - cpeUri\n                properties:\n                  architecture:\n                    type: string\n                    x-dcl-go-name: Architecture\n                    x-dcl-go-type: NotePackageDistributionArchitectureEnum\n                    description: 'The CPU architecture for which packages in this\n                      distribution channel were built Possible values: ARCHITECTURE_UNSPECIFIED,\n                      X86, X64'\n                    enum:\n                    - ARCHITECTURE_UNSPECIFIED\n                    - X86\n                    - X64\n                  cpeUri:\n                    type: string\n                    x-dcl-go-name: CpeUri\n                    description: The cpe_uri in [cpe format](https://cpe.mitre.org/specification/)\n                      denoting the package manager version distributing a package.\n                  description:\n                    type: string\n                    x-dcl-go-name: Description\n                    description: The distribution channel-specific description of\n                      this package.\n                  latestVersion:\n                    type: object\n                    x-dcl-go-name: LatestVersion\n                    x-dcl-go-type: NotePackageDistributionLatestVersion\n                    description: The latest available version of this package in this\n                      distribution channel.\n                    required:\n                    - kind\n                    properties:\n                      epoch:\n                        type: integer\n                        format: int64\n                        x-dcl-go-name: Epoch\n                        description: Used to correct mistakes in the version numbering\n                          scheme.\n                      fullName:\n                        type: string\n                        x-dcl-go-name: FullName\n                        description: Human readable version string. This string is\n                          of the form :- and is only set when kind is NORMAL.\n                      kind:\n                        type: string\n                        x-dcl-go-name: Kind\n                        x-dcl-go-type: NotePackageDistributionLatestVersionKindEnum\n                        description: 'Distinguish between sentinel MIN/MAX versions\n                          and normal versions. If kind is not NORMAL, then the other\n                          fields are ignored. Possible values: VERSION_KIND_UNSPECIFIED,\n                          NORMAL, MINIMUM, MAXIMUM'\n                        enum:\n                        - VERSION_KIND_UNSPECIFIED\n                        - NORMAL\n                        - MINIMUM\n                        - MAXIMUM\n                      name:\n                        type: string\n                        x-dcl-go-name: Name\n                        description: The main part of the version name.\n                      revision:\n                        type: string\n                        x-dcl-go-name: Revision\n                        description: The iteration of the package build from the above\n                          version.\n                  maintainer:\n                    type: string\n                    x-dcl-go-name: Maintainer\n                    description: A freeform string denoting the maintainer of this\n                      package.\n                  url:\n                    type: string\n                    x-dcl-go-name: Url\n                    description: The distribution channel-specific homepage for this\n                      package.\n            name:\n              type: string\n              x-dcl-go-name: Name\n              description: The name of the package.\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-parameter: true\n        relatedNoteNames:\n          type: array\n          x-dcl-go-name: RelatedNoteNames\n          description: Other notes related to this note.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n            x-dcl-references:\n            - resource: Containeranalysis/Note\n              field: name\n        relatedUrl:\n          type: array\n          x-dcl-go-name: RelatedUrl\n          description: URLs associated with this note.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: NoteRelatedUrl\n            properties:\n              label:\n                type: string\n                x-dcl-go-name: Label\n                description: Label to describe usage of the URL\n              url:\n                type: string\n                x-dcl-go-name: Url\n                description: Specific URL to associate with the note\n        shortDescription:\n          type: string\n          x-dcl-go-name: ShortDescription\n          description: A one sentence description of this note.\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time this note was last updated. This field\n            can be used as a filter in list requests.\n          x-kubernetes-immutable: true\n        vulnerability:\n          type: object\n          x-dcl-go-name: Vulnerability\n          x-dcl-go-type: NoteVulnerability\n          description: A note describing a package vulnerability.\n          x-dcl-conflicts:\n          - build\n          - image\n          - package\n          - deployment\n          - discovery\n          - attestation\n          properties:\n            cvssScore:\n              type: number\n              format: double\n              x-dcl-go-name: CvssScore\n              description: The CVSS score of this vulnerability. CVSS score is on\n                a scale of 0 - 10 where 0 indicates low severity and 10 indicates\n                high severity.\n            cvssV3:\n              type: object\n              x-dcl-go-name: CvssV3\n              x-dcl-go-type: NoteVulnerabilityCvssV3\n              description: The full description of the CVSSv3 for this vulnerability.\n              properties:\n                attackComplexity:\n                  type: string\n                  x-dcl-go-name: AttackComplexity\n                  x-dcl-go-type: NoteVulnerabilityCvssV3AttackComplexityEnum\n                  description: ' Possible values: ATTACK_COMPLEXITY_UNSPECIFIED, ATTACK_COMPLEXITY_LOW,\n                    ATTACK_COMPLEXITY_HIGH'\n                  enum:\n                  - ATTACK_COMPLEXITY_UNSPECIFIED\n                  - ATTACK_COMPLEXITY_LOW\n                  - ATTACK_COMPLEXITY_HIGH\n                attackVector:\n                  type: string\n                  x-dcl-go-name: AttackVector\n                  x-dcl-go-type: NoteVulnerabilityCvssV3AttackVectorEnum\n                  description: 'Base Metrics Represents the intrinsic characteristics\n                    of a vulnerability that are constant over time and across user\n                    environments. Possible values: ATTACK_VECTOR_UNSPECIFIED, ATTACK_VECTOR_NETWORK,\n                    ATTACK_VECTOR_ADJACENT, ATTACK_VECTOR_LOCAL, ATTACK_VECTOR_PHYSICAL'\n                  enum:\n                  - ATTACK_VECTOR_UNSPECIFIED\n                  - ATTACK_VECTOR_NETWORK\n                  - ATTACK_VECTOR_ADJACENT\n                  - ATTACK_VECTOR_LOCAL\n                  - ATTACK_VECTOR_PHYSICAL\n                availabilityImpact:\n                  type: string\n                  x-dcl-go-name: AvailabilityImpact\n                  x-dcl-go-type: NoteVulnerabilityCvssV3AvailabilityImpactEnum\n                  description: ' Possible values: IMPACT_UNSPECIFIED, IMPACT_HIGH,\n                    IMPACT_LOW, IMPACT_NONE'\n                  enum:\n                  - IMPACT_UNSPECIFIED\n                  - IMPACT_HIGH\n                  - IMPACT_LOW\n                  - IMPACT_NONE\n                baseScore:\n                  type: number\n                  format: double\n                  x-dcl-go-name: BaseScore\n                  description: The base score is a function of the base metric scores.\n                confidentialityImpact:\n                  type: string\n                  x-dcl-go-name: ConfidentialityImpact\n                  x-dcl-go-type: NoteVulnerabilityCvssV3ConfidentialityImpactEnum\n                  description: ' Possible values: IMPACT_UNSPECIFIED, IMPACT_HIGH,\n                    IMPACT_LOW, IMPACT_NONE'\n                  enum:\n                  - IMPACT_UNSPECIFIED\n                  - IMPACT_HIGH\n                  - IMPACT_LOW\n                  - IMPACT_NONE\n                exploitabilityScore:\n                  type: number\n                  format: double\n                  x-dcl-go-name: ExploitabilityScore\n                impactScore:\n                  type: number\n                  format: double\n                  x-dcl-go-name: ImpactScore\n                integrityImpact:\n                  type: string\n                  x-dcl-go-name: IntegrityImpact\n                  x-dcl-go-type: NoteVulnerabilityCvssV3IntegrityImpactEnum\n                  description: ' Possible values: IMPACT_UNSPECIFIED, IMPACT_HIGH,\n                    IMPACT_LOW, IMPACT_NONE'\n                  enum:\n                  - IMPACT_UNSPECIFIED\n                  - IMPACT_HIGH\n                  - IMPACT_LOW\n                  - IMPACT_NONE\n                privilegesRequired:\n                  type: string\n                  x-dcl-go-name: PrivilegesRequired\n                  x-dcl-go-type: NoteVulnerabilityCvssV3PrivilegesRequiredEnum\n                  description: ' Possible values: PRIVILEGES_REQUIRED_UNSPECIFIED,\n                    PRIVILEGES_REQUIRED_NONE, PRIVILEGES_REQUIRED_LOW, PRIVILEGES_REQUIRED_HIGH'\n                  enum:\n                  - PRIVILEGES_REQUIRED_UNSPECIFIED\n                  - PRIVILEGES_REQUIRED_NONE\n                  - PRIVILEGES_REQUIRED_LOW\n                  - PRIVILEGES_REQUIRED_HIGH\n                scope:\n                  type: string\n                  x-dcl-go-name: Scope\n                  x-dcl-go-type: NoteVulnerabilityCvssV3ScopeEnum\n                  description: ' Possible values: SCOPE_UNSPECIFIED, SCOPE_UNCHANGED,\n                    SCOPE_CHANGED'\n                  enum:\n                  - SCOPE_UNSPECIFIED\n                  - SCOPE_UNCHANGED\n                  - SCOPE_CHANGED\n                userInteraction:\n                  type: string\n                  x-dcl-go-name: UserInteraction\n                  x-dcl-go-type: NoteVulnerabilityCvssV3UserInteractionEnum\n                  description: ' Possible values: USER_INTERACTION_UNSPECIFIED, USER_INTERACTION_NONE,\n                    USER_INTERACTION_REQUIRED'\n                  enum:\n                  - USER_INTERACTION_UNSPECIFIED\n                  - USER_INTERACTION_NONE\n                  - USER_INTERACTION_REQUIRED\n            details:\n              type: array\n              x-dcl-go-name: Details\n              description: Details of all known distros and packages affected by this\n                vulnerability.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: NoteVulnerabilityDetails\n                required:\n                - affectedCpeUri\n                - affectedPackage\n                properties:\n                  affectedCpeUri:\n                    type: string\n                    x-dcl-go-name: AffectedCpeUri\n                    description: Required. The (https://cpe.mitre.org/specification/)\n                      this vulnerability affects.\n                  affectedPackage:\n                    type: string\n                    x-dcl-go-name: AffectedPackage\n                    description: Required. The package this vulnerability affects.\n                  affectedVersionEnd:\n                    type: object\n                    x-dcl-go-name: AffectedVersionEnd\n                    x-dcl-go-type: NoteVulnerabilityDetailsAffectedVersionEnd\n                    description: 'The version number at the end of an interval in\n                      which this vulnerability exists. A vulnerability can affect\n                      a package between version numbers that are disjoint sets of\n                      intervals (example: ) each of which will be represented in its\n                      own Detail. If a specific affected version is provided by a\n                      vulnerability database, affected_version_start and affected_version_end\n                      will be the same in that Detail.'\n                    required:\n                    - kind\n                    properties:\n                      epoch:\n                        type: integer\n                        format: int64\n                        x-dcl-go-name: Epoch\n                        description: Used to correct mistakes in the version numbering\n                          scheme.\n                      fullName:\n                        type: string\n                        x-dcl-go-name: FullName\n                        description: Human readable version string. This string is\n                          of the form :- and is only set when kind is NORMAL.\n                      kind:\n                        type: string\n                        x-dcl-go-name: Kind\n                        x-dcl-go-type: NoteVulnerabilityDetailsAffectedVersionEndKindEnum\n                        description: 'Required. Distinguishes between sentinel MIN/MAX\n                          versions and normal versions. Possible values: NOTE_KIND_UNSPECIFIED,\n                          VULNERABILITY, BUILD, IMAGE, PACKAGE, DEPLOYMENT, DISCOVERY,\n                          ATTESTATION, UPGRADE'\n                        enum:\n                        - NOTE_KIND_UNSPECIFIED\n                        - VULNERABILITY\n                        - BUILD\n                        - IMAGE\n                        - PACKAGE\n                        - DEPLOYMENT\n                        - DISCOVERY\n                        - ATTESTATION\n                        - UPGRADE\n                      name:\n                        type: string\n                        x-dcl-go-name: Name\n                        description: Required only when version kind is NORMAL. The\n                          main part of the version name.\n                      revision:\n                        type: string\n                        x-dcl-go-name: Revision\n                        description: The iteration of the package build from the above\n                          version.\n                  affectedVersionStart:\n                    type: object\n                    x-dcl-go-name: AffectedVersionStart\n                    x-dcl-go-type: NoteVulnerabilityDetailsAffectedVersionStart\n                    description: 'The version number at the start of an interval in\n                      which this vulnerability exists. A vulnerability can affect\n                      a package between version numbers that are disjoint sets of\n                      intervals (example: ) each of which will be represented in its\n                      own Detail. If a specific affected version is provided by a\n                      vulnerability database, affected_version_start and affected_version_end\n                      will be the same in that Detail.'\n                    required:\n                    - kind\n                    properties:\n                      epoch:\n                        type: integer\n                        format: int64\n                        x-dcl-go-name: Epoch\n                        description: Used to correct mistakes in the version numbering\n                          scheme.\n                      fullName:\n                        type: string\n                        x-dcl-go-name: FullName\n                        description: Human readable version string. This string is\n                          of the form :- and is only set when kind is NORMAL.\n                        x-dcl-server-default: true\n                      kind:\n                        type: string\n                        x-dcl-go-name: Kind\n                        x-dcl-go-type: NoteVulnerabilityDetailsAffectedVersionStartKindEnum\n                        description: 'Required. Distinguishes between sentinel MIN/MAX\n                          versions and normal versions. Possible values: NOTE_KIND_UNSPECIFIED,\n                          VULNERABILITY, BUILD, IMAGE, PACKAGE, DEPLOYMENT, DISCOVERY,\n                          ATTESTATION, UPGRADE'\n                        enum:\n                        - NOTE_KIND_UNSPECIFIED\n                        - VULNERABILITY\n                        - BUILD\n                        - IMAGE\n                        - PACKAGE\n                        - DEPLOYMENT\n                        - DISCOVERY\n                        - ATTESTATION\n                        - UPGRADE\n                      name:\n                        type: string\n                        x-dcl-go-name: Name\n                        description: Required only when version kind is NORMAL. The\n                          main part of the version name.\n                      revision:\n                        type: string\n                        x-dcl-go-name: Revision\n                        description: The iteration of the package build from the above\n                          version.\n                  description:\n                    type: string\n                    x-dcl-go-name: Description\n                    description: A vendor-specific description of this vulnerability.\n                  fixedCpeUri:\n                    type: string\n                    x-dcl-go-name: FixedCpeUri\n                    description: The distro recommended (https://cpe.mitre.org/specification/)\n                      to update to that contains a fix for this vulnerability. It\n                      is possible for this to be different from the affected_cpe_uri.\n                  fixedPackage:\n                    type: string\n                    x-dcl-go-name: FixedPackage\n                    description: The distro recommended package to update to that\n                      contains a fix for this vulnerability. It is possible for this\n                      to be different from the affected_package.\n                  fixedVersion:\n                    type: object\n                    x-dcl-go-name: FixedVersion\n                    x-dcl-go-type: NoteVulnerabilityDetailsFixedVersion\n                    description: The distro recommended version to update to that\n                      contains a fix for this vulnerability. Setting this to VersionKind.MAXIMUM\n                      means no such version is yet available.\n                    required:\n                    - kind\n                    properties:\n                      epoch:\n                        type: integer\n                        format: int64\n                        x-dcl-go-name: Epoch\n                        description: Used to correct mistakes in the version numbering\n                          scheme.\n                      fullName:\n                        type: string\n                        x-dcl-go-name: FullName\n                        description: Human readable version string. This string is\n                          of the form :- and is only set when kind is NORMAL.\n                      kind:\n                        type: string\n                        x-dcl-go-name: Kind\n                        x-dcl-go-type: NoteVulnerabilityDetailsFixedVersionKindEnum\n                        description: 'Required. Distinguishes between sentinel MIN/MAX\n                          versions and normal versions. Possible values: NOTE_KIND_UNSPECIFIED,\n                          VULNERABILITY, BUILD, IMAGE, PACKAGE, DEPLOYMENT, DISCOVERY,\n                          ATTESTATION, UPGRADE'\n                        enum:\n                        - NOTE_KIND_UNSPECIFIED\n                        - VULNERABILITY\n                        - BUILD\n                        - IMAGE\n                        - PACKAGE\n                        - DEPLOYMENT\n                        - DISCOVERY\n                        - ATTESTATION\n                        - UPGRADE\n                      name:\n                        type: string\n                        x-dcl-go-name: Name\n                        description: Required only when version kind is NORMAL. The\n                          main part of the version name.\n                      revision:\n                        type: string\n                        x-dcl-go-name: Revision\n                        description: The iteration of the package build from the above\n                          version.\n                  isObsolete:\n                    type: boolean\n                    x-dcl-go-name: IsObsolete\n                    description: Whether this detail is obsolete. Occurrences are\n                      expected not to point to obsolete details.\n                  packageType:\n                    type: string\n                    x-dcl-go-name: PackageType\n                    description: The type of package; whether native or non native\n                      (e.g., ruby gems, node.js packages, etc.).\n                  severityName:\n                    type: string\n                    x-dcl-go-name: SeverityName\n                    description: The distro assigned severity of this vulnerability.\n                  sourceUpdateTime:\n                    type: string\n                    format: date-time\n                    x-dcl-go-name: SourceUpdateTime\n                    description: The time this information was last changed at the\n                      source. This is an upstream timestamp from the underlying information\n                      source - e.g. Ubuntu security tracker.\n            severity:\n              type: string\n              x-dcl-go-name: Severity\n              x-dcl-go-type: NoteVulnerabilitySeverityEnum\n              description: 'The note provider assigned severity of this vulnerability.\n                Possible values: SEVERITY_UNSPECIFIED, MINIMAL, LOW, MEDIUM, HIGH,\n                CRITICAL'\n              enum:\n              - SEVERITY_UNSPECIFIED\n              - MINIMAL\n              - LOW\n              - MEDIUM\n              - HIGH\n              - CRITICAL\n            sourceUpdateTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: SourceUpdateTime\n              description: The time this information was last changed at the source.\n                This is an upstream timestamp from the underlying information source\n                - e.g. Ubuntu security tracker.\n            windowsDetails:\n              type: array\n              x-dcl-go-name: WindowsDetails\n              description: Windows details get their own format because the information\n                format and model don't match a normal detail. Specifically Windows\n                updates are done as patches, thus Windows vulnerabilities really are\n                a missing package, rather than a package being at an incorrect version.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: NoteVulnerabilityWindowsDetails\n                required:\n                - cpeUri\n                - name\n                - fixingKbs\n                properties:\n                  cpeUri:\n                    type: string\n                    x-dcl-go-name: CpeUri\n                    description: Required. The (https://cpe.mitre.org/specification/)\n                      this vulnerability affects.\n                  description:\n                    type: string\n                    x-dcl-go-name: Description\n                    description: The description of this vulnerability.\n                  fixingKbs:\n                    type: array\n                    x-dcl-go-name: FixingKbs\n                    description: Required. The names of the KBs which have hotfixes\n                      to mitigate this vulnerability. Note that there may be multiple\n                      hotfixes (and thus multiple KBs) that mitigate a given vulnerability.\n                      Currently any listed KBs presence is considered a fix.\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: object\n                      x-dcl-go-type: NoteVulnerabilityWindowsDetailsFixingKbs\n                      properties:\n                        name:\n                          type: string\n                          x-dcl-go-name: Name\n                          description: The KB name (generally of the form KB+ (e.g.,\n                            KB123456)).\n                        url:\n                          type: string\n                          x-dcl-go-name: Url\n                          description: A link to the KB in the (https://www.catalog.update.microsoft.com/).\n                  name:\n                    type: string\n                    x-dcl-go-name: Name\n                    description: Required. The name of this vulnerability.\n")

// 36395 bytes
// MD5: 8e9963b6ae1aff8fc551832bd3bb753b
