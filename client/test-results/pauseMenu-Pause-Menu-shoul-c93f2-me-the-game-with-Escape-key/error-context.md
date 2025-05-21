# Test info

- Name: Pause Menu >> should pause and resume the game with Escape key
- Location: C:\Users\mbobbitt\dev\type-defense\client\e2e\pauseMenu.spec.ts:7:5

# Error details

```
TimeoutError: page.waitForSelector: Timeout 5000ms exceeded.
Call log:
  - waiting for locator('[data-testid="pause-header"]') to be visible

    at C:\Users\mbobbitt\dev\type-defense\client\e2e\pauseMenu.spec.ts:43:20
```

# Page snapshot

```yaml
- heading "TypeDefense" [level=1]
```

# Test source

```ts
   1 | // Playwright test for Pause functionality in Type Defense
   2 | // This test assumes the dev server is running on http://localhost:5173
   3 | // Adjust the selectors as needed for your menu flow.
   4 | import { expect, test } from '@playwright/test';
   5 |
   6 | test.describe('Pause Menu', () => {
   7 |     test('should pause and resume the game with Escape key', async ({ page }) => {
   8 |         // Set viewport to a known size
   9 |         await page.setViewportSize({ width: 1280, height: 720 });
  10 |         await page.goto('http://localhost:5173');
  11 |
  12 |         // Wait for page to load and stabilize
  13 |         await page.waitForLoadState('networkidle');
  14 |         await page.waitForTimeout(1000); // Allow animations to complete
  15 |         console.log('Navigating through start screens...');
  16 |
  17 |         // All game interactions are via canvas
  18 |         const canvas = await page.locator('canvas').first();
  19 |
  20 |         // Click the Play button (MainMenuScene) at canvas center (x=400, y=320)
  21 |         await canvas.click({ position: { x: 400, y: 320 } });
  22 |         await page.waitForTimeout(1000);
  23 |
  24 |         // Click on first world (MenuScene) at x=400, y=120
  25 |         await canvas.click({ position: { x: 400, y: 120 } });
  26 |         await page.waitForTimeout(1000);
  27 |
  28 |         // Click on first level (LevelMenuScene) at x=400, y=120
  29 |         await canvas.click({ position: { x: 400, y: 120 } });
  30 |         await page.waitForTimeout(2000);
  31 |
  32 |         console.log('Game should be loaded now, pressing Escape to pause...');
  33 |
  34 |         // Refocus canvas before pressing Escape (x=400, y=300)
  35 |         await canvas.click({ position: { x: 400, y: 300 } });
  36 |         await page.waitForTimeout(300);
  37 |
  38 |         // Press Escape to open pause menu
  39 |         await page.keyboard.press('Escape');
  40 |
  41 |         // Wait for pause header element
  42 |         console.log('Waiting for pause menu to appear...');
> 43 |         await page.waitForSelector('[data-testid="pause-header"]', { timeout: 5000 });
     |                    ^ TimeoutError: page.waitForSelector: Timeout 5000ms exceeded.
  44 |
  45 |         // Take screenshot for debugging
  46 |         await page.screenshot({ path: 'pause-menu-visible.png' });
  47 |
  48 |         // Assert that the pause header is visible
  49 |         await expect(page.locator('[data-testid="pause-header"]')).toBeVisible();
  50 |
  51 |         // Press Escape again to close pause menu
  52 |         await page.keyboard.press('Escape');
  53 |
  54 |         // Wait for pause menu to disappear
  55 |         console.log('Waiting for pause menu to disappear...');
  56 |         await page.waitForSelector('[data-testid="pause-header"]', { state: 'detached', timeout: 5000 });
  57 |
  58 |         // Take screenshot for debugging
  59 |         await page.screenshot({ path: 'pause-menu-hidden.png' });
  60 |
  61 |         // Assert that the pause header is hidden (game is unpaused)
  62 |         await expect(page.locator('[data-testid="pause-header"]')).toBeHidden();
  63 |     });
  64 | });
  65 |
  66 | // Contains AI-generated edits.
  67 |
```