# Writing Web Applications In Go
build go binary
```go build -o bin/site```

run app
```./bin/site```

## Next Steps
* [X] Try to incoperate the makeHandler with the above.
* [X] Try to move most of this out of main. Make it a single line. Like the example.
* [] Read more about patterns around HanlderFunc and HandleFunc
* [] Try to make the handle func have a URL path related to the file.
* [] Try to make the files drive page path - // 	http.HandleFunc("/home/", makeHandler(viewHandler)) 
* [] Try  to use templating to generate headers, footers and components ( Nav, "ads") {{header}}
    * [] This might be called nested templates - Verify
	* [] https://www.calhoun.io/intro-to-templates-p2-actions/
* [] What is routing? Does this project need it?
* [] Much later ... autoreloading? Maybe?
* [] Think more about how to incoperate more elemets - use an API to generate content in a table

###### Build a Homepage using "custom" JS, HTML, CSS
include flexbox in CSS ```https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Flexible_Box_Layout/Basic_Concepts_of_Flexbox```


## Other Things To Read
###### Resource: https://golang.org/pkg/net/http/
###### Resource: https://golang.org/pkg/html/template/
###### Resource: https://golang.org/doc/faq#Pointers
###### Resource: https://www.calhoun.io/
###### Examples: https://course.gowebexamples.com/?utm_source=banner