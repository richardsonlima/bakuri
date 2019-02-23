// Copyright © 2019 Richardson Lima <contato@richardsonlima.com.br>
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

package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// listservicesCmd represents the listservices command
var listservicesCmd = &cobra.Command{
	Use:   "listservices",
	Short: "Listing k8s services",
	Long: `Description: xxx`,
	Run: func(cmd *cobra.Command, args []string) {
		c := color.New(color.FgCyan).Add(color.Underline)
		c.Println("[+] Listing Services on kube-system namespace: ")

		kubeconfig := os.Getenv("HOME") + "/.kube/config"
		  // uses the current context in kubeconfig
			config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
			if err != nil {
				panic(err.Error())
			}
		  // creates the clientset
			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				panic(err.Error())
			}
			// List all service on namespace
			namespace := "kube-system"
			services, err := clientset.CoreV1().Services(namespace).List(metav1.ListOptions{})
			if err != nil {
	            panic(err.Error())
			}
	        for _, service := range services.Items{
	            fmt.Printf("   Service: %s\n", service.GetName())
	        }
		},
	}

func init() {
	rootCmd.AddCommand(listservicesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listservicesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listservicesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
