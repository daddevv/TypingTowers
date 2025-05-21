/**
 * Manages level progression and tracking player progress through worlds.
 */

export interface LevelProgress {
    completed: boolean;
    highScore: number;
    bestWPM: number;
    bestAccuracy: number;
    attempts: number;
}

export default class LevelManager {
    private currentWorldId: number = 1;
    private currentLevelId: string = '1-1';
    private progress: Map<string, LevelProgress> = new Map();
    private storageKey: string = 'type-defense-progress';

    constructor() {
        this.loadProgress();
    }

    public getCurrentWorld(): number {
        return this.currentWorldId;
    }

    public getCurrentLevel(): string {
        return this.currentLevelId;
    }

    public setCurrentLevel(worldId: number, levelId: string) {
        this.currentWorldId = worldId;
        this.currentLevelId = levelId;
    }

    public getLevelProgress(levelId: string): LevelProgress | undefined {
        return this.progress.get(levelId);
    }

    public setLevelProgress(levelId: string, progress: Partial<LevelProgress>) {
        const prev = this.progress.get(levelId) || {
            completed: false,
            highScore: 0,
            bestWPM: 0,
            bestAccuracy: 0,
            attempts: 0,
        };
        this.progress.set(levelId, { ...prev, ...progress });
        this.saveProgress();
    }

    public isLevelUnlocked(levelId: string): boolean {
        // Unlock first level by default
        if (levelId === '1-1') return true;
        // Unlock if previous level is completed
        const [world, level] = levelId.split('-').map(Number);
        if (level > 1) {
            const prevLevelId = `${world}-${level - 1}`;
            return this.progress.get(prevLevelId)?.completed ?? false;
        } else if (world > 1) {
            const prevWorldLastLevel = `${world - 1}-3`;
            return this.progress.get(prevWorldLastLevel)?.completed ?? false;
        }
        return false;
    }

    public completeLevel(levelId: string, stats: { score: number; wpm: number; accuracy: number; }) {
        const prev = this.progress.get(levelId) || {
            completed: false,
            highScore: 0,
            bestWPM: 0,
            bestAccuracy: 0,
            attempts: 0,
        };
        this.progress.set(levelId, {
            completed: true,
            highScore: Math.max(prev.highScore, stats.score),
            bestWPM: Math.max(prev.bestWPM, stats.wpm),
            bestAccuracy: Math.max(prev.bestAccuracy, stats.accuracy),
            attempts: prev.attempts + 1,
        });
        this.saveProgress();
    }

    public resetProgress() {
        this.progress.clear();
        this.saveProgress();
    }

    private saveProgress() {
        const obj: Record<string, LevelProgress> = {};
        this.progress.forEach((v, k) => { obj[k] = v; });
        localStorage.setItem(this.storageKey, JSON.stringify(obj));
    }

    public loadProgress() {
        const raw = localStorage.getItem(this.storageKey);
        if (raw) {
            try {
                const obj = JSON.parse(raw);
                this.progress = new Map(Object.entries(obj));
            } catch {
                this.progress = new Map();
            }
        } else {
            this.progress = new Map();
        }
    }

    public getAllProgress(): Map<string, LevelProgress> {
        return this.progress;
    }
}

export const levelManager = new LevelManager();
