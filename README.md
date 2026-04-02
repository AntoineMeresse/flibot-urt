# flibot-urt

Administration and jump bot for Urban Terror servers.  
Developed by Antoine Méresse (IG: Flirow)

Reads the server log file, parses events and executes commands via RCON.

---

## Architecture

- **Log parsing** — `src/logs/` tails the server log and dispatches events to `actions.HandleAction`
- **Actions** — `src/actions/actions_list/` one file per URT event (say, ClientConnect, ClientDisconnect, etc.)
- **Commands** — `src/commands/commands_list/` one file per bot command, registered in `command_map.go`
- **DB** — PostgreSQL via sqlc. Interface: `src/db/interface.go`
- **Context** — `src/context/context.go` holds DB, Rcon, Players, Settings, Runs, Api
- **Models** — `src/models/` Player, RunsInfo, UrtConfig, etc.

---

## Configuration

Configuration can be provided via `config.yml` (placed next to the binary or in `/app`) or environment variables. Environment variables take precedence over the config file.

| Key                   | Description                                          | Default                        |
|-----------------------|------------------------------------------------------|--------------------------------|
| `serverip`            | Server IP address                                    | —                              |
| `serverport`          | Server port                                          | `27960`                        |
| `password`            | Rcon password                                        | —                              |
| `urtPath`             | Path to the UrbanTerror installation                 | —                              |
| `logFilePath`         | Path to the server log file                          | —                              |
| `botWorkerNumber`     | Number of log workers                                | `1`                            |
| `dbUri`               | PostgreSQL connection URI                            | —                              |
| `schemaPath`          | Path to the SQL schema file                          | `./sqlc/postgres/schema.sql`   |
| `urtRepo`             | Map repository base URL                              | —                              |
| `ujmUrl`              | UJM API base URL                                     | —                              |
| `ujmApiKey`           | UJM API key                                          | —                              |
| `discordWebhook`      | Discord webhook URL for notifications                | —                              |
| `bridgeUrl`           | Bridge API URL (multi-server messaging)              | —                              |
| `bridgeApiKey`        | Bridge API key                                       | —                              |
| `channelId`           | Discord channel ID for this server                   | `0`                            |
| `serverName`          | Display name for this server                         | `Server`                       |
| `translateUrl`        | LibreTranslate instance URL                          | —                              |
| `translateLangs`      | Comma-separated list of supported languages          | `fr,en,es,it,de`               |
| `welcomeMessage`      | Welcome message (`{name}`, `{id}` placeholders)      | —                              |
| `dailyPbPenCoinLimit` | Max pencoin per day                                  | `2`                            |
| `resetOptions`        | Rcon commands applied on every map load              | see below                      |

### resetOptions

A list of rcon commands sent automatically each time a map loads, before any map-specific options.

In `config.yml`:
```yaml
resetOptions:
  - "sv_fps 125"
  - "g_maxGameClients 0"
  - "g_oldtriggers 0"
  - "g_gear QS"
  - "g_allownoclip 1"
  - "g_flagreturntime 0"
  - "g_nodamage 1"
  - "g_novest 1"
```

As an environment variable (comma-separated):
```
resetOptions="sv_fps 125,g_maxGameClients 0,g_oldtriggers 0"
```

### LibreTranslate (Docker)

```yaml
services:
  libretranslate:
    image: libretranslate/libretranslate:latest
    restart: unless-stopped
    ports:
      - "5000:5000"
    environment:
      - LT_LOAD_ONLY=fr,en,es,it,de
      - LT_API_KEYS=false
      - LT_UPDATE_MODELS=true
```

---

## Command levels

| Level | Description  |
|-------|--------------|
| 0     | Everyone     |
| 20    | Registered   |
| 40    | Moderator    |
| 60    | Admin        |
| 80    | Senior admin |
| 90    | Head admin   |
| 100   | Super admin  |

---

## Commands

### General (level 0)

