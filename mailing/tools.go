package mailing

import (
	"path/filepath"
	"github.com/resurtm/boomak-server/tools"
	"bytes"
	ht "html/template"
	tt "text/template"
	"fmt"
)

func renderTextTemplate(name string, data interface{}) string {
	tpl, err := tt.ParseFiles(filepath.Join(tools.CurrentDir(), "templates", fmt.Sprintf("%s.txt", name)))
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
	tpl, err := ht.ParseFiles(filepath.Join(tools.CurrentDir(), "templates", fmt.Sprintf("%s.html", name)))
	if err != nil {
		panic(err)
	}

	var tplBytes bytes.Buffer
	if err := tpl.Execute(&tplBytes, data); err != nil {
		panic(err)
	}

	return tplBytes.String()
}
