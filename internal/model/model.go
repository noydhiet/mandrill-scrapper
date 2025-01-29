package model

import "time"

type Manufacturer struct {
	Name                  string
	Country               string
	RegisteredProducts    int
	Lawsuits              bool
	ProductRecalls        bool
	PatentExpiry          PatentExpiry
	Distributors          int
	GMPStandards          []string
	QualityCertifications []string
	RDType                string
	TherapeuticSpecialty  string
	ProductsInPipeline    int
	Products              []Product
}

type Product struct {
	BrandName           string
	Dosage              string
	ColdStorageRequired bool
	UsageFrequency      string
	GovernmentPurchase  int
	ProcurementPrice    float64
	RetailPrice         float64
	PatentExpiry        string
}

type PatentExpiry struct {
	Category   string
	ExpiryDate time.Time
	BrandName  string
}
