package main

import (
	"flag"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/c-bata/go-prompt"
	"github.com/gozeloglu/gitignore/internal/http"
	"log"
	"strings"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		var opts []string
		p := &survey.MultiSelect{
			Message:  "Which languages, frameworks, or OSs do you want to select?",
			Options:  promptOpts,
			PageSize: 20,
			Default:  []string{},
		}
		err := survey.AskOne(p, &opts)
		if err != nil {
			log.Fatalln(err)
		}

		optsWithComma := strings.Join(opts, ",")
		gitignoreFile, err := http.GetGitignoreFiles(optsWithComma)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(gitignoreFile)
	} else if flag.Arg(0) == "cli" {
		t := prompt.Input(">>> ", completer, prompt.OptionPrefixTextColor(prompt.Blue))
		inp := strings.Split(t, " ")
		var opts []string
		for _, s := range inp {
			s = strings.Trim(s, " ")
			if s == "" {
				continue
			}
			opts = append(opts, s)
		}
		optsWithComma := strings.Join(opts, ",")
		gitignore, err := http.GetGitignoreFiles(optsWithComma)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(gitignore)
	} else {
		fmt.Println(flag.Arg(0), "is wrong argument")
	}
}

// completer is a prompt.Completer type function that enables us to complete
// options.
func completer(d prompt.Document) []prompt.Suggest {
	s := prepSuggestions()
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursorUntilSeparator(" "), true)
}

// prepSuggestions creates a prompt.Suggest slice which contains techs.
func prepSuggestions() []prompt.Suggest {
	var promptSuggest []prompt.Suggest
	for _, opt := range promptOpts {
		promptSuggest = append(promptSuggest, prompt.Suggest{Text: opt})
	}
	return promptSuggest
}
