# HTML Parse Example

This Go program demonstrates how to use HTML templates to render web pages. The example includes two handlers: one for viewing a page and another for editing a page. Below is a detailed explanation of the code.

## Code Explanation
### Global Variables
```go
var (
    editHTMLFile = "edit.html"
    viewHTMLFile = "view.html"
    templates    = template.Must(template.ParseFiles(editHTMLFile, viewHTMLFile))
)
```
* editHTMLFile and viewHTMLFile store the filenames of the HTML templates.
* templates is a parsed template object that includes both HTML files.

### Page Struct
```go
type Page struct {
    Title string
    Body  string
}
```
The Page struct represents the data structure passed to the HTML templates. It contains:

* `Title`: The title of the page.
* `Body`: The body content of the page.

### Handlers
#### ViewHandler
```go
func viewHandler() {
    content := &Page{
        Title: "View Page",
        Body:  "This is view page",
    }
    renderTemplate(viewHTMLFile, content)
}
```
The viewHandler function creates a Page object with a title and body for the view page and calls renderTemplate to render the view.html template.

#### EditHandler
```go
func editHandler() {
    content := &Page{
        Title: "Edit Page",
        Body:  "This is edit page",
    }
    renderTemplate(editHTMLFile, content)
}
```
The editHandler function creates a Page object with a title and body for the edit page and calls renderTemplate to render the edit.html template.

### Render Template Function
```go
func renderTemplate(htmlFile string, content *Page) {
    if err := templates.ExecuteTemplate(os.Stdout, htmlFile, content); err != nil {
        log.Fatalf("Error executing template: %v", err)
    }
}
```
The renderTemplate function takes an HTML file name and a Page object as arguments. It executes the specified template and writes the output to os.Stdout. If an error occurs, it logs the error and exits the program.

### Main Function
```go
func main() {
    viewHandler()

    fmt.Println()
    fmt.Println()
    fmt.Println(strings.Repeat("-", 100))
    fmt.Println()

    editHandler()
}
```
The main function calls viewHandler to render the view page, prints a separator line, and then calls editHandler to render the edit page.

## Running the Program
```bash
go run main.go
```
This will output the rendered HTML content of both the view and edit pages to the console.

## Output
```
<h1>View Page</h1>

<p>[<a href="/edit/View%20Page">edit</a>]</p>

<div>This is view page</div>

----------------------------------------------------------------------------------------------------

<h1>Editing Edit Page</h1>

<form action="/save/Edit%20Page" method="POST">
<div><textarea name="body" rows="20" cols="80">This is edit page</textarea></div>
<div><input type="submit" value="Save"></div>
</form>
```

## Conclusion
This example demonstrates the basics of using HTML templates in Go to render web pages

## References
https://go.dev/doc/articles/wiki/