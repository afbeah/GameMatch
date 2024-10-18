# GAMESINSERTION

A GamesInsertion API é uma API RESTful desenvolvida em Go usando o framework Echo. Ela permite gerenciar uma lista de jogos assistidos, possibilitando operações como adicionar, atualizar, deletar e listar as partidas. A aplicação é ideal para servir como um backend para aplicativos de gerenciamento de competições e modalidades acompanhadss pelo usuário. Criado na intenção e listar/gerenciar os jogos que vi. 

# Utilização do GO

A principal intenção em utilizar GO é aplicar tudo que venho aprendendo e observando sobre a linguagem.

# Pré-requisitos

Go 1.20+
Cliente HTTP para testes: Insomnia ou Postman

# Inicialização da API

-> Instalação

1. Clone o repositório:

-- git clone git@github.com:afbeah/GameMatch.git --
cd GameMatch

2. Instale as dependências:

-- go mod tidy --

3. Inicie o servidor:

-- go run main.go --
O servidor estará disponível em http://localhost:8088.

# Endpoints

1. Health Check

-- GET /health
Retorna o status de saúde a API

2. Matchs

-- GET /matchs/:id
Retorna uma partida específica pelo ID
// GET http://localhost:8088/matchs/1

-- POST /matchs/
Retorna a inclusão de um jogo.

-- PUT /matchs/
Atualiza uma partida existente.

-- DELETE /matchs/:id
Deleta uma partida específica existente.

# Estrutura do Projeto 

GameMatch/
│
├── src/
│   ├── handler/
│   │   └── GameMatch.go
│   ├── model/
│   │   └── Match.go
│   ├── service/
│      └── GameMatch_service.go
│   
│
├── main.go
├── go.mod
├── go.sum
└── README.md