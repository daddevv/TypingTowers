// HeadlessGameEngine.unit.test.ts
// Unit tests for HeadlessGameEngine (headless core game logic)
import { beforeEach, describe, expect, it, vi } from 'vitest';
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
    // Initialize with a default winThreshold, can be overridden in specific tests via engine.reset(options)
    engine = new HeadlessGameEngine({ winThreshold: 3 }); 
    engine.reset(); // Resets with the initial options, including winThreshold: 3
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
    // engine.reset() was called in beforeEach with winThreshold: 3
    // Simulate mobs being defeated and removed, which increments engine.defeatedCount
    // This requires mobRemoved events to be emitted.
    // The engine listens to stateManager's 'mobRemoved' event.

    // Directly trigger 'mobRemoved' enough times to meet the winThreshold.
    // This simulates the effect of mobs being defeated and processed by MobSpawner.
    for (let i = 0; i < 3; i++) {
      // To ensure stateManager.removeMob doesn't fail if it expects the mob to exist,
      // we can add a placeholder mob first, then remove it.
      // However, for this test, we only care that the event fires and engine's count increments.
      // If removeMob is robust to non-existent IDs for eventing, this is simpler.
      // Let's assume stateManager.removeMob will emit 'mobRemoved' even if mobId wasn't in its list.
      // Or, more robustly, add and remove.
      const mobIdToRemove = `test_mob_for_removal_${i}`;
      stateManager.addMob({ id: mobIdToRemove, word: 'test', currentTypedIndex: 0, position: { x: 0, y: 0 }, speed: 0, type: 'normal', isDefeated: false });
      stateManager.removeMob(mobIdToRemove); // This should increment engine.defeatedCount
    }

    engine.step(16); // _checkWinLoss will use engine.defeatedCount

    const finalState = engine.getState();
    expect(finalState.level.levelStatus).toBe('complete');
    expect(finalState.gameStatus).toBe('levelComplete');
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
    // Default engine for these scenarios, winThreshold can be overridden in engine.reset() per test
    engine = new HeadlessGameEngine();
    // engine.reset() will be called in each test with specific options or defaults.
  });

  it('simulates a full game from start to win', () => {
    const testWinThreshold = 5;
    engine.reset({ winThreshold: testWinThreshold, spawnInterval: 200, mobsPerInterval: 1 }); // Faster spawning

    let state = engine.getState();
    let simulationSteps = 0;
    const maxSimulationSteps = testWinThreshold * 10 + 100; // Adjusted max steps

    // Loop until level is complete or max steps reached
    while (state.level.levelStatus !== 'complete' && simulationSteps < maxSimulationSteps) {
      let typedSomethingThisLoop = false;
      if (state.mobs.length > 0) {
        const mobToDefeat = state.mobs.find(m => !m.isDefeated && m.currentTypedIndex < m.word.length);
        if (mobToDefeat) {
          for (let i = mobToDefeat.currentTypedIndex; i < mobToDefeat.word.length; i++) {
            const char = mobToDefeat.word[i];
            engine.injectInput(char);
            engine.step(50); // Process input
          }
          typedSomethingThisLoop = true;
        }
      }
      engine.step(100); // Main step for game progression (spawning, mob removal, win/loss check)

      if (!typedSomethingThisLoop && state.mobs.length === 0 && state.level.levelStatus !== 'complete') {
        // If no mobs and not complete, step more to encourage spawning
        engine.step(200);
      }
      state = engine.getState();
      simulationSteps++;
    }
    if (state.level.levelStatus !== 'complete') {
      console.warn(`[Unit Test] 'simulates a full game' timed out. Status: ${state.level.levelStatus}, Defeated: ${(engine as any).defeatedCount}, Threshold: ${(engine as any).winThreshold}`);
    }
    // Give the engine a few more steps to process win condition if not yet complete
    let retries = 0;
    while (state.level.levelStatus !== 'complete' && retries < 5) {
      engine.step(16);
      state = engine.getState();
      retries++;
    }
    expect(state.level.levelStatus).toBe('complete');
    expect(state.gameStatus).toBe('levelComplete');
    expect(simulationSteps).toBeLessThan(maxSimulationSteps); // Ensure it didn't timeout
  });

  it('simulates a bot that types correct letters to defeat mobs', () => {
    const testWinThreshold = 5;
    engine.reset({ winThreshold: testWinThreshold, spawnInterval: 200, mobsPerInterval: 1 }); // Faster spawning

    let state = engine.getState();
    let steps = 0;
    const maxBotSteps = testWinThreshold * 10 + 100; // Adjusted max steps

    while (state.level.levelStatus !== 'complete' && steps < maxBotSteps) {
      let actionTakenThisStep = false;
      if (state.mobs.length > 0) {
        const mobToDefeat = state.mobs.find(m => !m.isDefeated && m.currentTypedIndex < m.word.length);
        if (mobToDefeat) {
          const nextLetter = mobToDefeat.word[mobToDefeat.currentTypedIndex];
          engine.injectInput(nextLetter);
          actionTakenThisStep = true;
          engine.step(50); // Process this specific input
        }
      }
      engine.step(100); // Main step for game progression

      if (!actionTakenThisStep && state.mobs.length === 0 && state.level.levelStatus !== 'complete') {
        engine.step(200); // Encourage spawning
      }
      state = engine.getState();
      steps++;
    }
    if (state.level.levelStatus !== 'complete') {
      console.warn(`[Unit Test] 'simulates a bot' timed out. Status: ${state.level.levelStatus}, Defeated: ${(engine as any).defeatedCount}, Threshold: ${(engine as any).winThreshold}`);
    }
    // Give the engine a few more steps to process win condition if not yet complete
    let retries = 0;
    while (state.level.levelStatus !== 'complete' && retries < 5) {
      engine.step(16);
      state = engine.getState();
      retries++;
    }
    expect(state.level.levelStatus).toBe('complete');
    expect(state.gameStatus).toBe('levelComplete');
    expect(steps).toBeLessThan(maxBotSteps);
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
    // Track how many mobs are spawned
    let mobsSpawned = 0;
    const origAddMob = stateManager.addMob;
    stateManager.addMob = function (...args: any[]) {
      mobsSpawned++;
      // @ts-ignore
      return origAddMob.apply(this, args);
    };
    fastForward(fastEngine, 500);
    // Restore original addMob to avoid side effects
    stateManager.addMob = origAddMob;
    expect(mobsSpawned).toBeGreaterThanOrEqual(10);
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
