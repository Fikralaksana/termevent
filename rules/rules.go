package rules

type Rule struct {
	Condition      []byte
	Accomplishment []byte
}

// Create creates a new rule
func Create(condition []byte) Rule {
	return Rule{Condition: condition, Accomplishment: []byte("")}
}
