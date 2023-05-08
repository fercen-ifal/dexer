package utils

import "github.com/fercen-ifal/dexer/models"

type ElectricityBillsByYear []models.ElectricityBill

func (items ElectricityBillsByYear) Len() int           { return len(items) }
func (items ElectricityBillsByYear) Less(i, j int) bool { return items[i].Year < items[j].Year }
func (items ElectricityBillsByYear) Swap(i, j int)      { items[i], items[j] = items[j], items[i] }
