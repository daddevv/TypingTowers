import Phaser from 'phaser';
import { levelManager } from './managers/levelManager';
import BootScene from './scenes/BootScene';
import GameScene from './scenes/GameScene';
import LevelMenuScene from './scenes/LevelMenuScene';
import MainMenuScene from './scenes/MainMenuScene';
import MenuScene from './scenes/MenuScene';

levelManager.loadProgress();

const config: Phaser.Types.Core.GameConfig = {
    type: Phaser.AUTO,
    width: window.innerWidth,
    height: window.innerHeight,
    backgroundColor: '#222',
    scene: [BootScene, MainMenuScene, MenuScene, LevelMenuScene, GameScene],
    parent: 'game-container',
    scale: {
        mode: Phaser.Scale.RESIZE,
        autoCenter: Phaser.Scale.CENTER_BOTH,
    },
};

window.addEventListener('DOMContentLoaded', () => {
    const game = new Phaser.Game(config);
    game.scene.start('BootScene');
});
