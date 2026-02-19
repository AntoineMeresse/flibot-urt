# Flibot Urban Terror

Administration bot for Urban Terror servers.
Developped mainly for Urban Terror by Antoine Méresse (IG: Flirow)

## Configuration

Configuration can be provided via `config.yml` (placed next to the binary or in `/app`) or environment variables. Environment variables take precedence over the config file.

| Key               | Env               | Description                                         | Default            |
|-------------------|-------------------|-----------------------------------------------------|--------------------|
| `serverip`        | `serverip`        | Server IP address                                   | —                  |
| `serverport`      | `serverport`      | Server port                                         | `27960`            |
| `password`        | `password`        | Rcon password                                       | —                  |
| `urtPath`         | `urtPath`         | Path to the UrbanTerror installation                | —                  |
| `logFilePath`     | `logFilePath`     | Path to the server log file                         | —                  |
| `botWorkerNumber` | `botWorkerNumber` | Number of log workers                               | `1`                |
| `dbUri`           | `dbUri`           | PostgreSQL connection URI                           | —                  |
| `urtRepo`         | `urtRepo`         | Map repository base URL                             | —                  |
| `ujmUrl`          | `ujmUrl`          | UJM API base URL                                    | —                  |
| `ujmApiKey`       | `ujmApiKey`       | UJM API key                                         | —                  |
| `discordWebhook`  | `discordWebhook`  | Discord webhook URL for notifications               | —                  |
| `resetOptions`    | `resetOptions`    | Rcon commands applied on every map load (see below) | see defaults below |

### resetOptions

A list of rcon commands sent automatically each time a map loads, before any map-specific options.

In `config.yml` (YAML list or comma-separated string, spaces after commas are fine):
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
resetOptions="sv_fps 125,g_maxGameClients 0,g_oldtriggers 0,g_gear QS,g_allownoclip 1,g_flagreturntime 0,g_nodamage 1,g_novest 1"
```

## Map options

Per-map rcon commands stored in the database, applied after `resetOptions` on every map load.

| Command                            | Level | Description                                        |
|------------------------------------|-------|----------------------------------------------------|
| `!mapoptions`                      | 0     | Display options set for the current map            |
| `!setmapoptions [opt1, opt2, ...]` | 80    | Save options for the current map (comma-separated) |
| `!removemapoptions`                | 80    | Delete options for the current map                 |

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
Stores `["g_stamina 2", "g_overbounces 0", "g_walljumps 5"]` and applies them when the map loads.

## Subprojects

- tail: https://github.com/nxadm/tail
- natural: https://github.com/maruel/natural
- quake3-rcon: https://github.com/AntoineMeresse/quake3-rcon-go