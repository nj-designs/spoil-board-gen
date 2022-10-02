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

	fmt.Println(*boardFilePath)

	job := cnc.NewJob()

	job.AddCommand("T1")  // tool 1
	job.AddCommand("G21") // Units are mm
	job.AddCommand("G17") // XY plane selection
	job.AddCommand("G90") // Abs positions

	job.AddMovement("G01", cnc.CommandParamsT{"F": float32(cnc.DefaultHorizontalFeedRate)}) // G01 feedrate

	job.Print()

}
