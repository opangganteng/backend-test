package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestSumEven(t *testing.T) {
	tests := []struct {
		name        string
		numbers     []int
		workerCount int
		want        int64
	}{
		{name: "mixed positive and negative", numbers: []int{1, 2, 3, 4, -6}, workerCount: 4, want: 0},
		{name: "uneven partitions", numbers: []int{2, 4, 6, 8, 10}, workerCount: 3, want: 30},
		{name: "more workers than numbers", numbers: []int{2, 3}, workerCount: 4, want: 2},
		{name: "empty slice", numbers: []int{}, workerCount: 4, want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := sumEven(test.numbers, test.workerCount)
			if err != nil {
				t.Fatalf("sumEven() returned an unexpected error: %v", err)
			}
			if got != test.want {
				t.Errorf("sumEven() = %d, want %d", got, test.want)
			}
		})
	}
}

func TestSumEvenRejectsInvalidWorkerCount(t *testing.T) {
	if _, err := sumEven([]int{2, 4}, 0); err == nil {
		t.Fatal("sumEven() expected an error for a non-positive worker count")
	}
}

func TestRunReadsWhitespaceSeparatedIntegers(t *testing.T) {
	var output bytes.Buffer
	err := run(strings.NewReader("1 2\n3\t4 -6"), &output, 4)
	if err != nil {
		t.Fatalf("run() returned an unexpected error: %v", err)
	}
	if got, want := output.String(), "0\n"; got != want {
		t.Errorf("run() output = %q, want %q", got, want)
	}
}

func TestRunRejectsInvalidInput(t *testing.T) {
	err := run(strings.NewReader("1 nope 4"), &bytes.Buffer{}, 4)
	if err == nil {
		t.Fatal("run() expected an error for invalid input")
	}
}
