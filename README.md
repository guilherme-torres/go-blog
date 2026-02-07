# Go Blog

Um blog desenvolvido em **Go**.

## 🎯 Features

- ✅ **Autenticação baseada em sessão**
- ✅ **Controle de acesso por roles**
- ✅ **Editor markdown**
- ✅ **Página de administração**

## 🛠️ Tecnologias Utilizadas

- **Go**
- **Docker & Docker Compose**
- **SQLite3**
- **Redis**
- **Bootstrap 5**
- **Makefile**

## 🚀 Instalação e Configuração

### 1. Clonar o repositório

```bash
git clone https://github.com/guilherme-torres/go-blog.git
cd go-blog
```

### 2. Instalar dependências

```bash
go mod tidy
```

### 3. Configurar banco de dados

As migrations são executadas automaticamente ao iniciar o servidor:

```bash
make migrate
```

### 4. Criar usuário administrador

```bash
./bin/create_admin
```

## 🏃 Execução

### Modo Desenvolvimento

```bash
make dev
```

Inicia o servidor em modo desenvolvimento:
- URL: `http://localhost:8080`
- Migrations são executadas automaticamente

### Modo Produção

```bash
make start
```

Compila o projeto e inicia o servidor:
- Binário compilado em `bin/app`
- Banco de dados em `blog.db`

### Docker Compose

```bash
docker-compose up -d --build
```

## 📁 Estrutura do Projeto

```
go-blog/
├── cmd/
│   ├── create_admin/        # CLI para criar admin
│   └── server/              # Servidor principal
├── internal/
│   ├── errors/              # Definições de erros
│   ├── handlers/            # Handlers HTTP
│   ├── middlewares/         # Middlewares (autenticação, etc)
│   ├── models/              # Estruturas de dados
│   ├── repositories/        # Acesso ao banco de dados
│   ├── services/            # Lógica de negócio
│   └── utils/               # Funções utilitárias
├── migrations/              # Migrations SQL
├── templates/               # Templates HTML
├── assets/                  # CSS e JavaScript
├── docker-compose.yml       # Configuração Docker
├── Dockerfile               # Imagem Docker
├── Makefile                 # Automatização
└── go.mod                   # Dependências Go
```

## 📝 Comandos Make

| Comando | Descrição |
|---------|-----------|
| `make dev` | Inicia servidor em desenvolvimento |
| `make start` | Build + migrations + servidor produção |
| `make build` | Compila o binário |
| `make migrate` | Executa migrations do banco de dados |
| `make tidy` | Limpa dependências do projeto |
| `make clean` | Remove binários compilados |
