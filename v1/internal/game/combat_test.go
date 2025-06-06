package game

// // helper to update orc slice and remove dead ones
// func updateOrcs(orcs []*entity.OrcGrunt, dt float64) []*entity.OrcGrunt {
// 	for i := 0; i < len(orcs); {
// 		o := orcs[i]
// 		o.Update(dt)
// 		if !o.Alive {
// 			orcs = append(orcs[:i], orcs[i+1:]...)
// 			continue
// 		}
// 		i++
// 	}
// 	return orcs
// }

// func TestFootmanSurvivesAfterKill(t *testing.T) {
// 	f := entity.NewFootman(0, 0)
// 	f.Speed = 0
// 	o := entity.NewOrcGrunt(0, 0)
// 	o.Speed = 0

// 	m := entity.NewMilitary()
// 	m.AddUnit(f)
// 	orcs := []*entity.OrcGrunt{o}

// 	for steps := 0; o.Alive && steps < 10; steps++ {
// 		orcs = updateOrcs(orcs, 0.1)
// 		orcs = m.Update(0.1, orcs)
// 	}

// 	if o.Alive {
// 		t.Fatalf("orc grunt should be defeated")
// 	}
// 	if !f.Alive || f.Health() <= 0 {
// 		t.Fatalf("footman should survive with HP > 0, hp=%d", f.Health())
// 	}
// }

// func TestFootmanDiesLethalDamage(t *testing.T) {
// 	f := entity.NewFootman(0, 0)
// 	f.Speed = 0
// 	f.Hp = 2
// 	o := entity.NewOrcGrunt(0, 0)
// 	o.Speed = 0
// 	o.Damage = 5

// 	m := entity.NewMilitary()
// 	m.AddUnit(f)
// 	orcs := []*entity.OrcGrunt{o}

// 	orcs = updateOrcs(orcs, 0.1)
// 	orcs = m.Update(0.1, orcs)

// 	if m.Count() != 0 {
// 		t.Fatalf("expected footman removed after death")
// 	}
// }

// func TestCombatMultipleCombinations(t *testing.T) {
// 	combos := []struct{ f, g int }{{1, 2}, {2, 1}, {2, 2}}
// 	for _, c := range combos {
// 		m := entity.NewMilitary()
// 		var orcs []*entity.OrcGrunt
// 		for i := 0; i < c.f; i++ {
// 			f := entity.NewFootman(0, 0)
// 			f.Speed = 0
// 			m.AddUnit(f)
// 		}
// 		for i := 0; i < c.g; i++ {
// 			o := entity.NewOrcGrunt(0, 0)
// 			o.Speed = 0
// 			orcs = append(orcs, o)
// 		}

// 		for steps := 0; steps < 20 && m.Count() > 0 && len(orcs) > 0; steps++ {
// 			orcs = updateOrcs(orcs, 0.1)
// 			orcs = m.Update(0.1, orcs)
// 		}

// 		if m.Count() == 0 && len(orcs) == 0 {
// 			// both sides died - acceptable
// 			continue
// 		}
// 		if m.Count() == 0 && len(orcs) > 0 {
// 			// orcs won
// 			continue
// 		}
// 		if m.Count() > 0 && len(orcs) == 0 {
// 			// footmen won
// 			continue
// 		}
// 		t.Fatalf("combat did not resolve for combo %dv%d", c.f, c.g)
// 	}
// }

// func TestSimultaneousDamage(t *testing.T) {
// 	f := entity.NewFootman(0, 0)
// 	f.Speed = 0
// 	f.Hp = 1
// 	f.Damage = 5
// 	o := entity.NewOrcGrunt(0, 0)
// 	o.Speed = 0
// 	o.Hp = 5
// 	o.Damage = 1

// 	m := entity.NewMilitary()
// 	m.AddUnit(f)
// 	orcs := []*entity.OrcGrunt{o}

// 	orcs = updateOrcs(orcs, 0.1)
// 	orcs = m.Update(0.1, orcs)

// 	if f.Alive || o.Alive {
// 		t.Fatalf("both units should die in the same tick")
// 	}
// }

