package main

import (
	"fmt"
	"testing"
)

func TestOr(t *testing.T) {
	tests := []struct {
		channel chan interface{}
	}{
		{
			make(chan interface{}),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", i), func(t *testing.T) {
			close(test.channel)
			done := or(test.channel)

			if _, ok := <-done; !ok {
				return
			} else {
				t.Errorf("Channel is not closed!")
				return
			}
		})
	}
}
