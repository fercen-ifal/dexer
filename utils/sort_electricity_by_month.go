package utils

import "github.com/fercen-ifal/dexer/models"

type ElectricityBillsByMonth []models.ElectricityBill

func (items ElectricityBillsByMonth) Len() int { return len(items) }
func (items ElectricityBillsByMonth) Less(i, j int) bool {
	return items[i].Month < items[j].Month
}
func (items ElectricityBillsByMonth) Swap(i, j int) { items[i], items[j] = items[j], items[i] }
