package pset

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/codegangsta/cli"
)

func GenerateFilesCmd(c *cli.Context) {
	psetSpec := c.Args()[0]
	err := generateFiles(psetSpec)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
}

func generateFiles(psetSpec string) error {
	specParts := strings.Split(psetSpec, ".")
	sectionIdParts := specParts[:3]
	sectionId := strings.Join(sectionIdParts, ".")
	err := os.Mkdir(sectionId, os.ModePerm)
	if err != nil {
		return err
	}

	exercisesFileName := sectionId + "-exercises.scm"
	testsFileName := sectionId + "-tests.scm"
	depsFileName := sectionId + "-deps.scm"

	chapter := strings.Split(sectionId, ".")[0]
	problemRange := strings.Split(specParts[3], "-")
	problems, err := makeProblemSlice(chapter, problemRange)
	if err != nil {
		return err
	}

	err = createExercisesFile(sectionId,
		exercisesFileName,
		testsFileName,
		depsFileName,
		problems)

	if err != nil {
		return err
	}

	err = createDepsFile(sectionId, depsFileName)
	if err != nil {
		return err
	}

	err = createTestsFile(sectionId, testsFileName, problems)
	if err != nil {
		return err
	}

	return nil
}

func createTestsFile(sectionId, testsFileName string, problems []Problem) error {

	bytes, err := dataTestsTemplateScmBytes()
	if err != nil {
		return err
	}

	t, err := template.New("tests").Parse(string(bytes))
	if err != nil {
		return err
	}

	dot := struct {
		Problems []Problem
	}{
		problems,
	}

	f, err := os.Create(filepath.Join(sectionId, testsFileName))
	if err != nil {
		return err
	}

	err = t.Execute(f, dot)

	if err != nil {
		return err
	}

	return nil
}

func createExercisesFile(sectionId, exercisesFileName, testsFileName, depsFileName string, problems []Problem) error {
	exercisesTemplate, err := dataExercisesTemplateScmBytes()
	if err != nil {
		return err
	}
	exercisesT, err := template.New("exercises").Parse(string(exercisesTemplate))
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(sectionId, exercisesFileName))
	if err != nil {
		return err
	}

	dot := struct {
		Problems  []Problem
		TestsFile string
		DepsFile  string
	}{
		Problems:  problems,
		TestsFile: filepath.Join(sectionId, testsFileName),
		DepsFile:  filepath.Join(sectionId, depsFileName),
	}

	err = exercisesT.Execute(f, dot)

	if err != nil {
		return err
	}

	return nil
}

func createDepsFile(sectionId, depsFileName string) error {
	_, err := os.Create(filepath.Join(sectionId, depsFileName))
	return err
}

type Problem struct {
	Chapter string
	Number  string
}

func makeProblemSlice(chapter string, pRange []string) ([]Problem, error) {
	upperBound, err := strconv.Atoi(pRange[1])
	if err != nil {
		return nil, err
	}
	lowerBoudn, err := strconv.Atoi(pRange[0])
	if err != nil {
		return nil, err
	}
	numRange := []Problem{}
	for i := lowerBoudn; i <= upperBound; i++ {
		numRange = append(numRange, Problem{chapter, strconv.Itoa(i)})
	}
	return numRange, nil
}
