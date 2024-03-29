package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/Gifted-s/sunkdb/core"
	prompt "github.com/c-bata/go-prompt"
	"github.com/marianogappa/sqlparser"
	"github.com/rs/zerolog/log"
)

func completer (d prompt.Document)  []prompt.Suggest{
	suggestions := []prompt.Suggest{}
	for _, s := range suggestionsMap{
		suggestions = append(suggestions, s...)
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func getExecutor(file string) func(string){
	// TODO: setup file for database
	t := &core.Table{}
	return func(s string) {
		s = strings.TrimSpace(s)
		s = strings.ToLower(s)

		switch s {
		case "":
			return 
		case ".quit", ".exit":
			log.Info().Msg("Goodbye!")
			os.Exit(0)
		default:
			// TODO Process command here
			if strings.HasPrefix(s, "."){
				log.Error().Msg(fmt.Sprintf("Unrecognized command '%s'", s))
				break
			}
			// TODO Prepare statement with sql compiler
			q, err := sqlparser.Parse(s)
			if err == nil {
				core.ExecuteStatement(q, t)
			}else{
				log.Error().Msg(fmt.Sprintf("unrecognized keyword atstmt of '%s'", s))
			}
		}
	}
}

func StartPrompt(file string) {
	 p:= prompt.New(
		getExecutor(file),
		completer,
		prompt.OptionTitle("Welcome to SQL DB"),
		prompt.OptionPrefix("sql > "),
	 )

	 p.Run()
}

