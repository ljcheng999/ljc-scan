package jfrogxray

import (
	"log"

	ljcscan "github.com/ljcheng999/ljc-scan/cmd/ljc-scan"
	client "github.com/ljcheng999/ljc-scan/pkg/client"
	"github.com/spf13/cobra"
)

var onlyDocker bool
var username string
var password string
var artifactoryUrl string

var jfrogXrayCmd = &cobra.Command{
	Use:     "jfrog-xray",
	Aliases: []string{"jf"},
	Short:   "Reverses a string",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Execute jfrog-xray")

		if onlyDocker {

			client.JfrogLogin(username, password, artifactoryUrl)
		}
	},
}

func init() {
	jfrogXrayCmd.Flags().BoolVarP(&onlyDocker, "docker", "d", false, "Scan docker image(s)")
	jfrogXrayCmd.Flags().StringVarP(&username, "username", "u", "", "Username to login Jfrog")
	jfrogXrayCmd.Flags().StringVarP(&password, "password", "p", "", "Password to login Jfrog")
	jfrogXrayCmd.Flags().StringVar(&artifactoryUrl, "url", "", "URL of the Jfrog")
	ljcscan.RootCmd.AddCommand(jfrogXrayCmd)
}
