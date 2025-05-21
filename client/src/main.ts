import Phaser from 'phaser';
import GameScene from './scenes/GameScene';

const config: Phaser.Types.Core.GameConfig = {
    type: Phaser.AUTO,
    width: 800,
    height: 600,
    backgroundColor: '#222',
    scene: [GameScene],
    parent: 'game-container',
};

window.addEventListener('DOMContentLoaded', () => {
    new Phaser.Game(config);
});
// Contains AI-generated edits.
