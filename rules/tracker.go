package rules

import (
	"fmt"
)

// Tracker is a struct that contains a list of rules
type Tracker struct {
	Rules []Rule
}

// Track tracks the input and checks if the condition is met
func (t *Tracker) Track(input []byte) {
	for i := range t.Rules {
		if string(t.Rules[i].Condition) == string(input) {
			t.Rules[i].Accomplishment = true
			fmt.Println("event triggered")
		}
	}
}
