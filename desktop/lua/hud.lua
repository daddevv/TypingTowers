-- Table-driven HUD configuration
HUD = {
    health = {
        x = 30,
        y = 900,
        font = "Game-Bold",
        fontSize = 32,
        color = {255, 80, 80, 255},
        format = "Health: %d"
    },
    stats = {
        x = 30,
        y = 940,
        font = "Game-Bold",
        fontSize = 24,
        color = {255, 255, 255, 255},
        fields = {
            {name = "Misses", key = "Misses"},
            {name = "Highest Streak", key = "HighestStreak"},
            {name = "Lucky Hits", key = "LuckyHits"}
        }
    }
}
