package main

import (
	"fmt"
	"log"
	"net/http" // Pacote para criar servidores HTTP
)

func main() {
	// Cria um FileServer que serve arquivos do diretório "static"
	// http.StripPrefix remove o "/static/" do caminho antes de procurar o arquivo no disco.
	// Por exemplo, uma requisição para /static/index.html buscará "index.html" dentro da pasta "static".
	fileServer := http.FileServer(http.Dir("./static"))

	// Configura o handler para servir arquivos estáticos na rota "/static/"
	// Qualquer requisição que comece com "/static/" será tratada por este FileServer.
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Rota para a página inicial
	// Quando alguém acessa a URL raiz ("/") do servidor, ele redireciona para "/static/index.html".
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Se a requisição não for para a raiz, pode ser um erro ou outro caminho não mapeado.
		// Para simplificar, redirecionamos tudo para a página principal.
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		// Redireciona o navegador para o arquivo HTML principal dentro de 'static'
		http.Redirect(w, r, "/static/index.html", http.StatusSeeOther)
	})

	// Inicia o servidor na porta 8080
	port := ":8080"
	fmt.Printf("Servidor Go rodando em http://localhost%s\n", port)
	fmt.Println("Pressione CTRL+C para sair.")

	// log.Fatal irá parar o programa se houver um erro ao iniciar o servidor
	log.Fatal(http.ListenAndServe(port, nil)) // 'nil' usa o Mux padrão do Go
}
