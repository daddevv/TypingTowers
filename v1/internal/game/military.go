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

// rectOverlap checks if two axis-aligned rectangles overlap.
func rectOverlap(ax, ay, aw, ah, bx, by, bw, bh int) bool {
	return ax < bx+bw && ax+aw > bx && ay < by+bh && ay+ah > by
}

// Update advances all units, resolves combat with orc grunts, and removes any
// that are no longer alive.
func (m *Military) Update(dt float64, orcs []*OrcGrunt) {
	for i := 0; i < len(m.units); {
		u := m.units[i]
		u.Update(dt)
		if !u.Alive() {
			m.units = append(m.units[:i], m.units[i+1:]...)
			continue
		}
		// Combat resolution against orc grunts
		fx, fy, fw, fh := u.Hitbox()
		for _, o := range orcs {
			if !o.Alive() {
				continue
			}
			ox, oy, ow, oh := o.Hitbox()
			if rectOverlap(fx, fy, fw, fh, ox, oy, ow, oh) {
				o.Damage(u.damage)
				u.Damage(o.AttackDamage())
			}
		}
		i++
	}
}

// Units returns the list of active Footmen.
func (m *Military) Units() []*Footman { return m.units }

// Count returns the number of active units.
func (m *Military) Count() int { return len(m.units) }
