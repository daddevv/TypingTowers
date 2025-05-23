import Phaser from 'phaser';
import { GameState } from '../state/gameState';
import { IRenderManager } from './RenderManager';

/**
 * PhaserRenderManager
 * Implements IRenderManager using Phaser for rendering.
 * All game logic/state is passed in via render(state).
 *
 * This implementation expects to be given a Phaser.Scene via setScene().
 * It does NOT create its own Phaser.Game or scene.
 */
export class PhaserRenderManager implements IRenderManager {
    private scene?: Phaser.Scene;
    private container?: HTMLElement;
    private width: number = 800;
    private height: number = 600;
    private initialized = false;

    // Internal references to Phaser objects for efficient updates
    private mobSprites: Map<string, Phaser.GameObjects.Arc> = new Map();
    private mobTexts: Map<string, Phaser.GameObjects.Text> = new Map();
    private playerSprite?: Phaser.GameObjects.Arc;
    private scoreText?: Phaser.GameObjects.Text;
    private comboText?: Phaser.GameObjects.Text;

    /**
     * Set the Phaser.Scene to use for rendering.
     */
    setScene(scene: Phaser.Scene) {
        this.scene = scene;
        this.width = scene.scale.width;
        this.height = scene.scale.height;
        this.initialized = true;
    }

    /**
     * Optionally store the container, but do not create a Phaser.Game or scene here.
     */
    init(container: HTMLElement): void {
        this.container = container;
        this.width = container.offsetWidth || 800;
        this.height = container.offsetHeight || 600;
        // Try to get the Phaser.Scene from the container or global game
        // @ts-ignore
        if (container && (container as any).__phaserScene) {
            // @ts-ignore
            this.scene = (container as any).__phaserScene;
        } else if (window && (window as any).Phaser && (window as any).Phaser.GAMES && (window as any).Phaser.GAMES.length > 0) {
            // Try to get the active scene from the global Phaser game
            const phaserGame = (window as any).Phaser.GAMES[0];
            if (phaserGame && phaserGame.scene && phaserGame.scene.keys) {
                // Pick the first running scene
                const keys = Object.keys(phaserGame.scene.keys);
                for (const key of keys) {
                    const s = phaserGame.scene.keys[key];
                    if (s && s.sys && s.sys.isActive()) {
                        this.scene = s;
                        break;
                    }
                }
            }
        }
        this.initialized = true;
    }

