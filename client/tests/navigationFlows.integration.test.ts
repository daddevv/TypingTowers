// Playwright E2E tests for navigation flows in TypeDefense
// Covers: main menu → world select → level select → game, level complete navigation, back navigation, and edge cases
import { test, expect } from '@playwright/test';

const BASE_URL = 'http://localhost:5173';

test.describe('Navigation Flows', () => {
  test('Main menu to world select to level select to game', async ({ page }) => {
    await page.goto(BASE_URL);
    await page.waitForLoadState('networkidle');
    await page.waitForTimeout(1000);
    // Main menu: Play button
    await expect(page.locator('text=Play')).toBeVisible({ timeout: 10000 });
    await page.click('text=Play');
    // World select
    await expect(page.locator('text=TypeDefense')).toBeVisible({ timeout: 10000 });
    // Select first world (Enter)
    await page.keyboard.press('Enter');
    // Level select
    await expect(page.locator('text=/Level 1-1/i')).toBeVisible({ timeout: 10000 });
    // Select first level (Enter)
    await page.keyboard.press('Enter');
    // Game scene: look for score UI
    await page.waitForTimeout(1000);
    await expect(page.locator('text=/Score:/i')).toBeVisible({ timeout: 10000 });
  });

  test('Back navigation: Level select → world select → main menu', async ({ page }) => {
    await page.goto(BASE_URL);
    await page.waitForLoadState('networkidle');
    await page.waitForTimeout(1000);
    await expect(page.locator('text=Play')).toBeVisible({ timeout: 10000 });
    await page.click('text=Play');
    await expect(page.locator('text=TypeDefense')).toBeVisible({ timeout: 10000 });
    await page.keyboard.press('Enter'); // World select
    await expect(page.locator('text=/Level 1-1/i')).toBeVisible({ timeout: 10000 });
    // Press Escape to go back to world select
    await page.keyboard.press('Escape');
    await expect(page.locator('text=TypeDefense')).toBeVisible({ timeout: 10000 });
    // Press Escape to go back to main menu
    await page.keyboard.press('Escape');
    await expect(page.locator('text=Play')).toBeVisible({ timeout: 10000 });
  });

  test('Level complete navigation: Continue and Back', async ({ page }) => {
    await page.goto(BASE_URL);
    await page.waitForLoadState('networkidle');
    await page.waitForTimeout(1000);
    await expect(page.locator('text=Play')).toBeVisible({ timeout: 10000 });
    await page.click('text=Play');
    await expect(page.locator('text=TypeDefense')).toBeVisible({ timeout: 10000 });
    await page.keyboard.press('Enter'); // World select
    await expect(page.locator('text=/Level 1-1/i')).toBeVisible({ timeout: 10000 });
    await page.keyboard.press('Enter'); // Level select
    await page.keyboard.press('Enter'); // Start level
    // Simulate level completion (cheat: set score high via window)
    await page.waitForTimeout(2000);
    await page.evaluate(() => {
      // @ts-ignore
      window.gameState.levelStatus = 'complete';
      // @ts-ignore
      window.dispatchEvent(new Event('gameStateChanged'));
    });
    // Wait for level complete UI
    await page.waitForSelector('text=Level Complete', { timeout: 10000 });
    // Continue (Enter)
    await page.keyboard.press('Enter');
    // Should start next level or return to menu
    await page.waitForTimeout(1000);
    // Back to level select (simulate level complete again)
    await page.evaluate(() => {
      // @ts-ignore
      window.gameState.levelStatus = 'complete';
      // @ts-ignore
      window.dispatchEvent(new Event('gameStateChanged'));
    });
    await page.waitForSelector('text=Level Complete', { timeout: 10000 });
    // Back (Escape)
    await page.keyboard.press('Escape');
    await expect(page.locator('text=/Level 1-1/i')).toBeVisible({ timeout: 10000 });
  });

  test('Edge: cannot select locked level', async ({ page }) => {
    await page.goto(BASE_URL);
    await page.waitForLoadState('networkidle');
    await page.waitForTimeout(1000);
    await expect(page.locator('text=Play')).toBeVisible({ timeout: 10000 });
    await page.click('text=Play');
    await expect(page.locator('text=TypeDefense')).toBeVisible({ timeout: 10000 });
    await page.keyboard.press('Enter'); // World select
    // Try to select a locked level (e.g., Level 1-7)
    await page.keyboard.press('ArrowDown');
    await page.keyboard.press('ArrowDown');
    await page.keyboard.press('ArrowDown');
    await page.keyboard.press('ArrowDown');
    await page.keyboard.press('ArrowDown');
    await page.keyboard.press('ArrowDown');
    // Try to select (Enter)
    await page.keyboard.press('Enter');
    // Should remain on level select (locked)
    await expect(page.locator('text=/Level 1-7 \\(Locked\\)/i')).toBeVisible({ timeout: 10000 });
  });
});
