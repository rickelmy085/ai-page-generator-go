# ai-page-generator-go
Projeto desenvolvido para criar páginas web em segundos usando Golang no Back-end e a API da Groq 

# Feedback para Desenvolvedores Go
sou estudante de Sistemas, e esse projeto é muito importante para o meu aprendizado com Golang. Se você for um dev mais experiente e notar algo que possa ser melhorado, sinta-se à vontade para:
*   Abrir uma Issue explicando a melhoria.
*   Enviar um Pull Request com sugestões.
*   Comentar sobre o que melhorar no meu código.

# Gerador de Sites com IA(Go + Groq API)
Este projeto é uma aplicação que utiliza **Go (Golang)** no Back-end para mandar as requisições entre o Front-end e a API de inteligência artificial da **Groq**. O objetivo é permitir que usuários gerem estruturas completas de sites apenas descrevendo sua ideia.

## 🛠️ Tecnologias Utilizadas
*   **Back-end:** Go (Golang)
*   **Front-end:** HTML5, CSS3, JavaScript
*   **IA:** Llama 3.3 pela Groq API
*   **Gestão de Dependências:** Go Modules (`go.mod`, `go.sum`)

## 📋 Funcionalidades
*   Lugar para entrada de prompts do usuário.
*   Servidor Go que processa requisições e gerencia segurança.
*   Injeção de `System Prompt` para garantir que a IA retorne apenas código funcional.
*   Exibição do site gerado em tempo real `iframe`.

## Como Executar o Projeto

### Pré-requisitos
*   Uma chave de API da [Groq](https://console.groq.com/).

### Configuração
1.  Clone o repositório.
2.  Crie um arquivo `.env` na raiz da pasta `Back-end` e adicione sua chave:
    ```env
    GROQ_API_KEY=sua_chave_aqui
    ```
3.  Instale as dependências:
    ```bash
    go mod tidy
    ```

### Execução
1.  Inicie o servidor Go:
    ```bash
    go run serve.go structs.go
    ```
2.  Abra o arquivo `index.html`
