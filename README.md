# flibot-urt

A bot for Urban Terror jump servers. Reads the server log file, parses events and executes commands via RCON.

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

Copy `config.yml` and fill in the values:

```yaml
serverip: 127.0.0.1
serverport: 27960
password: rconpassword
urtPath: /path/to/UrbanTerror43
logFilePath: /path/to/games.log
botWorkerNumber: 1

dbUri: postgresql://user:password@host:5432/dbname
schemaPath: ./sqlc/postgres/schema.sql   # optional, default: ./sqlc/postgres/schema.sql

# API
urtRepo: https://your-repo/q3ut4
ujmUrl: https://urtjumpmaps.com
ujmApiKey: YOUR_KEY

# Bridge (multi-server messaging)
bridgeUrl: https://your-bridge
bridgeApiKey: ""
channelId: 0
serverName: "Server"

# Translation (libretranslate)
translateUrl: http://your-host:5000
translateLangs: fr,en,es,it,de       # languages loaded in your libretranslate instance

# Misc
welcomeMessage: "Welcome back ^5{name}^3 [{id}]."
dailyPbPenCoinLimit: 2
```

### Environment variables

All config keys can be overridden via environment variables (same name).

Useful for Docker:
```yaml
environment:
  - dbUri=postgresql://...
  - translateUrl=http://libretranslate:5000
  - translateLangs=fr,en,es,it,de
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

| Level | Description |
|-------|-------------|
| 0 | Everyone |
| 20 | Registered |
| 40 | Moderator |
| 60 | Admin |
| 80 | Senior admin |
| 90 | Head admin |
| 100 | Super admin |

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
| `!trad [message]` | Toggle auto-translation / translate a message to English |
| `!tradto [lang\|src->tgt] [message]` | Translate a message to a specific language |
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

## Translation feature (`!trad` / `!tradto`)

- `!trad` — toggle auto-translation ON/OFF for yourself. When ON, every non-English message you send is automatically translated to English and sent as a PM to all players who also have `!trad` enabled.
- `!trad [message]` — manually translate a message to English and broadcast it globally.
- `!tradto [lang] [message]` — translate a message to the target language and broadcast globally (e.g. `!tradto fr Hello everyone`).
- `!tradto [src->tgt] [message]` — explicit source and target (e.g. `!tradto en->it This is a test`). No confidence check when source is explicit.

Requires a running [LibreTranslate](https://github.com/LibreTranslate/LibreTranslate) instance configured via `translateUrl`.  
Auto-translation is silently skipped if confidence < 70%. `!tradto` with auto-detect skips if confidence < 40%.

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
