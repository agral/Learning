To serve content as a web page use `http.ListenAndServe()`.
It takes port number to serve on, and options that can be `nil`.
So I've written `http.ListenAndServe(":8080", nil)` and it worked out of the box.

Multiple subpages can be provided via `http.HandleFunc()`.
I've provided:
```
http.HandleFunc("/", Home)
http.HandleFunc("/about", About)
```

Each of these `Home`, `About` functions have the following signature:
```
func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Content")
}
```
So it's pretty easy, just write any content to the writer, and it appears on the page.


## Working with templates.
To use templates, just import "html/template". Then there's `template.ParseFiles()`,
and each of such parsed tempates can be `Execute()`d.

Templates can use other templates. Variable blocks can be named as I wish: `{{block "NAME" .}} {{end}}`
Then the blocks get actually filled in: `{{define "NAME"}} Actual content {{end}}`. Easy.

### Template caching
Parsing and disk access might be expensive. Templates, once parsed, can be cached in memory.
This usually significantly improves the performance.
