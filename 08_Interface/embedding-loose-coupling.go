package interfaces

import "fmt"

type Tax interface {
	SetIncomeTax(float64)
	GetIncomeTax() float64
}

type IndividualTax struct {
	incomeTax float64
}

func (it *IndividualTax) SetIncomeTax(tax float64) {
	it.incomeTax = tax
}

func (it *IndividualTax) GetIncomeTax() float64 {
	return it.incomeTax
}

type CorporateTax struct {
	incomeTax float64
}

func (ct *CorporateTax) SetIncomeTax(tax float64) {
	ct.incomeTax = tax
}

func (ct *CorporateTax) GetIncomeTax() float64 {
	return ct.incomeTax
}

type Taxpayer struct {
	tax  Tax
	name string
}

func NewTaxpayer(name string, tax Tax) *Taxpayer {
	return &Taxpayer{
		tax:  tax,
		name: name,
	}
}

func (tp *Taxpayer) SetTax(taxToPay float64) {
	tp.tax.SetIncomeTax(taxToPay)
}

func (tp *Taxpayer) GetTax() float64 {
	return tp.tax.GetIncomeTax()
}

func Embedding() {
	individual := &IndividualTax{}
	corporate := &CorporateTax{}

	individualTaxpayer := NewTaxpayer("Wah Modi Ji", individual)
	corporateTaxpayer := NewTaxpayer("Nirmala Tai", corporate)

	individualTaxpayer.SetTax(99999.99999)
	corporateTaxpayer.SetTax(1.1)

	fmt.Println("Individual Income Tax:", individualTaxpayer.GetTax())
	fmt.Println("Corporate Income Tax:", corporateTaxpayer.GetTax())
}
