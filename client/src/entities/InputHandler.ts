// InputHandler.ts
// Handles keyboard input for the game.

export default class InputHandler {
    private currentInput: string = '';
    private scene: any;
    private keydownCallback: any;

    constructor(scene: any) {
        this.scene = scene;
        if (scene.input && scene.input.keyboard && typeof scene.input.keyboard.on === 'function') {
            this.keydownCallback = (event: any) => {
                if (event && event.key) {
                    this.currentInput += event.key;
                }
            };
            scene.input.keyboard.on('keydown', this.keydownCallback);
            // Remove listener on shutdown/destroy
            if (scene.events && typeof scene.events.once === 'function') {
                scene.events.once('shutdown', () => {
                    scene.input.keyboard.off('keydown', this.keydownCallback);
                });
                scene.events.once('destroy', () => {
                    scene.input.keyboard.off('keydown', this.keydownCallback);
                });
            }
        }
    }

    public getInput(): string {
        return this.currentInput;
    }

    public clearInput(): void {
        this.currentInput = '';
    }
}
