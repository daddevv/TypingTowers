// Player.ts
// Represents the player character as a pure data/model class (no Phaser dependency)
import stateManager from '../state/stateManager';

export default class Player {
    // Player state fields (mirrors GameState.player)
    health: number;
    maxHealth: number;
    score: number;
    combo: number;
    currentInput: string;
    position: { x: number; y: number };

    constructor() {
        const playerState = stateManager.getState().player;
        this.health = playerState.health;
        this.maxHealth = playerState.maxHealth;
        this.score = playerState.score;
        this.combo = playerState.combo;
        this.currentInput = playerState.currentInput;
        this.position = { ...playerState.position };
    }

    syncFromState() {
        const playerState = stateManager.getState().player;
        this.health = playerState.health;
        this.maxHealth = playerState.maxHealth;
        this.score = playerState.score;
        this.combo = playerState.combo;
        this.currentInput = playerState.currentInput;
        this.position = { ...playerState.position };
    }

    // Optionally, add methods for logic (not rendering)
    // Rendering is now handled in GameScene based on stateManager
    // No rendering logic here; handled by RenderManager.
}
