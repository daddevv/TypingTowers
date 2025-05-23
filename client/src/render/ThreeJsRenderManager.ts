import * as THREE from 'three';
import { IRenderManager } from './RenderManager';
import { GameState } from '../state/gameState';

export class ThreeJsRenderManager implements IRenderManager {
    private renderer?: THREE.WebGLRenderer;
    private scene?: THREE.Scene;
    private camera?: THREE.OrthographicCamera;
    private container?: HTMLElement;
    private width: number = 800;
    private height: number = 600;
    private mobMeshes: Map<string, THREE.Mesh> = new Map();
    private playerMesh?: THREE.Mesh;

    init(container: HTMLElement): void {
        this.container = container;
        this.width = container.offsetWidth || 800;
        this.height = container.offsetHeight || 600;

        // Set up renderer
        this.renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true });
        this.renderer.setSize(this.width, this.height);
        this.renderer.setClearColor(0x222222, 1);

        // Remove any previous canvas
        while (container.firstChild) container.removeChild(container.firstChild);
        container.appendChild(this.renderer.domElement);

        // Set up scene and camera
        this.scene = new THREE.Scene();
        this.camera = new THREE.OrthographicCamera(
            0, this.width, this.height, 0, -1000, 1000
        );
        this.camera.position.z = 10;
    }

    render(state: GameState): void {
        if (!this.scene || !this.camera || !this.renderer) return;

        // Remove any previous HTML overlays
        const prevMenu = document.getElementById('threejs-mainmenu');
        if (prevMenu) prevMenu.remove();

        // --- Main Menu ---
        if (state.gameStatus === 'mainMenu') {
            // Create HTML overlay for main menu
            const menuDiv = document.createElement('div');
            menuDiv.id = 'threejs-mainmenu';
            menuDiv.style.position = 'absolute';
            menuDiv.style.left = '0';
            menuDiv.style.top = '0';
            menuDiv.style.width = '100%';
            menuDiv.style.height = '100%';
            menuDiv.style.display = 'flex';
            menuDiv.style.flexDirection = 'column';
            menuDiv.style.alignItems = 'center';
            menuDiv.style.justifyContent = 'center';
            menuDiv.style.pointerEvents = 'auto';
            menuDiv.style.zIndex = '10';

            // Title
            const title = document.createElement('div');
            title.textContent = 'TypeDefense';
            title.style.fontSize = '48px';
            title.style.color = '#fff';
            title.style.marginBottom = '40px';
            title.style.fontWeight = 'bold';
            menuDiv.appendChild(title);

            // Start button
            const startBtn = document.createElement('button');
            startBtn.textContent = 'Start';
            startBtn.style.fontSize = '32px';
            startBtn.style.padding = '16px 48px';
            startBtn.style.background = '#007bff';
            startBtn.style.color = '#fff';
            startBtn.style.border = 'none';
            startBtn.style.borderRadius = '8px';
            startBtn.style.cursor = 'pointer';
            startBtn.style.marginBottom = '32px';
            startBtn.onmouseenter = () => startBtn.style.background = '#0056b3';
            startBtn.onmouseleave = () => startBtn.style.background = '#007bff';
            startBtn.onclick = () => {
                (window as any).stateManager?.setGameStatus('worldSelect');
            };
            menuDiv.appendChild(startBtn);

            // Instructions
            const instr = document.createElement('div');
            instr.textContent = 'Press Enter or click Start to begin';
            instr.style.fontSize = '20px';
            instr.style.color = '#aaa';
            menuDiv.appendChild(instr);

            // Keyboard support
            const keyHandler = (e: KeyboardEvent) => {
                if (e.key === 'Enter') {
                    (window as any).stateManager?.setGameStatus('worldSelect');
                }
            };
            window.addEventListener('keydown', keyHandler, { once: true });

            // Attach to container
            if (this.container) {
                this.container.appendChild(menuDiv);
            }

            // No 3D rendering needed for main menu
            this.renderer.setClearColor(0x222222, 1);
            this.renderer.clear();
            return;
        }

        // Remove menu if not in mainMenu
        const menu = document.getElementById('threejs-mainmenu');
        if (menu) menu.remove();

        // Clear previous frame
        // Remove all mob meshes
        for (const mesh of this.mobMeshes.values()) {
            this.scene.remove(mesh);
        }
        this.mobMeshes.clear();

        // Remove player mesh if exists
        if (this.playerMesh) {
            this.scene.remove(this.playerMesh);
            this.playerMesh = undefined;
        }

        // Render mobs as colored spheres
        if (Array.isArray(state.mobs)) {
            for (const mob of state.mobs) {
                if (mob.isDefeated) continue;
                const geometry = new THREE.SphereGeometry(18, 16, 16);
                const material = new THREE.MeshBasicMaterial({ color: 0xffaa00 });
                const mesh = new THREE.Mesh(geometry, material);
                mesh.position.set(mob.position.x, mob.position.y, 0);
                this.scene.add(mesh);
                this.mobMeshes.set(mob.id, mesh);
            }
        }

        // Render player as a blue sphere
        if (state.player && state.player.position) {
            const geometry = new THREE.SphereGeometry(22, 18, 18);
            const material = new THREE.MeshBasicMaterial({ color: 0x3399ff });
            const mesh = new THREE.Mesh(geometry, material);
            mesh.position.set(state.player.position.x, state.player.position.y, 0);
            this.scene.add(mesh);
            this.playerMesh = mesh;
        }

        // Optionally: render score/combo as overlay using HTML or Three.js sprites (not implemented in prototype)

        this.renderer.render(this.scene, this.camera);
    }

    destroy(): void {
        if (this.renderer) {
            this.renderer.dispose();
            if (this.renderer.domElement.parentNode) {
                this.renderer.domElement.parentNode.removeChild(this.renderer.domElement);
            }
        }
        this.scene = undefined;
        this.camera = undefined;
        this.mobMeshes.clear();
        this.playerMesh = undefined;
    }

    resize(width: number, height: number): void {
        this.width = width;
        this.height = height;
        if (this.renderer) {
            this.renderer.setSize(width, height);
        }
        if (this.camera) {
            this.camera.left = 0;
            this.camera.right = width;
            this.camera.top = height;
            this.camera.bottom = 0;
            this.camera.updateProjectionMatrix();
        }
    }
}
