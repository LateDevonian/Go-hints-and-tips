package main

import (
	"bytes"
	"testing"
	"text/template"
)

//parallel benchmark tests how a given peice of code performs when on goroutines
//test with variable cpus
//testing.B instance provides a RunParallel method.
func ParallellBenchmarkTemplates(b *testing.B) {
	tpl := "hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}

	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		//run the core of the test b.N times
		for pb.Next() {
			t.Execute(&buf, data)
			buf.Reset()
		}
	})

}
