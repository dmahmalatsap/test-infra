// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package openapi

import (
	strconf "github.com/gardener/test-infra/pkg/util/strconf"
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.ConfigElement":            schema_pkg_apis_testmachinery_v1beta1_ConfigElement(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.DAGStep":                  schema_pkg_apis_testmachinery_v1beta1_DAGStep(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.LocationSet":              schema_pkg_apis_testmachinery_v1beta1_LocationSet(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.Pause":                    schema_pkg_apis_testmachinery_v1beta1_Pause(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.StepDefinition":           schema_pkg_apis_testmachinery_v1beta1_StepDefinition(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.StepStatus":               schema_pkg_apis_testmachinery_v1beta1_StepStatus(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.StepStatusPosition":       schema_pkg_apis_testmachinery_v1beta1_StepStatusPosition(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.StepStatusTestDefinition": schema_pkg_apis_testmachinery_v1beta1_StepStatusTestDefinition(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestDefMetadata":          schema_pkg_apis_testmachinery_v1beta1_TestDefMetadata(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestDefSpec":              schema_pkg_apis_testmachinery_v1beta1_TestDefSpec(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestDefinition":           schema_pkg_apis_testmachinery_v1beta1_TestDefinition(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestLocation":             schema_pkg_apis_testmachinery_v1beta1_TestLocation(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.Testrun":                  schema_pkg_apis_testmachinery_v1beta1_Testrun(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestrunKubeconfigs":       schema_pkg_apis_testmachinery_v1beta1_TestrunKubeconfigs(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestrunList":              schema_pkg_apis_testmachinery_v1beta1_TestrunList(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestrunSpec":              schema_pkg_apis_testmachinery_v1beta1_TestrunSpec(ref),
		"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestrunStatus":            schema_pkg_apis_testmachinery_v1beta1_TestrunStatus(ref),
		"github.com/gardener/test-infra/pkg/util/strconf.ConfigSource":                           schema_test_infra_pkg_util_strconf_ConfigSource(ref),
		"github.com/gardener/test-infra/pkg/util/strconf.StringOrConfig":                         schema_test_infra_pkg_util_strconf_StringOrConfig(ref),
	}
}

func schema_pkg_apis_testmachinery_v1beta1_ConfigElement(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ConfigElement is a parameter of a certain type which is passed to TestDefinitions.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of the config value. For now only environment variables are supported.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the environment variable. Must be a C_IDENTIFIER.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"private": {
						SchemaProps: spec.SchemaProps{
							Description: "Private indicates if the config is shared with further steps",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"value": {
						SchemaProps: spec.SchemaProps{
							Description: "value of the environment variable.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"valueFrom": {
						SchemaProps: spec.SchemaProps{
							Description: "Fetches the value from a secret or configmap on the testmachinery cluster.",
							Ref:         ref("github.com/gardener/test-infra/pkg/util/strconf.ConfigSource"),
						},
					},
					"path": {
						SchemaProps: spec.SchemaProps{
							Description: "Only for type=file. Path where the file should be mounted.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type", "name"},
			},
		},
		Dependencies: []string{
			"github.com/gardener/test-infra/pkg/util/strconf.ConfigSource"},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_DAGStep(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"definition": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.StepDefinition"),
						},
					},
					"useGlobalArtifacts": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"dependsOn": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"artifactsFrom": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"pause": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.Pause"),
						},
					},
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.Pause", "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.StepDefinition"},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_LocationSet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "LocationSet defines a set of locations with a specific name and a flag marking the set as the default set.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Unique name of the set.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"default": {
						SchemaProps: spec.SchemaProps{
							Description: "default defines this location set as the default location set to search for TestDefinitions. Only one default location per Testrun is possible.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"locations": {
						SchemaProps: spec.SchemaProps{
							Description: "Locations defines all Locations corresponding to the set.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestLocation"),
									},
								},
							},
						},
					},
				},
				Required: []string{"name", "locations"},
			},
		},
		Dependencies: []string{
			"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestLocation"},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_Pause(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Description: "pauses before this step is executed",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"resumeTimeoutSeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "Resumes the workflow after specified time if it is not manually resumed",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_StepDefinition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StepDefinition is a reference to one or more TestDefinitions to execute in a series of steps.StepDefinition",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"label": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"continueOnError": {
						SchemaProps: spec.SchemaProps{
							Description: "Continue the execution of the workflow even when the step errors or fails.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"condition": {
						SchemaProps: spec.SchemaProps{
							Description: "Condition when the step should be executed. Only used if the step is in the onExit testflow.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"config": {
						SchemaProps: spec.SchemaProps{
							Description: "Step specific configuration.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.ConfigElement"),
									},
								},
							},
						},
					},
					"locationSet": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the configset to look for testDefinitions. If this is empty the default location set is used",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.ConfigElement"},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_StepStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StepStatus is the status of Testflow step",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"position": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.StepStatusPosition"),
						},
					},
					"testdefinition": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.StepStatusTestDefinition"),
						},
					},
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"phase": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"startTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"completionTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"duration": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
					"exportArtifactKey": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"podName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"name", "position", "exportArtifactKey", "podName"},
			},
		},
		Dependencies: []string{
			"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.StepStatusPosition", "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.StepStatusTestDefinition", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_StepStatusPosition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"dependsOn": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"flow": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"step": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_StepStatusTestDefinition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StepStatusTestDefinition holds information about the used testdefinition and its location.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"location": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestLocation"),
						},
					},
					"config": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.ConfigElement"),
									},
								},
							},
						},
					},
					"owner": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"recipientsOnFailure": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"activeDeadlineSeconds": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
				},
				Required: []string{"recipientsOnFailure", "activeDeadlineSeconds"},
			},
		},
		Dependencies: []string{
			"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.ConfigElement", "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestLocation"},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_TestDefMetadata(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TestDefMetadata holds the metadata of a testrun.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"name"},
			},
		},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_TestDefSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TestDefSpec is the actual description of the test.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"owner": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"recipientsOnFailure": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"behavior": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"activeDeadlineSeconds": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
					"command": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"args": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"image": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"config": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.ConfigElement"),
									},
								},
							},
						},
					},
				},
				Required: []string{"owner", "recipientsOnFailure", "description", "labels", "behavior", "activeDeadlineSeconds", "command", "args", "image"},
			},
		},
		Dependencies: []string{
			"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.ConfigElement"},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_TestDefinition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TestDefinition describes the execution of a test.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestDefMetadata"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestDefSpec"),
						},
					},
				},
				Required: []string{"kind", "spec"},
			},
		},
		Dependencies: []string{
			"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestDefMetadata", "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestDefSpec"},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_TestLocation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TestLocation describes a location to search for TestDefinitions.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"repo": {
						SchemaProps: spec.SchemaProps{
							Description: "Only for LocationType git",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"revision": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"hostPath": {
						SchemaProps: spec.SchemaProps{
							Description: "The absolute host on the minikube VM. Only for local",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type"},
			},
		},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_Testrun(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Testrun is the description of the testflow that should be executed.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestrunSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestrunStatus"),
						},
					},
				},
				Required: []string{"spec", "status"},
			},
		},
		Dependencies: []string{
			"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestrunSpec", "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestrunStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_TestrunKubeconfigs(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TestrunKubeconfigs are parameters where Shoot, Seed or a Gardener strconf for the Testrun can be specified.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"gardener": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gardener/test-infra/pkg/util/strconf.StringOrConfig"),
						},
					},
					"seed": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gardener/test-infra/pkg/util/strconf.StringOrConfig"),
						},
					},
					"shoot": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/gardener/test-infra/pkg/util/strconf.StringOrConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/gardener/test-infra/pkg/util/strconf.StringOrConfig"},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_TestrunList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TestrunList contains a list of Testruns",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.Testrun"),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.Testrun", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_TestrunSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TestrunSpec is the specification of a Testrun.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"creator": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"ttlSecondsAfterFinished": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"testLocations": {
						SchemaProps: spec.SchemaProps{
							Description: "TestLocation define repositories to look for TestDefinitions that are then executed in a workflow as specified in testflow. DEPRECATED: This field will be removed in a future version.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestLocation"),
									},
								},
							},
						},
					},
					"locationSets": {
						SchemaProps: spec.SchemaProps{
							Description: "LocationSet define location profiles with repositories to look for TestDefinitions.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.LocationSet"),
									},
								},
							},
						},
					},
					"kubeconfigs": {
						SchemaProps: spec.SchemaProps{
							Description: "Base64 encoded kubeconfigs that are mounted to every testflow step. They are available at $TM_KUBECONFIG_PATH/xxx.config, where xxx is either (gardener, seed or shoot).",
							Ref:         ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestrunKubeconfigs"),
						},
					},
					"config": {
						SchemaProps: spec.SchemaProps{
							Description: "Global config which is available to all tests in the testflow and onExit flow.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.ConfigElement"),
									},
								},
							},
						},
					},
					"testflow": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.DAGStep"),
									},
								},
							},
						},
					},
					"onExit": {
						SchemaProps: spec.SchemaProps{
							Description: "OnExit flow is called when the test is completed.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.DAGStep"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.ConfigElement", "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.DAGStep", "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.LocationSet", "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestLocation", "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.TestrunKubeconfigs"},
	}
}

