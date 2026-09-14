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
