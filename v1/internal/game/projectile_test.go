package game

import "testing"

func TestProjectileIntercept(t *testing.T) {
	mob := NewMob(200, 100, nil, 1, 1)
	g := &Game{mobs: []*Mob{mob}}
	p := NewProjectile(g, 100, 100, mob, 1, 5, 0)
	for i := 0; i < 200 && mob.alive && p.alive; i++ {
		mob.Update()
		p.Update()
	}
	if mob.alive {
		t.Errorf("projectile did not hit the mob")
	}
}
