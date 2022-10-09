package cnc

import (
	"fmt"
	"log"
	"strings"
)

const DefaultHorizontalFeedRate float64 = 2500 // mm/min
const DefaultVerticleFeedRate float64 = 1000   // mm/min
const DefaultSafeZ float64 = 5                 // mm

type CommandParamsT map[string]float64

// Job represents a CNC job
type Job struct {
	commmands          []string
	horizontalFeedRate float64
	verticleFeedRate   float64
	safeZ              float64
}

func NewJob() *Job {
	return &Job{commmands: make([]string, 0), horizontalFeedRate: DefaultHorizontalFeedRate, verticleFeedRate: DefaultVerticleFeedRate, safeZ: DefaultSafeZ}
}

func (j *Job) AddCommand(cmd_line string) {
	j.commmands = append(j.commmands, cmd_line)
}

func (j *Job) AddMovement(cmd string, params CommandParamsT) {

	parts := make([]string, 0)

	parts = append(parts, cmd)

	for _, param := range []string{"X", "Y", "Z", "I", "J", "F"} {

		if val32, pres := params[param]; pres {

			parts = append(parts, fmt.Sprintf("%s%.3f", param, val32))

		}

	}

	j.AddCommand(strings.Join(parts, " "))
}

func (j *Job) AddVerticleMovement(cmd string, params CommandParamsT) {
	parts := make([]string, 0)

	parts = append(parts, cmd)

	if _, pres := params["Z"]; !pres {
		log.Fatalf("Missing Z value for veritcle movement")
	}

	if _, pres := params["F"]; !pres {
		params["F"] = j.verticleFeedRate
	}

	for _, param := range []string{"Z", "F"} {
		val := params[param]
		parts = append(parts, fmt.Sprintf("%s%.3f", param, val))
	}

	j.AddCommand(strings.Join(parts, " "))
}

func (j *Job) AddHorizontalMovement(cmd string, params CommandParamsT) {

	parts := make([]string, 0)

	parts = append(parts, cmd)

	if _, pres := params["F"]; !pres {
		params["F"] = j.horizontalFeedRate
	}

	for _, param := range []string{"X", "Y", "I", "J", "F"} {

		if val32, pres := params[param]; pres {

			parts = append(parts, fmt.Sprintf("%s%.3f", param, val32))

		}

	}

	j.AddCommand(strings.Join(parts, " "))
}

func (j *Job) AddComment(comment string) {

	j.AddCommand(fmt.Sprintf("; %s", comment))
}

// drillHole will generate gcode to drill a hole.
// Hole cente: x_pos, y_pos
func (j *Job) DrillHole(x_pos, y_pos, hole_diameter, max_depth, bit_diameter, slice_depth float64) {
	// Goto safe Z
	j.AddVerticleMovement("G0", CommandParamsT{"Z": j.safeZ})

	bit_radius := bit_diameter / 2.0

	for d := slice_depth; d < max_depth; d += slice_depth {
		j.AddComment(fmt.Sprintf("Depth: %.3f", d))

		hole_radius := hole_diameter / 2.0

		// Locate start position
		j.AddHorizontalMovement("G0", CommandParamsT{"X": x_pos - hole_radius, "Y": y_pos})
		// Plunge
		j.AddVerticleMovement("G1", CommandParamsT{"Z": -d, "F": j.verticleFeedRate})

		for hole_radius >= bit_radius {
			i := hole_radius - bit_radius
			j.AddHorizontalMovement("G3", CommandParamsT{"X": x_pos - i, "Y": y_pos, "I": i, "J": 0})
			hole_radius -= bit_radius
		}
	}

	// Goto safe Z
	j.AddVerticleMovement("G0", CommandParamsT{"Z": j.safeZ})
}

func (j *Job) Print() {
	for _, l := range j.commmands {
		fmt.Printf("%s\n", l)
	}
}
