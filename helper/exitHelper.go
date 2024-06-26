package helper

import (
	"fmt"
	"io"
	"os"
)

var ExitFunc = os.Exit

func Exitf(w io.Writer, format string, params ...any) {
	_, err := w.Write([]byte(fmt.Sprintf(format, params...) + "\n"))
	if err != nil {
		panic(err)
	}

	ExitFunc(1)
}
