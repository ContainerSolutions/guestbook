package main

import (
	"fmt"
	"net/http"
)

func GETguestBook(res http.ResponseWriter, req *http.Request) {
	// TODO: Load messages from database and spit them out into the HTML
	form := `<form method="POST" action="/"><legend>Guestbook</legend><fieldset><label name="message">Leave your message:</label><textarea name="message" autofocus="yes" cols="100" rows="6"></textarea><button type="submit">Send</button></fieldset></form>`
	fmt.Fprintf(res, layout(form))
}

func POSTguestBook(res http.ResponseWriter, req *http.Request) {
	message := req.FormValue("message")
	if len(message) > 0 {
		// TODO: Save the message in database
	}
	http.Redirect(res, req, "/", 302)
}

func layout(content string) string {
	return fmt.Sprintf("<html><head><style>body { font-family: Century Gothic,CenturyGothic,AppleGothic,sans-serif; margin: 10em; } label, textarea { display: block; } textarea { margin: 1em 0; }</style></head><body>%s</body></html>", content)
}
