package game

// Military manages all player-controlled units such as Footmen.
type Military struct {
	units []*Footman
}

// NewMilitary creates an empty Military manager.
func NewMilitary() *Military {
	return &Military{units: make([]*Footman, 0)}
}

// AddUnit registers a new Footman with the military system.
func (m *Military) AddUnit(f *Footman) {
	if f != nil {
		m.units = append(m.units, f)
	}
}

// Update advances all units and removes any that are no longer alive.
func (m *Military) Update(dt float64) {
	for i := 0; i < len(m.units); {
		u := m.units[i]
		u.Update(dt)
		if !u.Alive() {
			m.units = append(m.units[:i], m.units[i+1:]...)
			continue
		}
		i++
	}
}

// Units returns the list of active Footmen.
func (m *Military) Units() []*Footman { return m.units }

// Count returns the number of active units.
func (m *Military) Count() int { return len(m.units) }
