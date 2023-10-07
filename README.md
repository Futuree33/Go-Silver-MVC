# Silver MVC: A Lightweight Golang Framework

## Introduction

Silver MVC is a lightweight and useful framework. I created this as i wanted a similar layout to my php projects (Silver MVC PHP).

Silver is not complete and not ready for production. The plan is to keep developing it into the future and making it a great alternative to other frameworks.

## Why Silver?

Silver keeps things simple. Silver keeps things structured. 

## How to Use Silver

Silver must stay as clean as possible! Therefore, we maintain a simple and organized file structure.

- Create Views (`./Views`)
- Create Controllers (`./Controllers`)

Keeping the file structure minimal ensures a clean and efficient development experience with Silver MVC.


## Creating a Silver App 

```go
func main() {
	silver := Silver.App{}

	silver.AddRoute(Silver.Get, "/", Controllers.Hello)

	silver.Listen(8080)
}
```

## Using Controllers

```go
func Hello(request Utils.Request, response Utils.Response) {
	response.SendView("view.html")
}
```

### Utils.Request
```go
// Getting params from the url which can be declared with {} on AddRoute, example : silver.AddRoute(Silver.Get, "/test/{somevar}", Controllers.Hello)
request.Params["somevar"]
```

### Utils.Response

```go
// Send a standard view
response.SendView("view.html")

// Send a templated view, you can put {{ somevar }} and it will be replaced before its sent to the client
response.SendTemplateView("view.html", Utils.ViewParams{
  "somevar": "hello world!",
})

// Send a string to the client
response.SendString("Hello World!")
```

