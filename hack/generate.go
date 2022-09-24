package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// https://github.com/wesbos/burner-email-providers
func main() {
	resp, err := http.Get("https://github.com/wesbos/burner-email-providers/raw/master/emails.txt")
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()

	var out bytes.Buffer
	out.WriteString(`package emailproviders

// https://github.com/wesbos/burner-email-providers
var disposableEmailServices = map[string]struct{}{
`)

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if line != "" {
			out.WriteString(fmt.Sprintf(`"%s": empty,`, line))
			out.WriteRune('\n')
		}
	}

	out.WriteString(`
}`)

	err = os.WriteFile("./disposable_email_services.go", out.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
