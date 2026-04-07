# Calculate Tax Go

Uma ferramenta CLI em Go para calcular impostos sobre operações de compra e venda de ações no mercado brasileiro.

## Descrição

Este projeto implementa um calculador de impostos para operações de ações, seguindo as regras fiscais brasileiras. Ele processa operações de compra e venda, calcula o preço médio ponderado, compensa prejuízos acumulados e aplica isenções e taxas de imposto sobre lucros.

## Funcionalidades

- **Processamento de Operações**: Suporta operações de compra (`buy`) e venda (`sell`).
- **Cálculo de Preço Médio**: Mantém o preço médio ponderado das ações.
- **Compensação de Prejuízos**: Deduz prejuízos acumulados dos lucros antes de calcular impostos.
- **Isenção Fiscal**: Isenta operações de venda com valor total até R$ 20.000,00 (exceto se houver compensação de prejuízo).
- **Taxa de Imposto**: Aplica 20% sobre o lucro tributável.
- **Interface CLI**: Lê operações em formato JSON via entrada padrão e retorna impostos calculados.

## Instalação

### Pré-requisitos

- Go 1.26.1 ou superior instalado no sistema.

### Clonando e Construindo

1. Clone o repositório:
   ```bash
   git clone https://github.com/jadersonmarc/calculate-tax-go.git
   cd calculate-tax-go
   ```

2. Baixe as dependências:
   ```bash
   go mod download
   ```

3. Construa o executável:
   ```bash
   go build -o calculate-tax main.go
   ```

## Uso

Execute o programa e forneça operações em formato JSON via entrada padrão. Cada linha deve conter um array JSON de operações.

### Exemplo de Entrada

```json
[{"operation":"buy", "unit-cost":"10.00", "quantity":100}, {"operation":"sell", "unit-cost":"15.00", "quantity":50}]
```

### Exemplo de Uso

```bash
./calculate-tax
> [{"operation":"buy", "unit-cost":"10.00", "quantity":100}, {"operation":"sell", "unit-cost":"15.00", "quantity":50}]
[{"tax":"0.00"}]
> exit
```

### Formato da Operação

Cada operação é um objeto JSON com:
- `operation`: `"buy"` ou `"sell"`
- `unit-cost`: String representando o custo unitário (ex: `"10.00"`)
- `quantity`: Inteiro representando a quantidade

### Saída

A saída é um array JSON de objetos com o imposto calculado para cada operação:
- `tax`: String com o valor do imposto (ex: `"100.00"`)

## Testes

Para executar os testes:

```bash
go test ./...
```

Os testes incluem casos para operações de compra/venda, compensação de prejuízos e isenções.

## Estrutura do Projeto

- `main.go`: Ponto de entrada da CLI.
- `operation.go`: Definição das estruturas de operação.
- `calculator.go`: Lógica de cálculo de impostos e preço médio.
- `processor.go`: Processamento das operações.
- `Tax.go`: Estruturas de saída de impostos.
- `money/`: Pacote para manipulação de valores monetários (atualmente não utilizado extensivamente).
- Arquivos de teste: `*_test.go` para cada módulo.

## Dependências

- [github.com/shopspring/decimal](https://github.com/shopspring/decimal): Para manipulação precisa de números decimais.

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.

## Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.