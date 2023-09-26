package cli

import "github.com/c-bata/go-prompt"

type suggestionType int

const (
	Commands suggestionType = iota
	Keywords
)

var commandSuggestions = []prompt.Suggest{
	{Text: ".quit", Description: "Quit/Exit application"},
	{Text: ".exit", Description: "Quit/Exit application"},
}

var keywordSuggestions = []prompt.Suggest{
	{Text: "select", Description: "Select statememt to read data from table"},
	{Text: "insert", Description: "Insert statememt to add data into table"},
}

var suggestionsMap = map[suggestionType][]prompt.Suggest{
	Commands : commandSuggestions,
	Keywords: keywordSuggestions,

}