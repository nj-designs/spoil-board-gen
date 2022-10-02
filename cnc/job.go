package cnc

import (
	"fmt"
	"strings"
)

const DefaultHorizontalFeedRate float64 = 2500 // mm/min
const DefaultVerticleFeedRate float64 = 1000   // mm/min

type CommandParamsT map[string]float32

// Job represents a CNC job
type Job struct {
	commmands          []string
	horizontalFeedRate float64
	verticleFeedRate   float64
}

func NewJob() *Job {
	return &Job{commmands: make([]string, 0), horizontalFeedRate: DefaultHorizontalFeedRate, verticleFeedRate: DefaultVerticleFeedRate}
}

func (j *Job) AddCommand(cmd_line string) {
	j.commmands = append(j.commmands, cmd_line)
}

func (j *Job) AddMovement(cmd string, params CommandParamsT) {

	parts := make([]string, 0)

	parts = append(parts, cmd)

	for _, param := range []string{"X", "Y", "Z", "F"} {

		if val32, pres := params[param]; pres {

			parts = append(parts, fmt.Sprintf("%s%.3f", param, val32))

		}

	}

	j.AddCommand(strings.Join(parts, " "))
}

// drillHole will generate gcode to drill a hole.
// Hole cente: x_pos, y_pos
func (j *Job) DrillHole(x_pos, y_pos, hole_diameter, hole_depth, bit_diameter, slice_depth float64) {

}

func (j *Job) Print() {
	for _, l := range j.commmands {
		fmt.Printf("%s\n", l)
	}
}
