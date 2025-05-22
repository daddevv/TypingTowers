// HeadlessGameEngine.unit.test.ts
// Unit tests for HeadlessGameEngine (headless core game logic)
import { describe, it, expect, beforeEach, vi } from 'vitest';
import HeadlessGameEngine from '../engine/HeadlessGameEngine';
import stateManager from '../state/stateManager';

function fastForward(engine: HeadlessGameEngine, ms: number, step: number = 16) {
  let elapsed = 0;
  while (elapsed < ms) {
    engine.step(step);
    elapsed += step;
  }
}

describe('HeadlessGameEngine', () => {
  let engine: HeadlessGameEngine;

  beforeEach(() => {
    engine = new HeadlessGameEngine({ winThreshold: 3 });
    engine.reset();
  });

  it('should initialize and expose state', () => {
    const state = engine.getState();
    expect(state.player).toBeDefined();
    expect(state.mobs).toBeInstanceOf(Array);
  });

  it('should step the game and spawn mobs', () => {
    const initialMobs = engine.getState().mobs.length;
    fastForward(engine, 3000);
    const afterMobs = engine.getState().mobs.length;
    expect(afterMobs).toBeGreaterThan(initialMobs);
  });

  it('should inject input and update player input in state', () => {
    engine.injectInput('f');
    expect(engine.getState().player.currentInput).toBe('f');
  });

  it('should emit events on state changes', () => {
    const handler = vi.fn();
    engine.on('playerInputChanged', handler);
    engine.injectInput('j');
    expect(handler).toHaveBeenCalledWith('j');
    engine.off('playerInputChanged', handler);
  });

  it('should trigger win condition when enough mobs are defeated', () => {
    // Simulate defeating mobs
    const state = engine.getState();
    // Add 3 defeated mobs
    stateManager.updateMobs([
      { id: '1', word: 'fj', currentTypedIndex: 2, position: { x: 0, y: 0 }, speed: 100, type: 'normal', isDefeated: true },
      { id: '2', word: 'fj', currentTypedIndex: 2, position: { x: 0, y: 0 }, speed: 100, type: 'normal', isDefeated: true },
      { id: '3', word: 'fj', currentTypedIndex: 2, position: { x: 0, y: 0 }, speed: 100, type: 'normal', isDefeated: true },
    ]);
    engine.step(16);
    expect(engine.getState().level.levelStatus).toBe('complete');
    expect(engine.getState().gameStatus).toBe('levelComplete');
  });

  it('should trigger loss condition when player health is zero', () => {
    stateManager.updatePlayerHealth(0);
    engine.step(16);
    expect(engine.getState().level.levelStatus).toBe('failed');
    expect(engine.getState().gameStatus).toBe('gameOver');
  });

  it('should reset state and engine', () => {
    engine.injectInput('f');
    engine.reset();
    expect(engine.getState().player.currentInput).toBe('');
  });
});

describe('HeadlessGameEngine (comprehensive scenarios)', () => {
  let engine: HeadlessGameEngine;

  beforeEach(() => {
    engine = new HeadlessGameEngine({ winThreshold: 5 });
    engine.reset();
  });

  it('simulates a full game from start to win', () => {
    // Fast-forward until mobs spawn
    fastForward(engine, 5000);
    let state = engine.getState();
    // Defeat mobs as they appear
    let defeated = 0;
    while (state.level.levelStatus !== 'complete' && defeated < 20) {
      // Mark all mobs as defeated
      const mobs = state.mobs.map(m => ({ ...m, isDefeated: true, currentTypedIndex: m.word.length }));
      stateManager.updateMobs(mobs);
      engine.step(16);
      state = engine.getState();
      defeated++;
    }
    expect(state.level.levelStatus).toBe('complete');
    expect(state.gameStatus).toBe('levelComplete');
  });

  it('simulates a bot that types correct letters to defeat mobs', () => {
    engine.reset();
    fastForward(engine, 2000);
    let state = engine.getState();
    let steps = 0;
    while (state.level.levelStatus !== 'complete' && steps < 1000) {
      // For each mob, type the next required letter
      for (const mob of state.mobs) {
        if (!mob.isDefeated && mob.currentTypedIndex < mob.word.length) {
          const nextLetter = mob.word[mob.currentTypedIndex];
          engine.injectInput(nextLetter);
        }
      }
      engine.step(16);
      state = engine.getState();
      steps++;
    }
    expect(state.level.levelStatus).toBe('complete');
    expect(state.gameStatus).toBe('levelComplete');
  });

  it('handles empty word list gracefully', () => {
    const emptyEngine = new HeadlessGameEngine({ wordList: [], winThreshold: 1 });
    emptyEngine.reset();
    fastForward(emptyEngine, 2000);
    const state = emptyEngine.getState();
    // Should not crash, and no mobs should spawn
    expect(state.mobs.length).toBe(0);
    // Should never reach win condition
    expect(state.level.levelStatus).not.toBe('complete');
  });

  it('handles extremely high spawn rate', () => {
    const fastEngine = new HeadlessGameEngine({ spawnInterval: 10, winThreshold: 10, mobsPerWave: 20 });
    fastEngine.reset();
    fastForward(fastEngine, 500);
    const state = fastEngine.getState();
    expect(state.mobs.length).toBeGreaterThanOrEqual(10);
  });

  it('handles extremely low spawn rate', () => {
    const slowEngine = new HeadlessGameEngine({ spawnInterval: 10000, winThreshold: 1 });
    slowEngine.reset();
    fastForward(slowEngine, 1000);
    const state = slowEngine.getState();
    // Should have at most 1 mob
    expect(state.mobs.length).toBeLessThanOrEqual(1);
  });

  it('handles player health edge case (negative health)', () => {
    engine.reset();
    stateManager.updatePlayerHealth(-5);
    engine.step(16);
    const state = engine.getState();
    expect(state.level.levelStatus).toBe('failed');
    expect(state.gameStatus).toBe('gameOver');
  });

  it('handles rapid input (spam)', () => {
    engine.reset();
    fastForward(engine, 2000);
    let state = engine.getState();
    for (let i = 0; i < 100; i++) {
      engine.injectInput('f');
      engine.injectInput('j');
      engine.step(1);
    }
    // Should not crash, and state should be valid
    state = engine.getState();
    expect(state.player.currentInput).toBe('j');
  });
});
