// HeadlessGameEngine.ts
// Headless, UI-agnostic core game logic for TypeDefense
// Provides a programmatic API for stepping the game, injecting input, and retrieving state

import { MobSystem } from '../entities/Mob';
import MobSpawner, { MobSpawnerConfig } from '../entities/MobSpawner';
import { IRenderAdapter } from '../render/IRenderAdapter';
import { GameState } from '../state/gameState';
import stateManager from '../state/stateManager';
import WordGenerator from '../utils/wordGenerator';

// --- Engine Options ---
export interface HeadlessGameEngineOptions {
  wordGenerator?: WordGenerator;
  availableKeys?: string[];
  wordList?: string[]; // If provided, MobSpawner will use this list instead of the WordGenerator
  spawnInterval?: number;
  mobsPerInterval?: number;
  mobBaseSpeed?: number;
  mobsPerWave?: number;
  winThreshold?: number; // Number of mobs to defeat
  scalingDuration?: number;
  // Options for MobSpawnerConfig can be added if needed, e.g. spawner width/height
}

// Default options for the engine
const defaultEngineOptions: Required<Omit<HeadlessGameEngineOptions, 'wordGenerator' | 'wordList' | 'mobsPerWave' | 'availableKeys'>> = {
  spawnInterval: 2000,
  mobsPerInterval: 1,
  mobBaseSpeed: 90,
  winThreshold: 50, // Default win condition if not specified by level
  scalingDuration: 120000, // 2 minutes
};


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
  private mobSpawner!: MobSpawner; // Definite assignment in configure
  private wordGenerator!: WordGenerator; // Definite assignment in configure
  private winThreshold!: number; // Definite assignment in configure
  private scalingDuration!: number; // Definite assignment in configure
  private elapsedTime = 0;
  private scalingProgression = 0;
  private lastTime = 0;
  private isInitialized = false;
  private defeatedCount = 0; // Track total defeated mobs

  private initialOptions: HeadlessGameEngineOptions;
  private renderAdapter?: IRenderAdapter;

  constructor(options: HeadlessGameEngineOptions = {}, renderAdapter?: IRenderAdapter) {
    this.initialOptions = { ...defaultEngineOptions, ...options };
    this.configure(this.initialOptions);
    this.renderAdapter = renderAdapter;

    // Listen for mobRemoved events to track defeated mobs
    stateManager.on('mobRemoved', (mobId: string) => {
      // Ensure this listener is only active if the engine instance is the one running
      // and the game is in a state where mobs can be defeated (e.g. 'playing')
      if (this.isInitialized && stateManager.getState().gameStatus === 'playing') {
        this.defeatedCount++;
        console.debug(`[HeadlessGameEngine] mobRemoved event: mobId=${mobId}, defeatedCount=${this.defeatedCount}`);
      }
    });

    // Expose stateManager globally for renderer access (for menu button)
    (window as any).stateManager = stateManager;
    console.debug('[HeadlessGameEngine] Initialized and stateManager exposed on window');
  }

  private configure(options: HeadlessGameEngineOptions) {
    console.log('[HeadlessGameEngine] Configuring with options:', options);
    const mergedOptions = { ...defaultEngineOptions, ...options };

    this.winThreshold = mergedOptions.winThreshold;
    this.scalingDuration = mergedOptions.scalingDuration;

    if (options.wordGenerator) {
      this.wordGenerator = options.wordGenerator;
    } else {
      this.wordGenerator = new WordGenerator(options.availableKeys || ['f', 'j'], true);
    }

    // Provide default width/height for headless engine (e.g., 800x600)
    const spawnerConfig: MobSpawnerConfig = { width: 800, height: 600 };
    this.mobSpawner = new MobSpawner(
      spawnerConfig,
      this.wordGenerator,
      mergedOptions.spawnInterval,
      mergedOptions.mobsPerInterval,
      mergedOptions.mobBaseSpeed,
      options.wordList, // Pass through: if undefined, MobSpawner uses WordGenerator
      options.mobsPerWave // Pass through
    );

    this.elapsedTime = 0;
    this.scalingProgression = 0;
    this.lastTime = 0;
    this.defeatedCount = 0;
    this.isInitialized = true;

    // Ensure mob spawner starts wave if game is in 'playing' state
    // This might be called during reset when game status is already 'playing'
    if (stateManager.getState().gameStatus === 'playing') {
      this.mobSpawner.startNextWave();
    }
    console.log('[HeadlessGameEngine] Configuration complete.');
  }

  /**
   * Step the game forward by delta ms.
   * Updates mobs, spawner, win/loss, and state.
   */
  step(delta: number, timestamp?: number) {
    console.log('[HeadlessGameEngine] step() called, delta:', delta, 'timestamp:', timestamp);
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

    // Check for defeated mobs and remove them
    // This ensures mobRemoved event fires and defeatedCount increments
    const currentMobs = stateManager.getState().mobs;
    let mobsRemovedThisStep = false;
    for (const mob of currentMobs) {
      if (mob.isDefeated) {
        // Check if mob still exists before trying to remove, to avoid duplicate removals if systems overlap
        if (stateManager.getState().mobs.find(m => m.id === mob.id)) {
          stateManager.removeMob(mob.id);
          mobsRemovedThisStep = true;
        }
      }
    }

    this.mobSpawner.update(timestamp ?? this.lastTime + delta, delta);
    this.lastTime = timestamp ?? this.lastTime + delta;
    console.debug(`[HeadlessGameEngine] step: delta=${delta}, timestamp=${timestamp}, elapsedTime=${this.elapsedTime}, scalingProgression=${this.scalingProgression}`);

    // Always check win/loss after all updates, regardless of whether mobs were removed
    this._checkWinLoss();
  }

  /**
   * Inject player input (simulate typing a key or string)
   */
  injectInput(input: string) {
    console.log('[HeadlessGameEngine] injectInput() called, input:', input);
    stateManager.updatePlayerInput(input);
  }

  /**
   * Get a deep copy of the current game state
   */
  getState(): GameState {
    const state = stateManager.getState();
    if (!state.player) {
      state.player = { position: { x: 400, y: 500 }, score: 0, combo: 0, health: 3, maxHealth: 3, currentInput: '' };
    }
    // Log state summary for debugging
    console.debug('[HeadlessGameEngine] getState:', {
      gameStatus: state.gameStatus,
      player: state.player,
      mobs: state.mobs?.length,
      level: state.level,
    });
    return state;
  }

  /**
   * Listen for game events (state changes)
   */
  on(event: string, handler: (...args: any[]) => void) {
    console.debug(`[HeadlessGameEngine] Registering event handler for "${event}"`);
    stateManager.on(event, handler);
  }

  off(event: string, handler: (...args: any[]) => void) {
    console.debug(`[HeadlessGameEngine] Removing event handler for "${event}"`);
    stateManager.off(event, handler);
  }

  /**
   * Reset the engine and state
   */
  reset(newOptions?: HeadlessGameEngineOptions) {
    console.debug('[HeadlessGameEngine] Resetting engine. New options:', newOptions);
    const optionsToUse = newOptions || this.initialOptions;

    // Reset global game state first
    stateManager.reset(); 
    // stateManager.reset() sets gameStatus to 'booting' and player health to default.
    // It also clears mobs, player input, etc.
    // Crucially, this.defeatedCount needs to be reset *before* configure might set a new winThreshold
    // and *after* stateManager.reset() so we don't lose its value if reset is called mid-game.
    // configure() will reset defeatedCount.

    // Apply new engine configuration
    this.configure(optionsToUse); // This will reset this.defeatedCount = 0;

    // Set the game status to playing *after* configuration and state reset
    // The currentLevelId and currentWorld should be set by the calling context (e.g., test or scene manager)
    // before reset, or updated after reset if stateManager.reset() clears them.
    // stateManager.reset() sets level to default, so we need to update it if newOptions imply a specific level.
    // However, the test/caller is responsible for setting currentLevelId in stateManager.
    // For now, assume stateManager has the correct currentLevelId after its reset or it's set by caller.
    const currentLevelState = stateManager.getState().level;
    stateManager.updateCurrentLevelContext({
      currentWorld: currentLevelState.currentWorld, // Preserve from default or caller
      currentLevelId: currentLevelState.currentLevelId, // Preserve from default or caller
      levelStatus: 'playing'
    });
    stateManager.setGameStatus('playing');
    // Player health is reset by stateManager.reset() to defaultGameState.player.health

    // Explicitly start the first wave after setting gameStatus to 'playing'.
    // The configure() method also attempts to start a wave if gameStatus is 'playing',
    // but at the time configure() is called within this reset() flow, gameStatus is 'booting'.
    if (this.mobSpawner) {
      this.mobSpawner.startNextWave();
    }
    // If configure was called with options that imply a new level,
    // mobSpawner should be ready for that level.
    console.debug('[HeadlessGameEngine] Engine reset complete.');
  }

  /**
   * Internal: check win/loss conditions and update state
   */
  private _checkWinLoss() {
    const state = stateManager.getState();
    // Use this.defeatedCount which tracks total mobs removed (defeated) via 'mobRemoved' event
    if (this.defeatedCount >= this.winThreshold && state.level.levelStatus !== 'complete' && state.level.levelStatus !== 'failed') {
      console.debug('[HeadlessGameEngine] Win condition met. Defeated (total via mobRemoved):', this.defeatedCount, 'Threshold:', this.winThreshold);
      stateManager.updateCurrentLevelContext({ levelStatus: 'complete' });
      stateManager.setGameStatus('levelComplete');
      // Immediately update local state to reflect changes for this tick
      const updatedState = stateManager.getState();
      updatedState.level.levelStatus = 'complete';
      updatedState.gameStatus = 'levelComplete';
    }
    if (state.player.health <= 0 && state.level.levelStatus !== 'failed' && state.level.levelStatus !== 'complete') {
      console.debug('[HeadlessGameEngine] Loss condition met. Player health:', state.player.health);
      stateManager.updateCurrentLevelContext({ levelStatus: 'failed' });
      stateManager.setGameStatus('gameOver');
      // Immediately update local state to reflect changes for this tick
      const updatedState = stateManager.getState();
      updatedState.level.levelStatus = 'failed';
      updatedState.gameStatus = 'gameOver';
    }
  }

  /**
   * Optionally, call the render adapter to render the current state.
   * This should be called by the game loop or test harness, not automatically in step().
   */
  render() {
    if (this.renderAdapter) {
      this.renderAdapter.render(this.getState());
    }
  }

  /**
   * Optionally, initialize the render adapter with width/height.
   * Should be called by the host if using a render adapter.
   */
  initRenderer(width: number, height: number) {
    if (this.renderAdapter) {
      this.renderAdapter.init(width, height);
    }
  }

  /**
   * Optionally, destroy the render adapter.
   */
  destroyRenderer() {
    if (this.renderAdapter) {
      this.renderAdapter.destroy();
    }
  }
}

export default HeadlessGameEngine;
