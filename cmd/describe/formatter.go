// This is just a file to store the printFormatted function
// that will be used throughout this command tree

package describe

import (
	"fmt"
	"strings"
)

func printFormatted(title string, keys []string, values []interface{}) {
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("-", len(title)+4))
	for i, key := range keys {
		var stringified string = fmt.Sprint(values[i])
		if strings.Contains(stringified, "\n") {
			stringified = formatMultiline(stringified)
		}
		fmt.Printf("  %s: %s\n", key, stringified)
	}
}

func formatMultiline(str string) string {
	var unpacked []string = strings.Split(str, "\n")
	var result string = "\n"
	for _, line := range unpacked {
		result += "    " + line + "\n"
	}
	return strings.TrimSuffix(result, "\n")
}
