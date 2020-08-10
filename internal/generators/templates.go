package generators

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// generateTemplatesCmd represents the generateTemplates command
var generateTemplatesCmd = &cobra.Command{
	Use:   "generateTemplates",
	Short: "Generates authboss templates",
	Long:  `Generates the default authboss templates into the specified folder`,
	Run: func(cmd *cobra.Command, args []string) {
		err := generateTemplates(cmd)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			os.Exit(-1)
		}
	},
}

func generateTemplates(cmd *cobra.Command) error {
	dst, _ := cmd.Flags().GetString("dst")

	if len(dst) == 0 {
		return fmt.Errorf("You must specify a destination path")
	}
	_, tplDir, _, _ := runtime.Caller(0)
	tplDir = strings.TrimSuffix(tplDir, "generators/templates.go") + "web/templates"
	err := CopyDir(tplDir, dst)
	if err != nil {
		return fmt.Errorf("Could not create the directory: %w", err)
	}
	fmt.Println("Templates generated into: " + dst)

	return nil
}

func init() {
	rootCmd.AddCommand(generateTemplatesCmd)
	generateTemplatesCmd.Flags().StringP("dst", "d", "", "the path the templates will be copied to")
}
