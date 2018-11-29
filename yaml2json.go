package yaml2json

import (
	"io"
	"io/ioutil"
)

func Convert(fn func(j []byte) ([]byte, error), in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)

	if err != nil {
		return err
	}

	output, err := fn(input)

	if err != nil {
		return err
	}

	_, err = out.Write(output)

	if err != nil {
		return err
	}

	return nil
}
