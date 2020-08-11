package generators

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/spf13/cobra"
)

// generateTemplatesCmd represents the generateTemplates command
var generateTemplatesCmd = &cobra.Command{
	Use:   "generate:templates [destination_path]",
	Short: "Generates authboss templates",
	Long:  `Generates the default authboss templates into the specified folder`,
	Run: func(cmd *cobra.Command, args []string) {
		err := generateTemplates(cmd, args)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			os.Exit(-1)
		}
	},
}

func generateTemplates(cmd *cobra.Command, args []string) error {
	dst := args[0]

	if len(dst) == 0 {
		return fmt.Errorf("You must specify a destination path")
	}
	_, c, _, _ := runtime.Caller(0)
	tplDir := path.Join(c, "../../../web/templates")
	err := CopyDir(tplDir, dst)
	if err != nil {
		return fmt.Errorf("Could not create the directory: %w", err)
	}
	fmt.Println("Templates generated into: " + dst)

	return nil
}

func init() {
	rootCmd.AddCommand(generateTemplatesCmd)
}
