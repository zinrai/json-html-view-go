# json-html-view-go

A sample Go application that fetches JSON data from an HTTP API endpoint specified as a command-line argument and renders it as HTML in a web browser.

Use [{JSON} Placeholder](https://jsonplaceholder.typicode.com/) as HTTP API.

## Installation

```bash
$ go build
```

## Usage

Run with default endpoint:

```bash
$ ./json-html-view-go
```

Specify a custom API endpoint:

```bash
$ ./json-html-view-go https://jsonplaceholder.typicode.com/todos
```

Open your browser and navigate to http://localhost:8000 to view the rendered data.

## License

This project is licensed under the MIT License - see the [LICENSE](https://opensource.org/license/mit) for details.
