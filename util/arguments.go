package util

import "strings"

/*
	Parses arguments in a similar fashion to *nix command line utils.

	argument forms:
		--<name> <value>
			set name to value
		-abcdef
			set flag a, b, c, d, e and f
		a single - stops the parsing and returns the rest in rest
		also something is in rest if it doesn't start with -
	named returns the named arguments and flags contains all set flags

	more advanced functions are to come.
 */
func ParseArguments(args []string) (named map[string]string, flags []string, unnamed []string) {
	named = make(map[string]string)
	flags = make([]string, 0)
	unnamed = make([]string, 0)
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if "-" == arg {
			unnamed = append(unnamed, arg[i+1:])
			return
		}
		if strings.HasPrefix(arg, "--") { // Named argument
			name := arg[2:]
			var value string = ""
			i++
			if i < len(args) { // No value left...
				value = args[i]
			}
			named[name] = value
		} else if strings.HasPrefix(arg, "-") { // some flags
			flags = append(flags, strings.Split(arg[1:], "")...)
		} else {
			unnamed = append(unnamed, arg)
		}
	}
	return
}
