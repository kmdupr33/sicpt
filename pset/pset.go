package pset

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

func GenerateFilesCmd(c *cli.Context) {
	psetSpec := c.Args()[0]
	err := generateFiles(psetSpec)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	// problemRange := strings.Split(specParts[3], "-")
	// rangeSlice := sliceFromProblemRange(problemRange)

}

func generateFiles(psetSpec string) error {
	specParts := strings.Split(psetSpec, ".")
	sectionId := specParts[:3]
	dirName := strings.Join(sectionId, ".")
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func sliceFromProblemRange([]string) []int {
	return nil
}
