# Test info

- Name: Pause Menu >> should pause and resume the game with Escape key
- Location: /home/bobbitt/projects/public/type-defense/client/tests/pauseMenu.spec.ts:7:5

# Error details

```
Error: page.waitForSelector: Test timeout of 30000ms exceeded.
Call log:
  - waiting for locator('#pause-header') to be visible

    at /home/bobbitt/projects/public/type-defense/client/tests/pauseMenu.spec.ts:32:20
```

# Page snapshot

```yaml
- heading "TypeDefense" [level=1]
```

# Test source

```ts
   1 | // Playwright test for Pause functionality in Type Defense
   2 | // This test assumes the dev server is running on http://localhost:5173
   3 | // Adjust the URL if needed.
   4 | import { expect, test } from '@playwright/test';
   5 |
   6 | test.describe('Pause Menu', () => {
   7 |     test('should pause and resume the game with Escape key', async ({ page }) => {
   8 |         await page.goto('http://localhost:5173');
   9 |
  10 |         // Wait for the game to load
  11 |         await page.waitForLoadState('networkidle');
  12 |         await page.waitForTimeout(1000);
  13 |
  14 |         // Click on the game canvas to ensure focus
  15 |         const canvas = await page.locator('canvas').first();
  16 |         await canvas.click({ position: { x: 640, y: 360 } });
  17 |
  18 |         // Navigate to the actual game
  19 |         await page.keyboard.press('Enter'); // Select world
  20 |         await page.waitForTimeout(500);
  21 |         await page.keyboard.press('Enter'); // Select level
  22 |         await page.waitForTimeout(500);
  23 |         await page.keyboard.press('Enter'); // Start level
  24 |         await page.waitForTimeout(2000);
  25 |
  26 |         // Click again to ensure focus
  27 |         await canvas.click({ position: { x: 640, y: 360 } });
  28 |
  29 |         // Press Escape to open pause menu
  30 |         await page.keyboard.press('Escape');
  31 |         // Check for pause menu UI
> 32 |         await page.waitForSelector('#pause-header');
     |                    ^ Error: page.waitForSelector: Test timeout of 30000ms exceeded.
  33 |         await expect(page.locator('#pause-header')).toBeVisible();
  34 |
  35 |         // Press Escape again to close pause menu
  36 |         await page.keyboard.press('Escape');
  37 |         // Pause menu should disappear
  38 |         await page.waitForSelector('#pause-header', { state: 'detached' });
  39 |         await expect(page.locator('#pause-header')).toBeHidden();
  40 |     });
  41 | });
  42 | // Contains AI-generated edits.
  43 |
```