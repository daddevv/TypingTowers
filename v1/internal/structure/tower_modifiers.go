package structure

// TowerModifiers defines adjustments applied globally to tower stats.
type TowerModifiers struct {
	DamageMult   float64
	RangeMult    float64
	FireRateMult float64
	AmmoAdd      int
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
