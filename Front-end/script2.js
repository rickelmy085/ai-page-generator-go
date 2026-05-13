// O endereço da API que criei em GO
const endereco = "http://localhost:8080/gerar";

async function gerarCodigo() {
    //  Pegar os elementos do HTML para usar depois
    let espacoCodigo = document.querySelector(".bloco-codigo");
    let espacoSite = document.querySelector(".bloco-site");
    let textarea = document.querySelector(".texto-pagina");

    const promptUser = textarea.value; //Aqui pega o que o usuario escreveu na caixa

    espacoCodigo.textContent = "Gerando código... aguarde.";

    try {
        let resposta = await fetch(endereco, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                messages: [
                    role: "user",
                    content: promptUser,
                ]
            })
        });

        let resultado = await resposta.text();

        espacoCodigo.textContent = resultado; //Vai colocar o codigo feito no site
        espacoSite.srcdoc = resultado; //O srcdoc vai pegar o codigo feito e renderiza no Iframe

    } catch (err) {
        console.err("Erro na comunicação com o servido:", err);
        espacoCodigo.textContent = "Erro ao conectar com o servido GO.";
    }
}