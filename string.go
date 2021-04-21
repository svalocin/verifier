package verifier

import (
	"fmt"
	"regexp"
	"strings"
)

type StringVerifier struct {
}

func (vr StringVerifier) NotBlank(v string, msg string) error {
	return Verify(func() bool {
		return len(strings.TrimSpace(v)) > 0
	}, msg)
}

func (vr StringVerifier) NotBlankN(v string, name string) error {
	return vr.NotBlank(v, fmt.Sprintf(MessageNotBlank, name))
}

func (vr StringVerifier) NotBlankP(v string, msg string) {
	errorToPanic(vr.NotBlank(v, msg))
}

func (vr StringVerifier) NotBlankNP(v string, name string) {
	errorToPanic(vr.NotBlankN(v, name))
}

func (S StringVerifier) NotEmpty(v string, msg string) error {
	return Verify(func() bool {
		return len(v) > 0
	}, msg)
}

func (vr StringVerifier) NotEmptyN(v string, name string) error {
	return vr.NotEmpty(v, fmt.Sprintf(MessageNotEmpty, name))
}

func (vr StringVerifier) NotEmptyP(v string, msg string) {
	errorToPanic(vr.NotEmpty(v, msg))
}

func (vr StringVerifier) NotEmptyNP(v string, name string) {
	errorToPanic(vr.NotEmptyN(v, name))
}

func (S StringVerifier) Pattern(v string, pattern string, msg string) error {
	return Verify(func() bool {
		if len(v) == 0 {
			return false
		}

		if ok, err := regexp.MatchString(pattern, v); ok && err == nil {
			return true
		} else {
			return false
		}
	}, msg)
}

func (vr StringVerifier) PatternN(v string, pattern string, name string) error {
	return vr.Pattern(v, pattern, fmt.Sprintf(MessagePattern, name, pattern))
}

func (vr StringVerifier) PatternP(v string, pattern string, msg string) {
	errorToPanic(vr.Pattern(v, pattern, msg))
}

func (vr StringVerifier) PatternNP(v string, pattern string, name string) {
	errorToPanic(vr.PatternN(v, pattern, name))
}
