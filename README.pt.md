
# `buffalos` Template

[Read in English](README.en.md)

## Sobre

O **`buffalos` Template** é uma base robusta e opinionada para construir aplicações backend e APIs utilizando **Go**. Ele segue as melhores práticas do ecossistema Go, adotando uma arquitetura limpa e modular que facilita a escalabilidade, testabilidade e manutenção do código.

Este template é ideal para desenvolvedores que procuram iniciar projetos rapidamente com uma estrutura já preparada para lidar com **banco de dados, rotas, configurações** e **lógica de domínio** bem definidas.

## Estrutura do Projeto

A arquitetura do `buffalos` é baseada no conceito de Separação de Preocupações, organizando o código em camadas lógicas:

```
├── .gitignore             # Arquivos ignorados pelo Git
├── go.mod / go.sum        # Gerenciamento de módulos e dependências Go
├── LICENSE                # Informações sobre a licença (Ex: MIT)
├── migrations             # Scripts SQL para gerenciamento de esquema de banco de dados
├── README.md              # Este arquivo
└── src
├── .env               # Variáveis de ambiente
├── internal           # Código não importável externamente (lógica da aplicação)
│   ├── configurations # Configurações de inicialização da aplicação (e.g., banco de dados, servidor)
│   ├── controllers    # Camada de entrada (manipuladores HTTP ou lógica de interface)
│   ├── domain         # Estruturas de dados e interfaces
│   ├── misc           # Funções utilitárias e código auxiliar
│   ├── repositories   # Camada de persistência (lógica de acesso ao banco de dados)
│   └── services       # Camada de serviços/lógica de aplicação (coordena domínio e persistência)
└── main.go            # Ponto de entrada da aplicação
```

## Tecnologias Principais

* **Linguagem:** Go (Golang)
* **Gerenciamento de Dependências:** Go Modules (`go.mod`)
* **Banco de Dados:** PostgreSQL
* **Migrações:** `golang-migrate/migrate`

## Primeiros Passos

### 1\. Pré-requisitos

Certifique-se de ter instalado em sua máquina:

* **Go (versão 1.24 ou superior)**
* **[Opcional: Docker/Docker Compose]**
* **[Opcional: O banco de dados principal (PostgreSQL)]**

### 2\. Instalação e Inicialização

Se você estiver usando o `buffalos-cli`:

```bash
buffalos install
```

Caso contrário, clone e inicialize manualmente:

```bash
# 1. Clone o repositório
git clone https://github.com/nes-xiof11/buffalos.git
cd buffalos

# 2. Configure o ambiente
# Edite src/.env com suas credenciais e configurações

# 3. Baixe as dependências
go mod tidy

# 4. Execute as migrações (se aplicável)
# [COMANDO PARA EXECUTAR AS MIGRAÇÕES]
#  migrate -database "postgres://..." -path migrations up

# 5. Inicie a aplicação
go run src/
```

## Migrações de Banco de Dados

A pasta `migrations` contém os arquivos SQL ordenados que definem o esquema do banco de dados.

* `000_init.sql`
* `001_create_table_users.sql`
* `002_create_table_project.sql`

É crucial executar essas migrações antes de iniciar a aplicação pela primeira vez.

## Contribuições

Contribuições são bem-vindas\! Se você encontrou um bug ou tem uma sugestão de melhoria, por favor, abra uma *issue* ou envie um *Pull Request*.
