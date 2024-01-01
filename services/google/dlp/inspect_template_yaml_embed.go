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
// gen_go_data -package dlp -var YAML_inspect_template blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/dlp/inspect_template.yaml

package dlp

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/dlp/inspect_template.yaml
var YAML_inspect_template = []byte("info:\n  title: Dlp/InspectTemplate\n  description: The Dlp InspectTemplate resource\n  x-dcl-struct-name: InspectTemplate\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a InspectTemplate\n    parameters:\n    - name: inspectTemplate\n      required: true\n      description: A full instance of a InspectTemplate\n  apply:\n    description: The function used to apply information about a InspectTemplate\n    parameters:\n    - name: inspectTemplate\n      required: true\n      description: A full instance of a InspectTemplate\n  delete:\n    description: The function used to delete a InspectTemplate\n    parameters:\n    - name: inspectTemplate\n      required: true\n      description: A full instance of a InspectTemplate\n  deleteAll:\n    description: The function used to delete all InspectTemplate\n    parameters:\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: parent\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many InspectTemplate\n    parameters:\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: parent\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    InspectTemplate:\n      title: InspectTemplate\n      x-dcl-id: '{{parent}}/inspectTemplates/{{name}}'\n      x-dcl-locations:\n      - region\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - parent\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The creation timestamp of an inspectTemplate.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Short description (max 256 chars).\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Display name (max 256 chars).\n        inspectConfig:\n          type: object\n          x-dcl-go-name: InspectConfig\n          x-dcl-go-type: InspectTemplateInspectConfig\n          description: The core content of the template. Configuration of the scanning\n            process.\n          properties:\n            contentOptions:\n              type: array\n              x-dcl-go-name: ContentOptions\n              description: List of options defining data content to scan. If empty,\n                text, images, and other content will be included.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: InspectTemplateInspectConfigContentOptionsEnum\n                enum:\n                - CONTENT_UNSPECIFIED\n                - CONTENT_TEXT\n                - CONTENT_IMAGE\n            customInfoTypes:\n              type: array\n              x-dcl-go-name: CustomInfoTypes\n              description: CustomInfoTypes provided by the user. See https://cloud.google.com/dlp/docs/creating-custom-infotypes\n                to learn more.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: InspectTemplateInspectConfigCustomInfoTypes\n                properties:\n                  dictionary:\n                    type: object\n                    x-dcl-go-name: Dictionary\n                    x-dcl-go-type: InspectTemplateInspectConfigCustomInfoTypesDictionary\n                    description: A list of phrases to detect as a CustomInfoType.\n                    x-dcl-conflicts:\n                    - regex\n                    - surrogateType\n                    - storedType\n                    properties:\n                      cloudStoragePath:\n                        type: object\n                        x-dcl-go-name: CloudStoragePath\n                        x-dcl-go-type: InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath\n                        description: Newline-delimited file of words in Cloud Storage.\n                          Only a single file is accepted.\n                        x-dcl-conflicts:\n                        - wordList\n                        properties:\n                          path:\n                            type: string\n                            x-dcl-go-name: Path\n                            description: 'A url representing a file or path (no wildcards)\n                              in Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt'\n                      wordList:\n                        type: object\n                        x-dcl-go-name: WordList\n                        x-dcl-go-type: InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList\n                        description: List of words or phrases to search for.\n                        x-dcl-conflicts:\n                        - cloudStoragePath\n                        properties:\n                          words:\n                            type: array\n                            x-dcl-go-name: Words\n                            description: Words or phrases defining the dictionary.\n                              The dictionary must contain at least one phrase and\n                              every phrase must contain at least 2 characters that\n                              are letters or digits. [required]\n                            x-dcl-send-empty: true\n                            x-dcl-list-type: list\n                            items:\n                              type: string\n                              x-dcl-go-type: string\n                  exclusionType:\n                    type: string\n                    x-dcl-go-name: ExclusionType\n                    x-dcl-go-type: InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum\n                    description: 'If set to EXCLUSION_TYPE_EXCLUDE this infoType will\n                      not cause a finding to be returned. It still can be used for\n                      rules matching. Possible values: EXCLUSION_TYPE_UNSPECIFIED,\n                      EXCLUSION_TYPE_EXCLUDE'\n                    enum:\n                    - EXCLUSION_TYPE_UNSPECIFIED\n                    - EXCLUSION_TYPE_EXCLUDE\n                  infoType:\n                    type: object\n                    x-dcl-go-name: InfoType\n                    x-dcl-go-type: InspectTemplateInspectConfigCustomInfoTypesInfoType\n                    description: CustomInfoType can either be a new infoType, or an\n                      extension of built-in infoType, when the name matches one of\n                      existing infoTypes and that infoType is specified in `InspectContent.info_types`\n                      field. Specifying the latter adds findings to the one detected\n                      by the system. If built-in info type is not specified in `InspectContent.info_types`\n                      list then the name is treated as a custom info type.\n                    properties:\n                      name:\n                        type: string\n                        x-dcl-go-name: Name\n                        description: Name of the information type. Either a name of\n                          your choosing when creating a CustomInfoType, or one of\n                          the names listed at https://cloud.google.com/dlp/docs/infotypes-reference\n                          when specifying a built-in type. When sending Cloud DLP\n                          results to Data Catalog, infoType names should conform to\n                          the pattern `[A-Za-z0-9$-_]{1,64}`.\n                  likelihood:\n                    type: string\n                    x-dcl-go-name: Likelihood\n                    x-dcl-go-type: InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum\n                    description: 'Likelihood to return for this CustomInfoType. This\n                      base value can be altered by a detection rule if the finding\n                      meets the criteria specified by the rule. Defaults to `VERY_LIKELY`\n                      if not specified. Possible values: LIKELIHOOD_UNSPECIFIED, VERY_UNLIKELY,\n                      UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY'\n                    enum:\n                    - LIKELIHOOD_UNSPECIFIED\n                    - VERY_UNLIKELY\n                    - UNLIKELY\n                    - POSSIBLE\n                    - LIKELY\n                    - VERY_LIKELY\n                  regex:\n                    type: object\n                    x-dcl-go-name: Regex\n                    x-dcl-go-type: InspectTemplateInspectConfigCustomInfoTypesRegex\n                    description: Regular expression based CustomInfoType.\n                    x-dcl-conflicts:\n                    - dictionary\n                    - surrogateType\n                    - storedType\n                    properties:\n                      groupIndexes:\n                        type: array\n                        x-dcl-go-name: GroupIndexes\n                        description: The index of the submatch to extract as findings.\n                          When not specified, the entire match is returned. No more\n                          than 3 may be included.\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: integer\n                          format: int64\n                          x-dcl-go-type: int64\n                      pattern:\n                        type: string\n                        x-dcl-go-name: Pattern\n                        description: Pattern defining the regular expression. Its\n                          syntax (https://github.com/google/re2/wiki/Syntax) can be\n                          found under the google/re2 repository on GitHub.\n                  storedType:\n                    type: object\n                    x-dcl-go-name: StoredType\n                    x-dcl-go-type: InspectTemplateInspectConfigCustomInfoTypesStoredType\n                    description: Load an existing `StoredInfoType` resource for use\n                      in `InspectDataSource`. Not currently supported in `InspectContent`.\n                    x-dcl-conflicts:\n                    - dictionary\n                    - regex\n                    - surrogateType\n                    properties:\n                      createTime:\n                        type: string\n                        format: date-time\n                        x-dcl-go-name: CreateTime\n                        readOnly: true\n                        description: Timestamp indicating when the version of the\n                          `StoredInfoType` used for inspection was created. Output-only\n                          field, populated by the system.\n                      name:\n                        type: string\n                        x-dcl-go-name: Name\n                        description: Resource name of the requested `StoredInfoType`,\n                          for example `organizations/433245324/storedInfoTypes/432452342`\n                          or `projects/project-id/storedInfoTypes/432452342`.\n                        x-dcl-references:\n                        - resource: Dlp/StoredInfoType\n                          field: name\n                          parent: true\n                  surrogateType:\n                    type: object\n                    x-dcl-go-name: SurrogateType\n                    x-dcl-go-type: InspectTemplateInspectConfigCustomInfoTypesSurrogateType\n                    description: Message for detecting output from deidentification\n                      transformations that support reversing.\n                    x-dcl-conflicts:\n                    - dictionary\n                    - regex\n                    - storedType\n            excludeInfoTypes:\n              type: boolean\n              x-dcl-go-name: ExcludeInfoTypes\n              description: When true, excludes type information of the findings.\n            includeQuote:\n              type: boolean\n              x-dcl-go-name: IncludeQuote\n              description: When true, a contextual quote from the data that triggered\n                a finding is included in the response; see Finding.quote.\n            infoTypes:\n              type: array\n              x-dcl-go-name: InfoTypes\n              description: Restricts what info_types to look for. The values must\n                correspond to InfoType values returned by ListInfoTypes or listed\n                at https://cloud.google.com/dlp/docs/infotypes-reference. When no\n                InfoTypes or CustomInfoTypes are specified in a request, the system\n                may automatically choose what detectors to run. By default this may\n                be all types, but may change over time as detectors are updated. If\n                you need precise control and predictability as to what detectors are\n                run you should specify specific InfoTypes listed in the reference,\n                otherwise a default list will be used, which may change over time.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: InspectTemplateInspectConfigInfoTypes\n                properties:\n                  name:\n                    type: string\n                    x-dcl-go-name: Name\n                    description: Name of the information type. Either a name of your\n                      choosing when creating a CustomInfoType, or one of the names\n                      listed at https://cloud.google.com/dlp/docs/infotypes-reference\n                      when specifying a built-in type. When sending Cloud DLP results\n                      to Data Catalog, infoType names should conform to the pattern\n                      `[A-Za-z0-9$-_]{1,64}`.\n            limits:\n              type: object\n              x-dcl-go-name: Limits\n              x-dcl-go-type: InspectTemplateInspectConfigLimits\n              description: Configuration to control the number of findings returned.\n              properties:\n                maxFindingsPerInfoType:\n                  type: array\n                  x-dcl-go-name: MaxFindingsPerInfoType\n                  description: Configuration of findings limit given for specified\n                    infoTypes.\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: object\n                    x-dcl-go-type: InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType\n                    properties:\n                      infoType:\n                        type: object\n                        x-dcl-go-name: InfoType\n                        x-dcl-go-type: InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType\n                        description: Type of information the findings limit applies\n                          to. Only one limit per info_type should be provided. If\n                          InfoTypeLimit does not have an info_type, the DLP API applies\n                          the limit against all info_types that are found but not\n                          specified in another InfoTypeLimit.\n                        properties:\n                          name:\n                            type: string\n                            x-dcl-go-name: Name\n                            description: Name of the information type. Either a name\n                              of your choosing when creating a CustomInfoType, or\n                              one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference\n                              when specifying a built-in type. When sending Cloud\n                              DLP results to Data Catalog, infoType names should conform\n                              to the pattern `[A-Za-z0-9$-_]{1,64}`.\n                      maxFindings:\n                        type: integer\n                        format: int64\n                        x-dcl-go-name: MaxFindings\n                        description: Max findings limit for the given infoType.\n                maxFindingsPerItem:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: MaxFindingsPerItem\n                  description: Max number of findings that will be returned for each\n                    item scanned. When set within `InspectJobConfig`, the maximum\n                    returned is 2000 regardless if this is set higher. When set within\n                    `InspectContentRequest`, this field is ignored.\n                maxFindingsPerRequest:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: MaxFindingsPerRequest\n                  description: Max number of findings that will be returned per request/job.\n                    When set within `InspectContentRequest`, the maximum returned\n                    is 2000 regardless if this is set higher.\n            minLikelihood:\n              type: string\n              x-dcl-go-name: MinLikelihood\n              x-dcl-go-type: InspectTemplateInspectConfigMinLikelihoodEnum\n              description: 'Only returns findings equal or above this threshold. The\n                default is POSSIBLE. See https://cloud.google.com/dlp/docs/likelihood\n                to learn more. Possible values: LIKELIHOOD_UNSPECIFIED, VERY_UNLIKELY,\n                UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY'\n              enum:\n              - LIKELIHOOD_UNSPECIFIED\n              - VERY_UNLIKELY\n              - UNLIKELY\n              - POSSIBLE\n              - LIKELY\n              - VERY_LIKELY\n            ruleSet:\n              type: array\n              x-dcl-go-name: RuleSet\n              description: Set of rules to apply to the findings for this InspectConfig.\n                Exclusion rules, contained in the set are executed in the end, other\n                rules are executed in the order they are specified for each info type.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: InspectTemplateInspectConfigRuleSet\n                properties:\n                  infoTypes:\n                    type: array\n                    x-dcl-go-name: InfoTypes\n                    description: List of infoTypes this rule set is applied to.\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: object\n                      x-dcl-go-type: InspectTemplateInspectConfigRuleSetInfoTypes\n                      properties:\n                        name:\n                          type: string\n                          x-dcl-go-name: Name\n                          description: Name of the information type. Either a name\n                            of your choosing when creating a CustomInfoType, or one\n                            of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference\n                            when specifying a built-in type. When sending Cloud DLP\n                            results to Data Catalog, infoType names should conform\n                            to the pattern `[A-Za-z0-9$-_]{1,64}`.\n                  rules:\n                    type: array\n                    x-dcl-go-name: Rules\n                    description: Set of rules to be applied to infoTypes. The rules\n                      are applied in order.\n                    x-dcl-send-empty: true\n                    x-dcl-list-type: list\n                    items:\n                      type: object\n                      x-dcl-go-type: InspectTemplateInspectConfigRuleSetRules\n                      properties:\n                        exclusionRule:\n                          type: object\n                          x-dcl-go-name: ExclusionRule\n                          x-dcl-go-type: InspectTemplateInspectConfigRuleSetRulesExclusionRule\n                          description: Exclusion rule.\n                          x-dcl-conflicts:\n                          - hotwordRule\n                          properties:\n                            dictionary:\n                              type: object\n                              x-dcl-go-name: Dictionary\n                              x-dcl-go-type: InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary\n                              description: Dictionary which defines the rule.\n                              x-dcl-conflicts:\n                              - regex\n                              - excludeInfoTypes\n                              properties:\n                                cloudStoragePath:\n                                  type: object\n                                  x-dcl-go-name: CloudStoragePath\n                                  x-dcl-go-type: InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath\n                                  description: Newline-delimited file of words in\n                                    Cloud Storage. Only a single file is accepted.\n                                  x-dcl-conflicts:\n                                  - wordList\n                                  properties:\n                                    path:\n                                      type: string\n                                      x-dcl-go-name: Path\n                                      description: 'A url representing a file or path\n                                        (no wildcards) in Cloud Storage. Example:\n                                        gs://[BUCKET_NAME]/dictionary.txt'\n                                wordList:\n                                  type: object\n                                  x-dcl-go-name: WordList\n                                  x-dcl-go-type: InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList\n                                  description: List of words or phrases to search\n                                    for.\n                                  x-dcl-conflicts:\n                                  - cloudStoragePath\n                                  properties:\n                                    words:\n                                      type: array\n                                      x-dcl-go-name: Words\n                                      description: Words or phrases defining the dictionary.\n                                        The dictionary must contain at least one phrase\n                                        and every phrase must contain at least 2 characters\n                                        that are letters or digits. [required]\n                                      x-dcl-send-empty: true\n                                      x-dcl-list-type: list\n                                      items:\n                                        type: string\n                                        x-dcl-go-type: string\n                            excludeInfoTypes:\n                              type: object\n                              x-dcl-go-name: ExcludeInfoTypes\n                              x-dcl-go-type: InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes\n                              description: Set of infoTypes for which findings would\n                                affect this rule.\n                              x-dcl-conflicts:\n                              - dictionary\n                              - regex\n                              properties:\n                                infoTypes:\n                                  type: array\n                                  x-dcl-go-name: InfoTypes\n                                  description: InfoType list in ExclusionRule rule\n                                    drops a finding when it overlaps or contained\n                                    within with a finding of an infoType from this\n                                    list. For example, for `InspectionRuleSet.info_types`\n                                    containing \"PHONE_NUMBER\"` and `exclusion_rule`\n                                    containing `exclude_info_types.info_types` with\n                                    \"EMAIL_ADDRESS\" the phone number findings are\n                                    dropped if they overlap with EMAIL_ADDRESS finding.\n                                    That leads to \"555-222-2222@example.org\" to generate\n                                    only a single finding, namely email address.\n                                  x-dcl-send-empty: true\n                                  x-dcl-list-type: list\n                                  items:\n                                    type: object\n                                    x-dcl-go-type: InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes\n                                    properties:\n                                      name:\n                                        type: string\n                                        x-dcl-go-name: Name\n                                        description: Name of the information type.\n                                          Either a name of your choosing when creating\n                                          a CustomInfoType, or one of the names listed\n                                          at https://cloud.google.com/dlp/docs/infotypes-reference\n                                          when specifying a built-in type. When sending\n                                          Cloud DLP results to Data Catalog, infoType\n                                          names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.\n                            matchingType:\n                              type: string\n                              x-dcl-go-name: MatchingType\n                              x-dcl-go-type: InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum\n                              description: 'How the rule is applied, see MatchingType\n                                documentation for details. Possible values: MATCHING_TYPE_UNSPECIFIED,\n                                MATCHING_TYPE_FULL_MATCH, MATCHING_TYPE_PARTIAL_MATCH,\n                                MATCHING_TYPE_INVERSE_MATCH'\n                              enum:\n                              - MATCHING_TYPE_UNSPECIFIED\n                              - MATCHING_TYPE_FULL_MATCH\n                              - MATCHING_TYPE_PARTIAL_MATCH\n                              - MATCHING_TYPE_INVERSE_MATCH\n                            regex:\n                              type: object\n                              x-dcl-go-name: Regex\n                              x-dcl-go-type: InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex\n                              description: Regular expression which defines the rule.\n                              x-dcl-conflicts:\n                              - dictionary\n                              - excludeInfoTypes\n                              properties:\n                                groupIndexes:\n                                  type: array\n                                  x-dcl-go-name: GroupIndexes\n                                  description: The index of the submatch to extract\n                                    as findings. When not specified, the entire match\n                                    is returned. No more than 3 may be included.\n                                  x-dcl-send-empty: true\n                                  x-dcl-list-type: list\n                                  items:\n                                    type: integer\n                                    format: int64\n                                    x-dcl-go-type: int64\n                                pattern:\n                                  type: string\n                                  x-dcl-go-name: Pattern\n                                  description: Pattern defining the regular expression.\n                                    Its syntax (https://github.com/google/re2/wiki/Syntax)\n                                    can be found under the google/re2 repository on\n                                    GitHub.\n                        hotwordRule:\n                          type: object\n                          x-dcl-go-name: HotwordRule\n                          x-dcl-go-type: InspectTemplateInspectConfigRuleSetRulesHotwordRule\n                          x-dcl-conflicts:\n                          - exclusionRule\n                          properties:\n                            hotwordRegex:\n                              type: object\n                              x-dcl-go-name: HotwordRegex\n                              x-dcl-go-type: InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex\n                              description: Regular expression pattern defining what\n                                qualifies as a hotword.\n                              properties:\n                                groupIndexes:\n                                  type: array\n                                  x-dcl-go-name: GroupIndexes\n                                  description: The index of the submatch to extract\n                                    as findings. When not specified, the entire match\n                                    is returned. No more than 3 may be included.\n                                  x-dcl-send-empty: true\n                                  x-dcl-list-type: list\n                                  items:\n                                    type: integer\n                                    format: int64\n                                    x-dcl-go-type: int64\n                                pattern:\n                                  type: string\n                                  x-dcl-go-name: Pattern\n                                  description: Pattern defining the regular expression.\n                                    Its syntax (https://github.com/google/re2/wiki/Syntax)\n                                    can be found under the google/re2 repository on\n                                    GitHub.\n                            likelihoodAdjustment:\n                              type: object\n                              x-dcl-go-name: LikelihoodAdjustment\n                              x-dcl-go-type: InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment\n                              description: Likelihood adjustment to apply to all matching\n                                findings.\n                              properties:\n                                fixedLikelihood:\n                                  type: string\n                                  x-dcl-go-name: FixedLikelihood\n                                  x-dcl-go-type: InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum\n                                  description: 'Set the likelihood of a finding to\n                                    a fixed value. Possible values: LIKELIHOOD_UNSPECIFIED,\n                                    VERY_UNLIKELY, UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY'\n                                  x-dcl-conflicts:\n                                  - relativeLikelihood\n                                  enum:\n                                  - LIKELIHOOD_UNSPECIFIED\n                                  - VERY_UNLIKELY\n                                  - UNLIKELY\n                                  - POSSIBLE\n                                  - LIKELY\n                                  - VERY_LIKELY\n                                relativeLikelihood:\n                                  type: integer\n                                  format: int64\n                                  x-dcl-go-name: RelativeLikelihood\n                                  description: Increase or decrease the likelihood\n                                    by the specified number of levels. For example,\n                                    if a finding would be `POSSIBLE` without the detection\n                                    rule and `relative_likelihood` is 1, then it is\n                                    upgraded to `LIKELY`, while a value of -1 would\n                                    downgrade it to `UNLIKELY`. Likelihood may never\n                                    drop below `VERY_UNLIKELY` or exceed `VERY_LIKELY`,\n                                    so applying an adjustment of 1 followed by an\n                                    adjustment of -1 when base likelihood is `VERY_LIKELY`\n                                    will result in a final likelihood of `LIKELY`.\n                                  x-dcl-conflicts:\n                                  - fixedLikelihood\n                            proximity:\n                              type: object\n                              x-dcl-go-name: Proximity\n                              x-dcl-go-type: InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity\n                              description: Proximity of the finding within which the\n                                entire hotword must reside. The total length of the\n                                window cannot exceed 1000 characters. Note that the\n                                finding itself will be included in the window, so\n                                that hotwords may be used to match substrings of the\n                                finding itself. For example, the certainty of a phone\n                                number regex \"(d{3}) d{3}-d{4}\" could be adjusted\n                                upwards if the area code is known to be the local\n                                area code of a company office using the hotword regex\n                                \"(xxx)\", where \"xxx\" is the area code in question.\n                              properties:\n                                windowAfter:\n                                  type: integer\n                                  format: int64\n                                  x-dcl-go-name: WindowAfter\n                                  description: Number of characters after the finding\n                                    to consider.\n                                windowBefore:\n                                  type: integer\n                                  format: int64\n                                  x-dcl-go-name: WindowBefore\n                                  description: Number of characters before the finding\n                                    to consider.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location of the resource\n          x-kubernetes-immutable: true\n        locationId:\n          type: string\n          x-dcl-go-name: LocationId\n          readOnly: true\n          description: Output only. The geographic location where this resource is\n            stored.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'The template name. The template will have one of the following\n            formats: `projects/PROJECT_ID/inspectTemplates/TEMPLATE_ID` OR `organizations/ORGANIZATION_ID/inspectTemplates/TEMPLATE_ID`;'\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        parent:\n          type: string\n          x-dcl-go-name: Parent\n          description: The parent of the resource\n          x-kubernetes-immutable: true\n          x-dcl-forward-slash-allowed: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Organization\n            field: name\n            parent: true\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The last update timestamp of an inspectTemplate.\n          x-kubernetes-immutable: true\n")

// 36418 bytes
// MD5: 35543edbc707a9256f3ceddde9c24d47
