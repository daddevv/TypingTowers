import Phaser from 'phaser';
import { levelManager } from './managers/levelManager';
import BootScene from './scenes/BootScene';
import GameScene from './scenes/GameScene';
import LevelMenuScene from './scenes/LevelMenuScene';
import MainMenuScene from './scenes/MainMenuScene';
import WorldSelectionScene from './scenes/WorldSelectionScene';
import stateManager from './state/stateManager';

// --- Render backend selection ---
import { PhaserRenderManager } from './render/PhaserRenderManager';
import { ThreeJsRenderManager } from './render/ThreeJsRenderManager';

// Type declarations for window globals (optional, for type safety)
declare global {
    interface Window {
        RENDER_BACKEND?: string;
        renderManager?: any;
    }
}

// Determine render backend: 'phaser' (default) or 'three'
const RENDER_BACKEND =
    window.RENDER_BACKEND ||
    // Type assertion to allow access to Vite's env
    ((import.meta as ImportMeta & { env: Record<string, any> }).env?.VITE_RENDER_BACKEND) ||
    'phaser';

console.log('[main.ts] RENDER_BACKEND:', RENDER_BACKEND);

if (RENDER_BACKEND === 'three') {
    window.renderManager = new ThreeJsRenderManager();
    console.log('[main.ts] Using Three.js RenderManager');
} else {
    window.renderManager = new PhaserRenderManager();
    console.log('[main.ts] Using Phaser RenderManager');
}

// Expose stateManager globally for renderer access (for menu button)
(window as any).stateManager = stateManager;

levelManager.loadProgress();

const config: Phaser.Types.Core.GameConfig = {
    type: Phaser.AUTO,
    width: window.innerWidth,
    height: window.innerHeight,
    backgroundColor: '#222',
    scene: [BootScene, MainMenuScene, WorldSelectionScene, LevelMenuScene, GameScene],
    parent: 'game-container',
    scale: {
        mode: Phaser.Scale.RESIZE,
        autoCenter: Phaser.Scale.CENTER_BOTH,
    },
};

window.addEventListener('DOMContentLoaded', () => {
    console.log('[main.ts] DOMContentLoaded, creating Phaser.Game');
    // Ensure the container exists
    let container = document.getElementById('game-container');
    if (!container) {
        container = document.createElement('div');
        container.id = 'game-container';
        document.body.appendChild(container);
        console.log('[main.ts] Created #game-container');
    }
    const game = new Phaser.Game(config);
    console.log('[main.ts] Phaser.Game created, starting BootScene');
    game.scene.start('BootScene');
});
