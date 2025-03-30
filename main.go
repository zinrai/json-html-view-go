package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getTodos(endpoint string) ([]Todo, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to access the endpoint: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var todos []Todo
	if err := json.Unmarshal(body, &todos); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return todos, nil
}

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Todos一覧</title>
    <link rel="icon" href="data:,">
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 20px;
        }
        h1 {
            color: #333;
        }
        table {
            border-collapse: collapse;
            width: 100%;
            margin-top: 20px;
        }
        th, td {
            border: 1px solid #ddd;
            padding: 8px;
            text-align: left;
        }
        th {
            background-color: #f2f2f2;
            font-weight: bold;
        }
        tr:nth-child(even) {
            background-color: #f9f9f9;
        }
        .completed {
            color: green;
            font-weight: bold;
        }
        .not-completed {
            color: red;
        }
    </style>
</head>
<body>
    <h1>Todos一覧</h1>
    <table>
    <tr>
        <th>ID</th>
        <th>ユーザーID</th>
        <th>タイトル</th>
        <th>状態</th>
    </tr>
    {{range .}}
    <tr>
        <td>{{.ID}}</td>
        <td>{{.UserID}}</td>
        <td>{{.Title}}</td>
        <td class="{{if .Completed}}completed{{else}}not-completed{{end}}">
            {{if .Completed}}完了{{else}}未完了{{end}}
        </td>
    </tr>
    {{end}}
    </table>
</body>
</html>
`

func main() {
	endpoint := "https://jsonplaceholder.typicode.com/todos"
	if len(os.Args) > 1 {
		endpoint = os.Args[1]
	}

	tmpl, err := template.New("todos").Parse(htmlTemplate)
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		todos, err := getTodos(endpoint)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.Execute(w, todos); err != nil {
			log.Printf("Template execution error: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	port := "8000"
	fmt.Printf("Serving HTTP on port %s...\n", port)
	fmt.Printf("API Endpoint: %s\n", endpoint)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
