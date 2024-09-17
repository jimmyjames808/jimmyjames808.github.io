package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	employees = Employees{}
)

func main() {
	employees.loadFromJson()

	http.HandleFunc("/employee/", allWebPage)
	http.HandleFunc("/images/", handleImages)

	log.Print("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleImages(w http.ResponseWriter, r *http.Request) {
	pathBase := filepath.Base(filepath.Clean(r.URL.Path))

	if strings.ToLower(pathBase) == "alt" {
		i, err := os.ReadFile("Website/Static/alt.png")
		if err != nil {
			log.Print(err.Error())
			http.Error(w, http.StatusText(500), 500)
		}
		_, err1 := w.Write(i)
		if err1 != nil {
			log.Print(err1.Error())
			http.Error(w, http.StatusText(500), 500)
		}
	}
}

func allWebPage(w http.ResponseWriter, r *http.Request) {
	pathBase := filepath.Base(filepath.Clean(r.URL.Path))
	params, _ := url.ParseQuery(r.URL.RawQuery)

	if strings.ToLower(pathBase) == "all" {
		employees.WebUpdateAllEndpoint()
		s, _ := os.ReadFile("Website/Static/All.html")
		_, err := w.Write(s)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, http.StatusText(500), 500)
		}
		return
	}

	if strings.ToLower(pathBase) == "createuserbackend" {
			pay, _ := strconv.ParseFloat(params.Get("pay"), 64)
			name := params.Get("name")
			name = url.QueryEscape(name)
			employees.NewEmployee(strings.ReplaceAll(name, "+", "%20"), pay)
		http.Redirect(w, r, "all", http.StatusSeeOther)
		return
	}

	if strings.Contains(strings.ToLower(filepath.Clean(r.URL.Path)), "checkinu") {
		name := url.QueryEscape(pathBase)
		u, err := employees.FindEmployeeByName(strings.ReplaceAll(name, "+", "%20"))
		if err != nil {
			http.Redirect(w, r, "../all", http.StatusSeeOther)
			return
		}
		u.StartShift()
		http.Redirect(w, r, "../all", http.StatusSeeOther)
	}
	if strings.Contains(strings.ToLower(filepath.Clean(r.URL.Path)), "checkoutu") {
		name := url.QueryEscape(pathBase)
		u, err := employees.FindEmployeeByName(strings.ReplaceAll(name, "+", "%20"))
		if err != nil {
			http.Redirect(w, r, "../all", http.StatusSeeOther)
			return
		}
		u.EndShift()
		http.Redirect(w, r, "../all", http.StatusSeeOther)
	}
	if strings.Contains(strings.ToLower(filepath.Clean(r.URL.Path)), "termu") {
		name := url.QueryEscape(pathBase)
		u, err := employees.FindEmployeeByName(strings.ReplaceAll(name, "+", "%20"))
		if err != nil {
			http.Redirect(w, r, "../all", http.StatusSeeOther)
			return
		}
			employees.TerminateEmployee(*u)
		http.Redirect(w, r, "../all", http.StatusSeeOther)
	}
	return
}
