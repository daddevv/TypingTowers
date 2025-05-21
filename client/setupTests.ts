// setupTests.ts
// Global test setup for Vitest: mock browser globals for Node.js

// Mock 'phaser3spectorjs' for Phaser compatibility in Vitest/Node.js
try {
    if (typeof require !== 'undefined') {
        require.cache = require.cache || {};
        const fakeModulePath = require.resolve ? require.resolve('phaser3spectorjs') : 'phaser3spectorjs';
        require.cache[fakeModulePath] = {
            id: fakeModulePath,
            filename: fakeModulePath,
            loaded: true,
            exports: {},
            children: [],
            paths: [],
            parent: null,
            require: require,
            isPreloading: false,
            path: '',
        };
    }
} catch (e) {
    // Ignore if require.resolve fails (module not found)
}

const g: any = typeof globalThis !== 'undefined' ? globalThis : (typeof window !== 'undefined' ? window : {});

if (typeof g !== 'undefined') {
    // Mock window for Phaser
    if (typeof g.window === 'undefined') {
        g.window = {};
    }
    // Mock document for Phaser
    if (typeof g.document === 'undefined') {
        g.document = {
            createElement: (type: string) => {
                if (type === 'canvas') {
                    // Minimal mock of CanvasRenderingContext2D for Phaser
                    return {
                        getContext: () => ({
                            fillRect: () => { },
                            clearRect: () => { },
                            getImageData: () => ({ data: [] }),
                            putImageData: () => { },
                            createImageData: () => [],
                            setTransform: () => { },
                            drawImage: () => { },
                            save: () => { },
                            restore: () => { },
                            beginPath: () => { },
                            moveTo: () => { },
                            lineTo: () => { },
                            closePath: () => { },
                            stroke: () => { },
                            translate: () => { },
                            scale: () => { },
                            rotate: () => { },
                            arc: () => { },
                            fill: () => { },
                            measureText: (text: string) => ({ width: text.length * 10 }),
                            transform: () => { },
                            rect: () => { },
                            clip: () => { },
                            // Add more stubs as needed by Phaser
                        })
                    };
                }
                return { getContext: () => ({}) };
            },
            addEventListener: () => { },
            removeEventListener: () => { },
            body: {},
            documentElement: {},
            head: {},
        };
    }
    // Mock HTMLCanvasElement and CanvasRenderingContext2D for Phaser
    if (typeof g.HTMLCanvasElement === 'undefined') {
        g.HTMLCanvasElement = function () { };
    }
    if (typeof g.CanvasRenderingContext2D === 'undefined') {
        g.CanvasRenderingContext2D = function () { };
    }
    // Mock localStorage for LevelManager and others
    if (typeof g.localStorage === 'undefined') {
        let store: Record<string, string> = {};
        g.localStorage = {
            getItem: (key: string) => store[key] || null,
            setItem: (key: string, value: string) => { store[key] = value; },
            removeItem: (key: string) => { delete store[key]; },
            clear: () => { store = {}; },
        } as any;
    }
    // Mock Image for Phaser (fixes ReferenceError: Image is not defined)
    if (typeof g.Image === 'undefined') {
        g.Image = class {
            // Optionally add minimal properties if needed
        };
    }
}
// Contains AI-generated edits.
