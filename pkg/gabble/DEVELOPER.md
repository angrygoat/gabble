Strategies and tools

### logging

see https://blog.boot.dev/golang/golang-logging-best-practices/
adopting zerolog: https://betterstack.com/community/guides/logging/zerolog/
formatting: https://gobyexample.com/string-formatting

### error handling

wrapping errors/custom errors: https://earthly.dev/blog/golang-errors/

```go
doc, err := html.Parse(resp.Body)
resp.Body.Close()
if err != nil {
    return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
}
```

### properties and config

adopting viper: https://github.com/spf13/viper
viper: https://pkg.go.dev/github.com/spf13/viper#section-readme  

## testing utils

random file trees: https://github.com/jbenet/go-random-files
more randomness: github.com/brianvoe/gofakeit/v6
setup and teardown: https://medium.com/nerd-for-tech/setup-and-teardown-unit-test-in-go-bd6fa1b785cd


### structuring code

https://www.toptal.com/go/golang-oop-tutorial#:~:text=Is%20Go%20object%2Doriented%3F,a%20clear%20and%20understandable%20way.

https://go.dev/blog/using-go-modules

context versus threadlocal: https://go.dev/blog/context

docs: https://go.dev/blog/godoc

style: https://google.github.io/styleguide/go/best-practices.html