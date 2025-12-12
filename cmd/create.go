package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// createCmd represents the file create command
var createCmd = &cobra.Command{
	Use:   "create <path>",
	Short: "Create a test file at the given path",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]

		filepath := filepath.Join(path, "testfile")
		content := []byte("test\n")

		if err := os.WriteFile(filepath, content, 0644); err != nil {
			return err
		}

		fmt.Println("file created", path)
		return nil
	},
}

func init() {
	fileCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
