package embedding_example

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
