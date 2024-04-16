package main

import "testing"

func TestGrep(t *testing.T) {
	tests := []struct {
		inFile  string
		outFile string
		params  *CmdFlags
		wantErr bool
	}{
		{
			"test_data/in.txt",
			"test_data/out_n__sint.txt",
			&CmdFlags{
				pattern:    "sint",
				after:      0,
				before:     0,
				count:      -1,
				ignoreCase: false,
				invert:     false,
				fixed:      false,
				lineNum:    true,
				color:      "",
			},
			false,
		},
		{
			"test_data/in.txt",
			"test_data/out_n_i__ut.txt",
			&CmdFlags{
				pattern:    "ut",
				after:      0,
				before:     0,
				count:      -1,
				ignoreCase: true,
				invert:     false,
				fixed:      false,
				lineNum:    true,
				color:      "",
			},
			false,
		},
		{
			"test_data/in.txt",
			"test_data/out_n_i_C_2__eu.txt",
			&CmdFlags{
				pattern:    "eu",
				after:      2,
				before:     2,
				count:      -1,
				ignoreCase: true,
				invert:     false,
				fixed:      false,
				lineNum:    true,
				color:      "",
			},
			false,
		},
		{
			"test_data/in.txt",
			"test_data/out_n_F_C_2_big.txt",
			&CmdFlags{
				pattern:    "nisi 3 ut aliquip ex ea commodo consequat. Duis aute irure dolor in",
				after:      2,
				before:     2,
				count:      -1,
				ignoreCase: false,
				invert:     false,
				fixed:      true,
				lineNum:    true,
				color:      "",
			},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.outFile, func(t *testing.T) {
			in, _ := openFile(test.inFile)
			got := Grep(in, test.params)
			out, _ := openFile(test.outFile)

			if got != out {
				t.Errorf("Grep() got string:\n\n%s\nwant:\n\n%s", got, out)
				return
			}
		})
	}
}
