package bar

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func MaybePanic() {
	rand.Seed(time.Now().Unix())
	r := rand.Intn(100)
	if r > 49 {
		panic("oh no!")
	}
}

func ServerTemplate(w http.ResponseWriter) {
	t, err := template.New("foo").Parse(`Hello, {{.}}!`)
	if err != nil {
		log.Println("Could not parse template.")
	}
	err = t.Execute(w, "<script>alert('you have been pwned')</script>")
}

func init() {
	http.HandleFunc("/maybe-panic", func(w http.ResponseWriter, r *http.Request) {
		MaybePanic()
	})
}
