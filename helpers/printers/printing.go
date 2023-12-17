package printers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"
)

var (
	jsonPrintFileIncrementer int
)

func PrettyPrintGird[V any](grid [][]V) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	for i := 0; i < len(grid); i++ {
		fmt.Fprint(w, i, ".\t")
		for j := 0; j < len(grid[i]); j++ {
			fmt.Fprint(w, "\t", grid[i][j])
		}

		fmt.Fprintln(w)
	}
	w.Flush()
}

func JsonPrint(val any, file bool) {
	marshallJson, _ := json.MarshalIndent(val, "", "  ")

	if file {
		folderPath := fmt.Sprintf("./output")

		if jsonPrintFileIncrementer == 0 {
			_ = os.RemoveAll(folderPath)
			_ = os.Mkdir(folderPath, os.ModePerm)
		}

		_ = ioutil.WriteFile(fmt.Sprintf("%s/out-%d.json", folderPath, jsonPrintFileIncrementer),
			marshallJson, 0644)

		jsonPrintFileIncrementer += 1
	} else {

		fmt.Println(string(marshallJson))
	}
}
