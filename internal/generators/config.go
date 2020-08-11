package generators

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/spf13/cobra"
)

// generateConfigCmd represents the generateConfig command
var generateConfigCmd = &cobra.Command{
	Use:   "generate:config [destination_path]",
	Short: "Generates the default config",
	Long: `Generates the default authboss config. 
	This is not neccessary, but useful if you want 
	to see all the available options`,
	Run: func(cmd *cobra.Command, args []string) {
		err := generateConfig(cmd, args)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			os.Exit(-1)
		}
	},
}

func generateConfig(cmd *cobra.Command, args []string) error {
	dst := args[0]

	if len(dst) == 0 {
		return fmt.Errorf("You must specify a destination path")
	}
	_, c, _, _ := runtime.Caller(0)
	userFile := path.Join(c, "../../../examples/config.go")
	err := CopyFile(userFile, dst)
	if err != nil {
		return fmt.Errorf("Could not create the directory: %w", err)
	}
	fmt.Println("User model generated into: " + dst)

	return nil
}

func init() {
	rootCmd.AddCommand(generateConfigCmd)
}
