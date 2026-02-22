package models

import (
	"strings"
	"testing"
)

func TestPlayers_GetPlayer(t *testing.T) {
	players := &Players{
		PlayerMap: map[string]*Player{
			"2": {Number: "2", Name: "Antoine"},
			"5": {Number: "5", Name: "ply2"},
		},
	}

	t.Run("Search by exact number", func(t *testing.T) {
		p, err := players.GetPlayer("2")
		if err != nil {
			t.Errorf("Expected to find player, got error: %v", err)
			return
		}
		if p.Number != "2" {
			t.Errorf("Expected player number 2, got %s", p.Number)
		}
	})

	t.Run("Ambiguous numeric search", func(t *testing.T) {
		// Add a player with "2" in their name
		players.PlayerMap["7"] = &Player{Number: "7", Name: "Antoine2"}
		
		p, err := players.GetPlayer("2")
		if err != nil {
			t.Errorf("Expected only one match (by number), but got error: %v", err)
		}
		if p != nil && p.Number != "2" {
			t.Errorf("Expected player number 2, got %s", p.Number)
		}
	})

	t.Run("Search by name containing number", func(t *testing.T) {
		p, err := players.GetPlayer("ply2")
		if err != nil {
			t.Errorf("Expected to find player, got error: %v", err)
			return
		}
		if p.Name != "ply2" {
			t.Errorf("Expected player name ply2, got %s", p.Name)
		}
	})

	t.Run("Ambiguous name search", func(t *testing.T) {
		players.PlayerMap["6"] = &Player{Number: "6", Name: "super-ply2"}
		
		_, err := players.GetPlayer("ply2")
		if err == nil {
			t.Errorf("Expected error (multiple match), got nil")
		} else if !strings.Contains(err.Error(), "multiple match") {
			t.Errorf("Expected 'multiple match' error, got: %v", err)
		}
	})
}
