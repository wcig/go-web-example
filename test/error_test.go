package test

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func TestErr(t *testing.T) {
	f := func() error {
		return errors.New("pkg error")
	}
	err := f()
	fmt.Printf("%+v", err)
}
