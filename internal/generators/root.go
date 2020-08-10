package generators

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "authboss",
	Short: "Built in authboss generators",
	Long:  "Authboss contains many generators to help get your application up and running!",
}

// Execute : Execute all child commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
