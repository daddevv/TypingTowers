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
        const keydownListener = (event: KeyboardEvent) => {
            this.currentInput += event.key;
            // You can add logic here to process input, check for word matches, etc.
        };
        this.scene.input.keyboard.on('keydown', keydownListener);

        // Unregister the listener when the scene is destroyed or shut down
        this.scene.events.once('shutdown', () => {
            this.scene.input.keyboard.off('keydown', keydownListener);
        });
        this.scene.events.once('destroy', () => {
            this.scene.input.keyboard.off('keydown', keydownListener);
        });
    }

    getInput(): string {
        return this.currentInput;
    }

    clearInput() {
        this.currentInput = '';
    }
}
