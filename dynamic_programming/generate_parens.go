package dynamicprogramming

import (
	"fmt"
)

type StrState struct {
	Str      string
	Count    int
	Openings int
}

func generateParenthesis(n int) []string {
	res := make([]string, 0, n)
	states := []StrState{{"(", 1, n - 1}}
	for range n*2 - 1 {
		newStates := make([]StrState, 0)
		for _, state := range states {
			if state.Openings > 0 {
				newStates = append(
					newStates,
					StrState{
						Str:      fmt.Sprintf("%s(", state.Str),
						Count:    state.Count + 1,
						Openings: state.Openings - 1,
					},
				)
			}
			if state.Count > 0 {
				newStates = append(
					newStates,
					StrState{
						Str:      fmt.Sprintf("%s)", state.Str),
						Count:    state.Count - 1,
						Openings: state.Openings,
					},
				)
			}
		}
		states = newStates
	}

	for _, state := range states {
		res = append(res, state.Str)
	}

	return res
}
