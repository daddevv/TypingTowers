import { resolve } from 'path';
import { defineConfig } from 'vitest/config';

export default defineConfig({
    test: {
        setupFiles: [resolve(__dirname, 'setupTests.ts')],
        environment: 'node',
        // Exclude Playwright E2E and Playwright-style test files from Vitest
        exclude: ['e2e/**', 'tests/**', '**/*.e2e.{ts,js}', '**/*.spec.{ts,js}'],
    },
});
// Contains AI-generated edits.
