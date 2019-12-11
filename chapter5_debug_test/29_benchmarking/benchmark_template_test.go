package main

import (
	"bytes"
	"testing"
	"text/template"
)

//run with go test -bench .

//want to know what path is faster
//this zeros in on one average time it takes to compile and run a text template
func BenchmarkTemplates(b *testing.B) {
	b.Logf("b.N is %d\n", b.N)
	tpl := "hello {{.Name}}"
	data := &map[string]string{
		"Name": "World",
	}

	var buf bytes.Buffer
	//run the core of the test b.N times
	for i := 0; i < b.N; i++ {
		t, _ := template.New("test").Parse(tpl)
		//excecute template and clears the buffer
		t.Execute(&buf, data)
		buf.Reset()
	}
}
func BenchmarkCompiledTemplates(b *testing.B) {
	b.Logf("b.N is %d\n", b.N)
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	//moves template comp out of loop
	data := &map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	//run the core of the test b.N times
	for i := 0; i < b.N; i++ {
		t.Execute(&buf, data)
		buf.Reset()
	}
}

//benchmark shows the comiled templates shavest time off average loop iteration.
