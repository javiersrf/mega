# MegaSena Optimizer API

API em Go para calcular a **melhor combinação de apostas da Mega-Sena** dentro de um orçamento informado, considerando:

- Probabilidade individual de cada jogo
- Combinação ótima sem repetir jogos (0/1 Knapsack)
- Probabilidade final de ganhar **ao menos uma vez**

---

## 🚀 Executando a API

Inicie a API via script:

```bash
./run-server.sh <PORT>
```

**Exemplo:**

```bash
./run-server.sh 8080
```

O servidor será acessível em `http://localhost:8080`

## 🧪 Executando os testes

```bash
./run-tests.sh <PORT>
```

A porta é propagada como variável de ambiente para os testes.

## 📡 Endpoint

**POST** `/api/v1/megasena/calculate`

### Request

```json
{
    "budget": 200.0,
    "games": [
        { "numbers": 6, "price": 5.0 },
        { "numbers": 7, "price": 35.0 },
        { "numbers": 8, "price": 140.0 }
    ]
}
```

### Response

```json
{
    "items": [
        { "game": 6, "quantity": 1, "amount": 5.0 },
        { "game": 7, "quantity": 1, "amount": 35.0 }
    ],
    "totalAmount": 40.0,
    "totalBenefit": 0.00000042,
    "finalProbability": 0.00000041
}
```

## 📁 Estrutura do Projeto

```
services/
    calculate.go
    calculate_best_combination_test.go
cmd/
    server/
        main.go
run-server.sh
run-tests.sh
README.md
```

## 🧮 Critério de Otimização

**Sem repetição de jogos** (0/1 knapsack)

**Função objetivo:** Soma das probabilidades individuais

**Probabilidade combinada:**

```
P(at least 1 win) = 1 − Π(1 − pᵢ)
```

## 🛠️ Requisitos

- Go 1.21+
- Linux / macOS / WSL

## 📝 Licença

MIT
