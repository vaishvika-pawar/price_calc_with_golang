package prices

import (
	"fmt"
	"price_calc_with_golang/conversion"
	"price_calc_with_golang/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"` //unserscore tells not to add this key
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

// func () - converts functions to a method
func (job *TaxIncludedPriceJob) Process() error {

	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	return job.IOManager.WriteResult(job)
}

// Output *TaxIncludedPriceJob,
// so that we can store value in same place in memory
func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	//Below line will return address of the pointer
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
