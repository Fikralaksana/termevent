package rules

// Tracker is a struct that contains a list of rules
type Tracker struct {
	Rules []Rule
}

// Track tracks the input and checks if the condition is met
func (t *Tracker) Track(input []byte) {
	for i := range t.Rules {
		if string(t.Rules[i].Pattern) == string(input) {
			t.Rules[i].Accomplishment = true
			for _, action := range t.Rules[i].Actions {
				action.GetSpecificAction().Execute()
			}
		}
	}
}
