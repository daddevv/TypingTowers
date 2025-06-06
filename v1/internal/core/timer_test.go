package core

import "testing"

func TestCooldownTimerProgress(t *testing.T) {
	timer := NewCooldownTimer(1.0)
	if p := timer.Progress(); p != 0 {
		t.Fatalf("expected progress 0 got %f", p)
	}
	timer.Tick(0.5)
	if p := timer.Progress(); p < 0.49 || p > 0.51 {
		t.Fatalf("expected progress around 0.5 got %f", p)
	}
	timer.Tick(0.6)
	if p := timer.Progress(); p != 1 {
		t.Fatalf("expected progress 1 got %f", p)
	}
	timer.Reset()
	if p := timer.Progress(); p != 0 {
		t.Fatalf("expected progress 0 after reset got %f", p)
	}
}

// func TestFarmerProgressReset(t *testing.T) {
// 	f := NewFarmer()
// 	f.SetInterval(0.1)
// 	f.SetCooldown(0.1)
// 	word := f.Update(0.11)
// 	if word == "" {
// 		t.Fatalf("expected word generated")
// 	}
// 	if p := f.CooldownProgress(); p != 1 {
// 		t.Fatalf("expected progress 1 after cooldown completed got %f", p)
// 	}
// 	f.OnWordCompleted(word, nil)
// 	if p := f.CooldownProgress(); p != 0 {
// 		t.Fatalf("expected progress reset to 0 got %f", p)
// 	}
// }
