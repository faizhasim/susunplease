package main

import (
	"github.com/faizhasim/susunplease/internal/service"
	se "github.com/faizhasim/susunplease/internal/sideEffect"
	"github.com/spf13/cobra"
	"io"
)

const processHelp = "process [flags] SRC_DIR DEST_DIR"

func newProcessCmd(out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "process SRC_DIR DEST_DIR",
		Short: "process [flags] SRC_DIR DEST_DIR",
		Long:  processHelp,
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runProcess(args)
		},
	}
	return cmd
}

func runProcess(args []string) error {
	srcGlob := args[0]
	destRootDir := args[1]
	pdfService := se.NewPdfService()
	rulesParser := service.NewRulesParser()
	fsFiling := se.NewFsFiling(destRootDir, se.NewFsOperation(), se.NewRandomGen())
	susunProcessor := service.NewSusunProcessor(pdfService, fsFiling, rulesParser)
	csvPath, err := rulesParser.GetCsvPath()
	if err != nil {
		return err
	}

	rules, err := rulesParser.ParseRulesFromCsv(csvPath)

	if err != nil {
		return err
	}

	if err := susunProcessor.Process(srcGlob, rules); err != nil {
		return err
	}

	return nil
}
