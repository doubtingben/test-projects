package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type version struct {
	Major         uint8
	Minor         uint8
	buildDate     string
	goVersion     string
	gitBranch     string
	gitCommitHash string
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the running utility version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		output := fmt.Sprintf("Version: %d.%d on Go Version: %s", currentVersion.Major, currentVersion.Minor, currentVersion.goVersion)
		if commandLineFlags.verbosity > 0 {
			output += fmt.Sprintf(" - Branch: %s commit: %s, built on: %s", currentVersion.gitBranch, currentVersion.gitCommitHash, currentVersion.buildDate)
		}
		color.White(output)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
