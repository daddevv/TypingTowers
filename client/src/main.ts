import Phaser from 'phaser';
import { WORLDS } from './curriculum/worldConfig';
import LevelManager from './managers/levelManager';
import GameScene from './scenes/GameScene';
import MenuScene from './scenes/MenuScene';

const levelManager = new LevelManager();
levelManager.loadProgress();

const config: Phaser.Types.Core.GameConfig = {
    type: Phaser.AUTO,
    width: 800,
    height: 600,
    backgroundColor: '#222',
    scene: [MenuScene, GameScene],
    parent: 'game-container',
};

window.addEventListener('DOMContentLoaded', () => {
    const game = new Phaser.Game(config);
    game.scene.start('MenuScene', { worlds: WORLDS, levelManager });
});
// Contains AI-generated edits.
