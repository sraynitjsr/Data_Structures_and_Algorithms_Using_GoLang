package embedding_example

type IndividualTax struct {
	incomeTax float64
}

func (it *IndividualTax) SetIncomeTax(tax float64) {
	it.incomeTax = tax
}

func (it *IndividualTax) GetIncomeTax() float64 {
	return it.incomeTax
}
