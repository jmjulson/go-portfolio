# Writing Web Applications In Go
build go binary
```go build -o bin/site```

run app
```./bin/site```

## Next Steps
* [X] Try to incoperate the makeHandler with the above.
* [X] Try to move most of this out of main. Make it a single line. Like the example.
* [] Read more about patterns around HanlderFunc and HandleFunc
* [] Try to make the files drive page path in HandleFunc()
* [] Try  to use templating to generate headers, footers and components ( Nav, "ads") {{header}}
    * [] This might be called nested templates - Verify
	* [] https://www.calhoun.io/intro-to-templates-p2-actions/
* [] Relearn marshelling and unmarshelling
* [] Fill templates using data/* files 🧁
* [] Try writing some content around what you have learned so far 🤔
* [] Start building the page layouts in CSS. 🖼
* [] Add some JS - Nav Accordian, Pagnation in dummy tables 🧬
* [] Add proper error handling 🤷‍♂️
* [] Figure out if there is an out of the box way to run tests 🧪
* [] Figure out where to store the files and hosting the site (Amazon Ecosystem) 💸
* [] Deploy site manually 🧱
* [] Try writing a CD pipeline 🚀
* [] Learn how to hotload content outside of the bundle 🌶
* [] Create s3 bucket. Store data/* files 🗄
* [] Try to refactor code and use hotloading from s3 🗂
* [] Refactor code to handle new files - blog'ish 📝
* [] Refactor layout so it displays "summary" from files - will require additional field in data/* files ✨
    * [] new structure in files for this type of metadata
* [] Update homepage so the files loaded become a "feed" 🍼
* [] Convert JS to Typescript 😅

## Discovery
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