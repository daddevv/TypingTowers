// Ensures headless engine and its tests are pure Node.js (no Phaser, no DOM)
// This test will fail if any import/require of 'phaser' is present in engine or its tests

// PhaserImportCheck.test.ts
import fs from 'fs';
import path from 'path';
import { describe, it, expect } from 'vitest';

describe('No Phaser import in headless engine or tests', () => {
  const files = [
    path.resolve(__dirname, 'HeadlessGameEngine.ts'),
    path.resolve(__dirname, 'HeadlessGameEngine.unit.test.ts'),
  ];
  for (const file of files) {
    it(`${path.basename(file)} should not import phaser`, () => {
      const content = fs.readFileSync(file, 'utf8');
      expect(content.includes("from 'phaser'"), `File ${file} should not import from 'phaser'`).toBe(false);
      expect(content.includes('from "phaser"'), `File ${file} should not import from \"phaser\"`).toBe(false);
      expect(content.includes('require("phaser")'), `File ${file} should not require(\"phaser\")`).toBe(false);
      expect(content.includes("require('phaser')"), `File ${file} should not require('phaser')`).toBe(false);
    });
  }
});
