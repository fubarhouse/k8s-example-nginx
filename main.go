package main

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	//KUBECONFIG = "/home/karl/.kube/config"
	KUBECONFIG = "/Users/karl/.kube/config"
)

func main() {

	config, _ := clientcmd.BuildConfigFromFlags("", KUBECONFIG)

	client, _ := kubernetes.NewForConfig(config)

	pod, e := client.CoreV1().Pods("default").Create(context.TODO(), &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx-pod",
			Namespace: "default",
			Labels:    map[string]string{
				"run": "nginx-pod",
			},
		},
		Spec: v1.PodSpec{
			Volumes: nil,
			Containers: []v1.Container{{
				Name: "nginx-pod",
				Image: "nginx",
			},
			},
			DNSPolicy: "ClusterFirst",
			RestartPolicy: "Always",
		},
	}, metav1.CreateOptions{})

	fmt.Sprintf("%v", pod)
	if e != nil {
		fmt.Errorf("error: %v", e)
	}

}
