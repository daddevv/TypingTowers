package game

import "testing"


func TestProjectileIntercept(t *testing.T) {
	mob := NewMob(200, 100, nil, 1)
	p := NewProjectile(100, 100, mob)
	for i := 0; i < 200 && mob.alive && p.alive; i++ {
		mob.Update()
		p.Update()
	}
	if mob.alive {
		t.Errorf("projectile did not hit the mob")
	}
}
