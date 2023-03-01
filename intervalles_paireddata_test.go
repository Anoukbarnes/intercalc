package main

import (
	"os"
	"testing"
)

func TestSaisieFloat64Positif(t *testing.T) {
	testCases := []struct {
		message string
		input   string
		want    float64
	}{
		{"Test saisie positive", "3\n", 3},
		{"Test saisie négative puis positive", "-2\n3\n", 3},
		{"Test saisie nulle puis positive", "0\n3\n", 3},
	}

	for _, tc := range testCases {
		t.Run(tc.message, func(t *testing.T) {
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("error creating pipe: %v", err)
			}
			origStdin := os.Stdin
			defer func() {
				os.Stdin = origStdin
				r.Close()
			}()
			_, err = w.WriteString(tc.input)
			if err != nil {
				t.Fatalf("error writing to pipe: %v", err)
			}
			w.Close()
			os.Stdin = r

			if got := saisieFloat64Positif("test"); got != tc.want {
				t.Errorf("saisieFloat64Positif() = %v, want %v", got, tc.want)
			}
		})
	}
}
