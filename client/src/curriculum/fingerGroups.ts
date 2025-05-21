/**
 * Defines the mapping of fingers to keys for the typing curriculum.
 * This structure is used to organize levels and track player progress.
 */

export enum FingerType {
    LEFT_INDEX = 'LEFT_INDEX',
    RIGHT_INDEX = 'RIGHT_INDEX',
    LEFT_MIDDLE = 'LEFT_MIDDLE',
    RIGHT_MIDDLE = 'RIGHT_MIDDLE',
    LEFT_RING = 'LEFT_RING',
    RIGHT_RING = 'RIGHT_RING',
    LEFT_PINKY = 'LEFT_PINKY',
    RIGHT_PINKY = 'RIGHT_PINKY',
    THUMB = 'THUMB'
}

export enum RowType {
    HOME_ROW = 'HOME_ROW',
    TOP_ROW = 'TOP_ROW',
    BOTTOM_ROW = 'BOTTOM_ROW'
}

export interface KeyMapping {
    key: string;
    finger: FingerType;
    row: RowType;
    displayName?: string;
}

export const FINGER_GROUP_KEYS: Record<string, KeyMapping[]> = {
    INDEX_FINGERS: [
        { key: 'f', finger: FingerType.LEFT_INDEX, row: RowType.HOME_ROW },
        { key: 'g', finger: FingerType.LEFT_INDEX, row: RowType.HOME_ROW },
        { key: 'r', finger: FingerType.LEFT_INDEX, row: RowType.TOP_ROW },
        { key: 't', finger: FingerType.LEFT_INDEX, row: RowType.TOP_ROW },
        { key: 'v', finger: FingerType.LEFT_INDEX, row: RowType.BOTTOM_ROW },
        { key: 'b', finger: FingerType.LEFT_INDEX, row: RowType.BOTTOM_ROW },
        { key: 'j', finger: FingerType.RIGHT_INDEX, row: RowType.HOME_ROW },
        { key: 'h', finger: FingerType.RIGHT_INDEX, row: RowType.HOME_ROW },
        { key: 'y', finger: FingerType.RIGHT_INDEX, row: RowType.TOP_ROW },
        { key: 'u', finger: FingerType.RIGHT_INDEX, row: RowType.TOP_ROW },
        { key: 'n', finger: FingerType.RIGHT_INDEX, row: RowType.BOTTOM_ROW },
        { key: 'm', finger: FingerType.RIGHT_INDEX, row: RowType.BOTTOM_ROW }
    ],
    MIDDLE_FINGERS: [
        { key: 'd', finger: FingerType.LEFT_MIDDLE, row: RowType.HOME_ROW },
        { key: 'e', finger: FingerType.LEFT_MIDDLE, row: RowType.TOP_ROW },
        { key: 'c', finger: FingerType.LEFT_MIDDLE, row: RowType.BOTTOM_ROW },
        { key: 'k', finger: FingerType.RIGHT_MIDDLE, row: RowType.HOME_ROW },
        { key: 'i', finger: FingerType.RIGHT_MIDDLE, row: RowType.TOP_ROW },
        { key: ',', finger: FingerType.RIGHT_MIDDLE, row: RowType.BOTTOM_ROW, displayName: 'comma' }
    ],
    RING_FINGERS: [
        { key: 's', finger: FingerType.LEFT_RING, row: RowType.HOME_ROW },
        { key: 'w', finger: FingerType.LEFT_RING, row: RowType.TOP_ROW },
        { key: 'x', finger: FingerType.LEFT_RING, row: RowType.BOTTOM_ROW },
        { key: 'l', finger: FingerType.RIGHT_RING, row: RowType.HOME_ROW },
        { key: 'o', finger: FingerType.RIGHT_RING, row: RowType.TOP_ROW },
        { key: '.', finger: FingerType.RIGHT_RING, row: RowType.BOTTOM_ROW, displayName: 'period' }
    ],
    PINKY_FINGERS: [
        { key: 'a', finger: FingerType.LEFT_PINKY, row: RowType.HOME_ROW },
        { key: 'q', finger: FingerType.LEFT_PINKY, row: RowType.TOP_ROW },
        { key: 'z', finger: FingerType.LEFT_PINKY, row: RowType.BOTTOM_ROW },
        { key: ';', finger: FingerType.RIGHT_PINKY, row: RowType.HOME_ROW, displayName: 'semicolon' },
        { key: 'p', finger: FingerType.RIGHT_PINKY, row: RowType.TOP_ROW },
        { key: '/', finger: FingerType.RIGHT_PINKY, row: RowType.BOTTOM_ROW, displayName: 'slash' }
    ],
    THUMBS: [
        { key: ' ', finger: FingerType.THUMB, row: RowType.HOME_ROW, displayName: 'space' }
    ]
};

// Helper function to get all available keys for a specific world
export function getKeysForWorld(worldNumber: number): string[] {
    const keys: string[] = [];

    // Add keys based on world progression
    if (worldNumber >= 1) {
        keys.push(...FINGER_GROUP_KEYS.INDEX_FINGERS.map(k => k.key));
    }
    if (worldNumber >= 2) {
        keys.push(...FINGER_GROUP_KEYS.MIDDLE_FINGERS.map(k => k.key));
    }
    if (worldNumber >= 3) {
        keys.push(...FINGER_GROUP_KEYS.RING_FINGERS.map(k => k.key));
    }
    if (worldNumber >= 4) {
        keys.push(...FINGER_GROUP_KEYS.PINKY_FINGERS.map(k => k.key));
    }

    return keys;
}

// Get finger and row information for a specific key
export function getKeyInfo(key: string): KeyMapping | undefined {
    const lowerKey = key.toLowerCase();
    for (const group of Object.values(FINGER_GROUP_KEYS)) {
        const keyInfo = group.find(k => k.key === lowerKey);
        if (keyInfo) {
            return keyInfo;
        }
    }
    return undefined;
}

//Contains AI - generated edits.
