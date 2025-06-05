package game

import "testing"

// helper to update orc slice and remove dead ones
func updateOrcs(orcs []*OrcGrunt, dt float64) []*OrcGrunt {
	for i := 0; i < len(orcs); {
		o := orcs[i]
		o.Update(dt)
		if !o.Alive() {
			orcs = append(orcs[:i], orcs[i+1:]...)
			continue
		}
		i++
	}
	return orcs
}

func TestFootmanSurvivesAfterKill(t *testing.T) {
	f := NewFootman(0, 0)
	f.speed = 0
	o := NewOrcGrunt(0, 0)
	o.speed = 0

	m := NewMilitary()
	m.AddUnit(f)
	orcs := []*OrcGrunt{o}

	for steps := 0; o.Alive() && steps < 10; steps++ {
		orcs = updateOrcs(orcs, 0.1)
		orcs = m.Update(0.1, orcs)
	}

	if o.Alive() {
		t.Fatalf("orc grunt should be defeated")
	}
	if !f.Alive() || f.Health() <= 0 {
		t.Fatalf("footman should survive with HP > 0, hp=%d", f.Health())
	}
}

func TestFootmanDiesLethalDamage(t *testing.T) {
	f := NewFootman(0, 0)
	f.speed = 0
	f.hp = 2
	o := NewOrcGrunt(0, 0)
	o.speed = 0
	o.damage = 5

	m := NewMilitary()
	m.AddUnit(f)
	orcs := []*OrcGrunt{o}

	orcs = updateOrcs(orcs, 0.1)
	orcs = m.Update(0.1, orcs)

	if m.Count() != 0 {
		t.Fatalf("expected footman removed after death")
	}
}

func TestCombatMultipleCombinations(t *testing.T) {
	combos := []struct{ f, g int }{{1, 2}, {2, 1}, {2, 2}}
	for _, c := range combos {
		m := NewMilitary()
		var orcs []*OrcGrunt
		for i := 0; i < c.f; i++ {
			f := NewFootman(0, 0)
			f.speed = 0
			m.AddUnit(f)
		}
		for i := 0; i < c.g; i++ {
			o := NewOrcGrunt(0, 0)
			o.speed = 0
			orcs = append(orcs, o)
		}

		for steps := 0; steps < 20 && m.Count() > 0 && len(orcs) > 0; steps++ {
			orcs = updateOrcs(orcs, 0.1)
			orcs = m.Update(0.1, orcs)
		}

		if m.Count() == 0 && len(orcs) == 0 {
			// both sides died - acceptable
			continue
		}
		if m.Count() == 0 && len(orcs) > 0 {
			// orcs won
			continue
		}
		if m.Count() > 0 && len(orcs) == 0 {
			// footmen won
			continue
		}
		t.Fatalf("combat did not resolve for combo %dv%d", c.f, c.g)
	}
}

func TestSimultaneousDamage(t *testing.T) {
	f := NewFootman(0, 0)
	f.speed = 0
	f.hp = 1
	f.damage = 5
	o := NewOrcGrunt(0, 0)
	o.speed = 0
	o.hp = 5
	o.damage = 1

	m := NewMilitary()
	m.AddUnit(f)
	orcs := []*OrcGrunt{o}

	orcs = updateOrcs(orcs, 0.1)
	orcs = m.Update(0.1, orcs)

	if f.Alive() || o.Alive() {
		t.Fatalf("both units should die in the same tick")
	}
}

func TestNoCombatWithoutOverlap(t *testing.T) {
	f := NewFootman(0, 0)
	f.speed = 0
	o := NewOrcGrunt(1000, 0)
	o.speed = 0

	m := NewMilitary()
	m.AddUnit(f)
	orcs := []*OrcGrunt{o}

	orcs = updateOrcs(orcs, 0.1)
	hpF := f.Health()
	hpO := o.Health()
	orcs = m.Update(0.1, orcs)

	if f.Health() != hpF || o.Health() != hpO {
		t.Fatalf("hp changed despite no overlap")
	}
}

func TestImmediateRemovalOfDeadUnits(t *testing.T) {
	f := NewFootman(0, 0)
	f.speed = 0
	o := NewOrcGrunt(0, 0)
	o.speed = 0

	m := NewMilitary()
	m.AddUnit(f)
	orcs := []*OrcGrunt{o}

	for steps := 0; steps < 10 && len(orcs) > 0; steps++ {
		orcs = updateOrcs(orcs, 0.1)
		orcs = m.Update(0.1, orcs)
	}

	if len(orcs) != 0 {
		t.Fatalf("orc not removed after death")
	}
	if m.Count() != 1 {
		// footman should survive this scenario
		t.Fatalf("footman unexpectedly removed")
	}
}

func TestDeadUnitsCannotAttack(t *testing.T) {
	f := NewFootman(0, 0)
	f.speed = 0
	o := NewOrcGrunt(0, 0)
	o.speed = 0
	o.hp = 1

	m := NewMilitary()
	m.AddUnit(f)
	orcs := []*OrcGrunt{o}

	orcs = updateOrcs(orcs, 0.1)
	orcs = m.Update(0.1, orcs)

	hp := f.Health()
	// update again with same (dead) orc in list
	orcs = updateOrcs(orcs, 0.1)
	orcs = m.Update(0.1, orcs)

	if f.Health() != hp {
		t.Fatalf("dead orc continued to deal damage")
	}
}

func TestNoCombatWithDeadUnit(t *testing.T) {
	f := NewFootman(0, 0)
	f.speed = 0
	f.alive = false
	o := NewOrcGrunt(0, 0)
	o.speed = 0

	m := NewMilitary()
	m.AddUnit(f)
	orcs := []*OrcGrunt{o}

	orcs = updateOrcs(orcs, 0.1)
	hp := o.Health()
	orcs = m.Update(0.1, orcs)

	if o.Health() != hp {
		t.Fatalf("combat occurred with dead footman")
	}
}

func TestBothUnitsDieSameTick(t *testing.T) {
	f := NewFootman(0, 0)
	f.speed = 0
	f.hp = 1
	f.damage = 5
	o := NewOrcGrunt(0, 0)
	o.speed = 0
	o.hp = 5
	o.damage = 5

	m := NewMilitary()
	m.AddUnit(f)
	orcs := []*OrcGrunt{o}

	orcs = updateOrcs(orcs, 0.1)
	orcs = m.Update(0.1, orcs)

	if m.Count() != 0 || len(orcs) != 0 {
		t.Fatalf("both units should be removed when they die in the same tick")
	}
}

func TestRemovalDuringIterationNoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panic during iteration: %v", r)
		}
	}()
	m := NewMilitary()
	var orcs []*OrcGrunt
	for i := 0; i < 3; i++ {
		f := NewFootman(0, 0)
		f.speed = 0
		f.hp = 1
		m.AddUnit(f)
	}
	for i := 0; i < 3; i++ {
		o := NewOrcGrunt(0, 0)
		o.speed = 0
		o.damage = 5
		orcs = append(orcs, o)
	}

	for steps := 0; steps < 5 && (m.Count() > 0 || len(orcs) > 0); steps++ {
		orcs = updateOrcs(orcs, 0.1)
		orcs = m.Update(0.1, orcs)
	}
}
