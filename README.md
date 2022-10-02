# SPOIL BOARD GENERATOR


## LINKS OF INTEREST

[GCODE G2/G3](https://youtu.be/Y7uJvO6-SSk)


## OLD CODE


```golang
func main() {

	const minX float32 = -(1.5 * end_mill_dia)
	const maxX float32 = spoil_board_width + (1.5 * end_mill_dia)
	const deltaY float32 = (step_over_percentage / 100.0) * end_mill_dia
	var headX float32 = minX
	var headY float32 = 0

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

}
```