package verifier

import (
	"errors"
	"fmt"
)

const (
	MessageLess            = "%v must be less than %v "
	MessageGreater         = "%v must be greater than %v"
	MessageLessAndEqual    = "%v must be less than or equal to %v"
	MessageGreaterAndEqual = "%v must be greater than or equal to %v"
	MessageNotNil          = "%v may not be nil"
	MessageNotBlank        = "%v may not be blank"
	MessageNotEmpty        = "%v may not be empty"
	MessagePattern         = "%v must match \"%v\""
)

func Verify(f func() bool, msg string) error {
	if f == nil {
		panic(fmt.Sprintf(MessageNotNil, "func"))
	}

	if !f() {
		return errors.New(msg)
	}

	return nil
}

func VerifyP(f func() bool, msg string) {
	errorToPanic(Verify(f, msg))
}

func errorToPanic(err error) {
	if err != nil {
		panic(err)
	}
}

//go:generate $GOPATH/bin/genny -in=number.gen -out=number.go gen "Number=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"

var stringVerifier = StringVerifier{}
var interfaceVerifier = InterfaceVerifier{}
var intVerifier = IntVerifier{}
var int8Verifier = Int8Verifier{}
var int16Verifier = Int16Verifier{}
var int32Verifier = Int32Verifier{}
var int64Verifier = Int64Verifier{}
var uintVerifier = UintVerifier{}
var uint8Verifier = Uint8Verifier{}
var uint16Verifier = Uint16Verifier{}
var uint32Verifier = Uint32Verifier{}
var uint64Verifier = Uint64Verifier{}
var float32Verifier = Float32Verifier{}
var float64Verifier = Float64Verifier{}

func S() StringVerifier {
	return stringVerifier
}
func IF() InterfaceVerifier {
	return interfaceVerifier
}
func I() IntVerifier {
	return intVerifier
}
func I8() Int8Verifier {
	return int8Verifier
}
func I16() Int16Verifier {
	return int16Verifier
}
func I32() Int32Verifier {
	return int32Verifier
}
func I64() Int64Verifier {
	return int64Verifier
}
func U() UintVerifier {
	return uintVerifier
}
func U8() Uint8Verifier {
	return uint8Verifier
}
func U16() Uint16Verifier {
	return uint16Verifier
}
func U32() Uint32Verifier {
	return uint32Verifier
}
func U64() Uint64Verifier {
	return uint64Verifier
}
func F32() Float32Verifier {
	return float32Verifier
}
func F64() Float64Verifier {
	return float64Verifier
}
