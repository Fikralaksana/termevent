package rules

import "log"

const (
	LogActionName = "log"
)

type Action struct {
	Name string                 `yaml:"name"`
	Data map[string]interface{} `yaml:",inline"`
}

type LogAction struct {
	Name    string `yaml:"name"`
	Message string `yaml:"message"`
}

func (a *LogAction) Execute() {
	log.Println(a.Message)
}

type ActionExecutor interface {
	Execute()
}

func (a *Action) GetSpecificAction() ActionExecutor {
	switch a.Name {
	case LogActionName:
		return &LogAction{
			Name:    a.Name,
			Message: a.Data["message"].(string),
		}
	}
	return nil
}
