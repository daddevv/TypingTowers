package game

import (
	"bytes"
	"math"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const sampleRate = 22050

var (
	audioCtx *audio.Context
)

func getAudioContext() *audio.Context {
	if audioCtx == nil {
		audioCtx = audio.NewContext(sampleRate)
	}
	return audioCtx
}

// SoundManager handles basic sound effects and music.
type SoundManager struct {
	ctx   *audio.Context
	mute  bool
	music *audio.Player
}

// NewSoundManager creates a SoundManager with an audio context.
func NewSoundManager() *SoundManager {
	return &SoundManager{ctx: getAudioContext()}
}

// StartMusic begins a simple looping tone for the title screen.
func (s *SoundManager) StartMusic() {
	if s == nil || s.mute || s.music != nil {
		return
	}
	buf := generateSineWave(110, 0.5)
	loop := audio.NewInfiniteLoop(bytes.NewReader(buf), sampleRate/2)
	p, _ := s.ctx.NewPlayer(loop)
	p.Play()
	s.music = p
}

// StopMusic stops any playing music.
func (s *SoundManager) StopMusic() {
	if s == nil || s.music == nil {
		return
	}
	s.music.Close()
	s.music = nil
}

// ToggleMute enables or disables all sound output.
func (s *SoundManager) ToggleMute() { s.mute = !s.mute }

// PlayBeep plays a short beep sound if not muted.
func (s *SoundManager) PlayBeep() {
	if s == nil || s.mute {
		return
	}
	buf := generateSineWave(440, 0.1)
	p := (*audio.Context).NewPlayerFromBytes(s.ctx, buf)
	p.Play()
}

// PlayClank plays a lower-pitched clank sound if not muted.
func (s *SoundManager) PlayClank() {
	if s == nil || s.mute {
		return
	}
	buf := generateSineWave(220, 0.1)
	p := (*audio.Context).NewPlayerFromBytes(s.ctx, buf)
	p.Play()
}

// generateSineWave returns raw PCM data for a sine wave.
func generateSineWave(freq float64, dur float64) []byte {
	length := int(float64(sampleRate) * dur)
	buf := bytes.NewBuffer(make([]byte, 0, length*2))
	for i := 0; i < length; i++ {
		v := math.Sin(2 * math.Pi * freq * float64(i) / float64(sampleRate))
		sample := int16(v * 0.3 * math.MaxInt16)
		buf.WriteByte(byte(sample))
		buf.WriteByte(byte(sample >> 8))
	}
	return buf.Bytes()
}
