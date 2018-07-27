package commands

type AutoCompleteConfig struct {
	Available bool
	OnlyDirs  bool
}

type Command interface {
	Supports(string) bool
	Handle([]string)
	Verify([]string) error
	CommandString() string
	GetAutoCompleteConfig() AutoCompleteConfig
}
