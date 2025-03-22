package singleton

import (
	"fmt"
	"sync"
)

// Shop represents a shop with inventory and name
type Shop struct {
	name      string
	inventory int
	mutex     sync.Mutex
}

// private variable to hold the single instance
var (
	shopInstance *Shop
	once         sync.Once
)

// GetShopInstance returns the single shop instance
func GetShopInstance() *Shop {
	once.Do(func() {
		shopInstance = &Shop{
			name:      "MyStore",
			inventory: 100,
		}
		fmt.Println("Created new shop instance")
	})
	return shopInstance
}

// SellItem decrements inventory
func (s *Shop) SellItem(quantity int) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.inventory >= quantity {
		s.inventory -= quantity
		fmt.Printf("Sold %d items. Remaining inventory: %d\n", quantity, s.inventory)
		return true
	}
	fmt.Printf("Cannot sell %d items. Only %d items in inventory\n", quantity, s.inventory)
	return false
}

// GetInventory returns current inventory
func (s *Shop) GetInventory() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.inventory
}
