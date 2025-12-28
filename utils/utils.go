package utils

import (
	"strings"

	"github.com/manifoldco/promptui"
)



var Colors = map[string]string{
	"White": "\033[97m",
	"Red": "\033[31m",
	"Green": "\033[32m",
	"Yellow": "\033[33m",
	"Blue": "\033[34m",
	"Magenta": "\033[35m",
	"Cyan": "\033[36m",
	"Gray": "\033[37m",
	"Reset": "\033[0m",
}

type PromptOpts struct {
	Message string
	Items []string
	UseSearcher bool
}

func PromptSelect(opts PromptOpts) (string, error){
	var prompt promptui.Select
	if (opts.UseSearcher){
		searcher := func (curr string, ind int) bool {
			curr = strings.ToLower(curr)
			selected := strings.ToLower(opts.Items[ind])
			return strings.Contains(selected, curr)
		}	
		prompt = promptui.Select{
			Label: opts.Message,
			Searcher: searcher,
			Items: opts.Items,
			StartInSearchMode:  true,
		}
	} else {
		prompt = promptui.Select{
			Label: opts.Message,
			Items: opts.Items,
		}
	}

	_, selected, err := prompt.Run()

	return selected, err
}

func PromptConfirm(message string) (error){
	prompt := promptui.Prompt{
		Label:     message + " (y/N)",
		IsConfirm: true,
	}

	_, err := prompt.Run()
	return err
}

func MakeColored(color string, text string) string {
	colorCode, exists := Colors[color]
	if !exists {
		return text
	}
	return colorCode + text + Colors["White"]
}