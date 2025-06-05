package game

// Gold tracks the player's gold resources.
type Gold struct {
	amount int
}

// Add increases the gold amount.
func (g *Gold) Add(n int) { g.amount += n }

// Spend subtracts the given amount if available and returns true.
func (g *Gold) Spend(n int) bool {
	if g.amount < n {
		return false
	}
	g.amount -= n
	return true
}

// Amount returns the current gold total.
func (g *Gold) Amount() int { return g.amount }

// Set sets the gold amount directly.
func (g *Gold) Set(n int) { g.amount = n }

// Wood tracks the player's wood resources.
type Wood struct {
	amount int
}

// Add increases the wood amount.
func (w *Wood) Add(n int) { w.amount += n }

// Spend subtracts the given amount if available and returns true.
func (w *Wood) Spend(n int) bool {
	if w.amount < n {
		return false
	}
	w.amount -= n
	return true
}

// Amount returns the current wood total.
func (w *Wood) Amount() int { return w.amount }

// Set sets the wood amount directly.
func (w *Wood) Set(n int) { w.amount = n }

// Stone tracks the player's stone resources.
type Stone struct {
	amount int
}

// Add increases the stone amount.
func (s *Stone) Add(n int) { s.amount += n }

// Spend subtracts the given amount if available and returns true.
func (s *Stone) Spend(n int) bool {
	if s.amount < n {
		return false
	}
	s.amount -= n
	return true
}

// Amount returns the current stone total.
func (s *Stone) Amount() int { return s.amount }

// Set sets the stone amount directly.
func (s *Stone) Set(n int) { s.amount = n }

// Iron tracks the player's iron resources.
type Iron struct {
	amount int
}

// Add increases the iron amount.
func (i *Iron) Add(n int) { i.amount += n }

// Spend subtracts the given amount if available and returns true.
func (i *Iron) Spend(n int) bool {
	if i.amount < n {
		return false
	}
	i.amount -= n
	return true
}

// Amount returns the current iron total.
func (i *Iron) Amount() int { return i.amount }

// Set sets the iron amount directly.
func (i *Iron) Set(n int) { i.amount = n }

// Food tracks the player's food resources.
type Food struct {
	amount int
}

// Add increases the food amount.
func (f *Food) Add(n int) { f.amount += n }

// Spend subtracts the given amount if available and returns true.
func (f *Food) Spend(n int) bool {
	if f.amount < n {
		return false
	}
	f.amount -= n
	return true
}

// Amount returns the current food total.
func (f *Food) Amount() int { return f.amount }

// Set sets the food amount directly.
func (f *Food) Set(n int) { f.amount = n }

// ResourcePool aggregates all resource types for the player.
type ResourcePool struct {
	Gold  Gold
	Food  Food
	Wood  Wood
	Stone Stone
	Iron  Iron
}

// AddGold adds the specified amount of gold.
func (r *ResourcePool) AddGold(n int) { r.Gold.Add(n) }

// AddFood adds the specified amount of food.
func (r *ResourcePool) AddFood(n int) { r.Food.Add(n) }

// GoldAmount returns the current gold total.
func (r *ResourcePool) GoldAmount() int { return r.Gold.Amount() }

// FoodAmount returns the current food total.
func (r *ResourcePool) FoodAmount() int { return r.Food.Amount() }
