package jfrogxray

import (
	"log"

	ljcscan "github.com/ljcheng999/ljc-scan/cmd/ljc-scan"
	"github.com/spf13/cobra"
)

var onlyDocker bool

var jfrogXrayCmd = &cobra.Command{
	Use:     "jfrog-xray",
	Aliases: []string{"jf"},
	Short:   "Reverses a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("jfrog-xray")
	},
}

func init() {
	jfrogXrayCmd.Flags().BoolVarP(&onlyDocker, "docker", "d", false, "Scan docker image(s)")
	ljcscan.RootCmd.AddCommand(jfrogXrayCmd)
}
