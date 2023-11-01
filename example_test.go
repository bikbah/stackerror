package stackerror_test

import (
	"errors"
	"fmt"

	"github.com/bikbah/stackerror"
)

func ExampleStackError() {
	e := errors.New("test error")

	es := stackerror.New(e)
	fmt.Println(es)
}
