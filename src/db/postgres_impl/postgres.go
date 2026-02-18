package postgres_impl

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	mydb "github.com/AntoineMeresse/flibot-urt/src/db"
	postgres_genererated "github.com/AntoineMeresse/flibot-urt/src/db/postgres_impl/generated"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const (
	dbTimeout = 2
)

type PostGresqlDB struct {
	ctx     context.Context
	conn    *pgx.Conn
	queries *postgres_genererated.Queries
}

func InitPostGresqlDb(ctx context.Context, uri string) (*PostGresqlDB, error) {
	conn, err := pgx.Connect(ctx, uri)
	if err != nil {
		return nil, err
	}

	queries := postgres_genererated.New(conn)
	schema := mydb.ReadSchema("./sqlc/postgres/schema.sql")
	_, err = conn.Exec(ctx, schema)

	if err != nil {
		return nil, err
	}

	slog.Debug("Schema loaded", "schema", schema)

	return &PostGresqlDB{conn: conn, queries: queries, ctx: ctx}, nil
}

func (db *PostGresqlDB) Close() {
	err := db.conn.Close(db.ctx)
	if err != nil {
		slog.Error("Error trying to close postgres connection", "err", err)
	}
}

func (db *PostGresqlDB) SaveNewPlayer(name string, guid string, ipAddress string) (int, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	p, err := db.queries.CreatePlayer(c, postgres_genererated.CreatePlayerParams{
		Name:       name,
		Guid:       guid,
		IpAddress:  ipAddress,
		TimeJoined: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return 0, err
	}
	slog.Debug("Player created", "player", p)
	return int(p.ID), nil
}

func (db *PostGresqlDB) UpdatePlayer() error {
	return fmt.Errorf("To implement")
}

func (db *PostGresqlDB) PenAdd(guid string, size float64) error {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	pen, err := db.queries.CreatePen(c, postgres_genererated.CreatePenParams{
		Guid: guid,
		Size: size,
		Date: pgtype.Date{Time: time.Now(), Valid: true},
	})
	slog.Debug("Pen created", "pen", pen)
	return err
}

func (db *PostGresqlDB) PenPenOfTheDay() (string, []mydb.PenData, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	potd, err := db.queries.GetAllPenByDate(c, postgres_genererated.GetAllPenByDateParams{
		Date:  pgtype.Date{Time: time.Now(), Valid: true},
		Limit: 50,
	})

	if err != nil {
		return utils.GetTodayDateFormated(), []mydb.PenData{}, nil
	}

	slog.Debug("PenOfTheDay", "potd", potd)
	res := make([]mydb.PenData, 0, len(potd))
	for _, v := range potd {
		res = append(res, mydb.PenData{Name: sql.NullString{String: v.Name, Valid: true}, Size: v.Size})
	}
	return utils.GetTodayDateFormated(), res, nil
}

func (db *PostGresqlDB) PenPenHallOfFame() ([]mydb.PenData, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	phof, err := db.queries.GetPensOrderBySizeDesc(c, 20)

	if err != nil {
		return []mydb.PenData{}, nil
	}

	slog.Debug("PenHallOfFame", "phof", phof)
	res := make([]mydb.PenData, 0, len(phof))
	for _, v := range phof {
		res = append(res, mydb.PenData{Name: sql.NullString{String: v.Name, Valid: true}, Size: v.Size, Date: v.Date.Time})
	}
	return res, nil
}

func (db *PostGresqlDB) PenPenHallOfShame() ([]mydb.PenData, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	phos, err := db.queries.GetPensOrderBySizeAsc(c, 20)

	if err != nil {
		return []mydb.PenData{}, nil
	}

	slog.Debug("PenHallOfShame", "phos", phos)
	res := make([]mydb.PenData, 0, len(phos))
	for _, v := range phos {
		res = append(res, mydb.PenData{Name: sql.NullString{String: v.Name, Valid: true}, Size: v.Size, Date: v.Date.Time})
	}
	return res, nil
}

func (db *PostGresqlDB) PenPlayerGetDailySize(guid string) (float64, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()

	return db.queries.GetPlayerPenByDate(c, pgtype.Date{Time: time.Now(), Valid: true})
}

func (db *PostGresqlDB) HandleRun(info models.PlayerRunInfo, checkpoints []int) error {
	slog.Debug("HandleRun", "info", info, "checkpoints", checkpoints)
	runtime, err := strconv.Atoi(info.Time)
	if err != nil {
		return err
	}

	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()

	previousTime, err := db.queries.GetRuntimeByMapWayUTJ(c, postgres_genererated.GetRuntimeByMapWayUTJParams{
		Guid:    info.Guid,
		Mapname: info.Mapname,
		Way:     info.Way,
		Utj:     info.Utj,
	})

	if err == nil {
		timeDiff := int(previousTime) - runtime
		slog.Debug("HandleRun time diff", "diff_ms", timeDiff)

		if timeDiff > 0 {
			if err = db.queries.UpdateRunByGuidAndUTJ(c, postgres_genererated.UpdateRunByGuidAndUTJParams{
				Runtime:     int32(runtime),
				Checkpoints: fmt.Sprintf("%v", checkpoints),
				RunDate:     pgtype.Timestamp{Time: time.Now(), Valid: true},
				Guid:        info.Guid,
				Utj:         info.Utj,
			}); err != nil {
				return err
			}
			slog.Debug("HandleRun: updated time", "runtime", runtime, "guid", info.Guid)
		} else {
			slog.Debug("HandleRun: not an improvement")
		}
	} else {
		slog.Debug("HandleRun: no existing run, creating new entry")
		if err = db.queries.CreateRun(c, postgres_genererated.CreateRunParams{
			Guid:        info.Guid,
			Utj:         info.Utj,
			Mapname:     info.Mapname,
			Way:         info.Way,
			Runtime:     int32(runtime),
			Checkpoints: fmt.Sprintf("%v", checkpoints),
			RunDate:     pgtype.Timestamp{Time: time.Now(), Valid: true},
			Demopath:    info.Demopath,
		}); err != nil {
			return err
		}
		slog.Debug("HandleRun: created new entry")
	}
	return nil
}

func (db *PostGresqlDB) GetPlayerByGuid(guid string) (models.Player, bool) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()

	if playerDb, err := db.queries.GetPLayerByGuid(c, guid); err != nil {
		slog.Error("GetPlayerByGuid error", "err", err)
		return models.Player{}, false
	} else {
		slog.Debug("Player found in db", "player", playerDb)
		return models.Player{
			Role: int(playerDb.Role),
			Name: playerDb.Name,
			Guid: guid,
			Id:   string(playerDb.ID),
			// Aliases: playerDb.Aliases,
		}, true
	}
}
