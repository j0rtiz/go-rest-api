package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var porta = ":8000"

type Pessoa struct {
	ID       string    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Telefone string    `json:"telefone,omitempty"`
	Email    string    `json:"email,omitempty"`
	Endereco *Endereco `json:"endereco,omitempty"`
}

type Endereco struct {
	Logradouro string `json:"logradouro,omitempty"`
	Numero     string `json:"numero,omitempty"`
	Bairro     string `json:"bairro,omitempty"`
	Cidade     string `json:"cidade,omitempty"`
	Estado     string `json:"estado,omitempty"`
}

var cadastro []Pessoa

func API(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Métodos:\n\n"))
	w.Write([]byte("GET    -> http://localhost:{porta}/cadastros     -> Lista todos os cadastros\n"))
	w.Write([]byte("GET    -> http://localhost:{porta}/cadastro/{id} -> Lista cadastro com o id que foi passado\n"))
	w.Write([]byte("POST   -> http://localhost:{porta}/cadastro/{id} -> Cria cadastro com o id que foi passado\n"))
	w.Write([]byte("DELETE -> http://localhost:{porta}/cadastro/{id} -> Exclui cadastro com o id que foi passado\n"))
}

func GetPessoaById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range cadastro {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Pessoa{})
}

func GetPessoa(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(cadastro)
}

func CreatePessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var pessoa Pessoa
	_ = json.NewDecoder(r.Body).Decode(&pessoa)
	pessoa.ID = params["id"]
	cadastro = append(cadastro, pessoa)
	json.NewEncoder(w).Encode(cadastro)
}

func DeletePessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range cadastro {
		if item.ID == params["id"] {
			cadastro = append(cadastro[:index], cadastro[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(cadastro)
	}
}

func main() {
	router := mux.NewRouter()
	cadastro = append(cadastro, Pessoa{ID: "1", Nome: "João Ninguém", Telefone: "51999999999", Email: "joao.ninguem@golang.io", Endereco: &Endereco{Logradouro: "Assis Brasil", Numero: "8450", Bairro: "Sarandi", Cidade: "Porto Algre", Estado: "RS"}})
	router.HandleFunc("/", API)
	router.HandleFunc("/cadastros", GetPessoa).Methods("GET")
	router.HandleFunc("/cadastro/{id}", GetPessoaById).Methods("GET")
	router.HandleFunc("/cadastro/{id}", CreatePessoa).Methods("POST")
	router.HandleFunc("/cadastro/{id}", DeletePessoa).Methods("DELETE")
	fmt.Printf("Web server listening at: http://localhost%s", porta)
	log.Fatal(http.ListenAndServe(porta, router))
}
