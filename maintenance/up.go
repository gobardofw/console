package maintenance

import (
	"fmt"

	"github.com/gobardofw/cache"
	"github.com/spf13/cobra"
)

// UpCommand get maintenance up command
func UpCommand(resolver func(driver string) cache.Cache, defDriver string) *cobra.Command {
	var cmd = new(cobra.Command)
	cmd.Use = "up"
	cmd.Short = "exist from maintenance mode"
	cmd.Run = func(cmd *cobra.Command, args []string) {
		var err error
		driver, err := cmd.Flags().GetString("driver")
		if err != nil {
			fmt.Printf("failed: %s\n", err.Error())
			return
		}

		cache := resolver(driver)
		if cache == nil {
			fmt.Printf("failed: %s cache driver not found\n", driver)
			return
		}
		cache.Forget("maintenance")
		fmt.Println("app is actived!")
	}
	cmd.Flags().StringP("driver", "d", defDriver, "cache driver name")
	return cmd
}
