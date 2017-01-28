# goboleto
[![Coverage Status](https://img.shields.io/badge/coverage-50%25-brightgreen.svg)]()

Gerador de boletos para diversos bancos usando Go.

[![golang.sh-600x600.jpg](https://s27.postimg.org/coqxnki9f/golang_sh_600x600.jpg)](https://postimg.org/image/yb5y4lgtr/)

## Features
* Gera linha digitavel e imagem do codigo;
* Apenas boletos registrados, conforme novas regras FEBRABAN;
* Gera remessas para os bancos;

## Usage

[Veja a documentação no GoDoc](https://godoc.org/github.com/kezzbr/goboleto)

## TODO
* Falta gerar nosso numero para: Bradesco, Caixa, Itau e Santander
* Gerar layout dos boletos;
* Gerar remessas dos boletos registrados;

## Dependencies
* [boombuler/barcode](github.com/boombuler/barcode)
