package main

import (
	"math"
	"sort"
	"strconv"
)

type SortParams struct {
	colon       int
	descending  bool
	numericSort bool
	unique      bool
}

// Return two strings at column index.
//
// If there is no such column in one of lines, then set "" to this string
//
// If both lines has no such columns, return first strings of lines
func getComparedStrings(data [][]string, params *SortParams, i, j int) (string, string) {
	var a, b string

	// If there is no column with given index, set "" value to string to place
	// lines with empty columns higher
	if len(data[i]) <= params.colon && len(data[j]) <= params.colon {
		a, b = data[i][0], data[j][0]
	} else if len(data[i]) <= params.colon {
		a, b = "", data[j][params.colon]
	} else if len(data[j]) <= params.colon {
		a, b = data[i][params.colon], ""
	} else {
		// Otherwise just return strings at given columns
		a, b = data[i][params.colon], data[j][params.colon]
	}

	return a, b
}

// Return compare decision based on parameters
func getCompareDecision(a, b string, params *SortParams) bool {
	var decision bool

	n1, err1 := strconv.Atoi(a)
	n2, err2 := strconv.Atoi(b)

	// We need string comparation if numeric sort is disabled or numeric sort is enabled,
	// but we can't do it because both values are strings, so compare it based on
	// string values
	if !params.numericSort ||
		(params.numericSort && err1 != nil && err2 != nil) {
		decision = a < b
	} else if params.numericSort {
		// If one of values is string, so make it negative infinity to place
		// number lower than string
		if err1 != nil {
			n1 = int(math.Inf(-1))
		} else if err2 != nil {
			n2 = int(math.Inf(-1))
		}
		decision = n1 < n2
	}

	return decision
}

// Sort strings based on sort parameters
//
// colon - key of sorting column
//
// descending - flag which will reverse sort
//
// numericSort - flag which will enable sort based on number value of strings
func sortData(data [][]string, params *SortParams) [][]string {
	sort.Slice(data, func(i, j int) bool {
		// Get two strings for compare based on sort params
		a, b := getComparedStrings(data, params, i, j)

		// Get compare decision based on parameters (numeric sort)
		decision := getCompareDecision(a, b, params)

		// Reverse decision if needed
		if params.descending {
			return !decision
		}

		return decision
	})

	return data
}
