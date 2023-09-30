package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Define uma função de manipulador (handler) para a rota "/"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Bem-vindo à minha API Go!")
    })

    // Inicia o servidor na porta 8080
    fmt.Println("Servidor está rodando em http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
