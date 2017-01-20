package goboleto

import (
	"time"
	"fmt"
)

func init() {
	BilletBB()
}

func BilletBB() {

	// static data, you should keep this configured in somewhere
	var bank Bank = BB{
		Account: 8888,
		Agency: 99999,
		Contrato: 12312351,
		Carteira: 15,
		Convenio: 1234,
		FormatacaoConvenio: 4,
		FormatacaoNossoNumero: 1,
		VariacaoCarteira: 6,
		Company: Company{
			Name: "Nome da empresa",
			LegalName: "Razao social",
			Address: "Endereço",
			Contact: "Email e telefone",
			Document: "CNPJ/CPF",
		},
	}

	// dynamic data, you should have this data coming from a database
	var document = Document{
		Id: 123,
		Value: 999.99,
		ValueTax: 1.00,
		ValueDiscount: 0.00,
		ValueForfeit: 199.99,
		OurNumber: 123,
		FebrabanType: "md",
		Date: time.Now(),
		DateDue: time.Now().AddDate(0, 0, 4),
		Instructions: [6]string{
			"Não receber após o vencimento",
			"Após vencimento, receber apenas no meu banco",
		},
		Payer: Payer{
			Name: "Nome completo",
			Address: "Endereço",
			Contact: "Email e telefone",
		},
	}
	
	
	// Print the layout
	//bank.Layout(document)
	
	// Optional, to use in your backend
	var barcode = bank.Barcode(document)
	fmt.Println(barcode.toString())
	//image := barcode.Image()
	//digitable := barcode.Digitable()
	
}