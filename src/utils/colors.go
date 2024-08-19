package utils

import "fmt"

const (
	black = "^0"
	red = "^1"
	green = "^2"
	yellow = "^3"
	blue = "^4"
	cyan = "^5"
	magenta = "^6"
	white = "^7"
	bronze = "^8"
)

func Black(str string) string {
	return fmt.Sprintf("%s%s%s", black, str, yellow)
}

func Red(str string) string {
	return fmt.Sprintf("%s%s%s", red, str, yellow)
}

func Green(str string) string {
	return fmt.Sprintf("%s%s%s", green, str, yellow)
}

func Yellow(str string) string {
	return fmt.Sprintf("%s%s%s", yellow, str, yellow)
}

func Blue(str string) string {
	return fmt.Sprintf("%s%s%s", blue, str, yellow)
}

func Cyan(str string) string {
	return fmt.Sprintf("%s%s%s", cyan, str, yellow)
}

func Magenta(str string) string {
	return fmt.Sprintf("%s%s%s", magenta, str, yellow)
}

func White(str string) string {
	return fmt.Sprintf("%s%s%s", white, str, yellow)
}