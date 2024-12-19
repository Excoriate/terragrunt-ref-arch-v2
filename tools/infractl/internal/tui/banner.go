package tui

import (
	"strings"

	"github.com/common-nighthawk/go-figure"
)

const (
	AppName = "InfraCTL"
)

// GetBanner generates a stylized ASCII art banner for the application.
//
// This function creates a large, visually appealing text representation of the application name
// using the "doom" font style from the go-figure library. It converts the application name to
// uppercase to enhance visibility and impact.
//
// The banner is typically used to display a prominent, eye-catching title when the CLI is launched,
// providing a professional and engaging user interface element.
//
// Returns:
//   - A string containing the ASCII art representation of the application name.
//
// Example:
//
//	banner := GetBanner()
//	fmt.Println(banner) // Prints a large, styled banner with the app name
func GetBanner() string {
	// Convert the application name to uppercase for better visual impact
	appNameUpper := strings.ToUpper(AppName)

	// Generate and return an ASCII art banner using the "doom" font style
	return figure.NewFigure(appNameUpper, "doom", true).String()
}
