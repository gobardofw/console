package base

import (
	"fmt"
	"path/filepath"

	"github.com/gobardofw/utils"
	"github.com/spf13/cobra"
)

// ClearCommand get clear command
func ClearCommand(storagePath string) *cobra.Command {
	var cmd = new(cobra.Command)
	cmd.Use = "clear"
	cmd.Short = "clear storage log directory"
	cmd.Args = cobra.MinimumNArgs(1)
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if isDir, err := utils.IsDirectory(storagePath); err != nil {
			fmt.Printf("failed: %s\n", err.Error())
			return
		} else if !isDir {
			fmt.Printf("failed: LOG_DIR is invalid or not found!\n")
			return
		}

		if args[0] == "all" {
			dirs, err := utils.GetSubDirectory(storagePath)
			if err != nil {
				fmt.Printf("failed: %s\n", err.Error())
				return
			}
			for _, dir := range dirs {
				if err := utils.ClearDirectory(filepath.Join(storagePath, dir)); err != nil {
					fmt.Printf("failed: %s\n", err.Error())
					return
				}
			}
		} else {
			if isDir, err := utils.IsDirectory(filepath.Join(storagePath, args[0])); err != nil {
				fmt.Printf("failed: %s\n", err.Error())
				return
			} else if !isDir {
				fmt.Printf("failed: %s log directory not found\n", args[0])
				return
			}

			if err := utils.ClearDirectory(filepath.Join(storagePath, args[0])); err != nil {
				fmt.Printf("failed: %s\n", err.Error())
				return
			}
		}
		fmt.Printf("cleared!\n")
	}
	return cmd
}
