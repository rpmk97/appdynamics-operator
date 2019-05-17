// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/Appdynamics/appdynamics-operator/pkg/apis/appdynamics/v1alpha1.Clusteragent":   schema_pkg_apis_appdynamics_v1alpha1_Clusteragent(ref),
		"github.com/Appdynamics/appdynamics-operator/pkg/apis/appdynamics/v1alpha1.InfraViz":       schema_pkg_apis_appdynamics_v1alpha1_InfraViz(ref),
		"github.com/Appdynamics/appdynamics-operator/pkg/apis/appdynamics/v1alpha1.InfraVizSpec":   schema_pkg_apis_appdynamics_v1alpha1_InfraVizSpec(ref),
		"github.com/Appdynamics/appdynamics-operator/pkg/apis/appdynamics/v1alpha1.InfraVizStatus": schema_pkg_apis_appdynamics_v1alpha1_InfraVizStatus(ref),
	}
}

func schema_pkg_apis_appdynamics_v1alpha1_Clusteragent(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Clusteragent is the Schema for the clusteragents API",
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
							Ref: ref("github.com/Appdynamics/appdynamics-operator/pkg/apis/appdynamics/v1alpha1.ClusteragentSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Appdynamics/appdynamics-operator/pkg/apis/appdynamics/v1alpha1.ClusteragentStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/Appdynamics/appdynamics-operator/pkg/apis/appdynamics/v1alpha1.ClusteragentSpec", "github.com/Appdynamics/appdynamics-operator/pkg/apis/appdynamics/v1alpha1.ClusteragentStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_appdynamics_v1alpha1_InfraViz(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InfraViz is the Schema for the infravizs API",
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
							Ref: ref("github.com/Appdynamics/appdynamics-operator/pkg/apis/appdynamics/v1alpha1.InfraVizSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/Appdynamics/appdynamics-operator/pkg/apis/appdynamics/v1alpha1.InfraVizStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/Appdynamics/appdynamics-operator/pkg/apis/appdynamics/v1alpha1.InfraVizSpec", "github.com/Appdynamics/appdynamics-operator/pkg/apis/appdynamics/v1alpha1.InfraVizStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_appdynamics_v1alpha1_InfraVizSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InfraVizSpec defines the desired state of InfraViz",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_appdynamics_v1alpha1_InfraVizStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InfraVizStatus defines the observed state of InfraViz",
				Type:        []string{"object"},
			},
		},
	}
}
