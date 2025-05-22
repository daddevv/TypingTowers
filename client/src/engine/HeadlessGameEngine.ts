// HeadlessGameEngine.ts
// Headless, UI-agnostic core game logic for TypeDefense
// Provides a programmatic API for stepping the game, injecting input, and retrieving state

import stateManager from '../state/stateManager';
import { MobSystem } from '../entities/Mob';
import MobSpawner, { MobSpawnerConfig } from '../entities/MobSpawner';
import WordGenerator from '../utils/wordGenerator';
import { GameState } from '../state/gameState';

// --- Engine contract/interface for renderer integration ---
/**
 * IGameEngine defines the contract for any TypeDefense game engine implementation.
 * Renderers (Phaser, Three.js, etc.) should depend on this interface.
 */
export interface IGameEngine {
  /**
   * Step the game forward by delta ms.
   * @param delta Time in ms to advance the simulation.
   * @param timestamp Optional absolute timestamp.
   */
  step(delta: number, timestamp?: number): void;

  /**
   * Inject player input (simulate typing a key or string).
   * @param input The input string or key.
   */
  injectInput(input: string): void;

  /**
   * Get a deep copy of the current game state.
   */
  getState(): GameState;

  /**
   * Listen for game events (state changes).
   * @param event Event name (e.g., 'gameStatusChanged', 'playerInputChanged', etc.)
   * @param handler Callback function.
   */
  on(event: string, handler: (...args: any[]) => void): void;

  /**
   * Remove a previously registered event handler.
   */
  off(event: string, handler: (...args: any[]) => void): void;

  /**
   * Reset the engine and state to initial values.
   */
  reset(): void;
}

/**
 * HeadlessGameEngine: UI-agnostic, testable core game logic for TypeDefense.
 * - Step the game forward (advance time, update mobs, win/loss, etc)
 * - Inject input (simulate typing)
 * - Retrieve current game state
 * - Listen for game events
 */
export class HeadlessGameEngine implements IGameEngine {
  private mobSpawner: MobSpawner;
  private wordGenerator: WordGenerator;
  private winThreshold: number;
  private scalingDuration: number;
  private elapsedTime = 0;
  private scalingProgression = 0;
  private lastTime = 0;
  private isInitialized = false;
  private defeatedCount = 0; // Track total defeated mobs

  constructor(options: HeadlessGameEngineOptions = {}) {
    this.wordGenerator = options.wordGenerator || new WordGenerator(['f', 'j'], true);
    // Provide default width/height for headless engine (e.g., 800x600)
    const spawnerConfig: MobSpawnerConfig = { width: 800, height: 600 };
    // Pass wordList as-is (even if empty) so MobSpawner disables spawning if []
    this.mobSpawner = new MobSpawner(
      spawnerConfig,
      this.wordGenerator,
      options.spawnInterval || 2000,
      options.mobsPerInterval || 1,
      options.mobBaseSpeed || 90,
      options.wordList,
      options.mobsPerWave // pass through
    );
    this.winThreshold = options.winThreshold || 50;
    this.scalingDuration = options.scalingDuration || 120000;
    this.isInitialized = true;
    this.mobSpawner.startNextWave(); // Start the first wave automatically

    this.defeatedCount = 0;

    // Listen for mobRemoved events to track defeated mobs
    stateManager.on('mobRemoved', (mobId: string) => {
      // Only increment if the mob was defeated (best effort: check last state)
      // This is a limitation, but for headless engine, we assume all removals are due to defeat
      this.defeatedCount++;
    });
  }

  /**
   * Step the game forward by delta ms.
   * Updates mobs, spawner, win/loss, and state.
   */
  step(delta: number, timestamp?: number) {
    if (!this.isInitialized) throw new Error('Engine not initialized');
    this.elapsedTime += delta;
    this.scalingProgression = Math.min(this.elapsedTime / this.scalingDuration, 1);
    this.mobSpawner.setProgression(this.scalingProgression);
    if (this.mobSpawner['wordGenerator'] && typeof this.mobSpawner['wordGenerator'].setWordLengthScaling === 'function') {
      this.mobSpawner['wordGenerator'].setWordLengthScaling(
        2 + Math.floor(2 * this.scalingProgression),
        5 + Math.floor(2 * this.scalingProgression)
      );
    }
    stateManager.updateTimestampAndDelta(timestamp ?? this.lastTime + delta, delta);
    MobSystem.updateAll(timestamp ?? this.lastTime + delta, delta);
    this.mobSpawner.update(timestamp ?? this.lastTime + delta, delta);
    this.lastTime = timestamp ?? this.lastTime + delta;
    this._checkWinLoss();
  }

  /**
   * Inject player input (simulate typing a key or string)
   */
  injectInput(input: string) {
    stateManager.updatePlayerInput(input);
  }

  /**
   * Get a deep copy of the current game state
   */
  getState(): GameState {
    return stateManager.getState();
  }

  /**
   * Listen for game events (state changes)
   */
  on(event: string, handler: (...args: any[]) => void) {
    stateManager.on(event, handler);
  }

  off(event: string, handler: (...args: any[]) => void) {
    stateManager.off(event, handler);
  }

  /**
   * Reset the engine and state
   */
  reset() {
    stateManager.reset();
    // Always set level status to 'playing' and gameStatus to 'playing' on reset
    stateManager.updateCurrentLevelContext({ levelStatus: 'playing' });
    stateManager.setGameStatus('playing');
    // Also ensure player health is reset
    stateManager.updatePlayerHealth(3);
    this.elapsedTime = 0;
    this.scalingProgression = 0;
    this.lastTime = 0;
    this.isInitialized = true;
    this.mobSpawner.startNextWave(); // Start the first wave on reset
    this.defeatedCount = 0;
  }

  /**
   * Internal: check win/loss conditions and update state
   */
  private _checkWinLoss() {
    // Use both defeatedCount and count of defeated mobs in state
    const state = stateManager.getState();
    const defeatedInState = state.mobs.filter(m => m.isDefeated).length;
    const defeated = Math.max(this.defeatedCount, defeatedInState);
    if (defeated >= this.winThreshold && state.level.levelStatus !== 'complete') {
      stateManager.updateCurrentLevelContext({ levelStatus: 'complete' });
      stateManager.setGameStatus('levelComplete');
    }
    if (state.player.health <= 0 && state.level.levelStatus !== 'failed') {
      stateManager.updateCurrentLevelContext({ levelStatus: 'failed' });
      stateManager.setGameStatus('gameOver');
    }
  }
}

export default HeadlessGameEngine;
