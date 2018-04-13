# Transactions Test

Lógica para cálculo de saldo atual em conta. 

## Decrição
Dado dois arquivos cvc's: contas.cvc, transacoes.cvc, cujo formatos são:

1. contas.cvc
```
idConta,saldoAtual
1,2000
2,2000

```

2. transacoes.cvc
```
idConta,valorTransacao
1,2012
2,210
```

O programa deve aplicar as transações na conta relacionada, e se, ao aplicar uma transação no saldo atual da conta, caso o saldo da conta se torne negativo, é necessário aplicar uma multa de 5 reais.

## Dependências

1. Garanta que [Golang](https://golang.org/dl/) está instalada e o $GOPATH setado
2. Lib [Decimal](https://github.com/shopspring/decimal/)

## Instalação

1. mkdir -p $GOPATH/src/ (Opcional)
2. cd $GOPATH/src/ (Opcional)
3. go get github.com/alanepontes/transactions-test
4. go get para pegar as dependencias
5. go build

## Testes

1. go test -v

## Uso

Após processo de instalação, faça
1. ./transactions-test data/contas.cvc data/transacoes.cvc

## TODO
1. Analisar o uso do padrão de projeto Money em relação a lib Decimal
2. Criar namespace e cli separado para facilitar o uso da aplicação: go get repo
3. Verificar a necessidade de usar a lib asserts ao invés da lib de teste default

