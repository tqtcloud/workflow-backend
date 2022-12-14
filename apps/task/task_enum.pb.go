// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package task

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseJenkinsEnvFromString Parse JenkinsEnv from string
func ParseJenkinsEnvFromString(str string) (JenkinsEnv, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := JenkinsEnv_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown JenkinsEnv: %s", str)
	}

	return JenkinsEnv(v), nil
}

// Equal type compare
func (t JenkinsEnv) Equal(target JenkinsEnv) bool {
	return t == target
}

// IsIn todo
func (t JenkinsEnv) IsIn(targets ...JenkinsEnv) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t JenkinsEnv) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *JenkinsEnv) UnmarshalJSON(b []byte) error {
	ins, err := ParseJenkinsEnvFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
