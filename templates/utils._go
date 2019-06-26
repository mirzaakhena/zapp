package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	gonanoid "github.com/matoous/go-nanoid"
	uuid "github.com/satori/go.uuid"
)

// GenID is
func GenID() string {
	x, _ := uuid.NewV4()
	return x.String()
}

// GetJSON is
func GetJSON(x interface{}) string {
	y, _ := json.Marshal(x)
	return string(y)
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// SnakeCase is
func SnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// PrintError is
func PrintError(code int, message string, args ...interface{}) error {
	return fmt.Errorf("Error (%d). %s", code, fmt.Sprintf(message, args...))
}

const nanoidstr = "A2BC3DE4FG5HJ6KL7MN8PQ9RSTUVWXYZ"

// GenerateUniqueID is
func GenerateUniqueID() string {
	// return ksuid.New().String()
	a, _ := gonanoid.Generate(nanoidstr, 12)
	return a
}
