// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package verifier

import (
	"fmt"
	"strconv"
	"testing"
)

func TestIntVerifier_Less(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: s, wantErr: true},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: s, wantErr: true},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: s, wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.Less(tt.args.v, tt.args.n, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("IntVerifier.Less() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("IntVerifier.Less() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestIntVerifier_LessN(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageLess, s, n), wantErr: true},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageLess, s, n), wantErr: true},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageLess, s, n), wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.LessN(tt.args.v, tt.args.n, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("IntVerifier.LessN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("IntVerifier.LessN() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestIntVerifier_LessP(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: s, wantErr: true},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: s, wantErr: true},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: s, wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("IntVerifier.LessP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("IntVerifier.LessP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("IntVerifier.LessP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("IntVerifier.LessP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()

			tt.vr.LessP(tt.args.v, tt.args.n, tt.args.s)
		})
	}
}

func TestIntVerifier_LessNP(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageLess, s, n), wantErr: true},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageLess, s, n), wantErr: true},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageLess, s, n), wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("IntVerifier.LessNP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("IntVerifier.LessNP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("IntVerifier.LessNP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("IntVerifier.LessNP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()
			tt.vr.LessNP(tt.args.v, tt.args.n, tt.args.s)
		})
	}
}

func TestIntVerifier_LessAndEqual(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: s, wantErr: true},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: s, wantErr: false},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: s, wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.LessAndEqual(tt.args.v, tt.args.n, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("IntVerifier.LessAndEqual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("IntVerifier.LessAndEqual() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestIntVerifier_LessAndEqualN(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageLessAndEqual, s, n), wantErr: true},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageLessAndEqual, s, n), wantErr: false},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageLessAndEqual, s, n), wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.LessAndEqualN(tt.args.v, tt.args.n, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("IntVerifier.LessAndEqualN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("IntVerifier.LessAndEqualN() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestIntVerifier_LessAndEqualP(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: s, wantErr: true},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: s, wantErr: false},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: s, wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("IntVerifier.LessAndEqualP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("IntVerifier.LessAndEqualP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("IntVerifier.LessAndEqualP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("IntVerifier.LessAndEqualP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()

			tt.vr.LessAndEqualP(tt.args.v, tt.args.n, tt.args.s)
		})
	}
}

func TestIntVerifier_LessAndEqualNP(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageLessAndEqual, s, n), wantErr: true},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageLessAndEqual, s, n), wantErr: false},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageLessAndEqual, s, n), wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("IntVerifier.LessAndEqualNP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("IntVerifier.LessAndEqualNP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("IntVerifier.LessAndEqualNP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("IntVerifier.LessAndEqualNP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()
			tt.vr.LessAndEqualNP(tt.args.v, tt.args.n, tt.args.s)
		})
	}
}

func TestIntVerifier_Greater(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: s, wantErr: false},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: s, wantErr: true},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: s, wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.Greater(tt.args.v, tt.args.n, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("IntVerifier.Greater() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("IntVerifier.Greater() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestIntVerifier_GreaterN(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageGreater, s, n), wantErr: false},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageGreater, s, n), wantErr: true},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageGreater, s, n), wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.GreaterN(tt.args.v, tt.args.n, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("IntVerifier.GreaterN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("IntVerifier.GreaterN() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestIntVerifier_GreaterP(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: s, wantErr: false},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: s, wantErr: true},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: s, wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("IntVerifier.GreaterP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("IntVerifier.GreaterP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("IntVerifier.GreaterP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("IntVerifier.GreaterP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()

			tt.vr.GreaterP(tt.args.v, tt.args.n, tt.args.s)
		})
	}
}

func TestIntVerifier_GreaterNP(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageGreater, s, n), wantErr: false},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageGreater, s, n), wantErr: true},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageGreater, s, n), wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("IntVerifier.GreaterNP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("IntVerifier.GreaterNP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("IntVerifier.GreaterNP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("IntVerifier.GreaterNP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()
			tt.vr.GreaterNP(tt.args.v, tt.args.n, tt.args.s)
		})
	}
}

func TestIntVerifier_GreaterAndEqual(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: s, wantErr: false},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: s, wantErr: false},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: s, wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.GreaterAndEqual(tt.args.v, tt.args.n, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("IntVerifier.GreaterAndEqual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("IntVerifier.GreaterAndEqual() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestIntVerifier_GreaterAndEqualN(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageGreaterAndEqual, s, n), wantErr: false},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageGreaterAndEqual, s, n), wantErr: false},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageGreaterAndEqual, s, n), wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.GreaterAndEqualN(tt.args.v, tt.args.n, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("IntVerifier.GreaterAndEqualN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("IntVerifier.GreaterAndEqualN() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestIntVerifier_GreaterAndEqualP(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: s, wantErr: false},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: s, wantErr: false},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: s, wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("IntVerifier.GreaterAndEqualP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("IntVerifier.GreaterAndEqualP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("IntVerifier.GreaterAndEqualP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("IntVerifier.GreaterAndEqualP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()

			tt.vr.GreaterAndEqualP(tt.args.v, tt.args.n, tt.args.s)
		})
	}
}

func TestIntVerifier_GreaterAndEqualNP(t *testing.T) {
	const (
		s = "x"
		n = 5
	)

	type args struct {
		v int
		n int
		s string
	}
	tests := []struct {
		vr         IntVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: IntVerifier{}, args: args{v: 6, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageGreaterAndEqual, s, n), wantErr: false},
		{vr: IntVerifier{}, args: args{v: 5, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageGreaterAndEqual, s, n), wantErr: false},
		{vr: IntVerifier{}, args: args{v: 4, n: n, s: s}, wantErrMsg: fmt.Sprintf(MessageGreaterAndEqual, s, n), wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("IntVerifier.GreaterAndEqualNP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("IntVerifier.GreaterAndEqualNP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("IntVerifier.GreaterAndEqualNP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("IntVerifier.GreaterAndEqualNP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()
			tt.vr.GreaterAndEqualNP(tt.args.v, tt.args.n, tt.args.s)
		})
	}
}
