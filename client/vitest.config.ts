import { resolve } from 'path';
import { defineConfig } from 'vitest/config';

export default defineConfig({
    test: {
        setupFiles: [resolve(__dirname, 'setupTests.ts')],
        environment: 'jsdom', // switched from 'node' to 'jsdom' for browser-like API
        globals: true,
        // Exclude Playwright E2E, Playwright-style test files, and all node_modules tests from Vitest
        exclude: [
            'e2e/**',
            'tests/**',
            '**/*.e2e.{ts,js}',
            '**/*.spec.{ts,js}',
            '**/node_modules/**'
        ],
        coverage: {
            reporter: ['text', 'json', 'html'],
        },
    },
});
// Contains AI-generated edits.
