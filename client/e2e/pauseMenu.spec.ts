// Playwright test for Pause functionality in Type Defense
// This test assumes the dev server is running on http://localhost:5173
// Adjust the selectors as needed for your menu flow.
import { expect, test } from '@playwright/test';

test.describe('Pause Menu', () => {
    test('should pause and resume the game with Escape key', async ({ page }) => {
        // Set viewport to a known size
        await page.setViewportSize({ width: 1280, height: 720 });
        await page.goto('http://localhost:5173');

        // Wait for page to load and stabilize
        await page.waitForLoadState('networkidle');
        await page.waitForTimeout(1000); // Allow animations to complete
        console.log('Navigating through start screens...');

        // All game interactions are via canvas
        const canvas = await page.locator('canvas').first();

        // Click the Play button (MainMenuScene) at canvas center (x=400, y=320)
        await canvas.click({ position: { x: 400, y: 320 } });
        await page.waitForTimeout(1000);

        // Click on first world (MenuScene) at x=400, y=120
        await canvas.click({ position: { x: 400, y: 120 } });
        await page.waitForTimeout(1000);

        // Click on first level (LevelMenuScene) at x=400, y=120
        await canvas.click({ position: { x: 400, y: 120 } });
        await page.waitForTimeout(2000);

        console.log('Game should be loaded now, pressing Escape to pause...');

        // Refocus canvas before pressing Escape (x=400, y=300)
        await canvas.click({ position: { x: 400, y: 300 } });
        await page.waitForTimeout(300);

        // Press Escape to open pause menu
        await page.keyboard.press('Escape');

        // Wait for pause header element
        console.log('Waiting for pause menu to appear...');
        await page.waitForSelector('[data-testid="pause-header"]', { timeout: 5000 });

        // Take screenshot for debugging
        await page.screenshot({ path: 'pause-menu-visible.png' });

        // Assert that the pause header is visible
        await expect(page.locator('[data-testid="pause-header"]')).toBeVisible();

        // Press Escape again to close pause menu
        await page.keyboard.press('Escape');

        // Wait for pause menu to disappear
        console.log('Waiting for pause menu to disappear...');
        await page.waitForSelector('[data-testid="pause-header"]', { state: 'detached', timeout: 5000 });

        // Take screenshot for debugging
        await page.screenshot({ path: 'pause-menu-hidden.png' });

        // Assert that the pause header is hidden (game is unpaused)
        await expect(page.locator('[data-testid="pause-header"]')).toBeHidden();
    });
});

// Contains AI-generated edits.
