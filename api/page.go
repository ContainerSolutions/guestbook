package api

import "github.com/ContainerSolutions/guestbook/entries"

type PageData struct {
	Entries  []entries.Entry
	Messages []string
}

var pagetemplate = `
<html>

<head>
    <style>
        body {
            font-family: Century Gothic, CenturyGothic, AppleGothic, sans-serif;
            margin: 10em;
        }

        label,
        textarea {
            display: block;
        }

        textarea {
            margin: 1em 0;
        }
    </style>
</head>

<body>
    {{range .Messages }}
    <div id="message"> {{ . }} </div>
    {{ end }}
    <form method="POST" action="/">
        <legend>Guestbook</legend>
        <fieldset>
		<label name="name">Your name:</label>
		<input type="text" name="name">
		<label name="message">Leave your message:</label>
		<textarea name="message" autofocus="yes" cols="100" rows="6"></textarea>
		<button type="submit">Send</button></fieldset>
    </form>
    <div id="submissions">
	{{range .Entries }}
           <div>{{.Name}} - {{ .Date }} - {{ .Message }}</div>  
	{{ end }}
    </div>
</body>

</html>
`
