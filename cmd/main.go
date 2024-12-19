package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"openbce.io/kube-mds2/pkg/apiserver"
	_ "openbce.io/kube-mds2/pkg/storage/engine"
)

var option = apiserver.MdsBridgeConfig{}

func init() {
	rootCmd.PersistentFlags().StringVarP(&option.Backend, "backend", "b", "", "The backend of KMDS")
	rootCmd.PersistentFlags().StringVarP(&option.Endpoint, "endpoint", "p", "", "The endpoint of KMDS")
	rootCmd.PersistentFlags().StringVarP(&option.Engine, "engine", "e", "gorm", "The engine of KMDS")
}

var rootCmd = &cobra.Command{
	Use:   "kmds",
	Short: "Kubernetes Metadata Server",
	Long:  `Kubernetes Metadata Server`,
	Run: func(cmd *cobra.Command, args []string) {
		bridge, err := apiserver.NewMdsBridage(&option)
		if err != nil {
			panic(err)
		}

		if err := bridge.Run(); err != nil {
			panic(err)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
