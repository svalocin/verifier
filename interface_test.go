package verifier

import (
	"fmt"
	"strconv"
	"testing"
)

func TestInterfaceVerifier_NotNil(t *testing.T) {
	type i interface{}
	var emptyPoit *int
	var emptyIntf1 i = emptyPoit
	var emptyIntf2 i
	const s = "x"

	type args struct {
		v interface{}
		s string
	}
	tests := []struct {
		vr         InterfaceVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: InterfaceVerifier{}, args: args{v: 1, s: s}, wantErrMsg: s, wantErr: false},
		{vr: InterfaceVerifier{}, args: args{v: nil, s: s}, wantErrMsg: s, wantErr: true},
		{vr: InterfaceVerifier{}, args: args{v: emptyPoit, s: s}, wantErrMsg: s, wantErr: true},
		{vr: InterfaceVerifier{}, args: args{v: emptyIntf1, s: s}, wantErrMsg: s, wantErr: true},
		{vr: InterfaceVerifier{}, args: args{v: emptyIntf2, s: s}, wantErrMsg: s, wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.NotNil(tt.args.v, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("InterfaceVerifier.NotNil() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("InterfaceVerifier.NotNil() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestInterfaceVerifier_NotNilN(t *testing.T) {
	type i interface{}
	var emptyPoit *int
	var emptyIntf1 i = emptyPoit
	var emptyIntf2 i
	const s = "x"

	type args struct {
		v interface{}
		s string
	}
	tests := []struct {
		vr         InterfaceVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: InterfaceVerifier{}, args: args{v: 1, s: s}, wantErrMsg: fmt.Sprintf(MessageNotNil, s), wantErr: false},
		{vr: InterfaceVerifier{}, args: args{v: nil, s: s}, wantErrMsg: fmt.Sprintf(MessageNotNil, s), wantErr: true},
		{vr: InterfaceVerifier{}, args: args{v: emptyPoit, s: s}, wantErrMsg: fmt.Sprintf(MessageNotNil, s), wantErr: true},
		{vr: InterfaceVerifier{}, args: args{v: emptyIntf1, s: s}, wantErrMsg: fmt.Sprintf(MessageNotNil, s), wantErr: true},
		{vr: InterfaceVerifier{}, args: args{v: emptyIntf2, s: s}, wantErrMsg: fmt.Sprintf(MessageNotNil, s), wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.NotNilN(tt.args.v, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("InterfaceVerifier.NotNilN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("InterfaceVerifier.NotNilN() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestInterfaceVerifier_NotNilP(t *testing.T) {
	type i interface{}
	var emptyPoit *int
	var emptyIntf1 i = emptyPoit
	var emptyIntf2 i
	const s = "x"

	type args struct {
		v interface{}
		s string
	}
	tests := []struct {
		vr         InterfaceVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: InterfaceVerifier{}, args: args{v: 1, s: s}, wantErrMsg: s, wantErr: false},
		{vr: InterfaceVerifier{}, args: args{v: nil, s: s}, wantErrMsg: s, wantErr: true},
		{vr: InterfaceVerifier{}, args: args{v: emptyPoit, s: s}, wantErrMsg: s, wantErr: true},
		{vr: InterfaceVerifier{}, args: args{v: emptyIntf1, s: s}, wantErrMsg: s, wantErr: true},
		{vr: InterfaceVerifier{}, args: args{v: emptyIntf2, s: s}, wantErrMsg: s, wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("InterfaceVerifier.NotNilP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("InterfaceVerifier.NotNilP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("InterfaceVerifier.NotNilP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("InterfaceVerifier.NotNilP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()

			tt.vr.NotNilP(tt.args.v, tt.args.s)
		})
	}
}

func TestInterfaceVerifier_NotNilNP(t *testing.T) {
	type i interface{}
	var emptyPoit *int
	var emptyIntf1 i = emptyPoit
	var emptyIntf2 i
	const s = "x"

	type args struct {
		v interface{}
		s string
	}
	tests := []struct {
		vr         InterfaceVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: InterfaceVerifier{}, args: args{v: 1, s: s}, wantErrMsg: fmt.Sprintf(MessageNotNil, s), wantErr: false},
		{vr: InterfaceVerifier{}, args: args{v: nil, s: s}, wantErrMsg: fmt.Sprintf(MessageNotNil, s), wantErr: true},
		{vr: InterfaceVerifier{}, args: args{v: emptyPoit, s: s}, wantErrMsg: fmt.Sprintf(MessageNotNil, s), wantErr: true},
		{vr: InterfaceVerifier{}, args: args{v: emptyIntf1, s: s}, wantErrMsg: fmt.Sprintf(MessageNotNil, s), wantErr: true},
		{vr: InterfaceVerifier{}, args: args{v: emptyIntf2, s: s}, wantErrMsg: fmt.Sprintf(MessageNotNil, s), wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("InterfaceVerifier.NotNilNP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("InterfaceVerifier.NotNilNP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("InterfaceVerifier.NotNilNP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("InterfaceVerifier.NotNilNP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()

			tt.vr.NotNilNP(tt.args.v, tt.args.s)
		})
	}
}
