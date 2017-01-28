package goboleto

import (
	"html/template"
	"net/http"
)

// Bradesco
// Source: (https://banco.bradesco/assets/pessoajuridica/pdf/4008-524-0121-08-layout-cobranca-versao-portuguesSS28785.pdf)
type Bradesco struct {
	Agency                int
	Account               int
	Convenio              int
	Contrato              int
	Carteira              int
	VariacaoCarteira      int
	FormatacaoConvenio    int
	FormatacaoNossoNumero int
	Company               Company
}

// configBradesco is a global for this bank configs
var configBradesco = bankConfig{
	Id:           237,
	Aceite:       "N",
	Currency:     9,
	CurrencyName: "R$",
}

// Barcode Get the Barcode, creating a BarcodeNumber
func (b Bradesco) Barcode(d Document) Barcode {

	// TODO, bank numbers (nosso numero, de acordo com a carteira e convenio)

	// Create a new Barcode
	var n Barcode = &BarcodeNumber{
		BankId:        configBradesco.Id,
		CurrencyId:    configBradesco.Currency,
		DateDueFactor: dateDueFactor(d.DateDue),
		Value:         formatValue(d.Value),
		//BankNumbers: fmt.Sprintf("%0"+strconv.Itoa(bankNumbersSize)+"s", bn),
	}
	n.verification()
	return n
}

// Transference Return the transference file (arquivo de remessa)
func (b Bradesco) Transference(d Document) {
	// TODO
}

// Layout return a HTML template using a document
func (b Bradesco) Layout(w http.ResponseWriter, d Document) {
	var barcode Barcode = b.Barcode(d)
	layout, _ := template.ParseFiles("templates/bradesco.html")
	layout.ExecuteTemplate(w, "bradesco", map[string]interface{}{
		"Barcode": barcode,
		"Document": d,
		"Bank": b,
	})
}
