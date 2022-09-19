package main

import (
	"fmt"

	"strings"
)

type paramT map[string]float32

const spoil_board_width float32 = 660 // mm

const spoil_board_height float32 = 485 // mm

const cut_depth float32 = 1 // mm

const end_mill_dia float32 = 24 // mm

const feed_rate float32 = 2500 // mm/min

const step_over_percentage float32 = 16.5 // % of dia to step over

const safeZ float32 = 5 // mm

var commmands []string

func add_command(cmd_line string) {

	commmands = append(commmands, cmd_line)

}

func generate_script() {

	for _, l := range commmands {

		fmt.Printf("%s\n", l)

	}

}

func add_movement(cmd string, params paramT) {

	parts := make([]string, 0)

	parts = append(parts, cmd)

	for _, param := range []string{"X", "Y", "Z", "F"} {

		if val32, pres := params[param]; pres {

			parts = append(parts, fmt.Sprintf("%s%.3f", param, val32))

		}

	}

	add_command(strings.Join(parts, " "))

}

func main() {

	commmands = make([]string, 0)

	const minX float32 = -(1.5 * end_mill_dia)
	const maxX float32 = spoil_board_width + (1.5 * end_mill_dia)
	const deltaY float32 = (step_over_percentage / 100.0) * end_mill_dia
	var headX float32 = minX
	var headY float32 = 0

	add_command("T1")                           // tool 1
	add_command("G21")                          // Units are mm
	add_command("G17")                          // XY plane selection
	add_command("G90")                          // Abs positions
	add_movement("G01", paramT{"F": feed_rate}) // G01 feedrate

	add_movement("G00", paramT{"Z": safeZ}) //
	add_movement("G00", paramT{"X": minX, "Y": 0.0})
	add_movement("G01", paramT{"Z": -cut_depth})

	for {

		// Setup pass direction

		if headX < maxX {

			headX = maxX

		} else {

			headX = minX

		}

		add_movement("G01", paramT{"X": headX})

		headY += deltaY

		if headY > spoil_board_height {

			break

		}

		add_movement("G01", paramT{"Y": headY})

	}

	add_movement("G00", paramT{"Z": safeZ})

	generate_script()

}
