package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/angrycub/go-javaprops/pkg/fatal"
	jp "github.com/angrycub/go-javaprops/pkg/javaprops"
)

func main() {
	if err := Main(); err != nil {
		ec := 1
		var fe fatal.FatalError
		if errors.As(err, &fe) {
			err = fe.Unwrap()
			ec = fe.ExitCode()
		}
		fmt.Fprintf(os.Stderr, "ERROR: %v", err.Error())
		os.Exit(ec)
	}
}

func Main() error {
	var err error
	var p jp.Props

	j := jp.New()

	if p, err = j.GetProps(); err != nil {
		return fatal.NewError(err, 1)
	}

	fmt.Println(p.AsJSON())
	return nil
}
