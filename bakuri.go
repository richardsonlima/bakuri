package main

import (
    "flag"
    "fmt"

    "github.com/chzyer/readline"
    "k8s.io/client-go/kubernetes"
    //"k8s.io/client-go/pkg/api/v1"
    //"github.com/kubernetes/api/core/v1"
    //"k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/tools/clientcmd"
)

func main() {
    kubeconfig := flag.String("kubeconfig", "/home/richardson/.kube/config", "this is the absolute path to the kubeconfig file")
    flag.Parse()
    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        panic(err.Error())
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }

    rl, err := readline.New("k8s> ")
    if err != nil {
        panic(err)
    }
    defer rl.Close()

    for {
        line, err := rl.Readline()
        if err != nil || line == "exit" {
            break
        }
        if line == "ps" {
            pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{}) //./k8s-conn.go:40:59: undefined: "k8s.io/api/core/v1".ListOptions
            if err != nil {
                panic(err.Error())
            }
            for _, pod := range pods.Items{
                fmt.Printf("%s %s\n", pod.GetName(), pod.GetCreationTimestamp())
            }

        } else {
            fmt.Printf("unknown command\n")
        }
    }
}
