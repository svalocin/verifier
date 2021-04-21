package verifier

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStringVerifier_NotBlank(t *testing.T) {
	const s = "x"

	type args struct {
		v string
		s string
	}
	tests := []struct {
		vr         StringVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: StringVerifier{}, args: args{v: "", s: s}, wantErrMsg: s, wantErr: true},
		{vr: StringVerifier{}, args: args{v: "   ", s: s}, wantErrMsg: s, wantErr: true},
		{vr: StringVerifier{}, args: args{v: "x", s: s}, wantErrMsg: s, wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.NotBlank(tt.args.v, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("StringVerifier.NotBlank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("StringVerifier.NotBlank() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestStringVerifier_NotBlankN(t *testing.T) {
	const s = "x"

	type args struct {
		v string
		s string
	}
	tests := []struct {
		vr         StringVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: StringVerifier{}, args: args{v: "", s: s}, wantErrMsg: fmt.Sprintf(MessageNotBlank, s), wantErr: true},
		{vr: StringVerifier{}, args: args{v: "   ", s: s}, wantErrMsg: fmt.Sprintf(MessageNotBlank, s), wantErr: true},
		{vr: StringVerifier{}, args: args{v: "x", s: s}, wantErrMsg: fmt.Sprintf(MessageNotBlank, s), wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.NotBlankN(tt.args.v, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("StringVerifier.NotBlankN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("StringVerifier.NotBlankN() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestStringVerifier_NotBlankP(t *testing.T) {
	const s = "x"

	type args struct {
		v string
		s string
	}
	tests := []struct {
		vr         StringVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: StringVerifier{}, args: args{v: "", s: s}, wantErrMsg: s, wantErr: true},
		{vr: StringVerifier{}, args: args{v: "   ", s: s}, wantErrMsg: s, wantErr: true},
		{vr: StringVerifier{}, args: args{v: "x", s: s}, wantErrMsg: s, wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("StringVerifier.NotBlankP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("StringVerifier.NotBlankP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("StringVerifier.NotBlankP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("StringVerifier.NotBlankP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()

			tt.vr.NotBlankP(tt.args.v, tt.args.s)
		})
	}
}

func TestStringVerifier_NotBlankNP(t *testing.T) {
	const s = "x"

	type args struct {
		v string
		s string
	}
	tests := []struct {
		vr         StringVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: StringVerifier{}, args: args{v: "", s: s}, wantErrMsg: fmt.Sprintf(MessageNotBlank, s), wantErr: true},
		{vr: StringVerifier{}, args: args{v: "   ", s: s}, wantErrMsg: fmt.Sprintf(MessageNotBlank, s), wantErr: true},
		{vr: StringVerifier{}, args: args{v: "x", s: s}, wantErrMsg: fmt.Sprintf(MessageNotBlank, s), wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("StringVerifier.NotBlankNP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("StringVerifier.NotBlankNP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("StringVerifier.NotBlankNP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("StringVerifier.NotBlankNP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()

			tt.vr.NotBlankNP(tt.args.v, tt.args.s)
		})
	}
}

func TestStringVerifier_NotEmpty(t *testing.T) {
	const s = "x"

	type args struct {
		v string
		s string
	}
	tests := []struct {
		vr         StringVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: StringVerifier{}, args: args{v: "", s: s}, wantErrMsg: s, wantErr: true},
		{vr: StringVerifier{}, args: args{v: "   ", s: s}, wantErrMsg: s, wantErr: false},
		{vr: StringVerifier{}, args: args{v: "x", s: s}, wantErrMsg: s, wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.NotEmpty(tt.args.v, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("StringVerifier.NotEmpty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("StringVerifier.NotEmpty() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestStringVerifier_NotEmptyN(t *testing.T) {
	const s = "x"

	type args struct {
		v string
		s string
	}
	tests := []struct {
		vr         StringVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: StringVerifier{}, args: args{v: "", s: s}, wantErrMsg: fmt.Sprintf(MessageNotEmpty, s), wantErr: true},
		{vr: StringVerifier{}, args: args{v: "   ", s: s}, wantErrMsg: fmt.Sprintf(MessageNotEmpty, s), wantErr: false},
		{vr: StringVerifier{}, args: args{v: "x", s: s}, wantErrMsg: fmt.Sprintf(MessageNotEmpty, s), wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.NotEmptyN(tt.args.v, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("StringVerifier.NotEmptyN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("StringVerifier.NotEmptyN() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestStringVerifier_NotEmptyP(t *testing.T) {
	const s = "x"

	type args struct {
		v string
		s string
	}
	tests := []struct {
		vr         StringVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: StringVerifier{}, args: args{v: "", s: s}, wantErrMsg: s, wantErr: true},
		{vr: StringVerifier{}, args: args{v: "   ", s: s}, wantErrMsg: s, wantErr: false},
		{vr: StringVerifier{}, args: args{v: "x", s: s}, wantErrMsg: s, wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("StringVerifier.NotEmptyP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("StringVerifier.NotEmptyP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("StringVerifier.NotEmptyP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("StringVerifier.NotEmptyP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()

			tt.vr.NotEmptyP(tt.args.v, tt.args.s)
		})
	}
}

func TestStringVerifier_NotEmptyNP(t *testing.T) {
	const s = "x"

	type args struct {
		v string
		s string
	}
	tests := []struct {
		vr         StringVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: StringVerifier{}, args: args{v: "", s: s}, wantErrMsg: fmt.Sprintf(MessageNotEmpty, s), wantErr: true},
		{vr: StringVerifier{}, args: args{v: "   ", s: s}, wantErrMsg: fmt.Sprintf(MessageNotEmpty, s), wantErr: false},
		{vr: StringVerifier{}, args: args{v: "x", s: s}, wantErrMsg: fmt.Sprintf(MessageNotEmpty, s), wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("StringVerifier.NotEmptyNP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("StringVerifier.NotEmptyNP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("StringVerifier.NotEmptyNP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("StringVerifier.NotEmptyNP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()

			tt.vr.NotEmptyNP(tt.args.v, tt.args.s)
		})
	}
}

func TestStringVerifier_Pattern(t *testing.T) {
	const s = "x"
	const p = "^[0-9]*$"

	type args struct {
		v string
		p string
		s string
	}
	tests := []struct {
		vr         StringVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: StringVerifier{}, args: args{v: "", p: p, s: s}, wantErrMsg: s, wantErr: true},
		{vr: StringVerifier{}, args: args{v: "   ", p: p, s: s}, wantErrMsg: s, wantErr: true},
		{vr: StringVerifier{}, args: args{v: "x", p: p, s: s}, wantErrMsg: s, wantErr: true},
		{vr: StringVerifier{}, args: args{v: "1", p: p, s: s}, wantErrMsg: s, wantErr: false},
		{vr: StringVerifier{}, args: args{v: " 1 ", p: p, s: s}, wantErrMsg: s, wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.Pattern(tt.args.v, tt.args.p, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("StringVerifier.Pattern() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("StringVerifier.Pattern() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestStringVerifier_PatternN(t *testing.T) {
	const s = "x"
	const p = "^[0-9]*$"

	type args struct {
		v string
		p string
		s string
	}
	tests := []struct {
		vr         StringVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: StringVerifier{}, args: args{v: "", p: p, s: s}, wantErrMsg: fmt.Sprintf(MessagePattern, s, p), wantErr: true},
		{vr: StringVerifier{}, args: args{v: "   ", p: p, s: s}, wantErrMsg: fmt.Sprintf(MessagePattern, s, p), wantErr: true},
		{vr: StringVerifier{}, args: args{v: "x", p: p, s: s}, wantErrMsg: fmt.Sprintf(MessagePattern, s, p), wantErr: true},
		{vr: StringVerifier{}, args: args{v: "1", p: p, s: s}, wantErrMsg: fmt.Sprintf(MessagePattern, s, p), wantErr: false},
		{vr: StringVerifier{}, args: args{v: " 1 ", p: p, s: s}, wantErrMsg: fmt.Sprintf(MessagePattern, s, p), wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := tt.vr.PatternN(tt.args.v, tt.args.p, tt.args.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("StringVerifier.PatternN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.wantErrMsg {
				t.Errorf("StringVerifier.PatternN() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
				return
			}
		})
	}
}

func TestStringVerifier_PatternP(t *testing.T) {
	const s = "x"
	const p = "^[0-9]*$"

	type args struct {
		v string
		p string
		s string
	}
	tests := []struct {
		vr         StringVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: StringVerifier{}, args: args{v: "", p: p, s: s}, wantErrMsg: s, wantErr: true},
		{vr: StringVerifier{}, args: args{v: "   ", p: p, s: s}, wantErrMsg: s, wantErr: true},
		{vr: StringVerifier{}, args: args{v: "x", p: p, s: s}, wantErrMsg: s, wantErr: true},
		{vr: StringVerifier{}, args: args{v: "1", p: p, s: s}, wantErrMsg: s, wantErr: false},
		{vr: StringVerifier{}, args: args{v: " 1 ", p: p, s: s}, wantErrMsg: s, wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("StringVerifier.PatternP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("StringVerifier.PatternP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("StringVerifier.PatternP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("StringVerifier.PatternP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()

			tt.vr.PatternP(tt.args.v, tt.args.p, tt.args.s)
		})
	}
}

func TestStringVerifier_PatternNP(t *testing.T) {
	const s = "x"
	const p = "^[0-9]*$"

	type args struct {
		v string
		p string
		s string
	}
	tests := []struct {
		vr         StringVerifier
		args       args
		wantErrMsg string
		wantErr    bool
	}{
		{vr: StringVerifier{}, args: args{v: "", p: p, s: s}, wantErrMsg: fmt.Sprintf(MessagePattern, s, p), wantErr: true},
		{vr: StringVerifier{}, args: args{v: "   ", p: p, s: s}, wantErrMsg: fmt.Sprintf(MessagePattern, s, p), wantErr: true},
		{vr: StringVerifier{}, args: args{v: "x", p: p, s: s}, wantErrMsg: fmt.Sprintf(MessagePattern, s, p), wantErr: true},
		{vr: StringVerifier{}, args: args{v: "1", p: p, s: s}, wantErrMsg: fmt.Sprintf(MessagePattern, s, p), wantErr: false},
		{vr: StringVerifier{}, args: args{v: " 1 ", p: p, s: s}, wantErrMsg: fmt.Sprintf(MessagePattern, s, p), wantErr: true},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					if tt.wantErr {
						t.Errorf("StringVerifier.PatternNP() no panic")
					}
					return
				}

				err, ok := r.(error)
				if !ok {
					t.Errorf("StringVerifier.PatternNP() recover not an error type: %#v", r)
					return
				}

				if (err != nil) != tt.wantErr {
					t.Errorf("StringVerifier.PatternNP() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err != nil && err.Error() != tt.wantErrMsg {
					t.Errorf("StringVerifier.PatternNP() error = %v, wantMsg %v", err.Error(), tt.wantErrMsg)
					return
				}
			}()

			tt.vr.PatternNP(tt.args.v, tt.args.p, tt.args.s)
		})
	}
}
