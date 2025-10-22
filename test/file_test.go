package test

import (
	"testing"

	"oot.me/link-parser/utils"
)

func TestSuccessReadHTMLFile(t *testing.T) {
	_, err := utils.ReadHTMLFile("html_files/ex1.html")
	if err != nil {
		t.Errorf("Cannot open fil html_files/ex1.html")
	}
}

func TestFailedReadHTMLFile(t *testing.T) {
	_, err := utils.ReadHTMLFile("html_files/xxx.html")
	if err == nil {
		t.Errorf("Error should raise")
	}
}
