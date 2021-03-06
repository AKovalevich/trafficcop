package entrypoint

import (
	"strings"
	"regexp"
	"fmt"

	"github.com/AKovalevich/trafficcop/pkg/entrypoint/api"
	"github.com/AKovalevich/trafficcop/pkg/entrypoint/admin"
	"github.com/AKovalevich/trafficcop/pkg/route"
)

// Base Entrypoint interface
type Entrypoint interface {
	// Start entrypoint
	Start()
	// Stop enptrypoint
	Stop()
	// Initialize entrypoint
	Init()
	// Get entrypoint name
	String() string
	// Get entrypoint routes list
	RoutesList() []route.Route
}

type EntrypointList []Entrypoint

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (e *EntrypointList) Set(value string) error {
	entrypoints := strings.Split(parseEntryPoints(value), ",")

	if len(entrypoints) == 0 {
		return fmt.Errorf("bad EntryPointList format: %s", value)
	}

	for _, entrypointName := range entrypoints {
		// Try to create entrypoint
		switch entrypointName {
		case "api":
			textClassifierEntrypoint := api.New()
			textClassifierEntrypoint.Name = "api"
			*e = append(*e, textClassifierEntrypoint)
			break
		case "admin":
			profanityEntrypoint := admin.New()
			profanityEntrypoint.Name = "admin"
			*e = append(*e, profanityEntrypoint)
		}
	}

	return nil
}

// Get return the EntryPoints map
func (e *EntrypointList) Get() interface{} {
	return EntrypointList(*e)
}

// SetValue sets the EntryPoints map with val
func (e *EntrypointList) SetValue(val interface{}) {
	*e = EntrypointList(val.(EntrypointList))
}

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (e *EntrypointList) String() string {
	var entrypoints []string
	for _, entrypoint := range *e {
		// Try to create entrypoint
		entrypoints = append(entrypoints, entrypoint.String())
	}

	return strings.Join(entrypoints, ", ")
}

// Type is type of the struct
func (dep *EntrypointList) Type() string {
	return "defaultentrypoints"
}

func parseEntryPoints(values string) string {
	valuesRegexp := regexp.MustCompile(`\'(.*$)`)
	return strings.Replace(valuesRegexp.FindString(values), "'", "", -1)
}
