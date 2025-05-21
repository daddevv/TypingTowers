import { beforeEach, describe, expect, it } from 'vitest';

// Minimal Combo logic for unit test
class ComboSystem {
    combo = 0;
    score = 0;
    basePoints = 10;

    correctKeystroke() {
        this.combo++;
        this.score += this.basePoints * this.combo;
    }

    incorrectKeystroke() {
        this.combo = 0;
    }
}

describe('ComboSystem Unit', () => {
    let comboSys: ComboSystem;
    beforeEach(() => {
        comboSys = new ComboSystem();
    });

    it('starts with combo and score at 0', () => {
        expect(comboSys.combo).toBe(0);
        expect(comboSys.score).toBe(0);
    });

    it('increments combo and score on correct keystroke', () => {
        comboSys.correctKeystroke();
        expect(comboSys.combo).toBe(1);
        expect(comboSys.score).toBe(10);
        comboSys.correctKeystroke();
        expect(comboSys.combo).toBe(2);
        expect(comboSys.score).toBe(30); // 10 + 20
    });

    it('resets combo on incorrect keystroke', () => {
        comboSys.correctKeystroke();
        comboSys.correctKeystroke();
        comboSys.incorrectKeystroke();
        expect(comboSys.combo).toBe(0);
    });

    it('score calculation uses combo multiplier', () => {
        comboSys.combo = 2;
        comboSys.correctKeystroke();
        expect(comboSys.score).toBe(30); // 0 + 10*3
    });
});

// Contains AI-generated edits.
