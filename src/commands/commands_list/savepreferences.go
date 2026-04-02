package commandslist

import (
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
)

func SavePreferences(cmd *appcontext.CommandsArgs) {
	player, err := cmd.Context.Players.GetPlayer(cmd.PlayerId)
	if err != nil {
		cmd.RconText("^1Could not find your player info.")
		return
	}

	if len(cmd.Params) == 0 {
		saved, found, _ := cmd.Context.DB.GetPreferences(player.Guid)
		if !found || len(saved) == 0 {
			cmd.RconText("^7No preferences saved.")
		} else {
			cmd.RconText("^7Saved preferences: ^5!%s", strings.Join(saved, "^7, ^5!"))
		}
		return
	}

	chunks := parseCommandChunks(cmd.Params, cmd.CommandExists)
	if len(chunks) == 0 {
		cmd.RconUsage()
		return
	}

	if err := cmd.Context.DB.UpsertPreferences(player.Guid, chunks); err != nil {
		cmd.RconText("^1Error saving preferences: %s", err.Error())
		return
	}
	cmd.RconText("^7Preferences saved: ^5!%s", strings.Join(chunks, "^7, ^5!"))
}

// parseCommandChunks splits params by '!' prefix into full command strings.
// e.g. ["!lo", "!say", "abc"] → ["lo", "say abc"]
// A token starting with '!' that is not a known command is treated as a param of the current chunk.
func parseCommandChunks(params []string, isKnownCmd func(string) bool) []string {
	var chunks []string
	current := ""
	for _, p := range params {
		if strings.HasPrefix(p, "!") {
			name := strings.TrimPrefix(p, "!")
			if isKnownCmd != nil && isKnownCmd(name) {
				if current != "" {
					chunks = append(chunks, current)
				}
				current = name
			} else if current != "" {
				current += " " + p
			}
		} else if current != "" {
			current += " " + p
		}
	}
	if current != "" {
		chunks = append(chunks, current)
	}
	return chunks
}
