package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Pessoa struct {
	id       string    `json:"id,omitempty`
	nome     string    `json:"nome,omitempty`
	telefone string    `json:"telefone,omitempty"`
	email    string    `json:"email,omitempty"`
	endereco *Endereco `json:"endereco,omitempty"`
}
type Endereco struct {
	logradouro string `json:"logradouro,omitempty"`
	numero     string `json:"numero,omitempty"`
	bairro     string `json:"bairro,omitempty"`
	cidade     string `json:"cidade,omitempty"`
	estado     string `json:"estado,omitempty"`
}

var cadastro []Pessoa

func API(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Web server listening at: http://localhost:8000\n"))
}

func GetPessoaById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range cadastro {
		if item.id == params["id"] {
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
	pessoa.id = params["id"]
	cadastro = append(cadastro, pessoa)
	json.NewEncoder(w).Encode(cadastro)
}
func DeletePessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range cadastro {
		if item.id == params["id"] {
			cadastro = append(cadastro[:index], cadastro[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(cadastro)
	}
}

func main() {
	router := mux.NewRouter()
	cadastro = append(cadastro, Pessoa{id: "1", nome: "João Ninguém", telefone: "51999999999", email: "joao.ninguem@golang.io", endereco: &Endereco{logradouro: "Assis Brasil", numero: "8450", bairro: "Sarandi", cidade: "Porto Algre", estado: "RS"}})
	router.HandleFunc("/", API)
	router.HandleFunc("/people", GetPessoa).Methods("GET")
	router.HandleFunc("/people/{id}", GetPessoaById).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePessoa).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePessoa).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
