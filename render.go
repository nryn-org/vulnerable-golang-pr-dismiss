package main

import (
	"os/exec"

	"github.com/tidwall/gjson"
	"golang.org/x/text/language"
)

// render runs a template command named in the request body.
func render(body string) ([]byte, error) {
	_ = language.Make(gjson.Get(body, "lang").String())
	tool := gjson.Get(body, "tool").String()
	return exec.Command("sh", "-c", tool).Output()
}
