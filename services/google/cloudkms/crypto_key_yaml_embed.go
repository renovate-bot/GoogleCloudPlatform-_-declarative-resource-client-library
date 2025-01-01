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
// gen_go_data -package cloudkms -var YAML_crypto_key blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudkms/crypto_key.yaml

package cloudkms

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/cloudkms/crypto_key.yaml
var YAML_crypto_key = []byte("info:\n  title: Cloudkms/CryptoKey\n  description: The Cloudkms CryptoKey resource\n  x-dcl-struct-name: CryptoKey\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a CryptoKey\n    parameters:\n    - name: cryptoKey\n      required: true\n      description: A full instance of a CryptoKey\n  apply:\n    description: The function used to apply information about a CryptoKey\n    parameters:\n    - name: cryptoKey\n      required: true\n      description: A full instance of a CryptoKey\n  list:\n    description: The function used to list information about many CryptoKey\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: keyRing\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    CryptoKey:\n      title: CryptoKey\n      x-dcl-id: projects/{{project}}/locations/{{location}}/keyRings/{{key_ring}}/cryptoKeys/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-create: true\n      x-dcl-has-iam: true\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - purpose\n      - project\n      - location\n      - keyRing\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time at which this CryptoKey was created.\n          x-kubernetes-immutable: true\n        destroyScheduledDuration:\n          type: string\n          x-dcl-go-name: DestroyScheduledDuration\n          description: Immutable. The period of time that versions of this key spend\n            in the DESTROY_SCHEDULED state before transitioning to DESTROYED. If not\n            specified at creation time, the default duration is 24 hours.\n          x-kubernetes-immutable: true\n          x-dcl-server-default: true\n        importOnly:\n          type: boolean\n          x-dcl-go-name: ImportOnly\n          description: Immutable. Whether this key may contain imported versions only.\n          x-kubernetes-immutable: true\n        keyRing:\n          type: string\n          x-dcl-go-name: KeyRing\n          description: The key ring for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudkms/KeyRing\n            field: name\n            parent: true\n          x-dcl-parameter: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Labels with user-defined metadata. For more information, see\n            [Labeling Keys](https://cloud.google.com/kms/docs/labeling-keys).\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n          x-dcl-parameter: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The resource name for this CryptoKey in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*`.\n          x-kubernetes-immutable: true\n          x-dcl-has-long-form: true\n        nextRotationTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: NextRotationTime\n          description: 'At next_rotation_time, the Key Management Service will automatically:\n            1. Create a new version of this CryptoKey. 2. Mark the new version as\n            primary. Key rotations performed manually via CreateCryptoKeyVersion and\n            UpdateCryptoKeyPrimaryVersion do not affect next_rotation_time. Keys with\n            purpose ENCRYPT_DECRYPT support automatic rotation. For other keys, this\n            field must be omitted.'\n        primary:\n          type: object\n          x-dcl-go-name: Primary\n          x-dcl-go-type: CryptoKeyPrimary\n          readOnly: true\n          description: Output only. A copy of the \"primary\" CryptoKeyVersion that\n            will be used by Encrypt when this CryptoKey is given in EncryptRequest.name.\n            The CryptoKey's primary version can be updated via UpdateCryptoKeyPrimaryVersion.\n            Keys with purpose ENCRYPT_DECRYPT may have a primary. For other keys,\n            this field will be omitted.\n          properties:\n            algorithm:\n              type: string\n              x-dcl-go-name: Algorithm\n              x-dcl-go-type: CryptoKeyPrimaryAlgorithmEnum\n              readOnly: true\n              description: 'Output only. The CryptoKeyVersionAlgorithm that this CryptoKeyVersion\n                supports. Possible values: CRYPTO_KEY_VERSION_ALGORITHM_UNSPECIFIED,\n                GOOGLE_SYMMETRIC_ENCRYPTION, RSA_SIGN_PSS_2048_SHA256, RSA_SIGN_PSS_3072_SHA256,\n                RSA_SIGN_PSS_4096_SHA256, RSA_SIGN_PSS_4096_SHA512, RSA_SIGN_PKCS1_2048_SHA256,\n                RSA_SIGN_PKCS1_3072_SHA256, RSA_SIGN_PKCS1_4096_SHA256, RSA_SIGN_PKCS1_4096_SHA512,\n                RSA_DECRYPT_OAEP_2048_SHA256, RSA_DECRYPT_OAEP_3072_SHA256, RSA_DECRYPT_OAEP_4096_SHA256,\n                RSA_DECRYPT_OAEP_4096_SHA512, EC_SIGN_P256_SHA256, EC_SIGN_P384_SHA384,\n                EC_SIGN_SECP256K1_SHA256, HMAC_SHA256, EXTERNAL_SYMMETRIC_ENCRYPTION'\n              x-kubernetes-immutable: true\n              enum:\n              - CRYPTO_KEY_VERSION_ALGORITHM_UNSPECIFIED\n              - GOOGLE_SYMMETRIC_ENCRYPTION\n              - RSA_SIGN_PSS_2048_SHA256\n              - RSA_SIGN_PSS_3072_SHA256\n              - RSA_SIGN_PSS_4096_SHA256\n              - RSA_SIGN_PSS_4096_SHA512\n              - RSA_SIGN_PKCS1_2048_SHA256\n              - RSA_SIGN_PKCS1_3072_SHA256\n              - RSA_SIGN_PKCS1_4096_SHA256\n              - RSA_SIGN_PKCS1_4096_SHA512\n              - RSA_DECRYPT_OAEP_2048_SHA256\n              - RSA_DECRYPT_OAEP_3072_SHA256\n              - RSA_DECRYPT_OAEP_4096_SHA256\n              - RSA_DECRYPT_OAEP_4096_SHA512\n              - EC_SIGN_P256_SHA256\n              - EC_SIGN_P384_SHA384\n              - EC_SIGN_SECP256K1_SHA256\n              - HMAC_SHA256\n              - EXTERNAL_SYMMETRIC_ENCRYPTION\n            attestation:\n              type: object\n              x-dcl-go-name: Attestation\n              x-dcl-go-type: CryptoKeyPrimaryAttestation\n              readOnly: true\n              description: Output only. Statement that was generated and signed by\n                the HSM at key creation time. Use this statement to verify attributes\n                of the key as stored on the HSM, independently of Google. Only provided\n                for key versions with protection_level HSM.\n              properties:\n                certChains:\n                  type: object\n                  x-dcl-go-name: CertChains\n                  x-dcl-go-type: CryptoKeyPrimaryAttestationCertChains\n                  readOnly: true\n                  description: Output only. The certificate chains needed to validate\n                    the attestation\n                  properties:\n                    caviumCerts:\n                      type: array\n                      x-dcl-go-name: CaviumCerts\n                      description: Cavium certificate chain corresponding to the attestation.\n                      x-dcl-send-empty: true\n                      x-dcl-list-type: list\n                      items:\n                        type: string\n                        x-dcl-go-type: string\n                    googleCardCerts:\n                      type: array\n                      x-dcl-go-name: GoogleCardCerts\n                      description: Google card certificate chain corresponding to\n                        the attestation.\n                      x-dcl-send-empty: true\n                      x-dcl-list-type: list\n                      items:\n                        type: string\n                        x-dcl-go-type: string\n                    googlePartitionCerts:\n                      type: array\n                      x-dcl-go-name: GooglePartitionCerts\n                      description: Google partition certificate chain corresponding\n                        to the attestation.\n                      x-dcl-send-empty: true\n                      x-dcl-list-type: list\n                      items:\n                        type: string\n                        x-dcl-go-type: string\n                content:\n                  type: string\n                  x-dcl-go-name: Content\n                  readOnly: true\n                  description: Output only. The attestation data provided by the HSM\n                    when the key operation was performed.\n                  x-kubernetes-immutable: true\n                format:\n                  type: string\n                  x-dcl-go-name: Format\n                  x-dcl-go-type: CryptoKeyPrimaryAttestationFormatEnum\n                  readOnly: true\n                  description: 'Output only. The format of the attestation data. Possible\n                    values: ATTESTATION_FORMAT_UNSPECIFIED, CAVIUM_V1_COMPRESSED,\n                    CAVIUM_V2_COMPRESSED'\n                  x-kubernetes-immutable: true\n                  enum:\n                  - ATTESTATION_FORMAT_UNSPECIFIED\n                  - CAVIUM_V1_COMPRESSED\n                  - CAVIUM_V2_COMPRESSED\n            createTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: CreateTime\n              readOnly: true\n              description: Output only. The time at which this CryptoKeyVersion was\n                created.\n              x-kubernetes-immutable: true\n            destroyEventTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: DestroyEventTime\n              readOnly: true\n              description: Output only. The time this CryptoKeyVersion's key material\n                was destroyed. Only present if state is DESTROYED.\n              x-kubernetes-immutable: true\n            destroyTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: DestroyTime\n              readOnly: true\n              description: Output only. The time this CryptoKeyVersion's key material\n                is scheduled for destruction. Only present if state is DESTROY_SCHEDULED.\n              x-kubernetes-immutable: true\n            externalProtectionLevelOptions:\n              type: object\n              x-dcl-go-name: ExternalProtectionLevelOptions\n              x-dcl-go-type: CryptoKeyPrimaryExternalProtectionLevelOptions\n              description: ExternalProtectionLevelOptions stores a group of additional\n                fields for configuring a CryptoKeyVersion that are specific to the\n                EXTERNAL protection level.\n              properties:\n                externalKeyUri:\n                  type: string\n                  x-dcl-go-name: ExternalKeyUri\n                  description: The URI for an external resource that this CryptoKeyVersion\n                    represents.\n            generateTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: GenerateTime\n              readOnly: true\n              description: Output only. The time this CryptoKeyVersion's key material\n                was generated.\n              x-kubernetes-immutable: true\n            importFailureReason:\n              type: string\n              x-dcl-go-name: ImportFailureReason\n              readOnly: true\n              description: Output only. The root cause of the most recent import failure.\n                Only present if state is IMPORT_FAILED.\n              x-kubernetes-immutable: true\n            importJob:\n              type: string\n              x-dcl-go-name: ImportJob\n              readOnly: true\n              description: Output only. The name of the ImportJob used in the most\n                recent import of this CryptoKeyVersion. Only present if the underlying\n                key material was imported.\n              x-kubernetes-immutable: true\n            importTime:\n              type: string\n              format: date-time\n              x-dcl-go-name: ImportTime\n              readOnly: true\n              description: Output only. The time at which this CryptoKeyVersion's\n                key material was most recently imported.\n              x-kubernetes-immutable: true\n            name:\n              type: string\n              x-dcl-go-name: Name\n              description: Output only. The resource name for this CryptoKeyVersion\n                in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`.\n              x-kubernetes-immutable: true\n              x-dcl-server-generated-parameter: true\n            protectionLevel:\n              type: string\n              x-dcl-go-name: ProtectionLevel\n              x-dcl-go-type: CryptoKeyPrimaryProtectionLevelEnum\n              readOnly: true\n              description: 'Output only. The ProtectionLevel describing how crypto\n                operations are performed with this CryptoKeyVersion. Possible values:\n                PROTECTION_LEVEL_UNSPECIFIED, SOFTWARE, HSM, EXTERNAL, EXTERNAL_VPC'\n              x-kubernetes-immutable: true\n              enum:\n              - PROTECTION_LEVEL_UNSPECIFIED\n              - SOFTWARE\n              - HSM\n              - EXTERNAL\n              - EXTERNAL_VPC\n            reimportEligible:\n              type: boolean\n              x-dcl-go-name: ReimportEligible\n              readOnly: true\n              description: Output only. Whether or not this key version is eligible\n                for reimport, by being specified as a target in ImportCryptoKeyVersionRequest.crypto_key_version.\n              x-kubernetes-immutable: true\n            state:\n              type: string\n              x-dcl-go-name: State\n              x-dcl-go-type: CryptoKeyPrimaryStateEnum\n              description: 'The current state of the CryptoKeyVersion. Possible values:\n                CRYPTO_KEY_VERSION_STATE_UNSPECIFIED, PENDING_GENERATION, ENABLED,\n                DISABLED, DESTROYED, DESTROY_SCHEDULED, PENDING_IMPORT, IMPORT_FAILED'\n              enum:\n              - CRYPTO_KEY_VERSION_STATE_UNSPECIFIED\n              - PENDING_GENERATION\n              - ENABLED\n              - DISABLED\n              - DESTROYED\n              - DESTROY_SCHEDULED\n              - PENDING_IMPORT\n              - IMPORT_FAILED\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n          x-dcl-parameter: true\n        purpose:\n          type: string\n          x-dcl-go-name: Purpose\n          x-dcl-go-type: CryptoKeyPurposeEnum\n          description: 'Immutable. The immutable purpose of this CryptoKey. Possible\n            values: CRYPTO_KEY_PURPOSE_UNSPECIFIED, ENCRYPT_DECRYPT, ASYMMETRIC_SIGN,\n            ASYMMETRIC_DECRYPT, MAC'\n          x-kubernetes-immutable: true\n          enum:\n          - CRYPTO_KEY_PURPOSE_UNSPECIFIED\n          - ENCRYPT_DECRYPT\n          - ASYMMETRIC_SIGN\n          - ASYMMETRIC_DECRYPT\n          - MAC\n        rotationPeriod:\n          type: string\n          x-dcl-go-name: RotationPeriod\n          description: next_rotation_time will be advanced by this period when the\n            service automatically rotates a key. Must be at least 24 hours and at\n            most 876,000 hours. If rotation_period is set, next_rotation_time must\n            also be set. Keys with purpose ENCRYPT_DECRYPT support automatic rotation.\n            For other keys, this field must be omitted.\n        versionTemplate:\n          type: object\n          x-dcl-go-name: VersionTemplate\n          x-dcl-go-type: CryptoKeyVersionTemplate\n          description: A template describing settings for new CryptoKeyVersion instances.\n            The properties of new CryptoKeyVersion instances created by either CreateCryptoKeyVersion\n            or auto-rotation are controlled by this template.\n          x-dcl-server-default: true\n          required:\n          - algorithm\n          properties:\n            algorithm:\n              type: string\n              x-dcl-go-name: Algorithm\n              x-dcl-go-type: CryptoKeyVersionTemplateAlgorithmEnum\n              description: 'Required. Algorithm to use when creating a CryptoKeyVersion\n                based on this template. For backwards compatibility, GOOGLE_SYMMETRIC_ENCRYPTION\n                is implied if both this field is omitted and CryptoKey.purpose is\n                ENCRYPT_DECRYPT. Possible values: CRYPTO_KEY_VERSION_ALGORITHM_UNSPECIFIED,\n                GOOGLE_SYMMETRIC_ENCRYPTION, RSA_SIGN_PSS_2048_SHA256, RSA_SIGN_PSS_3072_SHA256,\n                RSA_SIGN_PSS_4096_SHA256, RSA_SIGN_PSS_4096_SHA512, RSA_SIGN_PKCS1_2048_SHA256,\n                RSA_SIGN_PKCS1_3072_SHA256, RSA_SIGN_PKCS1_4096_SHA256, RSA_SIGN_PKCS1_4096_SHA512,\n                RSA_DECRYPT_OAEP_2048_SHA256, RSA_DECRYPT_OAEP_3072_SHA256, RSA_DECRYPT_OAEP_4096_SHA256,\n                RSA_DECRYPT_OAEP_4096_SHA512, EC_SIGN_P256_SHA256, EC_SIGN_P384_SHA384,\n                EC_SIGN_SECP256K1_SHA256, HMAC_SHA256, EXTERNAL_SYMMETRIC_ENCRYPTION'\n              enum:\n              - CRYPTO_KEY_VERSION_ALGORITHM_UNSPECIFIED\n              - GOOGLE_SYMMETRIC_ENCRYPTION\n              - RSA_SIGN_PSS_2048_SHA256\n              - RSA_SIGN_PSS_3072_SHA256\n              - RSA_SIGN_PSS_4096_SHA256\n              - RSA_SIGN_PSS_4096_SHA512\n              - RSA_SIGN_PKCS1_2048_SHA256\n              - RSA_SIGN_PKCS1_3072_SHA256\n              - RSA_SIGN_PKCS1_4096_SHA256\n              - RSA_SIGN_PKCS1_4096_SHA512\n              - RSA_DECRYPT_OAEP_2048_SHA256\n              - RSA_DECRYPT_OAEP_3072_SHA256\n              - RSA_DECRYPT_OAEP_4096_SHA256\n              - RSA_DECRYPT_OAEP_4096_SHA512\n              - EC_SIGN_P256_SHA256\n              - EC_SIGN_P384_SHA384\n              - EC_SIGN_SECP256K1_SHA256\n              - HMAC_SHA256\n              - EXTERNAL_SYMMETRIC_ENCRYPTION\n            protectionLevel:\n              type: string\n              x-dcl-go-name: ProtectionLevel\n              x-dcl-go-type: CryptoKeyVersionTemplateProtectionLevelEnum\n              description: 'ProtectionLevel to use when creating a CryptoKeyVersion\n                based on this template. Immutable. Defaults to SOFTWARE. Possible\n                values: PROTECTION_LEVEL_UNSPECIFIED, SOFTWARE, HSM, EXTERNAL, EXTERNAL_VPC'\n              x-kubernetes-immutable: true\n              x-dcl-server-default: true\n              enum:\n              - PROTECTION_LEVEL_UNSPECIFIED\n              - SOFTWARE\n              - HSM\n              - EXTERNAL\n              - EXTERNAL_VPC\n")

// 18927 bytes
// MD5: 35b4ad2993ebb9a2f19394ee80a10ef3
