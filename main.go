package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

const htmlPage = `
<!DOCTYPE html>
<html>
<head>
    <title>go-magic</title>
</head>
<body>
    <h3 align=center>Render-go-backend</h3>
    <form action="/" method="post">
        <label for="myText01">Enter Text:</label>
        <input type="text" id="myText01" name="myText01" value="caleb">
        <input type="submit" value="Submit">
    </form>
    {{.Result}}
</body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	result := `<span style='color:red'>Try the magic word "caleb"</span>`
	if r.Method == "POST" && r.FormValue("myText01") == "caleb" {
		result = `<b style='color:green'>Cool!</b>`
	}

	t := template.Must(template.New("page").Parse(htmlPage))
	t.Execute(w, map[string]string{"Result": result})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	http.HandleFunc("/", handler)

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}