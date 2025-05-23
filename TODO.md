# TypeDefense v2 — Headless Integration-Test Roadmap

This document tracks everything required to *prove* in CI that the game works and is balanceable without manual play-testing.  It assumes the engine stays renderer-agnostic and that future renderers (PixiJS, Three.js, etc.) can be swapped in by writing a new **RenderAdapter**.

---

## 1  Testing Runtime & Toolchain

- [x] **Adopt Vitest + jsdom as the primary test runner**  
      *Add `vitest` and `@vitest/ui` dev-deps; extend `vite.config.ts` with a `test` block.*  
      See Vitest config guide for canonical setup :contentReference[oaicite:0]{index=0}
- [x] **Polyfill browser APIs once**  
      *Install `jsdom`, `node-canvas` (2-D Canvas), and `headless-gl` (WebGL) and register them in `setupTests.ts`.*  
      Phaser & Pixi both compile against these shims :contentReference[oaicite:1]{index=1}
- [x] **Mock non-deterministic or heavy subsystems**  
      Use `vi.mock()` for loaders, audio and `Math.random`-driven helpers so frame-to-frame snapshots remain stable :contentReference[oaicite:2]{index=2}
- [ ] **Enable branch-level coverage (`c8`)** and fail PRs when < 90 % on `client/src/engine/**`
      The GitHub Action *vitest-coverage-report* emits PR comments and badges :contentReference[oaicite:3]{index=3}

---

## 2  Engine / Renderer Contract

- [x] **Extract an `IRenderAdapter` interface** if not already done.  
      Methods: `init(width,height)`, `render(state)`, `destroy()`.  
      The engine must depend only on this contract, never on Phaser classes.  
- [x] **NullRenderAdapter**: no-ops every call; used in all integration tests.  
- [ ] **Plugin detection**: pick adapter from `process.env.RENDER_BACKEND || 'phaser'` at runtime.  

---

## 3  Headless-Only Integration Suite (`client/__tests__/integration/`)  

> *Every test instantiates **HeadlessGameEngine** with `NullRenderAdapter`, steps the
simulation, injects input, then asserts against `engine.getState()`.*

| ID | Category | What we prove | Key APIs |
|----|----------|---------------|----------|
| IT-01 | Boot-strap | `reset()` yields pristine state, no leaks across tests | `reset` |
| IT-02 | Wave flow | After N `step()` calls spawn ≥ expected mobs, wave counter increments | `getState().spawners` |
| IT-03 | Input kill | `injectInput(word)` removes correct mob, awards score & combo | `injectInput`, `getState().player` |
| IT-04 | Lose / Win | Deplete `player.health` → `gameStatus='gameOver'`; defeat boss → `gameStatus='victory'` | `step`, `getState` |
| IT-05 | Progression | Simulate finishing L-1-3 and assert L-1-4 unlock flag toggles | `StateManager` |
| IT-06 | Dynamic diff. | Parameterised test (`it.each`) over spawn-rate & word-length tables; assert TTK within 5 % margin | `spawnConfig`, `it.each` :contentReference[oaicite:4]{index=4} |
| IT-07 | Edge lists | Empty word list or 1-char words do **not** crash engine | word-fixtures |
| IT-08 | Extreme fps | `step(1)` in a 10 000-iteration loop finishes in < 200 ms on CI box | performance |
| IT-09 | Save/load | Serialise + deserialise state; content round-trips equal | `StateManager.save/load` |
| IT-10 | Regression snapshots (optional) | Serialise minimal JSON of `gameState` every 100 frames and `expect(...).toMatchSnapshot()` | Vitest snapshots |

---

## 4  Alternative Renderer Experiments (spike stories)  

> *These do not block CI but give you confidence that the adapter boundary is solid.*

- [ ] **Phaser HEADLESS** with `@geckos.io/phaser-on-nodejs` for physics stress tests :contentReference[oaicite:5]{index=5}  
- [ ] **PixiJS + node-canvas** demo scene (render 100 sprites, export PNG) :contentReference[oaicite:6]{index=6}  
- [ ] **Three.js + headless-gl** proof: render coloured cubes, assert pixel buffer average ≠ 0 :contentReference[oaicite:7]{index=7}  
- [ ] Record startup time & memory, compare to Phaser; decide long-term renderer.

---

## 5  Continuous Integration  

- [ ] `.github/workflows/ci.yml`  

  ```yaml
  steps:
    - uses: actions/checkout@v4
    - uses: pnpm/action-setup@v3
    - run: pnpm install --frozen-lockfile
    - run: pnpm run test -- --run
