package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

// Color control sequences for terminal emulator
const (
	Default = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Blue    = "\033[34m"
	Yellow  = "\033[33m"
)

// Get all line indexes with matched pattern.
// Result is 2D array, where x - index of line, y - start index of matched string
func getMatchedStringsIdx(text, pattern string, fixedMatch, ignoreCase bool) [][]int {
	splitted := strings.Split(text, "\n")
	result := make([][]int, len(splitted))

	if fixedMatch {
		// If fixed search set to true, then just compare entire lines with the pattern
		for i, line := range splitted {
			if ignoreCase {
				line = strings.ToLower(line)
			}
			if line == pattern {
				result[i] = append(result[i], 0)
			}
		}
	} else {
		// Otherwise use regex to find all patterns in the line
		if ignoreCase {
			pattern = fmt.Sprintf("(?i)%s", pattern)
		}

		re := regexp.MustCompile(pattern)

		for i, line := range splitted {
			matches := re.FindAllStringIndex(line, -1)

			for _, match := range matches {
				result[i] = append(result[i], match[0])
			}
		}
	}

	return result
}

// Update bool map to set range of items to be added to the result.
// This needs to print before and after lines near matched line
func updateBoolMap(boolMap []bool, start, end int) {
	// If start index out of range, then use the first index
	start = max(start, 0)
	// If end index out of range, then use the last index
	end = min(end, len(boolMap)-1)

	for i := start; i < end; i++ {
		boolMap[i] = true
	}
}

// Highlight match with given color and return new text with found=true.
// If there is no matched patterns, return unchanged text and found=false.
func highlightMatch(text string, indexes []int, length int, color string) (string, bool) {
	defaultColor := Default
	if color == "" {
		defaultColor = ""
	}
	found := false
	// We need to reverse slice of indexes because we expanding the text while
	// highlighting => if we do this without reverse, indexes will incorrect
	slices.Reverse(indexes)
	for _, idx := range indexes {
		// We build string with such pattern:
		// <string_before_match><color_control_sequence><match><default_control_sequence><string_after_match>
		text = fmt.Sprintf("%s%s%s%s%s",
			text[0:idx],
			color,
			text[idx:idx+length],
			defaultColor,
			text[idx+length:],
		)
		found = true
	}

	return text, found
}

// Get match bool map for text, where 'true' means that this line need to be printed
func getMatchBoolMap(splitted []string, indexes [][]int, params *CmdFlags) []bool {
	boolMap := make([]bool, len(splitted))
	count := params.count

	// Algorithm is pretty simple:
	// 1) Try to highlight all mathes in the line
	// 2) Check if mathes was found
	// 3) If found, then replace line in the slice, reduce count of matches and
	// update bool map with context lines
	// 4) If count == 0, then stop loop and return bool map
	for i, line := range splitted {
		newLine, found := highlightMatch(line, indexes[i], len(params.pattern), params.color)
		if found {
			splitted[i] = newLine
			count--
			updateBoolMap(boolMap, i-params.before, i+params.after+1)
		}
		if count == 0 {
			return boolMap
		}
	}

	return boolMap
}

// Return formatted text considering invert and lineNum flags
// Result is formatted string
func getFormattedText(splitted []string, boolMap []bool, params *CmdFlags) string {
	var builder strings.Builder
	for i := range splitted {
		decision := boolMap[i]

		if params.invert {
			decision = !decision
		}

		lineNum := ""
		if params.lineNum {
			lineNum = fmt.Sprintf("%d) ", i+1)
		}

		if decision {
			builder.WriteString(fmt.Sprintf("%s%s\n", lineNum, splitted[i]))
		}
	}

	return builder.String()
}

// Grep will find all matches in the text with given pattern and parameters
func Grep(text string, params *CmdFlags) string {
	indexes := getMatchedStringsIdx(text, params.pattern, params.fixed, params.ignoreCase)

	splitted := strings.Split(text, "\n")

	boolMap := getMatchBoolMap(splitted, indexes, params)

	return getFormattedText(splitted, boolMap, params)
}
