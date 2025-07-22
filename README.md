# 🔗 Encurtador de URL em Go

    Este projeto é um encurtador de URLs simples, desenvolvido em Go. Ele gera versões curtas de URLs longas utilizando
    hash e possui uma estrutura modular que facilita a manutenção e evolução do código.

## 📁 Estrutura do Projeto

    encurtador-url/            : Nome do projeto
    ├── go.mod                 : Arquivo de gerenciamento de dependências e módulos do Go.
    ├── cmd/                   : Pasta para a entrada do app
    │ └── main.go              : Ponto de entrada da aplicação. Inicializa o servidor HTTP e registra os endpoints.
    ├── internal/              : Agrupa os pacotes internos da aplicação, separado por responsabilidades.
    │ ├── crypto/              : Responsável pela lógica de geração de hashes
    │ │ └── crypto.go          : Implementa a função de geração de hash que transforma URLs longas em códigos curtos.
    │ ├── handler/             : Define os manipuladores das requisições HTTP
    │ │ └── handler.go         : Define os endpoints da API
    │ └── utils/               : Contém funções auxiliares e genéricas que dão suporte a outras partes da aplicação.
    │   └── utils.go           : Funções auxiliares reutilizáveis

## 🚀 Funcionalidades

    - Geração de URLs curtas a partir de URLs longas.
    - Uso de hash para criação de identificadores únicos.
    - Arquitetura limpa e modular com separação por responsabilidade.

## ▶️ Como executar o projeto

    1. Clone o repositório:
       git clone https://github.com/pedrohfz/encurtador-url.git
       cd encurtador-url
    
    2. Execute a aplicação:
        go run ./cmd/main.go

## 🧪 Exemplos de uso

    1. Encurtar uma URL:
        curl "http://localhost:8080/shorten?url=http://www.google.com"

        - Resposta esperada:
            A URL encurtada deste url original é http://localhost:8080/tKAn3w

    2. Acessar uma URL encurtada:
        curl http://localhost:8080/tKAn3w

        - Respostas esperada:
            <a href="http://www.google.com">Found</a>

## 🧰 Tecnologias utilizadas

    - Go 1.20+.
    - Módulos nativos (net/http, crypto/sha256, etc.).

## 📌 Melhorias futuras

    - Redirecionamento automático da URL curta para a original.
    - Armazenamento em banco de dados.
    - Interface web para gerar e visualizar URLs encurtadas.
    - Testes automatizados.

## 📎 Autor

    Desenvolvido por Pedro Henrique Leite como uma forma de estudar Go.

### Feito com ❤️ em Go.