func schema_pkg_apis_testmachinery_v1beta1_TestrunStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TestrunStatus is the status of the Testrun.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is the summary of all executed steps.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "State is a string that represents the actual state and/or process of the testrun.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"startTime": {
						SchemaProps: spec.SchemaProps{
							Description: "StartTime is the time when the argo workflow starts executing the steps.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"completionTime": {
						SchemaProps: spec.SchemaProps{
							Description: "CompletionTime is the time when the argo workflow is completed.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"duration": {
						SchemaProps: spec.SchemaProps{
							Description: "Duration represents the overall duration of the argo workflow. This value is calculated by (CompletionTime - StartTime)",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"workflow": {
						SchemaProps: spec.SchemaProps{
							Description: "Workflow is the name of the generated argo workflow",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"steps": {
						SchemaProps: spec.SchemaProps{
							Description: "Steps is the detailed summary of every step. It also shows all specific executed tests.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.StepStatus"),
									},
								},
							},
						},
					},
					"ingested": {
						SchemaProps: spec.SchemaProps{
							Description: "Ingested states whether the result of a testrun is already ingested into a persistant storage (db).",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"uploadedToGithub": {
						SchemaProps: spec.SchemaProps{
							Description: "UploadedToGithub states whether the status of a testrun is already uploaded to github component",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"ingested", "uploadedToGithub"},
			},
		},
		Dependencies: []string{
			"github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1.StepStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_test_infra_pkg_util_strconf_ConfigSource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ConfigSource represents a source for the value of a config element.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"configMapKeyRef": {
						SchemaProps: spec.SchemaProps{
							Description: "Selects a key of a ConfigMap.",
							Ref:         ref("k8s.io/api/core/v1.ConfigMapKeySelector"),
						},
					},
					"secretKeyRef": {
						SchemaProps: spec.SchemaProps{
							Description: "Selects a key of a secret in the pod's namespace",
							Ref:         ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.ConfigMapKeySelector", "k8s.io/api/core/v1.SecretKeySelector"},
	}
}

func schema_test_infra_pkg_util_strconf_StringOrConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StringOrConfig represents a type that could be from a string or a configuration",
				Type:        strconf.StringOrConfig{}.OpenAPISchemaType(),
				Format:      strconf.StringOrConfig{}.OpenAPISchemaFormat(),
			},
		},
	}
}
