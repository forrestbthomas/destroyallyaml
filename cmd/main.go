package main

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/*
Via a cli
- Pick a language (eg, Golang)
- Pick a use case (eg, data eng, back end api, batch processor, etc)
- Pick a name (cannot gurantee to be changed non-destructively)
These options create an initial docker image and an initial CRD for this controller and apply it. The controller creates a set of basic resources including:

- Service
- Deployment
- Prometheus
- Istio
- Auth integration
- Standard Ports

The repo should already exist. There should be no assumptions about the project. This is not a bootstrapping project.
K8s resources should be standardized by the companies standards, with the option to override the defaults
Do the basic example first: A simple golang API, named foo
Overrides should be var args of functions

Questions:
- what will this interface to the developer look like?
- how will updates occur?
- what happens when there is drift?
- can existing resources be imported?
- can created resources be decoupled from the controller?

*/

func main() {
	svc := v1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                       "foo",
			GenerateName:               "",
			Namespace:                  "",
			SelfLink:                   "",
			UID:                        "",
			ResourceVersion:            "",
			Generation:                 0,
			CreationTimestamp:          metav1.Time{},
			DeletionTimestamp:          &metav1.Time{},
			DeletionGracePeriodSeconds: new(int64),
			Labels:                     map[string]string{},
			Annotations:                map[string]string{},
			OwnerReferences:            []metav1.OwnerReference{},
			Finalizers:                 []string{},
			ManagedFields:              []metav1.ManagedFieldsEntry{},
		},
		Spec: v1.ServiceSpec{
			Ports:                         []v1.ServicePort{},
			Selector:                      map[string]string{},
			ClusterIP:                     "",
			ClusterIPs:                    []string{},
			Type:                          "",
			ExternalIPs:                   []string{},
			SessionAffinity:               "",
			LoadBalancerIP:                "",
			LoadBalancerSourceRanges:      []string{},
			ExternalName:                  "",
			ExternalTrafficPolicy:         "",
			HealthCheckNodePort:           0,
			PublishNotReadyAddresses:      false,
			SessionAffinityConfig:         &v1.SessionAffinityConfig{},
			IPFamilies:                    []v1.IPFamily{},
			AllocateLoadBalancerNodePorts: new(bool),
			LoadBalancerClass:             new(string),
		},
		Status: v1.ServiceStatus{
			LoadBalancer: v1.LoadBalancerStatus{},
			Conditions:   []metav1.Condition{},
		},
	}

	fmt.Printf("Hello, %s\n", svc.Name)
}
