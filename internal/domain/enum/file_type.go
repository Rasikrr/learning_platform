package enum

//go:generate enumer -type=FileType -json -trimprefix FileType -transform=snake -output file_type_enumer.go
type FileType uint8

const (
	FileTypeUnknown FileType = iota
	FileTypeAvatar
	FileTypeCourse
	FileTypeFaq
)
