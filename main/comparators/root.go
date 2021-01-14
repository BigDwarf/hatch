package comparators

import (
	"github.com/spf13/cobra"
	"hatch/main/comparators/json"
)

func NewCompareCommand() *cobra.Command {

	cmds := &cobra.Command{
		Use:           "hatch SUBCOMMAND",
		Short:         "starts hatch binary",
		Long:          "starts orchestrator service",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmds.ResetFlags()

	cmds.AddCommand(json.NewCompareJsonCommand())

	return cmds
}