| Command | Description |
|---------|-------------|
| `!play` / `!pl` | Join the game |
| `!spec` / `!sp` | Go spectator |
| `!ready` / `!start` / `!run` | Mark yourself ready |
| `!goto` | Teleport to a goto point |
| `!invisible` / `!invi` | Toggle invisibility |
| `!loadonce` / `!lo` | Toggle loadonce |
| `!map [search] [index] (-f)` | Change map |
| `!maps` | List available maps |
| `!currentmap` / `!current` | Show current map |
| `!setnextmap [search]` / `!snm` | Set next map |
| `!nextmap` | Show next map |
| `!cyclemap (-f)` / `!cycle` | Cycle to next map |
| `!roll` | Roll the dice |
| `!stamina` | Toggle stamina |
| `!callvote` / `!cv` | Call a vote |
| `!+` / `!-` | Vote yes / no |
| `!help` / `!h` | Show available commands |
| `!mapinfo` / `!mi` | Show map info |
| `!topruns` / `!tr` | Show top runs for current map |
| `!serverruns [map]` / `!sr` | Show server runs |
| `!ahead [map]` | Show players ahead of you |
| `!missing [maxlvl]` | Show missing maps |
| `!similar [map]` | Show similar maps |
| `!top` | Show top players |
| `!latestruns` / `!lr` | Show latest runs |
| `!latestmaps` / `!lm` | Show latest maps |
| `!coin` / `!flip` | Flip a coin |
| `!birthday` | Birthday info |
| `!bounties` | Show bounties |
| `!pen` | Pencoin info |
| `!potd` | Pen of the day |
| `!phof` | Pen hall of fame |
| `!phos` | Pen hall of shame |
| `!pb` | Personal best |
| `!ignore [player]` | Ignore a player |
| `!unignore [id]` | Unignore a player |
| `!compare [-list \| 1-10...]` / `!cp` | Checkpoint comparison |
| `!time` | Show current time |
| `!tp [player]` | Teleport to player |
| `!status [p/e]` | Server status |
| `!quit [reason]` / `!ff` | Quit the game |
| `!admins` | List online admins |
| `!level` / `!lvl` | Show your level |
| `!locate [player]` | Locate a player |
| `!website` / `!site` | Show website |
| `!discord` | Show Discord link |
| `!quote [id]` | Show a quote |
| `!findquote [text]` | Search quotes |
| `!all [message]` | Send message to all servers (bridge) |
| `!translate [message]` / `!trad` | Toggle auto-translation / translate a message to English |
| `!translateto [lang\|src->tgt] [message]` / `!tradto` | Translate a message to a specific language |
| `!extend [1-999]` | Extend timelimit |
| `!redo [extra params]` / `!lc` | Repeat last command |

### Moderator (level 40)

| Command | Description |
|---------|-------------|
| `!afk [player] (-f)` | Mark player AFK |
| `!nuke [player]` | Nuke a player |
| `!veto` | Veto current vote |

### Admin (level 60)

| Command | Description |
|---------|-------------|
| `!slap [player] (-f)` | Slap a player |
| `!setgoto` | Set goto point |
| `!removegoto` / `!rmgoto` | Remove goto point |
| `!force [player] [team] (-f)` | Force player to team |
| `!mapget` / `!dl` | Download a map |
| `!mapremove` / `!mremove` | Remove a map |
| `!timelimit [1-999]` | Set timelimit |
| `!putgroup [player] [-1..100]` / `!rights` | Set player rights |

### Senior admin (level 80)

| Command | Description |
|---------|-------------|
| `!kick [player] [reason] (-f)` | Kick a player |
| `!moveplayer [p1] [p2] (-f)` | Move player slot |
| `!mute [player]` | Mute a player |
| `!addquote [text]` | Add a quote |
| `!removequote [id]` | Remove a quote |
| `!portgotos [sourcemap]` | Port gotos from map |
| `!portmapoptions [sourcemap]` | Port map options |
| `!overbounces [0\|1]` / `!ob` | Toggle overbounces |
| `!setmapoptions [opts]` / `!smo` | Set map options |
| `!resetoptions` | Reset server options |
| `!mapoptions` / `!mo` | Show map options |
| `!removemapoptions` / `!rmmo` | Remove map options |

