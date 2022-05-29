//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour   int
	minute int
	second int
}

type TimeParseError struct {
	errMsg string
	input  string
}

func (err *TimeParseError) Error() string {
	return fmt.Sprintf("%v : %v", err.errMsg, err.input)
}

func ParseTime(input string) (Time, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{}, &TimeParseError{"Invalid format. Please enter HH:MM:SS", input}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{"Invalid hour", input}
		}
		minute, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &TimeParseError{"Invalid minute", input}
		}
		second, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{"Invalid second", input}
		}
		if hour > 23 && hour < 0 {
			return Time{}, &TimeParseError{"Hour out of range", input}
		}
		if minute > 59 && minute < 0 {
			return Time{}, &TimeParseError{"Minute out of range", input}
		}
		if second > 59 && second < 0 {
			return Time{}, &TimeParseError{"Seconds out of range", input}
		}
		return Time{hour: hour, minute: minute, second: second}, nil
	}
}
