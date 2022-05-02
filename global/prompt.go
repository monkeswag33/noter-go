package global

import (
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
)

var Templates *promptui.PromptTemplates = &promptui.PromptTemplates{
	Prompt:  "{{ . }} ",
	Valid:   "{{ . | green }} ",
	Invalid: "{{ . | red }} ",
	Success: "{{ . | bold }} ",
}

func Prompt(config promptui.Prompt, label string, validator func(string) error) string {
	config.Label = label
	config.Templates = Templates
	config.Validate = validator
	result, err := config.Run()
	if err != nil {
		logrus.Fatalf("Error while getting user input: %v\n", err)
	}
	return result
}
