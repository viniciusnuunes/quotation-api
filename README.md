# Quotation API

API criada para estudos do Golang

### Tecnologias utilizadas

- Go
- Gin
- Gorm
- Postgres
- Docker
- Docker Compose
- Redis

### Rodando o Projeto

É necessário ter instalado as seguintes ferramentas

| Tool           | Version  |
| -------------- | -------- |
| Docker         | 20.10.21 |
| Docker Compose | 2.13.0   |

Para instalar as depencências, fazer build do docker e iniciar o projeto, basta o único comando abaixo
```sh
docker compose up --build
```

Caso já esteja com o build completo, utilizar o comando
```sh
docker compose up
```

O projeto estará rodando em

```sh
localhost:8080
```

### Materiais de Estudo para a criação do Projeto

#### Padrões de Projeto
[Boas práticas de Nomenclatura](https://go.dev/talks/2014/names.slide#1) | Andrew Gerrand - Google Inc - October 2014

[Layout padrão de projetos em Go](https://github.com/golang-standards/project-layout/blob/master/README_ptBR.md) | Golang-standards - Github

Criado por [Vinicius Nunes](https://github.com/viniciusnuunes) e [Igor Silva](https://github.com/igorverse)



