package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErroAPI representa a resposta de erro da API
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON retorna uma respotas em formato JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}

}

// TratarStatusCodeDeErro trata as requisições com status code 400 ou superior
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
