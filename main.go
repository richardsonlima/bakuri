package main

import (
	"fmt"
	"os"
    "flag"
	"github.com/chzyer/readline"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "github.com/richardsonlima/bakuri/packages"
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
    	
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
