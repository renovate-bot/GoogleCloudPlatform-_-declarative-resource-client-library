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
package alpha

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLFeatureSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "GkeHub/Feature",
			Description: "The GkeHub Feature resource",
			StructName:  "Feature",
			Mutex:       "{{project}}/{{location}}/{{feature}}",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Feature",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "feature",
						Required:    true,
						Description: "A full instance of a Feature",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Feature",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "feature",
						Required:    true,
						Description: "A full instance of a Feature",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Feature",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "feature",
						Required:    true,
						Description: "A full instance of a Feature",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Feature",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many Feature",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"Feature": &dcl.Component{
					Title:           "Feature",
					ID:              "projects/{{project}}/locations/{{location}}/features/{{name}}",
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. When the Feature resource was created.",
								Immutable:   true,
							},
							"deleteTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "DeleteTime",
								ReadOnly:    true,
								Description: "Output only. When the Feature resource was deleted.",
								Immutable:   true,
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "GCP labels for this Feature.",
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
								Parameter:   true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "The full, unique name of this Feature resource",
								Immutable:   true,
								HasLongForm: true,
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "The project for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
								Parameter: true,
							},
							"resourceState": &dcl.Property{
								Type:        "object",
								GoName:      "ResourceState",
								GoType:      "FeatureResourceState",
								ReadOnly:    true,
								Description: "State of the Feature resource itself.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"hasResources": &dcl.Property{
										Type:        "boolean",
										GoName:      "HasResources",
										ReadOnly:    true,
										Description: "Whether this Feature has outstanding resources that need to be cleaned up before it can be disabled.",
										Immutable:   true,
									},
									"state": &dcl.Property{
										Type:        "string",
										GoName:      "State",
										GoType:      "FeatureResourceStateStateEnum",
										ReadOnly:    true,
										Description: "The current state of the Feature resource in the Hub API. Possible values: STATE_UNSPECIFIED, ENABLING, ACTIVE, DISABLING, UPDATING, SERVICE_UPDATING",
										Immutable:   true,
										Enum: []string{
											"STATE_UNSPECIFIED",
											"ENABLING",
											"ACTIVE",
											"DISABLING",
											"UPDATING",
											"SERVICE_UPDATING",
										},
									},
								},
							},
							"spec": &dcl.Property{
								Type:        "object",
								GoName:      "Spec",
								GoType:      "FeatureSpec",
								Description: "Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.",
								Properties: map[string]*dcl.Property{
									"cloudauditlogging": &dcl.Property{
										Type:        "object",
										GoName:      "Cloudauditlogging",
										GoType:      "FeatureSpecCloudauditlogging",
										Description: "Cloud Audit Logging-specific spec.",
										Properties: map[string]*dcl.Property{
											"allowlistedServiceAccounts": &dcl.Property{
												Type:        "array",
												GoName:      "AllowlistedServiceAccounts",
												Description: "Service account that should be allowlisted to send the audit logs; eg cloudauditlogging@gcp-project.iam.gserviceaccount.com. These accounts must already exist, but do not need to have any permissions granted to them. The customer's entitlements will be checked prior to allowlisting (i.e. the customer must be an Anthos customer.)",
												SendEmpty:   true,
												ListType:    "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
										},
									},
									"fleetobservability": &dcl.Property{
										Type:        "object",
										GoName:      "Fleetobservability",
										GoType:      "FeatureSpecFleetobservability",
										Description: "Fleet Observability spec.",
										Properties: map[string]*dcl.Property{
											"loggingConfig": &dcl.Property{
												Type:        "object",
												GoName:      "LoggingConfig",
												GoType:      "FeatureSpecFleetobservabilityLoggingConfig",
												Description: "Fleet Observability Logging-specific spec.",
												Properties: map[string]*dcl.Property{
													"defaultConfig": &dcl.Property{
														Type:        "object",
														GoName:      "DefaultConfig",
														GoType:      "FeatureSpecFleetobservabilityLoggingConfigDefaultConfig",
														Description: "Specified if applying the default routing config to logs not specified in other configs.",
														Properties: map[string]*dcl.Property{
															"mode": &dcl.Property{
																Type:        "string",
																GoName:      "Mode",
																GoType:      "FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum",
																Description: "The logs routing mode Possible values: MODE_UNSPECIFIED, COPY, MOVE",
																Enum: []string{
																	"MODE_UNSPECIFIED",
																	"COPY",
																	"MOVE",
																},
															},
														},
													},
													"fleetScopeLogsConfig": &dcl.Property{
														Type:        "object",
														GoName:      "FleetScopeLogsConfig",
														GoType:      "FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig",
														Description: "Specified if applying the routing config to all logs for all fleet scopes.",
														Properties: map[string]*dcl.Property{
															"mode": &dcl.Property{
																Type:        "string",
																GoName:      "Mode",
																GoType:      "FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum",
																Description: "The logs routing mode Possible values: MODE_UNSPECIFIED, COPY, MOVE",
																Enum: []string{
																	"MODE_UNSPECIFIED",
																	"COPY",
																	"MOVE",
																},
															},
														},
													},
												},
											},
										},
									},
									"multiclusteringress": &dcl.Property{
										Type:        "object",
										GoName:      "Multiclusteringress",
										GoType:      "FeatureSpecMulticlusteringress",
										Description: "Multicluster Ingress-specific spec.",
										Required: []string{
											"configMembership",
										},
										Properties: map[string]*dcl.Property{
											"configMembership": &dcl.Property{
												Type:        "string",
												GoName:      "ConfigMembership",
												Description: "Fully-qualified Membership name which hosts the MultiClusterIngress CRD. Example: `projects/foo-proj/locations/global/memberships/bar`",
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Gkehub/Membership",
														Field:    "name",
													},
												},
											},
										},
									},
								},
							},
							"state": &dcl.Property{
								Type:        "object",
								GoName:      "State",
								GoType:      "FeatureState",
								ReadOnly:    true,
								Description: "Output only. The Hub-wide Feature state",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"servicemesh": &dcl.Property{
										Type:        "object",
										GoName:      "Servicemesh",
										GoType:      "FeatureStateServicemesh",
										ReadOnly:    true,
										Description: "Service Mesh-specific state.",
										Immutable:   true,
										Properties: map[string]*dcl.Property{
											"analysisMessages": &dcl.Property{
												Type:        "array",
												GoName:      "AnalysisMessages",
												ReadOnly:    true,
												Description: "Output only. Results of running Service Mesh analyzers.",
												Immutable:   true,
												ListType:    "list",
												Items: &dcl.Property{
													Type:   "object",
													GoType: "FeatureStateServicemeshAnalysisMessages",
													Properties: map[string]*dcl.Property{
														"args": &dcl.Property{
															Type: "object",
															AdditionalProperties: &dcl.Property{
																Type: "string",
															},
															GoName:      "Args",
															ReadOnly:    true,
															Description: "A UI can combine these args with a template (based on messageBase.type) to produce an internationalized message.",
															Immutable:   true,
														},
														"description": &dcl.Property{
															Type:        "string",
															GoName:      "Description",
															ReadOnly:    true,
															Description: "A human readable description of what the error means. It is suitable for non-internationalize display purposes.",
															Immutable:   true,
														},
														"messageBase": &dcl.Property{
															Type:        "object",
															GoName:      "MessageBase",
															GoType:      "FeatureStateServicemeshAnalysisMessagesMessageBase",
															ReadOnly:    true,
															Description: "Details common to all types of Istio and ServiceMesh analysis messages.",
															Immutable:   true,
															Properties: map[string]*dcl.Property{
																"documentationUrl": &dcl.Property{
																	Type:        "string",
																	GoName:      "DocumentationUrl",
																	ReadOnly:    true,
																	Description: "A url pointing to the Service Mesh or Istio documentation for this specific error type.",
																	Immutable:   true,
																},
																"level": &dcl.Property{
																	Type:        "string",
																	GoName:      "Level",
																	GoType:      "FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum",
																	ReadOnly:    true,
																	Description: "Represents how severe a message is. Possible values: LEVEL_UNSPECIFIED, ERROR, WARNING, INFO",
																	Immutable:   true,
																	Enum: []string{
																		"LEVEL_UNSPECIFIED",
																		"ERROR",
																		"WARNING",
																		"INFO",
																	},
																},
																"type": &dcl.Property{
																	Type:        "object",
																	GoName:      "Type",
																	GoType:      "FeatureStateServicemeshAnalysisMessagesMessageBaseType",
																	ReadOnly:    true,
																	Description: "Represents the specific type of a message.",
																	Immutable:   true,
																	Properties: map[string]*dcl.Property{
																		"code": &dcl.Property{
																			Type:        "string",
																			GoName:      "Code",
																			ReadOnly:    true,
																			Description: "A 7 character code matching ^IST[0-9]{4}$ or ^ASM[0-9]{4}$, intended to uniquely identify the message type. (e.g. \"IST0001\" is mapped to the \"InternalError\" message type.)",
																			Immutable:   true,
																		},
																		"displayName": &dcl.Property{
																			Type:        "string",
																			GoName:      "DisplayName",
																			ReadOnly:    true,
																			Description: "A human-readable name for the message type. e.g. \"InternalError\", \"PodMissingProxy\". This should be the same for all messages of the same type. (This corresponds to the name field in open-source Istio.)",
																			Immutable:   true,
																		},
																	},
																},
															},
														},
														"resourcePaths": &dcl.Property{
															Type:        "array",
															GoName:      "ResourcePaths",
															ReadOnly:    true,
															Description: "A list of strings specifying the resource identifiers that were the cause of message generation. A \"path\" here may be: * MEMBERSHIP_ID if the cause is a specific member cluster * MEMBERSHIP_ID/(NAMESPACE/)?RESOURCETYPE/NAME if the cause is a resource in a cluster",
															Immutable:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "string",
																GoType: "string",
															},
														},
													},
												},
											},
										},
									},
									"state": &dcl.Property{
										Type:        "object",
										GoName:      "State",
										GoType:      "FeatureStateState",
										ReadOnly:    true,
										Description: "Output only. The \"running state\" of the Feature in this Hub.",
										Immutable:   true,
										Properties: map[string]*dcl.Property{
											"code": &dcl.Property{
												Type:        "string",
												GoName:      "Code",
												GoType:      "FeatureStateStateCodeEnum",
												ReadOnly:    true,
												Description: "The high-level, machine-readable status of this Feature. Possible values: CODE_UNSPECIFIED, OK, WARNING, ERROR",
												Immutable:   true,
												Enum: []string{
													"CODE_UNSPECIFIED",
													"OK",
													"WARNING",
													"ERROR",
												},
											},
											"description": &dcl.Property{
												Type:        "string",
												GoName:      "Description",
												ReadOnly:    true,
												Description: "A human-readable description of the current status.",
												Immutable:   true,
											},
											"updateTime": &dcl.Property{
												Type:        "string",
												GoName:      "UpdateTime",
												ReadOnly:    true,
												Description: "The time this status and any related Feature-specific details were updated. A timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\"",
												Immutable:   true,
											},
										},
									},
								},
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. When the Feature resource was last updated.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
