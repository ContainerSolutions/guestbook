package api

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ContainerSolutions/guestbook/entries"
	"github.com/prometheus/common/log"
)

type Api struct {
	r entries.Retriever
	t *template.Template
}

func New(r entries.Retriever) (*Api, error) {
	t := template.New("page")
	var err error
	t, err = t.Parse(pagetemplate)
	a := &Api{
		r: r,
		t: t,
	}
	return a, err
}

func (a *Api) Serve() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s\t%s\n", r.Method, r.URL)
		switch r.Method {
		case "GET":
			a.getPage(w, r)
		case "POST":
			a.setEntry(w, r)
		}
	})
	return http.ListenAndServe(":8080", nil)
}

func (a *Api) getPage(w http.ResponseWriter, r *http.Request) {
	a.pageRender(w, []string{""})
}

func (a *Api) setEntry(w http.ResponseWriter, r *http.Request) {
	e := entries.New(
		r.FormValue("name"),
		r.FormValue("message"),
	)
	if len(e.Message) > 0 && len(e.Name) > 0 {
		err := a.r.Write(e)
		if err != nil {
			log.Info(err)
			a.pageRender(w, []string{"Could not save entry"})
			return
		}
		a.pageRender(w, []string{""})
		return
	}
	a.pageRender(w, []string{"Invalid name or message"})
}

func (a *Api) pageRender(w http.ResponseWriter, messages []string) {
	es, err := a.r.GetAll()
	if err != nil {
		messages = append(messages, err.Error())
	}
	p := PageData{
		Entries:  es,
		Messages: messages,
	}
	a.t.Execute(w, p)
}
