// InputHandler.ts
// Handles keyboard input for the game.
import Phaser from 'phaser';

export default class InputHandler {
    private scene: Phaser.Scene;

    constructor(scene: Phaser.Scene) {
        this.scene = scene;
        // InputHandler is deprecated. Input is now handled by InputSystem.
    }
}
