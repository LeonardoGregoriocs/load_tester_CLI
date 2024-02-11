# Load Tester CLI em Go

Este é um sistema CLI desenvolvido em Go para realizar testes de carga em serviços web. Com esta ferramenta, você pode testar o desempenho de qualquer serviço HTTP, fornecendo a URL do serviço, o número total de requests desejados e a quantidade de chamadas simultâneas.

## Funcionalidades

### Entrada de Parâmetros via CLI

Você pode configurar os testes fornecendo os seguintes parâmetros via linha de comando:

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requests a serem enviados.
- `--concurrency`: Número de chamadas simultâneas.

### Execução do Teste

O sistema realizará requests HTTP para a URL especificada, distribuindo os requests de acordo com o nível de concorrência definido, garantindo que o número total de requests seja cumprido.

### Geração de Relatório

Após a execução dos testes, o sistema irá gerar um relatório detalhado contendo as seguintes informações:

- Tempo total gasto na execução.
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200 (OK).
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Execução da Aplicação

Você pode executar a aplicação facilmente utilizando Docker. Basta seguir o exemplo abaixo:


```bash
docker build -t loadtester .
``` 
```bash
docker run loadtester --url=http://google.com --requests=1000 --concurrency=10
``` 