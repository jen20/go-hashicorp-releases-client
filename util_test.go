package releases_test

import (
	"fmt"
	"iter"
	"reflect"
	"testing"
)

func requireNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected no error, got: %s", err)
	}
}

func requireEqual(t *testing.T, expected any, actual any) {
	if !reflect.DeepEqual(expected, actual) {
		message := fmt.Sprintf("Not equal:\n"+
			"expected: %T(%#v)\n"+
			"  actual: %T(%#v)\n", expected, expected, actual, actual)
		t.Fatal(message)
	}
}

func collectResults[T any](t *testing.T, seq iter.Seq2[T, error]) []T {
	var result []T

	for item, err := range seq {
		requireNoError(t, err)
		result = append(result, item)
	}

	return result
}
