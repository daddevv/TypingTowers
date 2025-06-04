package game

import "testing"

func TestProjectileIntercept(t *testing.T) {
	mob := NewMob(200, 100, nil, 1, 1)
	g := &Game{mobs: []Enemy{mob}}
	p := NewProjectile(g, 100, 100, mob, 1, 5, 0)
	for i := 0; i < 200 && mob.Alive() && p.alive; i++ {
		mob.Update(0.016)
		p.Update(0.016)
	}
	if mob.Alive() {
		t.Errorf("projectile did not hit the mob")
	}
}
