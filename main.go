package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nj-designs/spoil-board-gen/cnc"
)

type BoardDimensions struct {
	Width     float32 `json:"width`
	Height    float32 `json:"height`
	Thickness float32 `json:"thickness`
}

type BoardDescriptor struct {
	Name       string          `json:"name"`
	Dimensions BoardDimensions `json:"dimensions"`
}

func printUsageAndQuit() {

	fmt.Println("Usage: spg --board <path to board.json file>")
	os.Exit(1)
}

func main() {

	boardFilePath := flag.String("board", "", "Spoil Board description file")
	flag.Parse()

	if len(*boardFilePath) == 0 {
		printUsageAndQuit()
	}

	job := cnc.NewJob()

	job.AddCommand("T1")  // tool 1
	job.AddCommand("G21") // Units are mm
	job.AddCommand("G17") // XY plane selection
	job.AddCommand("G90") // Abs positions

	job.AddMovement("G01", cnc.CommandParamsT{"F": 1600.0}) // G01 feedrate

	// job.DrillHole(100, 50, 19.5, 2.5, 6.35, 0.5)

	job.GenerateSurfaceCommands(0.0, 0, 215, 230, 0.5)

	job.Print()
}
