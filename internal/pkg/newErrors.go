package pkg

import "errors"

var (
	invalidNumberOfLines = errors.New("No required lines in file")
	BadFormatOfLine      = errors.New("Bad format of file on line:")
	invalidNumberOfArgs  = errors.New("invalid number of arguments in command line")
	invalidFileFormat    = errors.New("file must has \"txt\" format")
)
