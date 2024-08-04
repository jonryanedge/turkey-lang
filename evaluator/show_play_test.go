package evaluator

import (
	"testing"

	"turkey-lang/object"
)

func TestShow(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`show(5)`,
			`5`,
		},
		{
			`show(5 + 8)`,
			`(5 + 8)`,
		},
		{
			`show(foobar)`,
			`foobar`,
		},
		{
			`show(foobar + barfoo)`,
			`(foobar + barfoo)`,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		show, ok := evaluated.(*object.Show)
		if !ok {
			t.Fatalf("expected *object.Show. got=%T (%+v)", evaluated, evaluated)
		}
		if show.Node == nil {
			t.Fatalf("show.Node is nil")
		}
		if show.Node.String() != tt.expected {
			t.Errorf("not equal. got=%q, want=%q", show.Node.String(), tt.expected)
		}
	}
}

func TestShowPlay(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`show(play(4))`,
			`4`,
		},
		{
			`show(play(4 + 4))`,
			`8`,
		},
		{
			`show(8 + play(4 + 4))`,
			`(8 + 8)`,
		},
		{
			`show(play(4 + 4) + 8)`,
			`(8 + 8)`,
		},
		{
			`let foobar = 8; show(foobar)`,
			`foobar`,
		},
		{
			`let foobar = 8; show(play(foobar))`,
			`8`,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		show, ok := evaluated.(*object.Show)
		if !ok {
			t.Fatalf("expected *object.Show. got=%T (%+v)", evaluated, evaluated)
		}
		if show.Node == nil {
			t.Fatalf("show.Node is nil")
		}
		if show.Node.String() != tt.expected {
			t.Errorf("not equal. got=%q, want=%q", show.Node.String(), tt.expected)
		}
	}
}
