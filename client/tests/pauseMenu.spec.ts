// Playwright test for Pause functionality in Type Defense
// This test assumes the dev server is running on http://localhost:5173
// Adjust the URL if needed.
import { expect, test } from '@playwright/test';

test.describe('Pause Menu', () => {
    test('should pause and resume the game with Escape key', async ({ page }) => {
        await page.goto('http://localhost:5173');

        // Wait for the game to load
        await page.waitForLoadState('networkidle');
        await page.waitForTimeout(1000);

        // Click on the game canvas to ensure focus
        const canvas = await page.locator('canvas').first();
        await canvas.click({ position: { x: 640, y: 360 } });

        // Navigate to the actual game
        await page.keyboard.press('Enter'); // Select world
        await page.waitForTimeout(500);
        await page.keyboard.press('Enter'); // Select level
        await page.waitForTimeout(500);
        await page.keyboard.press('Enter'); // Start level
        await page.waitForTimeout(2000);

        // Click again to ensure focus
        await canvas.click({ position: { x: 640, y: 360 } });

        // Press Escape to open pause menu
        await page.keyboard.press('Escape');
        // Check for pause menu UI
        await page.waitForSelector('#pause-header');
        await expect(page.locator('#pause-header')).toBeVisible();

        // Press Escape again to close pause menu
        await page.keyboard.press('Escape');
        // Pause menu should disappear
        await page.waitForSelector('#pause-header', { state: 'detached' });
        await expect(page.locator('#pause-header')).toBeHidden();
    });
});
// Contains AI-generated edits.