    render(state: GameState): void {
        // Always try to get a valid scene
        if (!this.scene) {
            // Try to get the scene from the global Phaser game
            if (window && (window as any).Phaser && (window as any).Phaser.GAMES && (window as any).Phaser.GAMES.length > 0) {
                const phaserGame = (window as any).Phaser.GAMES[0];
                if (phaserGame && phaserGame.scene && phaserGame.scene.keys) {
                    const keys = Object.keys(phaserGame.scene.keys);
                    for (const key of keys) {
                        const s = phaserGame.scene.keys[key];
                        if (s && s.sys && s.sys.isActive()) {
                            this.scene = s;
                            break;
                        }
                    }
                }
            }
        }
        const scene = this.scene;
        if (!scene) return;

        // Remove any previous HTML overlays
        const prevMenu = document.getElementById('phaser-mainmenu');
        if (prevMenu) prevMenu.remove();

        // --- Main Menu ---
        if (state.gameStatus === 'mainMenu') {
            // Remove all Phaser objects (if any) so only menu is visible
            if (scene.children && typeof scene.children.removeAll === 'function') {
                scene.children.removeAll();
            }

            // --- Draw Phaser-based main menu ---
            // Title
            scene.add.text(this.width / 2, this.height / 2 - 80, 'TypeDefense', {
                fontSize: '64px',
                color: '#fff',
                fontStyle: 'bold',
                fontFamily: 'sans-serif',
                stroke: '#222',
                strokeThickness: 6,
                align: 'center'
            }).setOrigin(0.5);

            // Start button (visual only, not interactive)
            const startBtn = scene.add.rectangle(this.width / 2, this.height / 2 + 10, 320, 64, 0x007bff, 1)
                .setStrokeStyle(4, 0xffffff)
                .setInteractive({ useHandCursor: true });
            scene.add.text(this.width / 2, this.height / 2 + 10, 'Start', {
                fontSize: '36px',
                color: '#fff',
                fontFamily: 'sans-serif',
                fontStyle: 'bold'
            }).setOrigin(0.5);

            // Instructions
            scene.add.text(this.width / 2, this.height / 2 + 90, 'Press Enter or click Start to begin', {
                fontSize: '22px',
                color: '#aaa',
                fontFamily: 'sans-serif'
            }).setOrigin(0.5);

            // Version text
            scene.add.text(this.width / 2, this.height / 2 + 160, 'Version 2.0', {
                fontSize: '16px',
                color: '#888',
                fontFamily: 'sans-serif'
            }).setOrigin(0.5);

            // Credits text
            scene.add.text(this.width / 2, this.height / 2 + 190, '© TypeDefense Team', {
                fontSize: '14px',
                color: '#666',
                fontFamily: 'sans-serif'
            }).setOrigin(0.5);

            // Button interaction (update gameStatus)
            startBtn.on('pointerdown', () => {
                (window as any).stateManager?.setGameStatus('worldSelect');
            });

            // Keyboard support (Phaser input)
            scene.input.keyboard?.once('keydown-ENTER', () => {
                (window as any).stateManager?.setGameStatus('worldSelect');
            });

            // No HTML overlay for main menu
            return;
        }

        // Remove menu if not in mainMenu
        const menu = document.getElementById('phaser-mainmenu');
        if (menu) menu.remove();

        // --- Clear all objects if switching to a new gameStatus (not mainMenu) ---
        if (scene.children && typeof scene.children.removeAll === 'function') {
            scene.children.removeAll();
        }

        // --- Player ---
        if (!this.playerSprite) {
            this.playerSprite = scene.add.circle(
                state.player.position.x,
                state.player.position.y,
                22,
                0x3399ff
            );
        } else {
            this.playerSprite.setPosition(state.player.position.x, state.player.position.y);
        }

        // --- Mobs ---
        // Remove sprites/texts for mobs no longer present
        for (const [id, sprite] of this.mobSprites.entries()) {
            if (!state.mobs.some(m => m.id === id && !m.isDefeated)) {
                sprite.destroy();
                this.mobSprites.delete(id);
                const text = this.mobTexts.get(id);
                if (text) {
                    text.destroy();
                    this.mobTexts.delete(id);
                }
            }
        }
        // Add/update sprites/texts for current mobs
        for (const mob of state.mobs) {
            if (mob.isDefeated) continue;
            let sprite = this.mobSprites.get(mob.id);
            let text = this.mobTexts.get(mob.id);
            if (!sprite) {
                sprite = scene.add.circle(
                    mob.position.x,
                    mob.position.y,
                    18,
                    0xffaa00
                );
                this.mobSprites.set(mob.id, sprite);
            } else {
                sprite.setPosition(mob.position.x, mob.position.y);
            }
            if (!text) {
                text = scene.add.text(
                    mob.position.x,
                    mob.position.y - 28,
                    mob.word,
                    { fontSize: '18px', color: '#fff', fontStyle: 'bold', align: 'center' }
                ).setOrigin(0.5);
                this.mobTexts.set(mob.id, text);
            } else {
                text.setPosition(mob.position.x, mob.position.y - 28);
                text.setText(mob.word.substring(mob.currentTypedIndex));
            }
        }

        // --- Score/Combo UI ---
        if (!this.scoreText) {
            this.scoreText = scene.add.text(24, 16, `Score: ${state.player.score}`, {
                fontSize: '22px', color: '#fff', fontStyle: 'bold'
            }).setOrigin(0, 0);
        } else {
            this.scoreText.setText(`Score: ${state.player.score}`);
        }
        if (!this.comboText) {
            this.comboText = scene.add.text(24, 44, `Combo x${state.player.combo}`, {
                fontSize: '18px', color: '#44ff44'
            }).setOrigin(0, 0);
        }
        if (state.player.combo > 1) {
            this.comboText.setText(`Combo x${state.player.combo}`);
            this.comboText.setVisible(true);
        } else {
            this.comboText.setVisible(false);
        }
    }

    destroy(): void {
        this.mobSprites.clear();
        this.mobTexts.clear();
        this.playerSprite = undefined;
        this.scoreText = undefined;
        this.comboText = undefined;
        this.initialized = false;
        this.scene = undefined;
    }

    resize(width: number, height: number): void {
        this.width = width;
        this.height = height;
        if (this.scene) {
            this.scene.scale.resize(width, height);
        }
    }
}
