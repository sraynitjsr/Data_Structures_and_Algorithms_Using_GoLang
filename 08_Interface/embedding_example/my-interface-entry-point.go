// Final Commit Done

package embedding_example

import "fmt"

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
