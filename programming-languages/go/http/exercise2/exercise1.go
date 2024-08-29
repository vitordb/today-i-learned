package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
}

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "CEP não informado", http.StatusBadRequest)
		return
	}

	c, err := BuscaCep(cep)
	if err != nil {
		http.Error(w, "Erro ao buscar o CEP", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(c)
}

func BuscaCep(cep string) (*ViaCEP, error) {

	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var c ViaCEP
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil

}