// func TestNoCombatWithoutOverlap(t *testing.T) {
// 	f := entity.NewFootman(0, 0)
// 	f.Speed = 0
// 	o := entity.NewOrcGrunt(1000, 0)
// 	o.Speed = 0

// 	m := entity.NewMilitary()
// 	m.AddUnit(f)
// 	orcs := []*entity.OrcGrunt{o}

// 	orcs = updateOrcs(orcs, 0.1)
// 	hpF := f.Health()
// 	hpO := o.Health()
// 	orcs = m.Update(0.1, orcs)

// 	if f.Health() != hpF || o.Health() != hpO {
// 		t.Fatalf("hp changed despite no overlap")
// 	}
// }

// func TestImmediateRemovalOfDeadUnits(t *testing.T) {
// 	f := entity.NewFootman(0, 0)
// 	f.Speed = 0
// 	o := entity.NewOrcGrunt(0, 0)
// 	o.Speed = 0

// 	m := entity.NewMilitary()
// 	m.AddUnit(f)
// 	orcs := []*entity.OrcGrunt{o}

// 	for steps := 0; steps < 10 && len(orcs) > 0; steps++ {
// 		orcs = updateOrcs(orcs, 0.1)
// 		orcs = m.Update(0.1, orcs)
// 	}

// 	if len(orcs) != 0 {
// 		t.Fatalf("orc not removed after death")
// 	}
// 	if m.Count() != 1 {
// 		// footman should survive this scenario
// 		t.Fatalf("footman unexpectedly removed")
// 	}
// }

// func TestDeadUnitsCannotAttack(t *testing.T) {
// 	f := entity.NewFootman(0, 0)
// 	f.Speed = 0
// 	o := entity.NewOrcGrunt(0, 0)
// 	o.Speed = 0
// 	o.Hp = 1

// 	m := entity.NewMilitary()
// 	m.AddUnit(f)
// 	orcs := []*entity.OrcGrunt{o}

// 	orcs = updateOrcs(orcs, 0.1)
// 	orcs = m.Update(0.1, orcs)

// 	hp := f.Hp
// 	// update again with same (dead) orc in list
// 	orcs = updateOrcs(orcs, 0.1)
// 	orcs = m.Update(0.1, orcs)

// 	if f.Hp != hp {
// 		t.Fatalf("dead orc continued to deal damage")
// 	}
// }

// func TestNoCombatWithDeadUnit(t *testing.T) {
// 	f := entity.NewFootman(0, 0)
// 	f.Speed = 0
// 	f.Alive = false
// 	o := entity.NewOrcGrunt(0, 0)
// 	o.Speed = 0

// 	m := entity.NewMilitary()
// 	m.AddUnit(f)
// 	orcs := []*entity.OrcGrunt{o}

// 	orcs = updateOrcs(orcs, 0.1)
// 	hp := o.Hp
// 	orcs = m.Update(0.1, orcs)

// 	if o.Hp != hp {
// 		t.Fatalf("combat occurred with dead footman")
// 	}
// }

// func TestBothUnitsDieSameTick(t *testing.T) {
// 	f := mob.NewFootman(0, 0)
// 	f.Speed = 0
// 	f.Hp = 1
// 	f.Damage = 5
// 	o := mob.NewOrcGrunt(0, 0)
// 	o.Speed = 0
// 	o.Hp = 5
// 	o.Damage = 5

// 	orcs := []*mob.OrcGrunt{o}

// 	orcs = updateOrcs(orcs, 0.1)
// }

// func TestRemovalDuringIterationNoPanic(t *testing.T) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Fatalf("panic during iteration: %v", r)
// 		}
// 	}()
// 	var orcs []*mob.OrcGrunt
// 	for i := 0; i < 3; i++ {
// 		f := mob.NewFootman(0, 0)
// 		f.Speed = 0
// 		f.Hp = 5
// 	}
// 	for i := 0; i < 3; i++ {
// 		o := mob.NewOrcGrunt(0, 0)
// 		o.Speed = 0
// 		o.Hp = 5
// 		o.Damage = 5
// 		orcs = append(orcs, o)
// 	}

// 	for steps := 0; steps < 5 && (m.Count() > 0 || len(orcs) > 0); steps++ {
// 		orcs = updateOrcs(orcs, 0.1)
// 	}
// }
