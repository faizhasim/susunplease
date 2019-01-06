package main

import (
	"fmt"
	"github.com/faizhasim/susunplease/internal/service"
	"github.com/mitchellh/go-homedir"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func rulesHelp() string {
	home, _ := homedir.Dir()
	return fmt.Sprintf(`init

Will create sample 'rules.csv' in '%s'
`, home)
}

func newRulesCmd(out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "rules",
		Short: "Open rules CSV using OS default editor",
		Long:  rulesHelp(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInit(out)
		},
	}
	return cmd
}

var sampleCsvRule = []byte(`documentType,targetDir,matchRegex
sugar-high-inc,receipt/food,sugar.*high
`)

func runInit(out io.Writer) error {
	rulesParser := service.NewRulesParser()
	csvPath, err := rulesParser.GetCsvPath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		csvDir, _ := path.Split(csvPath)
		if os.MkdirAll(csvDir, os.ModePerm) != nil {
			return err
		}
		if err := ioutil.WriteFile(csvPath, sampleCsvRule, 0644); err != nil {
			return err
		}
	}

	if err := open.Run(csvPath); err != nil {
		return err
	}

	return nil
}
