// @ts-check
/** @type {import('@playwright/test').PlaywrightTestConfig} */
const config = {
    testDir: './tests',
    webServer: {
        command: 'npm run dev',
        port: 5173,
        reuseExistingServer: true,
        timeout: 120 * 1000,
    },
    use: {
        baseURL: 'http://localhost:5173',
        headless: true,
        viewport: { width: 1280, height: 720 },
        ignoreHTTPSErrors: true,
        video: 'on',
        launchOptions: {
            args: ['--window-size=1280,720'],
            env: {
                PLAYWRIGHT: 'true',
            },
        },
        slowMo: 500,
    },
};

export default config;
// Contains AI-generated edits.
