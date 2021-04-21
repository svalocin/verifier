package verifier

import (
	"fmt"
	"reflect"
)

type InterfaceVerifier struct {
}

func (vr InterfaceVerifier) NotNil(v interface{}, msg string) error {
	return Verify(func() bool {
		if v == nil {
			return false
		}

		vi := reflect.ValueOf(v)
		if vi.Kind() == reflect.Ptr {
			return !vi.IsNil()
		}

		return true
	}, msg)
}

func (vr InterfaceVerifier) NotNilN(v interface{}, name string) error {
	return vr.NotNil(v, fmt.Sprintf(MessageNotNil, name))
}

func (vr InterfaceVerifier) NotNilP(v interface{}, msg string) {
	errorToPanic(vr.NotNil(v, msg))
}

func (vr InterfaceVerifier) NotNilNP(v interface{}, name string) {
	errorToPanic(vr.NotNilN(v, name))
}
