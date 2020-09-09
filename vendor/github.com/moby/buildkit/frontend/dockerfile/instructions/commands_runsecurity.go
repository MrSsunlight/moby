// +build dfrunsecurity

package instructions

import (
<<<<<<< HEAD
	"encoding/csv"
	"strings"

=======
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
	"github.com/pkg/errors"
)

const (
	SecurityInsecure = "insecure"
	SecuritySandbox  = "sandbox"
)

var allowedSecurity = map[string]struct{}{
	SecurityInsecure: {},
	SecuritySandbox:  {},
}

func isValidSecurity(value string) bool {
	_, ok := allowedSecurity[value]
	return ok
}

<<<<<<< HEAD
type securityKeyT string

var securityKey = securityKeyT("dockerfile/run/security")
=======
var securityKey = "dockerfile/run/security"
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375

func init() {
	parseRunPreHooks = append(parseRunPreHooks, runSecurityPreHook)
	parseRunPostHooks = append(parseRunPostHooks, runSecurityPostHook)
}

func runSecurityPreHook(cmd *RunCommand, req parseRequest) error {
	st := &securityState{}
<<<<<<< HEAD
	st.flag = req.flags.AddStrings("security")
=======
	st.flag = req.flags.AddString("security", SecuritySandbox)
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
	cmd.setExternalValue(securityKey, st)
	return nil
}

func runSecurityPostHook(cmd *RunCommand, req parseRequest) error {
<<<<<<< HEAD
	st := getSecurityState(cmd)
=======
	st := cmd.getExternalValue(securityKey).(*securityState)
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
	if st == nil {
		return errors.Errorf("no security state")
	}

<<<<<<< HEAD
	for _, value := range st.flag.StringValues {
		csvReader := csv.NewReader(strings.NewReader(value))
		fields, err := csvReader.Read()
		if err != nil {
			return errors.Wrap(err, "failed to parse csv security")
		}

		for _, field := range fields {
			if !isValidSecurity(field) {
				return errors.Errorf("security %q is not valid", field)
			}

			st.security = append(st.security, field)
		}
	}

	return nil
}

func getSecurityState(cmd *RunCommand) *securityState {
	v := cmd.getExternalValue(securityKey)
	if v == nil {
		return nil
	}
	return v.(*securityState)
}

func GetSecurity(cmd *RunCommand) []string {
	return getSecurityState(cmd).security
=======
	value := st.flag.Value
	if !isValidSecurity(value) {
		return errors.Errorf("security %q is not valid", value)
	}

	st.security = value

	return nil
}

func GetSecurity(cmd *RunCommand) string {
	return cmd.getExternalValue(securityKey).(*securityState).security
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
}

type securityState struct {
	flag     *Flag
<<<<<<< HEAD
	security []string
=======
	security string
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
}
