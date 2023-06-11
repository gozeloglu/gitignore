package main

import (
	"flag"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/c-bata/go-prompt"
	"github.com/gozeloglu/gitignore/internal/file"
	"github.com/gozeloglu/gitignore/internal/http"
	"log"
	"os"
	"path"
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
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}

		err = file.Save(path.Join(dir, ".gitignore"), []byte(gitignoreFile))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("The .gitignore file has been created and saved in %s.\n", dir)
	} else if flag.Arg(0) == "cli" {
		t := prompt.Input(">>> ", completer, prompt.OptionPrefixTextColor(prompt.Blue))
		if strings.Trim(t, "") == "" {
			fmt.Println("No input provided. Please enter at least one programming language, operating system, framework, IDE, editor, or any other relevant information.")
			os.Exit(1)
		}
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
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}

		err = file.Save(path.Join(dir, ".gitignore"), []byte(gitignore))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("The .gitignore file has been created and saved in %s.\n", dir)
	} else {
		fmt.Println(flag.Arg(0), "is wrong argument")
		os.Exit(1)
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
