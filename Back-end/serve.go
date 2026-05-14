package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	http.HandleFunc("/gerar", HandleAIQroq)
	http.ListenAndServe(":8080", nil)

}

func HandleAIQroq(w http.ResponseWriter, r *http.Request) {
	// Sem isso, o navegador bloqueia a requisição vinda do seu HTML
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 2. Define quais métodos HTTP são permitidos
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	// 3. Permite que o cabeçalho 'Content-Type' seja enviado pelo JavaScript
	// Como enviamos JSON no JS, o navegador precisa dessa autorização
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		w.Write([]byte("Tudo ok"))
		return
	}

	godotenv.Load()
	API_qror := os.Getenv("GROQ_API_KEY")

	//O qrorRequest vai receber a requisicao do usuários
	var requisicao qrorRequest
	err := json.NewDecoder(r.Body).Decode(&requisicao)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Deu erro ao ler"))
		return
	}

	//Definimos o modelo dele
	requisicao.Modelo = "llama-3.3-70b-versatile"

	mensagemSystem := mensagemUserSystem{
		Role: "system",
		Conteudo: "Você é um programdor que recebera uma ideia de site e tera que fazer ele usando HTML,CSS e JS. Mas tem que ser feito em apenas UM ARQUIVO" +
			"Coloque o CSS dentro do <style> e o js dentro do <script>" +
			"IMPORTANTE: Não use nenhuma imagem, so diga que ali era para ter uma imagem usando o comentario, use apenas retângulos cinzas (divs com cor de fundo) como placeholders" +
			"e escreva um texto dentro deles dizendo 'Espaço para Foto'",
	}

	//Oq acontece aqui é que o requisicao.Mensagem(O mensagem é da struct qrorRequest)
	//Ele vai colocar no Mensagem a mensagem do system
	requisicao.Mensagem = append([]mensagemUserSystem{mensagemSystem}, requisicao.Mensagem...)

	//O Marshal transforma em JSON o requisicao
	corpoJSON, err := json.Marshal(requisicao)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Deu erro ao ler"))
		return
	}

	//Esse codigo tem no proprio site do QROQ, na parte de CURL
	URL := "https://api.groq.com/openai/v1/chat/completions"
	reqSite, err := http.NewRequest("POST", URL, bytes.NewBuffer(corpoJSON))
	reqSite.Header.Set("Content-Type", "application/json")
	reqSite.Header.Set("Authorization", fmt.Sprintf("Bearer %s", API_qror))

	//Aqui ele vai enviar o que o usuários escreveu para o servido da Qroq
	client := &http.Client{}
	resp, err := client.Do(reqSite)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro com o servido"))
		return
	}
	defer resp.Body.Close()

	var respostaQROR qroqResponse
	err = json.NewDecoder(resp.Body).Decode(&respostaQROR)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro do servido"))
		return
	}

	resultadoParaJS := respostaQROR.Choice[0].Mensagem.Conteudo
	w.Write([]byte(resultadoParaJS))
}
