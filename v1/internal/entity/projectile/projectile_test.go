package projectile

// func TestProjectileIntercept(t *testing.T) {
// 	base := NewBase(400, 100, 10)
// 	mob := NewMob(200, 100, base, 1, 0) // stationary mob
// 	g := &Game{mobs: []Enemy{mob}, input: NewInput(), typing: NewTypingStats()}
// 	p := NewProjectile(g, 100, 100, mob, 1, 50, 0) // Increased speed from 5 to 50
// 	for i := 0; i < 200 && mob.Alive() && p.alive; i++ {
// 		mob.Update(0.016)
// 		p.Update(0.016)
// 	}
// 	if mob.Alive() {
// 		t.Errorf("projectile did not hit the mob")
// 	}
// }
