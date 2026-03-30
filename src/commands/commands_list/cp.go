package commandslist

import (
	"fmt"
	"strconv"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	"github.com/AntoineMeresse/flibot-urt/src/models"
)

const cpListLimit = 10
const cpMaxCompare = 3

func fetchTargets(cmd *appcontext.CommandsArgs, way string, limit int) []models.CompareTarget {
	rows, err := cmd.Context.DB.GetTopCheckpoints(cmd.Context.GetCurrentMap(), way, limit)
	if err != nil {
		return nil
	}
	targets := make([]models.CompareTarget, 0, len(rows))
	for _, row := range rows {
		targets = append(targets, models.CompareTarget{
			Name:        row.Name,
			Runtime:     row.Runtime,
			Checkpoints: row.Checkpoints,
		})
	}
	return targets
}

func Compare(cmd *appcontext.CommandsArgs) {
	if len(cmd.Params) == 0 {
		enabled := cmd.Context.Runs.ToggleCp(cmd.PlayerId)
		if enabled {
			if way, running := cmd.Context.Runs.GetCurrentWay(cmd.PlayerId); running {
				cmd.RconText("^7Compare checkpoints: ^2On ^7(^3!cp -list^7 for way (^5%s^7) checkpoints)", way)
			} else {
				cmd.RconText("^7Compare checkpoints: ^2On ^7(^3!cp -list^7 to see checkpoints)")
			}
		} else {
			cmd.RconText("^7Compare checkpoints: ^1Off")
		}
		return
	}

	if cmd.Params[0] == "-list" || cmd.Params[0] == "-l" {
		way, running := cmd.Context.Runs.GetCurrentWay(cmd.PlayerId)
		if !running {
			way = "1"
		}
		targets := fetchTargets(cmd, way, cpListLimit)
		if len(targets) == 0 {
			cmd.RconText("^7No times found for way ^3%s", way)
			return
		}
		currentIdxs := cmd.Context.Runs.GetCompareIdxs(cmd.PlayerId)
		selectedSet := make(map[int]bool, len(currentIdxs))
		for _, i := range currentIdxs {
			selectedSet[i] = true
		}
		if running {
			cmd.Context.Runs.UpdateCompareTargets(cmd.PlayerId, targets, currentIdxs)
		}
		for i, t := range targets {
			marker := "  "
			if selectedSet[i] {
				marker = "^2> "
			}
			cmd.RconText("%s^7[^3%d^7] ^8%s^7 - %s", marker, i+1, t.Name, models.FormatMs(t.Runtime))
		}
		return
	}

	// parse all params as rank indices
	if len(cmd.Params) > cpMaxCompare {
		cmd.RconText("^1Max %d comparison targets.", cpMaxCompare)
		return
	}
	idxs := make([]int, 0, len(cmd.Params))
	for _, p := range cmd.Params {
		n, err := strconv.Atoi(p)
		if err != nil || n < 1 || n > cpListLimit {
			cmd.RconUsage()
			return
		}
		idxs = append(idxs, n-1) // convert to 0-based
	}

	cmd.Context.Runs.EnableCp(cmd.PlayerId)

	way, running := cmd.Context.Runs.GetCurrentWay(cmd.PlayerId)
	if !running {
		cmd.Context.Runs.UpdateCompareTargets(cmd.PlayerId, nil, idxs)
		parts := make([]string, len(idxs))
		for i, idx := range idxs {
			parts[i] = fmt.Sprintf("^3#%d^7", idx+1)
		}
		cmd.RconText("^7Compare checkpoints: ^2On ^7- saved: %s ^7(applied on next run)", strings.Join(parts, "^7, "))
		return
	}

	maxIdx := 0
	for _, idx := range idxs {
		if idx > maxIdx {
			maxIdx = idx
		}
	}
	targets := fetchTargets(cmd, way, maxIdx+1)
	// filter out indices that exceed available targets
	valid := make([]int, 0, len(idxs))
	for _, idx := range idxs {
		if idx < len(targets) {
			valid = append(valid, idx)
		}
	}
	if len(valid) == 0 {
		cmd.RconText("^1No runs found for the requested ranks.")
		return
	}
	cmd.Context.Runs.UpdateCompareTargets(cmd.PlayerId, targets, valid)
	for _, idx := range valid {
		t := targets[idx]
		cmd.RconText("^7[^3#%d^7] ^8%s^7 - %s", idx+1, t.Name, models.FormatMs(t.Runtime))
	}
}
