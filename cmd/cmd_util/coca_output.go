package cmd_util

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/phodal/igso/cmd/config"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func NewOutput(output io.Writer) *tablewriter.Table {
	table := tablewriter.NewWriter(output)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetColWidth(80)
	return table
}

var reporterPath = config.CmdConfig.ReporterPath

func WriteToCocaFile(fileName string, payload string) {
	if _, err := os.Stat(reporterPath); os.IsNotExist(err) {
		mkdirErr := os.Mkdir(reporterPath, os.ModePerm)
		if mkdirErr != nil {
			fmt.Println(mkdirErr)
		}
	}
	_ = ioutil.WriteFile(filepath.FromSlash(reporterPath+"/"+fileName), []byte(payload), os.ModePerm)
}

func ConvertToSvg(name string) {
	reporter_path := config.CmdConfig.ReporterPath
	cmd := exec.Command("dot", "-Tsvg", reporter_path+"/"+name+".dot", "-o", reporter_path+"/"+name+".svg")
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("cmd.Run() failed with:", err)
	}
}
