package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Verbose bool
var Pods bool

func init() {
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().String("lang", "en", "language to use")

	//RootCmd.AddCommand(X)
	//RootCmd.AddCommand(Y)
}

var RootCmd = &cobra.Command{
	Use: "pods",
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			fmt.Println("Listing pods...")
		}

        pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{}) // https://godoc.org/k8s.io/apimachinery/pkg/apis/meta/v1
        if err != nil {
           panic(err.Error())
        }
        for _, pod := range pods.Items{
            mt.Printf("%s %s\n", pod.GetName(), pod.GetCreationTimestamp())
        }
	},
}	
