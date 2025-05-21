import Phaser from 'phaser';
import { getKeyInfo } from '../curriculum/fingerGroups';
import InputHandler from '../entities/InputHandler';
import MobSpawner from '../entities/MobSpawner';
import Player from '../entities/Player';
import FingerGroupManager from '../managers/fingerGroupManager';

export default class GameScene extends Phaser.Scene {
    private player!: Player;
    private inputHandler!: InputHandler;
    private mobSpawner!: MobSpawner;
    private fingerGroupManager!: FingerGroupManager;

    constructor() {
        super('GameScene');
    }

    preload() {
        // Load assets here
        // Example: this.load.image('player', 'assets/images/player.png');
        // Example: this.load.image('mob', 'assets/images/mob.png');
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
        this.fingerGroupManager = new FingerGroupManager();

        // Initialize MobSpawner with a sample word list
        const words = ['type', 'defense', 'phaser', 'enemy', 'challenge'];
        this.mobSpawner = new MobSpawner(this, words, 2000);
    }

    update(time: number, delta: number) {
        // Core game loop logic
        if (this.player) {
            this.player.update(time, delta);
        }
        // Update MobSpawner and its mobs
        if (this.mobSpawner) {
            this.mobSpawner.update(time, delta);
        }
        // Check player input against mobs
        const input = this.inputHandler.getInput();
        if (input.length > 0) {
            // Record each key press in FingerGroupManager
            for (const char of input) {
                const keyInfo = getKeyInfo(char);
                if (keyInfo) {
                    this.fingerGroupManager.recordKeyPress(char, true, time); // 'true' for correct finger (future: detect real finger)
                }
            }
            const mobs = this.mobSpawner.getMobs();
            for (const mob of mobs) {
                if (!mob.isDefeated && input.trim().toLowerCase() === mob.word.toLowerCase()) {
                    mob.defeat();
                    this.inputHandler.clearInput();
                    // TODO: Add visual/audio feedback here
                    break;
                }
            }
        }
    }
}
// Contains AI-generated edits.
