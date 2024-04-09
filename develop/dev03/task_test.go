package main

import "testing"

func TestParseFile(t *testing.T) {
	tests := []struct {
		inFile  string
		want    [][]string
		wantErr bool
	}{
		{
			"test_data/in_1.txt",
			[][]string{
				{"2", "4", "hgg"},
				{"3", "sdf", "3"},
				{"f"},
			},
			false,
		},
		{
			"test_data/unexisted_file.txt",
			nil,
			true,
		},
	}
	for _, test := range tests {
		t.Run(test.inFile, func(t *testing.T) {
			got, err := ParseFile(test.inFile)
			if (err != nil) != test.wantErr {
				t.Errorf("parseFile() error = %v, wantErr = %v", err, test.wantErr)
				return
			}

			if len(got) != len(test.want) {
				t.Errorf("parseFile() got len = %d, want len = %d", len(got), len(test.want))
				return
			}

			for i := 0; i < len(got); i++ {
				if len(got[i]) != len(test.want[i]) {
					t.Errorf("parseFile() %d line got len = %d, want len = %d", i, len(got[i]), len(test.want[i]))
					return
				}
				for j := 0; j < len(got[i]); j++ {
					if got[i][j] != test.want[i][j] {
						t.Errorf("UnpackString() got = %s, want = %s", got[i][j], test.want[i][j])
					}
				}
			}
		})
	}
}

func TestSortData(t *testing.T) {
	tests := []struct {
		inFile     string
		outFile    string
		sortParams *SortParams
		wantErr    bool
	}{
		{
			"test_data/in_1.txt",
			"test_data/out_1.txt",
			&SortParams{
				0,
				false,
				false,
				false,
			},
			false,
		},
		{
			"test_data/in_1.txt",
			"test_data/out_1_r.txt",
			&SortParams{
				0,
				true,
				false,
				false,
			},
			false,
		},
		{
			"test_data/in_1.txt",
			"test_data/out_1_n.txt",
			&SortParams{
				0,
				false,
				true,
				false,
			},
			false,
		},
		{
			"test_data/in_2.txt",
			"test_data/out_2.txt",
			&SortParams{
				0,
				false,
				false,
				false,
			},
			false,
		},
		{
			"test_data/in_2.txt",
			"test_data/out_2_r_n.txt",
			&SortParams{
				0,
				true,
				true,
				false,
			},
			false,
		},
		{
			"test_data/in_3.txt",
			"test_data/out_3.txt",
			&SortParams{
				0,
				false,
				false,
				false,
			},
			false,
		},
		{
			"test_data/in_3.txt",
			"test_data/out_3_k1_r.txt",
			&SortParams{
				1,
				true,
				false,
				false,
			},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.outFile, func(t *testing.T) {
			in, _ := ParseFile(test.inFile)
			got := sortData(in, test.sortParams)
			out, _ := ParseFile(test.outFile)

			if len(got) != len(out) {
				t.Errorf("sortData() got len = %d, want len = %d", len(got), len(out))
				return
			}

			for i := 0; i < len(got); i++ {
				if len(got[i]) != len(out[i]) {
					t.Errorf("sortData() %d line got len = %d, want len = %d", i, len(got[i]), len(out[i]))
					return
				}
				for j := 0; j < len(got[i]); j++ {
					if got[i][j] != out[i][j] {
						t.Errorf("sortData() got = %s, want = %s", got[i][j], out[i][j])
					}
				}
			}
		})
	}
}
