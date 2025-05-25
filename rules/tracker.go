package rules

import (
	"bytes"
	"fmt"
)

// Tracker is a struct that contains a list of rules
type Tracker struct {
	Rules []Rule
}

// Track tracks the input and checks if the condition is met
func (t *Tracker) Track(input []byte) {

	for i := range t.Rules {
		fmt.Println(input)
		// append input to accomplishment
		t.Rules[i].Accomplishment = append(t.Rules[i].Accomplishment, input...)

		// check if condition and accomplishment are equal
		fmt.Println(string(t.Rules[i].Condition))
		if bytes.Equal(t.Rules[i].Condition, t.Rules[i].Accomplishment) {
			fmt.Println("event triggered")
		}
	}
}
