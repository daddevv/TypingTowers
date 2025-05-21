// InputHandler.ts
// Handles keyboard input for the game.
import Phaser from 'phaser';

export default class InputHandler {
    private scene: Phaser.Scene;
    private currentInput: string = '';

    constructor(scene: Phaser.Scene) {
        this.scene = scene;
        this.registerEvents();
    }

    private registerEvents() {
        if (!this.scene.input.keyboard) return;
        this.scene.input.keyboard.on('keydown', (event: KeyboardEvent) => {
            this.currentInput += event.key;
            // You can add logic here to process input, check for word matches, etc.
        });
    }

    getInput(): string {
        return this.currentInput;
    }

    clearInput() {
        this.currentInput = '';
    }
}
