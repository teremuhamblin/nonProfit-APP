package storage

import "github.com/teremuhamblin/nonProfit-APP/internal/domain"

type MemoryStore struct {
	Donors    []domain.Donor
	Donations []domain.Donation
	Campaigns []domain.Campaign
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		Donors:    []domain.Donor{},
		Donations: []domain.Donation{},
		Campaigns: []domain.Campaign{},
	}
}
