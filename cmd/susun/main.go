package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const globalUsage = `🗄  SUSUN 🗄

To start working with Susun, run the 'susun rules' command to edit the rules:

	$ susun rules

After the rules has been configured, the processing can be run via:

	$ susun process SRC_DIR DEST_DIR

`

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "susun",
		Short:        "Sort out PDF documents",
		Long:         globalUsage,
		SilenceUsage: true,
	}

	out := cmd.OutOrStdout()

	cmd.AddCommand(newProcessCmd(out), newRulesCmd(out))

	return cmd
}

func main() {
	cmd := newRootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
