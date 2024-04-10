package main

import (
	"fmt"
	"testing"
)

func TestGetAnagramSet(t *testing.T) {
	tests := []struct {
		arr     []string
		want    map[string][]string
		wantErr bool
	}{
		{
			nil,
			map[string][]string{},
			false,
		},
		{
			[]string{"abc", "cba", "cbb", "asd"},
			map[string][]string{
				"abc": {"abc", "cba"},
			},
			false,
		},
		{
			[]string{"123", "231", "231", "kEk", "eKk", "single", "Mmm", "mMm", "mmm"},
			map[string][]string{
				"123": {"123", "231"},
				"kek": {"ekk", "kek"},
			},
			false,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("Test case #%d", i), func(t *testing.T) {
			got := GetAnagramSet(test.arr)

			if len(got) != len(test.want) {
				t.Errorf("GetAnagramSet() got len = %d, want len= %d", len(got), len(test.want))
				return
			}

			for key, val := range got {
				wantSet, ok := test.want[key]
				if !ok {
					t.Errorf("GetAnagramSet() got key = %s doesn't exists in want", key)
					return
				}

				if len(val) != len(wantSet) {
					t.Errorf("GetAnagramSet() got set len = %d, want set len = %d", len(val), len(wantSet))
					return
				}

				for i := 0; i < len(val); i++ {
					if val[i] != wantSet[i] {
						t.Errorf("GetAnagramSet() got set value = %s, want set value = %s", val[i], wantSet[i])
						return
					}
				}
			}
		})
	}
}
