package mailing

import (
	"path/filepath"
	"github.com/resurtm/boomak-server/common"
	"bytes"
	html "html/template"
	text "text/template"
	"fmt"
)

const emailTemplatesDir = "templates"

func renderTextTemplate(name string, data interface{}) string {
	tpl, err := text.ParseFiles(
		filepath.Join(common.CurrentDir(), emailTemplatesDir, fmt.Sprintf("%s.txt", name)),
	)
	if err != nil {
		panic(err)
	}

	var tplBytes bytes.Buffer
	if err := tpl.Execute(&tplBytes, data); err != nil {
		panic(err)
	}
	return tplBytes.String()
}

func renderHtmlTemplate(name string, data interface{}) string {
	tpl, err := html.ParseFiles(
		filepath.Join(common.CurrentDir(), emailTemplatesDir, fmt.Sprintf("%s.html", name)),
	)
	if err != nil {
		panic(err)
	}

	var tplBytes bytes.Buffer
	if err := tpl.Execute(&tplBytes, data); err != nil {
		panic(err)
	}
	return tplBytes.String()
}
