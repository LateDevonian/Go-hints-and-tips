package main

import (
	"html/template"
	"net/http"
	"time"
)

var tpl = `<!DOCTYPE HTML>   
<html>
	<head>
		<meta charset="utf-8">
		<title>Date Examle</title>
	</head>
	<body>
		<p>{{.Date | dateFormat "Jan 2, 2006"}}</p>
	</body>
</html>`

//map go function to template function
var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

//covert a time to a formatted string
func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}

func serveTemplate(res http.ResponseWriter, req *http.Request) {
	t := template.New("date") //create template instance
	t.Funcs(funcMap)          //pass additional functions in map to template engine
	t.Parse(tpl)              //parse template string into template engine
	data := struct{ Date time.Time }{
		Date: time.Now(),
	}
	t.Execute(res, data) //send template with data to output response
}

func main() {
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":8080", nil)
}
