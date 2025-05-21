import Phaser from 'phaser';
import { levelManager } from './managers/levelManager';
import GameScene from './scenes/GameScene';
import LevelMenuScene from './scenes/LevelMenuScene';
import MainMenuScene from './scenes/MainMenuScene';
import MenuScene from './scenes/MenuScene';

levelManager.loadProgress();

const config: Phaser.Types.Core.GameConfig = {
    type: Phaser.AUTO,
    width: 800,
    height: 600,
    backgroundColor: '#222',
    scene: [MainMenuScene, MenuScene, LevelMenuScene, GameScene],
    parent: 'game-container',
};

window.addEventListener('DOMContentLoaded', () => {
    const game = new Phaser.Game(config);
    game.scene.start('MainMenuScene');
});
