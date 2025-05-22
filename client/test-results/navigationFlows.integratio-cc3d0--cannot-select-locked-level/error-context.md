# Test info

- Name: Navigation Flows >> Edge: cannot select locked level
- Location: /home/bobbitt/projects/public/type-defense/client/tests/navigationFlows.integration.test.ts:83:3

# Error details

```
Error: Timed out 10000ms waiting for expect(locator).toBeVisible()

Locator: locator('text=Play')
Expected: visible
Received: <element(s) not found>
Call log:
  - expect.toBeVisible with timeout 10000ms
  - waiting for locator('text=Play')

    at /home/bobbitt/projects/public/type-defense/client/tests/navigationFlows.integration.test.ts:87:45
```

# Page snapshot

```yaml
- heading "TypeDefense" [level=1]
```

# Test source

```ts
   1 | // Playwright E2E tests for navigation flows in TypeDefense
   2 | // Covers: main menu → world select → level select → game, level complete navigation, back navigation, and edge cases
   3 | import { test, expect } from '@playwright/test';
   4 |
   5 | const BASE_URL = 'http://localhost:5173';
   6 |
   7 | test.describe('Navigation Flows', () => {
   8 |   test('Main menu to world select to level select to game', async ({ page }) => {
   9 |     await page.goto(BASE_URL);
   10 |     await page.waitForLoadState('networkidle');
   11 |     await page.waitForTimeout(1000);
   12 |     // Main menu: Play button
   13 |     await expect(page.locator('text=Play')).toBeVisible({ timeout: 10000 });
   14 |     await page.click('text=Play');
   15 |     // World select
   16 |     await expect(page.locator('text=TypeDefense')).toBeVisible({ timeout: 10000 });
   17 |     // Select first world (Enter)
   18 |     await page.keyboard.press('Enter');
   19 |     // Level select
   20 |     await expect(page.locator('text=/Level 1-1/i')).toBeVisible({ timeout: 10000 });
   21 |     // Select first level (Enter)
   22 |     await page.keyboard.press('Enter');
   23 |     // Game scene: look for score UI
   24 |     await page.waitForTimeout(1000);
   25 |     await expect(page.locator('text=/Score:/i')).toBeVisible({ timeout: 10000 });
   26 |   });
   27 |
   28 |   test('Back navigation: Level select → world select → main menu', async ({ page }) => {
   29 |     await page.goto(BASE_URL);
   30 |     await page.waitForLoadState('networkidle');
   31 |     await page.waitForTimeout(1000);
   32 |     await expect(page.locator('text=Play')).toBeVisible({ timeout: 10000 });
   33 |     await page.click('text=Play');
   34 |     await expect(page.locator('text=TypeDefense')).toBeVisible({ timeout: 10000 });
   35 |     await page.keyboard.press('Enter'); // World select
   36 |     await expect(page.locator('text=/Level 1-1/i')).toBeVisible({ timeout: 10000 });
   37 |     // Press Escape to go back to world select
   38 |     await page.keyboard.press('Escape');
   39 |     await expect(page.locator('text=TypeDefense')).toBeVisible({ timeout: 10000 });
   40 |     // Press Escape to go back to main menu
   41 |     await page.keyboard.press('Escape');
   42 |     await expect(page.locator('text=Play')).toBeVisible({ timeout: 10000 });
   43 |   });
   44 |
   45 |   test('Level complete navigation: Continue and Back', async ({ page }) => {
   46 |     await page.goto(BASE_URL);
   47 |     await page.waitForLoadState('networkidle');
   48 |     await page.waitForTimeout(1000);
   49 |     await expect(page.locator('text=Play')).toBeVisible({ timeout: 10000 });
   50 |     await page.click('text=Play');
   51 |     await expect(page.locator('text=TypeDefense')).toBeVisible({ timeout: 10000 });
   52 |     await page.keyboard.press('Enter'); // World select
   53 |     await expect(page.locator('text=/Level 1-1/i')).toBeVisible({ timeout: 10000 });
   54 |     await page.keyboard.press('Enter'); // Level select
   55 |     await page.keyboard.press('Enter'); // Start level
   56 |     // Simulate level completion (cheat: set score high via window)
   57 |     await page.waitForTimeout(2000);
   58 |     await page.evaluate(() => {
   59 |       // @ts-ignore
   60 |       window.gameState.levelStatus = 'complete';
   61 |       // @ts-ignore
   62 |       window.dispatchEvent(new Event('gameStateChanged'));
   63 |     });
   64 |     // Wait for level complete UI
   65 |     await page.waitForSelector('text=Level Complete', { timeout: 10000 });
   66 |     // Continue (Enter)
   67 |     await page.keyboard.press('Enter');
   68 |     // Should start next level or return to menu
   69 |     await page.waitForTimeout(1000);
   70 |     // Back to level select (simulate level complete again)
   71 |     await page.evaluate(() => {
   72 |       // @ts-ignore
   73 |       window.gameState.levelStatus = 'complete';
   74 |       // @ts-ignore
   75 |       window.dispatchEvent(new Event('gameStateChanged'));
   76 |     });
   77 |     await page.waitForSelector('text=Level Complete', { timeout: 10000 });
   78 |     // Back (Escape)
   79 |     await page.keyboard.press('Escape');
   80 |     await expect(page.locator('text=/Level 1-1/i')).toBeVisible({ timeout: 10000 });
   81 |   });
   82 |
   83 |   test('Edge: cannot select locked level', async ({ page }) => {
   84 |     await page.goto(BASE_URL);
   85 |     await page.waitForLoadState('networkidle');
   86 |     await page.waitForTimeout(1000);
>  87 |     await expect(page.locator('text=Play')).toBeVisible({ timeout: 10000 });
      |                                             ^ Error: Timed out 10000ms waiting for expect(locator).toBeVisible()
   88 |     await page.click('text=Play');
   89 |     await expect(page.locator('text=TypeDefense')).toBeVisible({ timeout: 10000 });
   90 |     await page.keyboard.press('Enter'); // World select
   91 |     // Try to select a locked level (e.g., Level 1-7)
   92 |     await page.keyboard.press('ArrowDown');
   93 |     await page.keyboard.press('ArrowDown');
   94 |     await page.keyboard.press('ArrowDown');
   95 |     await page.keyboard.press('ArrowDown');
   96 |     await page.keyboard.press('ArrowDown');
   97 |     await page.keyboard.press('ArrowDown');
   98 |     // Try to select (Enter)
   99 |     await page.keyboard.press('Enter');
  100 |     // Should remain on level select (locked)
  101 |     await expect(page.locator('text=/Level 1-7 \\(Locked\\)/i')).toBeVisible({ timeout: 10000 });
  102 |   });
  103 | });
  104 |
```