### Head admin (level 90)

| Command | Description |
|---------|-------------|
| `!ban [player/@dbId] [reason]` | Ban a player (by slot, name or DB id) |
| `!unban [@dbId\|banId]` | Unban a player |
| `!lookup [name/@id]` / `!l` | Lookup a player |
| `!resetgotos [map] (-f)` | Reset gotos for map |

### Super admin (level 100)

| Command | Description |
|---------|-------------|
| `!pencoin [player]` / `!pcoin` | Give pencoin |
| `!resetpencoin [amount]` | Reset pencoin |
| `!password [password]` | Set server password |
| `!rcon [command]` | Execute raw RCON command |
| `!playerid [id]` | Show player by id |
| `!globalpen` | Global pen action |

---

## Map options

Per-map rcon commands stored in the database, applied after `resetOptions` on every map load.

| Command | Level | Description |
|---------|-------|-------------|
| `!mapoptions` | 0 | Display options set for the current map |
| `!setmapoptions [opt1, opt2, ...]` | 80 | Save options for the current map (comma-separated) |
| `!removemapoptions` | 80 | Delete options for the current map |

Options support shortcuts:

| Shortcut | Expands to        |
|----------|-------------------|
| `fstam`  | `g_stamina 2`     |
| `nstam`  | `g_stamina 1`     |
| `noob`   | `g_overbounces 0` |
| `ob`     | `g_overbounces 1` |

Example:
```
!setmapoptions fstam, noob, g_walljumps 5
```

---

## Translation (`!translate` / `!translateto`)

- `!translate` / `!trad` — toggle auto-translation ON/OFF. When ON, every non-English message you send is automatically translated to English and sent as a PM to all players who also have it enabled.
- `!translate [message]` / `!trad [message]` — manually translate a message to English and broadcast globally.
- `!translateto [lang] [message]` / `!tradto` — translate to a target language and broadcast globally (e.g. `!translateto fr Hello`).
- `!translateto [src->tgt] [message]` — explicit source and target (e.g. `!translateto en->it This is a test`). No confidence check when source is explicit.

Translated text has accents stripped for compatibility with URT's chat.  
Auto-translation is silently skipped if confidence < 70%. `!translateto` with auto-detect skips if confidence < 40%.

Requires a running [LibreTranslate](https://github.com/LibreTranslate/LibreTranslate) instance.

---

## Checkpoint comparison (`!cp`)

- `!cp` — toggle checkpoint comparison ON/OFF.
- `!cp 1 2 3` — set comparison targets by rank (up to 3). Also activates CP.
- `!cp -list` / `!cp -l` — list available checkpoint runs for the current way.

When enabled, a PM is sent after each checkpoint with the time difference vs selected targets.

---

## Ban system

Bans are checked by **GUID or IP** on each connection. Both are stored at ban time.

- `!ban [slot/name]` — ban an online player.
- `!ban @[dbId]` — ban a player by database id (works offline).
- `!unban` — list current bans.
- `!unban @[dbId]` or `!unban [banId]` — unban.

---

## Run system

Tracks jump runs per player. On run start, loads best checkpoints from DB. On run stop, saves to DB if it's a new best.

Flow: `ClientJumpRunStarted` → `RunStart` → checkpoints → `ClientJumpRunStopped` → `RunStopped` → `RunLog` → DB save.

---

## Bridge

Multi-server messaging via an external bridge API.

- `!all [message]` — broadcast to all servers.
- Server info (map, players) is pushed every 10s to the bridge.
- Commands are forwarded to the bridge unless marked `excludeFromBridge`.

---

## Subprojects

- [tail](https://github.com/nxadm/tail)
- [quake3-rcon](https://github.com/AntoineMeresse/quake3-rcon-go)
- [LibreTranslate](https://github.com/LibreTranslate/LibreTranslate)
