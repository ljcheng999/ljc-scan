package ljcscan

import (
	"fmt"
	"os"

	"github.com/ljcheng999/ljc-scan/pkg/constant"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "ljc-scan",
	Version: constant.LJC_SCAN_CLI_VERSION,
	Short:   "ljc-scan - a simple CLI to find vulnerability",
	Long: `ljc-scan is a script built on top of multiple different tools - Jfrog/Docker/Sonarqube/Veracode (in progress)
   
One can use ljc-scan to find out any risk from the code/image`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
