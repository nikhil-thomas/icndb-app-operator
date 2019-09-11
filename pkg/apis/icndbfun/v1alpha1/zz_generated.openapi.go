// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/nikhil-thomas/icndb-app-operator/pkg/apis/icndbfun/v1alpha1.FunApp":       schema_pkg_apis_icndbfun_v1alpha1_FunApp(ref),
		"github.com/nikhil-thomas/icndb-app-operator/pkg/apis/icndbfun/v1alpha1.FunAppSpec":   schema_pkg_apis_icndbfun_v1alpha1_FunAppSpec(ref),
		"github.com/nikhil-thomas/icndb-app-operator/pkg/apis/icndbfun/v1alpha1.FunAppStatus": schema_pkg_apis_icndbfun_v1alpha1_FunAppStatus(ref),
	}
}

func schema_pkg_apis_icndbfun_v1alpha1_FunApp(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FunApp is the Schema for the funapps API",
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
							Ref: ref("github.com/nikhil-thomas/icndb-app-operator/pkg/apis/icndbfun/v1alpha1.FunAppSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/nikhil-thomas/icndb-app-operator/pkg/apis/icndbfun/v1alpha1.FunAppStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/nikhil-thomas/icndb-app-operator/pkg/apis/icndbfun/v1alpha1.FunAppSpec", "github.com/nikhil-thomas/icndb-app-operator/pkg/apis/icndbfun/v1alpha1.FunAppStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_icndbfun_v1alpha1_FunAppSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FunAppSpec defines the desired state of FunApp",
				Properties: map[string]spec.Schema{
					"funpods": {
						SchemaProps: spec.SchemaProps{
							Description: "Funpods specify number of replicas in the deployment created",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"params": {
						SchemaProps: spec.SchemaProps{
							Description: "Params specify additional configuration if required",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/nikhil-thomas/icndb-app-operator/pkg/apis/icndbfun/v1alpha1.Param"),
									},
								},
							},
						},
					},
				},
				Required: []string{"funpods", "params"},
			},
		},
		Dependencies: []string{
			"github.com/nikhil-thomas/icndb-app-operator/pkg/apis/icndbfun/v1alpha1.Param"},
	}
}

func schema_pkg_apis_icndbfun_v1alpha1_FunAppStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FunAppStatus defines the observed state of FunApp",
				Properties: map[string]spec.Schema{
					"podnames": {
						SchemaProps: spec.SchemaProps{
							Description: "Podnames list all the pods created for FunApp",
							Type:        []string{"array"},
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
				},
				Required: []string{"podnames"},
			},
		},
		Dependencies: []string{},
	}
}
