// setupTests.ts
// Global test setup for Vitest: mock browser globals for Node.js

// Polyfill browser APIs for tests using jsdom, node-canvas, and headless-gl
try {
    // jsdom: create a global window/document if not present
    if (typeof globalThis.window === 'undefined' || typeof globalThis.document === 'undefined') {
        const { JSDOM } = require('jsdom');
        const dom = new JSDOM('<!doctype html><html><body></body></html>');
        globalThis.window = dom.window;
        globalThis.document = dom.window.document;
        globalThis.navigator = dom.window.navigator;
        globalThis.HTMLElement = dom.window.HTMLElement;
        globalThis.Node = dom.window.Node;
        globalThis.getComputedStyle = dom.window.getComputedStyle;
    }
    // node-canvas: patch window.CanvasRenderingContext2D and HTMLCanvasElement
    if (typeof globalThis.window !== 'undefined' && typeof globalThis.window.HTMLCanvasElement === 'undefined') {
        try {
            const canvas = require('canvas');
            globalThis.window.HTMLCanvasElement = canvas.Canvas;
            globalThis.window.CanvasRenderingContext2D = canvas.CanvasRenderingContext2D;
            globalThis.HTMLCanvasElement = canvas.Canvas;
            globalThis.CanvasRenderingContext2D = canvas.CanvasRenderingContext2D;
        } catch (e) {
            // node-canvas not available, warn and continue (tests that require Canvas2D will fail)
            // eslint-disable-next-line no-console
            console.warn('[setupTests] node-canvas not installed or failed to load. Canvas 2D APIs will not be polyfilled.');
        }
    }
    // headless-gl: patch window.WebGLRenderingContext
    if (typeof globalThis.window !== 'undefined' && typeof globalThis.window.WebGLRenderingContext === 'undefined') {
        try {
            const gl = require('gl');
            globalThis.window.WebGLRenderingContext = gl.WebGLRenderingContext || function () { };
            globalThis.WebGLRenderingContext = globalThis.window.WebGLRenderingContext;
        } catch (e) {
            // headless-gl not available, skip
        }
    }
} catch (e) {
    // Ignore errors if polyfills are already present or modules not found
}

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
