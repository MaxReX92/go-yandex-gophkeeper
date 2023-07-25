package commands

const (
	allFullArgName        = "--all"
	commentFullArgName    = "--comment"
	commentShortArgName   = "-c"
	cvvFullArgName        = "--cvv"
	helpFullArgName       = "--help"
	helpShortArgName      = "-h"
	idFullArgName         = "--identity"
	idShortArgName        = "-id"
	passFullArgName       = "--password"
	passShortArgName      = "-p"
	numFullArgName        = "--number"
	numShortArgName       = "-n"
	revealFullArgName     = "--reveal"
	revealShortArgName    = "-r"
	validThruFullArgName  = "--validThru"
	validThruShortArgName = "-v"
	verboseFullArgName    = "--verbose"
	verboseShortArgName   = "-v"
	versionFullArgName    = "--version"
	userFullArgName       = "--user"
	userShortArgName      = "-u"
)

type argument struct {
	keys           []string
	description    string
	nextArgIsValue bool
}

func newArgument(description string, nextArgIsValue bool, keyNames ...string) *argument {
	return &argument{
		keys:           keyNames,
		description:    description,
		nextArgIsValue: nextArgIsValue,
	}
}

func (a *argument) Keys() []string {
	return a.keys
}

func (a *argument) Description() string {
	return a.description
}

func (a *argument) NextArgIsValue() bool {
	return a.nextArgIsValue
}

func argValue(args map[string]string, keyNames ...string) (string, bool) {
	var value string
	var ok bool

	for _, keyName := range keyNames {
		value, ok = args[keyName]
		if ok {
			break
		}
	}

	return value, ok
}
