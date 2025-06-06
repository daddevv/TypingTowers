package event

// Event is a generic interface for all events.
type Event interface{}

// EntityEvent represents events related to entities (e.g., spawn, death).
type EntityEvent struct {
	Type    string
	Payload interface{}
}

// UIEvent represents UI-related events (e.g., notification, panel toggle).
type UIEvent struct {
	Type    string
	Payload interface{}
}

// TechEvent represents tech tree unlocks or changes.
type TechEvent struct {
	Type    string
	Payload interface{}
}

// TowerEvent represents tower-specific events (e.g., upgrade, fire).
type TowerEvent struct {
	Type    string
	Payload interface{}
}

// PhaseEvent represents game phase/state transitions.
type PhaseEvent struct {
	Type    string
	Payload interface{}
}

// EconEvent represents economy-related events.
type EconEvent struct {
	Type    string
	Payload interface{}
}

// SpriteEvent represents sprite/image-related events.
type SpriteEvent struct {
	Type    string
	Payload interface{}
}

// Each handler should expose a channel for its event type for pub/sub (T-007).
// Example: EntityHandler exposes EntityEvents chan Event, etc.

// Example: Add more event types as needed for other modules.

// -----------------------------------------------------------------------------
// Event name constants

// Building events cover construction, upgrades and cooldowns.
type BuildingEventType string

const (
	BuildingPlaced    BuildingEventType = "BuildingPlaced"
	BuildingDestroyed BuildingEventType = "BuildingDestroyed"
	BuildingPaused    BuildingEventType = "BuildingPaused"
	BuildingUnpaused  BuildingEventType = "BuildingUnpaused"
	BuildingLeveledUp BuildingEventType = "BuildingLeveledUp"
	CooldownComplete  BuildingEventType = "CooldownComplete"
)

// Typing events track player input and global queue status.
type TypingEventType string

const (
	WordQueued           TypingEventType = "WordQueued"
	WordLetterTyped      TypingEventType = "WordLetterTyped"
	WordMistyped         TypingEventType = "WordMistyped"
	WordBackspaced       TypingEventType = "WordBackspaced"
	WordCompleted        TypingEventType = "WordCompleted"
	QueueBacklogOverload TypingEventType = "QueueBacklogOverload"
)

// Combat events represent unit actions and damage resolution.
type CombatEventType string

const (
	UnitSpawned        CombatEventType = "UnitSpawned"
	UnitMoved          CombatEventType = "UnitMoved"
	UnitTargetAcquired CombatEventType = "UnitTargetAcquired"
	UnitTakesDamage    CombatEventType = "UnitTakesDamage"
	UnitDied           CombatEventType = "UnitDied"
	CombatTick         CombatEventType = "CombatTick"
	CritHit            CombatEventType = "CritHit"
)

// Tech events are emitted from the tech and skill tree systems.
type TechEventType string

const (
	TechNodeUnlocked TechEventType = "TechNodeUnlocked"
	TechNodeFailed   TechEventType = "TechNodeFailed"
	SkillUnlocked    TechEventType = "SkillUnlocked"
	LetterUnlocked   TechEventType = "LetterUnlocked"
)

// Econ events deal with resource generation and spending.
type EconomyEventType string

const (
	ResourceGained    EconomyEventType = "ResourceGained"
	ResourceSpent     EconomyEventType = "ResourceSpent"
	KingPointsChanged EconomyEventType = "KingPointsChanged"
)

// Special events are miscellaneous gameplay triggers.
type SpecialEventType string

const (
	SpellCast      SpecialEventType = "SpellCast"
	CritTriggered  SpecialEventType = "CritTriggered"
	EventTriggered SpecialEventType = "EventTriggered"
)

// Testing events are used by automated or headless runs.
type TestingEventType string

const (
	SimulationStart        TestingEventType = "SimulationStart"
	SimulationStepComplete TestingEventType = "SimulationStepComplete"
	GameStateSnapshot      TestingEventType = "GameStateSnapshot"
)

// Persistence events are fired when saving/loading game state.
type PersistenceEventType string

const (
	GameSaved       PersistenceEventType = "GameSaved"
	GameLoaded      PersistenceEventType = "GameLoaded"
	SaveSlotCorrupt PersistenceEventType = "SaveSlotCorrupt"
)
