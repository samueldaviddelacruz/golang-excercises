package choose_your_own_adventure

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))

}

var defaultHandlerTemplate = `
<html>
    <head>
        <meta charset="utf-8">
        <title>Choose your own adventure</title>
    </head>
<body>
<section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{else}}
        <h3>The End</h3>
      {{end}}
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>

</body>

</html>`

type HandlerOption func(h *handler)

func WithTemplate(template *template.Template) HandlerOption {
	return func(h *handler) {
		h.template = template
	}
}

func WithPathFunc(fn func(request *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl, defaultPathFn}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	story    Story
	template *template.Template
	pathFn   func(request *http.Request) string
}

func defaultPathFn(request *http.Request) string {
	path := strings.TrimSpace(request.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"

	}

	path = path[1:]
	return path
}
func (h handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := h.pathFn(request)
	if chapter, ok := h.story[path]; ok {
		err := h.template.Execute(writer, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(writer, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(writer, "Chapter not found...", http.StatusNotFound)
}

func JsonStory(reader io.Reader) (Story, error) {
	decoder := json.NewDecoder(reader)
	var story Story

	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
