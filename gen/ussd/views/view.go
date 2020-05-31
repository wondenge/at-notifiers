// Code generated by goa v3.1.3, DO NOT EDIT.
//
// ussd views
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// USSDResponse is the viewed result type that is projected based on a view.
type USSDResponse struct {
	// Type to project
	Projected *USSDResponseView
	// View to render
	View string
}

// USSDResponseView is a type that runs validations on a projected type.
type USSDResponseView struct {
	// Plain text response back to AT gateway
	Response *string
}

var (
	// USSDResponseMap is a map of attribute names in result type USSDResponse
	// indexed by view name.
	USSDResponseMap = map[string][]string{
		"default": []string{
			"response",
		},
	}
)

// ValidateUSSDResponse runs the validations defined on the viewed result type
// USSDResponse.
func ValidateUSSDResponse(result *USSDResponse) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateUSSDResponseView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateUSSDResponseView runs the validations defined on USSDResponseView
// using the "default" view.
func ValidateUSSDResponseView(result *USSDResponseView) (err error) {

	return
}
