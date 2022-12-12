#####Kubernetes operator in Golang that will monitor the number of pods running in a cluster..


package main

import (
    "context"
    "fmt"
    "time"

    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
)

func main() {
    // Create a new Kubernetes client
    config, err := rest.InClusterConfig()
    if err != nil {
        panic(err.Error())
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }

    // Monitor the number of pods every 5 minutes
    ticker := time.NewTicker(5 * time.Minute)
    for range ticker.C {
        pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
        if err != nil {
            panic(err.Error())
        }

        fmt.Printf("There are currently %d pods running in the cluster\n", len(pods.Items))
    }
}
