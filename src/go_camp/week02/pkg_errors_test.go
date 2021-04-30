package go_camp

import (
	"errors"
	"fmt"
	"testing"

	xerrors "github.com/pkg/errors"
)

var errMy = errors.New("my")

func TestErrorsAndPKGErrors(t *testing.T) {
	err := test2()
	fmt.Printf("test: %+v\n", err)
}

func test0() error {
	return xerrors.Wrapf(errMy, "test0 failed")
}

func test1() error {
	return test0()
}

func test2() error {
	return test1()
}

var ErrPermission = errors.New("permission denied")
var user = User{}
// DoSomething returns an error wrapping ErrPermission if the user
// does not have permission to do something
func DoSomething() error {
	if !userHasPermission() {
		// If we return ErrPermission directly, callers might come to depend on
		// the exact error value, writing code like this:
		//
		// if err := pkg.DoSomething(); err == pkg.ErrPermission { ... }
		//
		// This will cause problems if we want to add additional
		// context to the error in the future. To avoid this, we
		// return an error wrapping the sentinel so that users must
		// always unwrap it:
		// if err := pkg.DoSomething(); errors.Is(err, pkgErrPermission) { ... }
		return fmt.Errorf("%w", ErrPermission)
	}
	return nil
}

func userHasPermission() bool {
	return user.Name == "per"
}

type User struct {
	Name string
}
