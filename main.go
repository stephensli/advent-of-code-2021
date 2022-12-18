package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/urfave/cli/v2"
)

type params struct {
	year int
	day  int
}

func main() {
	inputs := params{}

	app := &cli.App{
		Name:  "aoc",
		Usage: "aoc supporting tool",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "year",
				Usage:       "The year of the newly started task",
				Required:    true,
				Destination: &inputs.year,
			},
			&cli.IntFlag{
				Name:        "day",
				Usage:       "The day that is being started",
				Required:    true,
				Destination: &inputs.day,
			},
		},
		Action: func(context *cli.Context) error {
			return action(context, inputs)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func action(ctx *cli.Context, inputs params) error {
	path, _ := os.Getwd()

	// create the day directory if it does not exist.
	dayDirectory := filepath.Join(path, strconv.Itoa(inputs.year), fmt.Sprintf("day%d", inputs.day))

	if _, err := os.Stat(dayDirectory); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(dayDirectory, 0777); err != nil {
			return err
		}
	}

	// create the input directories if it does not exist.
	inputsDirectory := filepath.Join(path, strconv.Itoa(inputs.year), "inputs")
	if _, err := os.Stat(inputsDirectory); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(inputsDirectory, 0777); err != nil {
			return err
		}
	}

	// create the example and standard txt files for reading the input for
	// the advent of code execution.
	for _, fileName := range []string{
		filepath.Join(inputsDirectory, fmt.Sprintf("day%d.txt", inputs.day)),
		filepath.Join(inputsDirectory, fmt.Sprintf("day%d.example.txt", inputs.day)),
	} {
		if file, err := os.Create(fileName); err != nil {
			return err
		} else {
			file.Close()
		}
	}

	// finally create the main.go file that will be used to run the program.
	// This will be based on the template.
	// variables
	vars := map[string]any{
		"Year": strconv.Itoa(inputs.year),
		"Day":  strconv.Itoa(inputs.day),
	}

	templateDirectory := filepath.Join(path, "aoc.tmpl")
	mainDirectory := filepath.Join(dayDirectory, "main.go")

	tmpl, _ := template.ParseFiles(templateDirectory)
	file, _ := os.Create(mainDirectory)
	defer file.Close()

	// apply the template to the vars map and write the result to file.
	return tmpl.Execute(file, vars)
}
