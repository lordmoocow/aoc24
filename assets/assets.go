package assets

import (
	"embed"
	"fmt"
)

//go:embed *
var assets embed.FS

func TestData(day int) string {
	bytes, _ := assets.ReadFile(filePath(day, "test"))
	return string(bytes)
}

func InputData(day int) string {
	bytes, _ := assets.ReadFile(filePath(day, "input"))
	return string(bytes)
}

func filePath(day int, dataType string) string {
	return fmt.Sprintf("days/%02d/%s", day, dataType)
}
