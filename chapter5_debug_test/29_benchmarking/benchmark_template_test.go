package main

import (
	"bytes"
	"testing"
	"text/template"
)

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
		t.Excecute(&buf, data)
		buf.Reset()
	}
}
