package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:      "pset",
			Usage:     "generate template files for a problem set",
			UsageText: "takes a string that specifies the chapter, section, and number of problems in a pset. For example, '1.1.6.1-5' creates files for the pset containing problems 1.1-1.5 in section 1.1.6. It uses this string to generates 4 files: <pset-name>-tests.scm, <pset-name>-deps.scm, <pset-name>-exercises.scm, <pset-name>-solutions.scm.",
			Action:    generatePsetFiles,
		},
		{
			Name:   "getdeps",
			Usage:  "get the definitions of a undefined symbols and their dependencies within an exercises file in a pset and add it to a dependencies template file for a pset",
			Action: getDeps,
		},
	}
	app.Run(os.Args)
}

func generatePsetFiles(c *cli.Context) {

}

func getDeps(c *cli.Context) {

}
