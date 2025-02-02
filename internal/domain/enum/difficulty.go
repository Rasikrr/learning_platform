package enum

//go:generate enumer -type=Difficulty -json -trimprefix Difficulty -transform=snake -output difficulty_enumer.go

type Difficulty uint8

const (
	DifficultyEasy Difficulty = iota
	DifficultyMedium
	DifficultyHard
)
