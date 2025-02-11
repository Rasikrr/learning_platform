package enum

//go:generate enumer -type=ProgrammingLanguage -json -trimprefix ProgrammingLanguage -transform=snake -output programming_language_enumer.go

type ProgrammingLanguage uint8

const (
	ProgrammingLanguageGo ProgrammingLanguage = iota
	ProgrammingLanguagePython3
	ProgrammingLanguageJava
	ProgrammingLanguageRuby
	ProgrammingLanguageSQL
	ProgrammingLanguageNodejs
	ProgrammingLanguageRust
)
