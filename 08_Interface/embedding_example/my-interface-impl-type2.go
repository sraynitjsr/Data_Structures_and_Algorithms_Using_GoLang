package embedding_example

type CorporateTax struct {
	incomeTax float64
}

func (ct *CorporateTax) SetIncomeTax(tax float64) {
	ct.incomeTax = tax
}

func (ct *CorporateTax) GetIncomeTax() float64 {
	return ct.incomeTax
}
