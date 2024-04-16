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
			"test_data/out_f_2_d__p.txt",
			&CmdFlags{
				fields:    []int{2},
				delimiter: "p",
				separated: false,
			},
			false,
		},
		{
			"test_data/in.txt",
			"test_data/out_f_123_d__space.txt",
			&CmdFlags{
				fields:    []int{1, 2, 3},
				delimiter: " ",
				separated: false,
			},
			false,
		},
		{
			"test_data/in.txt",
			"test_data/out_f_12_s_d__p.txt",
			&CmdFlags{
				fields:    []int{1, 2},
				delimiter: "p",
				separated: true,
			},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.outFile, func(t *testing.T) {
			in, _ := openFile(test.inFile)
			got := Cut(in, test.params)
			out, _ := openFile(test.outFile)

			if got != out {
				t.Errorf("Cut() got string:\n\n%s\nwant:\n\n%s", got, out)
				return
			}
		})
	}
}
