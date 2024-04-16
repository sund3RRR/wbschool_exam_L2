package main

import (
	"slices"
	"strings"
)

// Split text on fields with given delimiter and return 2D string array
func splitText(text string, params *CmdFlags) [][]string {
	result := make([][]string, 0)

	splitted := strings.Split(text, "\n")

	for _, str := range splitted {
		fields := strings.Split(str, params.delimiter)

		// If we need to print only strings with separator, then
		// skip this iteration if there is no separator in this line
		if len(fields) == 1 && params.separated {
			continue
		}

		result = append(result, fields)
	}

	return result
}

// Format text with given delimiter and field numbers, return result string
func getFormattedText(splitted [][]string, cmdFlags *CmdFlags) string {
	var textBuilder strings.Builder
	for _, line := range splitted {
		var lineBuilder strings.Builder

		for j, str := range line {
			if slices.Contains(cmdFlags.fields, j+1) {
				lineBuilder.WriteString(str)
				lineBuilder.WriteString(cmdFlags.delimiter)
			}
		}
		str := strings.TrimSuffix(lineBuilder.String(), cmdFlags.delimiter) + "\n"
		textBuilder.WriteString(str)
	}

	return textBuilder.String()
}

// Cut will return text separated with given delimiter and given field numbers
func Cut(text string, params *CmdFlags) string {
	splitted := splitText(text, params)

	return getFormattedText(splitted, params)
}
