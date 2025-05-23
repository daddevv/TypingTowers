import Phaser from 'phaser';
import stateManager from '../state/stateManager';
import { IRenderManager } from '../render/RenderManager';

export default class MainMenuScene extends Phaser.Scene {
    private onGameStatusChanged?: (status: string) => void;
    private renderManager!: IRenderManager;

    constructor() {
        super('MainMenuScene');
    }

    create(data?: { renderManager?: IRenderManager }) {
        console.log('[MainMenuScene] create() called');
        // Remove all children to clear any lingering error messages or UI
        this.children.removeAll();
        const { width, height } = this.scale;

        // Inject or create the renderManager
        this.renderManager = data?.renderManager || (window as any).renderManager;
        if (!this.renderManager) {
            throw new Error('RenderManager instance must be provided to MainMenuScene');
        }
        console.log('[MainMenuScene] Initializing RenderManager');
        this.renderManager.init(this.game.canvas.parentElement as HTMLElement);
        this.renderManager.render(stateManager.getState());

        // Listen for gameStatus changes and transition if needed
        this.onGameStatusChanged = (status: string) => {
            console.log('[MainMenuScene] gameStatusChanged:', status);
            if (status !== 'mainMenu') {
                this.scene.stop();
            }
        };
        stateManager.on('gameStatusChanged', this.onGameStatusChanged);

        // Clean up event listener when scene is destroyed
        this.events.once('shutdown', () => {
            if (this.onGameStatusChanged) {
                stateManager.off('gameStatusChanged', this.onGameStatusChanged);
            }
        });
    }
}
