package rules

type Rule struct {
	Condition      []byte
	Accomplishment bool
}

func CollectRules() []Rule {
	return []Rule{create([]byte("hello"))}
}

// Create creates a new rule
func create(condition []byte) Rule {
	return Rule{Condition: condition, Accomplishment: false}
}
