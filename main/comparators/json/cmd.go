package json

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"hatch/pkg/comparators"
	"time"
)

var (
	firstFile  string
	secondFile string
)

func NewCompareJsonCommand() *cobra.Command {
	err := flag.CommandLine.Parse([]string{})
	if err != nil {
		logrus.Fatal(err)
	}

	cmd := &cobra.Command{
		Use:   "json",
		Short: "Starts Json Comparator",
		Long:  "Starts Json Comparator",
		Run: func(cmd *cobra.Command, args []string) {
			err := runComparatorCmd()
			if err != nil {
				logrus.Error(err)
			}
		},
	}

	cmd.PersistentFlags().StringVar(&firstFile, "firstFilePath", "/Users/arslanusifov/go/src/hatch/bin/input1.json", "Path to first file to compare")
	cmd.PersistentFlags().StringVar(&secondFile, "secondFilePath", "/Users/arslanusifov/go/src/hatch/bin/input2.json", "Path to second file to compare")

	return cmd
}

func runComparatorCmd() error {
	logrus.Infof("Starting comparing json files: %s and %s", firstFile, secondFile)
	tStart := time.Now()
	comparator := comparators.NewJsonComparator()
	err, res := comparator.Compare(firstFile, secondFile)
	duration := time.Since(tStart)

	if err != nil {
		logrus.Errorf("Error happened during file comparison: %s", err.Error())
	} else {
		if res {
			logrus.Info("Json files are identical")
		} else {
			logrus.Info("Json files are NOT identical")
		}
		logrus.Infof("Finished comparing files in %s", duration)
	}
	return nil
}
