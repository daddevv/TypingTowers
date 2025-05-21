import Phaser from 'phaser';
import InputHandler from '../entities/InputHandler';
import Player from '../entities/Player';

export default class GameScene extends Phaser.Scene {
    private player!: Player;
    private inputHandler!: InputHandler;

    constructor() {
        super('GameScene');
    }

    preload() {
        // Load assets here
        // Example: this.load.image('player', 'assets/images/player.png');
    }

    create() {
        // Set up the main game scene here
        this.add.text(400, 300, 'Type Defense', {
            fontSize: '48px',
            color: '#fff',
        }).setOrigin(0.5);

        // Initialize Player and InputHandler
        this.player = new Player(this, 100, 300);
        this.inputHandler = new InputHandler(this);
    }

    update(time: number, delta: number) {
        // Core game loop logic
        if (this.player) {
            this.player.update(time, delta);
        }
        // InputHandler could process input here if needed
        // Placeholder: enemy spawning and word challenge logic will go here
    }
}
// Contains AI-generated edits.
