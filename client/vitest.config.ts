import { resolve } from 'path';
import { defineConfig } from 'vitest/config';

export default defineConfig({
    test: {
        setupFiles: [resolve(__dirname, 'setupTests.ts')],
        environment: 'node',
    },
});
// Contains AI-generated edits.
