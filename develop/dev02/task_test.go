package main

import "testing"

func TestUnpackString(t *testing.T) {
	tests := []struct {
		str     string
		want    string
		wantErr bool
	}{
		{
			"",
			"",
			false,
		},
		{
			"a4bc2d5e",
			"aaaabccddddde",
			false,
		},
		{
			"abcd",
			"abcd",
			false,
		},
		{
			"45",
			"",
			true,
		},
		{
			"a4bd5",
			"aaaabddddd",
			false,
		},
	}
	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			got, err := UnpackString(test.str)
			if (err != nil) != test.wantErr {
				t.Errorf("UnpackString() error = #{err}, wantErr = #{test.wantErr}")
				return
			}
			if got != test.want {
				t.Errorf("UnpackString() got = #{got}, want = #{test.want}")
			}
		})
	}
}
