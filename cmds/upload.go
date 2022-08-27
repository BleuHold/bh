package cmds

import (
	"errors"
	"fmt"
	"github.com/bleuhold/bh/ecsv"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"os"
	"strings"
)

// uploadExecute is the function executed when the upload command is called.
func uploadExecute(cmd *cli.Command) error {
	// since both -f and -file point to variable s1
	switch {
	case help:
		cmd.PrintHelp()
		return nil
	default:
		xb, err := validateCSV(&s1)
		if err != nil {
			return err
		}
		err = marshalCSV(xb)
		return err
	}
}

// validateCSV validates that the path points to a CSV file.
//
// Validates:
// 1. The path is a path to a file.
// 2. The file extension is csv.
// 3. Finally, read the file and return the []bytes or the error.
func validateCSV(path *string) ([]byte, error) {
	fileInfo, err := os.Stat(*path)
	if err != nil {
		return []byte{}, err
	}
	if fileInfo.IsDir() {
		return []byte{}, errors.New("invalid path: points to a directory not a file")
	}
	s := strings.Split(fileInfo.Name(), ".")
	// get the file extension
	ext := s[len(s)-1]
	ext = strings.ToLower(ext)
	if ext != "csv" {
		return []byte{}, fmt.Errorf("invalid file extension: expected '%s' got '%s'", "csv", ext)
	}
	xb, err := os.ReadFile(*path)
	return xb, err
}

func marshalCSV(xb []byte) error {
	c := ecsv.CSV{
		StartOffset: 2, // for Investec CSV files
	}
	c.ReadData(xb)
	xt := make(Transactions, 0)
	err := xt.MarshalCSV("investec", uuid.New(), &c)
	if err != nil {
		return err
	}
	fmt.Println(xt.String())
	//for _, r := range c.Records {
	//	fmt.Printf("* %v *\n", r)
	//}
	return nil
}
