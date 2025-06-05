package game

// TowerModifiers defines adjustments applied globally to tower stats.
type TowerModifiers struct {
	DamageMult   float64 `json:"damage_mult"`
	RangeMult    float64 `json:"range_mult"`
	FireRateMult float64 `json:"fire_rate_mult"`
	AmmoAdd      int     `json:"ammo_add"`
}

// Merge combines two modifier sets multiplicatively for multipliers
// and additively for ammo capacity increases.
func (m TowerModifiers) Merge(o TowerModifiers) TowerModifiers {
	if m.DamageMult == 0 {
		m.DamageMult = 1
	}
	if m.RangeMult == 0 {
		m.RangeMult = 1
	}
	if m.FireRateMult == 0 {
		m.FireRateMult = 1
	}
	if o.DamageMult != 0 {
		m.DamageMult *= o.DamageMult
	}
	if o.RangeMult != 0 {
		m.RangeMult *= o.RangeMult
	}
	if o.FireRateMult != 0 {
		m.FireRateMult *= o.FireRateMult
	}
	m.AmmoAdd += o.AmmoAdd
	return m
}
