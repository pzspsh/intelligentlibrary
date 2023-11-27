package flags

var (
	optionMap           map[*StringSlice]Options
	optionDefaultValues map[*StringSlice][]string
)

func init() {
	optionMap = make(map[*StringSlice]Options)
	optionDefaultValues = make(map[*StringSlice][]string)
}

// StringSlice is a slice of strings
type StringSlice []string

// Set appends a value to the string slice.
func (stringSlice *StringSlice) Set(value string) error {
	option, ok := optionMap[stringSlice]
	if !ok {
		option = StringSliceOptions
	}
	values, err := ToStringSlice(value, option)
	if err != nil {
		return err
	}
	// if new values are provided, we remove default ones
	if defaultValue, ok := optionDefaultValues[stringSlice]; ok {
		if Equal(*stringSlice, defaultValue) {
			*stringSlice = []string{}
		}
	}

	*stringSlice = append(*stringSlice, values...)
	return nil
}

func (stringSlice StringSlice) String() string {
	return ToString(stringSlice)
}

func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for idx := range s1 {
		if s1[idx] != s2[idx] {
			return false
		}
	}

	return true
}
