package main

type mensagemUserSystem struct {
	Role     string `json:"role"`
	Conteudo string `json:"content"` //A qroq so vai receber o JSON se tiver os nomes em ingles
}

type qrorRequest struct {
	//No mensagem ele recebe a struct mensagemUser em um slice
	//você precisa enviar pelo menos duas mensagens para ela entender o contexto, uma do system dizendo oq ela vai ser, e do user
	Mensagem []mensagemUserSystem `json:"messages"`
	Modelo   string               `json:"model"`
}

type qroqResponse struct {
	Choice []struct { //Ele pode te entregar varias respostas, por isso o slice
		Mensagem mensagemUserSystem `json:"message"`
	} `json:"choices"`
}
