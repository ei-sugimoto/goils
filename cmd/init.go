package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the Air configuration",
	Run: func(cmd *cobra.Command, args []string) {
		err := runAirInit()
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runAirInit() error {
	cmd := exec.Command("air", "init")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run air init: %w", err)
	}
	fmt.Println(string(output))
	return nil
}
