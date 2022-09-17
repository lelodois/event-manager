# App para criar tickets de eventos

> O objetivo desta app é explorarmos as facilidades do golang
> para criações de apps de back-end :)
> Build fast, reliable, and efficient software at scale

---
#### Ferramentas
- Go - Golang
- Mysql
- Backend
- Gorm / Gin
- Docker compose
- Postman

---
### Funcionalidades

- Cria / Lista - Usuários
- Cria / Lista - Eventos
- Cria / Lista - Tickets

---
### Modelagem
![Modelagem](doc/mer.png)

### Regras de negócio

- Não se pode criar um evento com data retroativa
- Um usuário pode comprar apenas um ticket por evento
- O usuário deve ter saldo maior que o valor do ticket do evento
- Não pode exceder tickets a mais que a capacidade do evento

---
### Instruções para testar usando o postman

- 1 - Inicie o docker
- 2 - Sugiro importar o projeto com o Goland
- 3 - Na raiz do projeto execute: `docker-compose up -d`
- 4 - Vá até o arquivo `main/main.go` clique no `botão direito` e depois em `debug`
- 5 - Importe a coleção do postam na pasta `docs`
- 6 - Crie na sequência um Usuário, um evento
- 7 - No list Recupere os ids (usuário e evento) e crie um Ticket
- 
---

### Instruções para usar o teste integrado

- 1 - Inicie o docker
- 2 - Sugiro importar o projeto com o Goland
- 3 - Na raiz do projeto execute: `docker-compose up -d`
- 4 - Vá até o arquivo `main/main_test.go` clique no `botão direito` e depois em `debug`
- 5 - O teste integrado criará um user, event e um ticket, acompanhe no log
```
=== RUN   TestIntegrationTest
2022/09/16 23:19:20 Starting app
2022/09/16 23:19:20 Database connected
[GIN-debug] POST   /user                     --> eventManager/user.Create (3 handlers)
[GIN-debug] GET    /user                     --> eventManager/user.List (3 handlers)
[GIN-debug] POST   /event                    --> eventManager/event.Create (3 handlers)
[GIN-debug] GET    /event                    --> eventManager/event.List (3 handlers)
[GIN-debug] POST   /ticket                   --> eventManager/ticket.Create (3 handlers)
[GIN-debug] GET    /ticket                   --> eventManager/ticket.List (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2022/09/16 - 23:19:22 | 201 |    6.701341ms |             ::1 | POST     "/user"
[GIN] 2022/09/16 - 23:19:22 | 200 |    3.003981ms |             ::1 | GET      "/user"
new user: [JarrettHudson] with id: [39]
[GIN] 2022/09/16 - 23:19:22 | 201 |    5.409329ms |             ::1 | POST     "/event"
[GIN] 2022/09/16 - 23:19:22 | 200 |    1.807638ms |             ::1 | GET      "/event"
new event: [ZPdFrHb@xyxJPag.com] with id: [13]
2022/09/16 23:19:22 decrease available eventId: [13], decrease amount: [10] of user balance: [39], because create ticket: [12]
[GIN] 2022/09/16 - 23:19:23 | 201 |   16.335612ms |             ::1 | POST     "/ticket"
[GIN] 2022/09/16 - 23:19:23 | 200 |   52.374754ms |             ::1 | GET      "/ticket"
new ticket: [12]
--- PASS: TestIntegrationTest (2.10s)
```


### Conceitos

| Nome | Detalhe |
| ------ | ------ |
| Anti-corruption layer | Nas apis rest, é trafegado uma estrutura diferente dos modelos do domínio |
| ORM | Usa framework para traduzir modelos em instruções sql no mysql|
| Migrate | Cria a base de dados e as tabelas no início da app |
| Transação | Multiplas instruções sql garantidas em apenas uma transação (reservation) |

---
### Arquivos 

![Modelagem](doc/files.png)


[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

