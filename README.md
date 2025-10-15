# Middleware de Rate Limiter em Go

[![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![Licença](https://img.shields.io/badge/Licença-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/knipers/rate-limiter)](https://goreportcard.com/report/github.com/knipers/rate-limiter)

Um middleware de limitação de taxa de requisições em Go, otimizado para desempenho e utilizando Redis como backend. Projetado para controlar o acesso a APIs e serviços web de forma eficiente e personalizável.

## ✨ Funcionalidades

- **Limitação Flexível**: Configure limites personalizados por token de API
- **Dual-Modo**: Limitação por IP (padrão) ou por token (prioritário)
- **Backend em Redis**: Suporte a ambientes distribuídos com Redis
- **Fácil Integração**: Compatível com qualquer roteador que siga o padrão `http.Handler`
- **Altamente Configurável**: Personalize limites e durações de bloqueio

## 🚀 Começando

### Pré-requisitos
-
- Go 1.24 ou superior
- Servidor Redis (para ambiente de produção)

## 🔧 Opções de Configuração

| Variável de Ambiente | Padrão | Descrição |
|----------------------|--------|------------|
| `RATE_LIMIT_DEFAULT` | 10 | Limite padrão de requisições por IP |
| `BLOCK_DURATION` | 120 | Duração do bloqueio em segundos |
| `REDIS_ADDR` | - | Endereço do servidor Redis (formato: host:porta) |
| `REDIS_PASSWORD` | - | Senha de autenticação do Redis |
| `REDIS_DB` | 0 | Número do banco de dados Redis |
| `TOKENS` | - | Lista de tokens e seus limites (formato: token:limite) |

## 📚 Como Funciona

1. **Processamento de Requisições**:
  - Verifica o cabeçalho `API_KEY`
  - Se existir, aplica o limite específico do token
  - Caso contrário, usa o limite baseado em IP

2. **Controle de Acesso**:
  - Cada requisição incrementa um contador no Redis
  - Ao exceder o limite, retorna status 429 (Muitas Requisições)
  - Clientes bloqueados permanecem assim pelo tempo configurado

3. **Sistema de Prioridades**:
  - Tokens têm precedência sobre limites de IP
  - Um token válido pode contornar restrições de IP
  - Exemplo: Se um IP estiver bloqueado, um token válido ainda permite o acesso

---

## Exemplo de `.env`

```env
# Limite padrão por IP
RATE_LIMIT_DEFAULT=10

# Tempo de bloqueio (em segundos)
BLOCK_DURATION=30

# Conexão Redis
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# Lista de tokens (token:limite), separados por ";"
TOKENS=token123:100;vipUser:500
```


## Executando o Projeto

1. Suba o Redis com Docker
```docker-compose up -d```
2. Instale as dependências
```go mod tidy```
3. Execute o servidor
```go run cmd/main.go```


## Fazendo Requisições

#### Com token`
```bash
curl -H "API_KEY: token123" http://localhost:8080
```

#### Sem Token (usa o IP)
```bash
curl http://localhost:8080
```

#### Adicionando tokens

- Você pode adicionar tokens manualmente no **.env**;
  - Cada token pode ter seu valor para limite, seguindo o padrão __${token}:${limite}__, seguindo o exemplo:  ```TOKENS=token1:100;token2:50```

- Caso você tenha o **CLI** do `Redis` instalado e conectado ao servidor do docker-compose que está rodando, basta executar o comando abaixo:

```bash
redis-cli set API_KEY_LIMIT:tokenEspecial 1000
```
- O prefixo no Redis deve ser __API_KEY_LIMIT:${TOKEN} {$LIMIT}__
