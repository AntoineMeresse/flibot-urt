package postgres_impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	mydb "github.com/AntoineMeresse/flibot-urt/src/db"
	postgres_genererated "github.com/AntoineMeresse/flibot-urt/src/db/postgres_impl/generated"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
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

	logrus.Debugf("Schema: %s", schema)

	return &PostGresqlDB{conn: conn, queries: queries, ctx: ctx}, nil
}

func (db *PostGresqlDB) Close() {
	err := db.conn.Close(db.ctx)
	if err != nil {
		logrus.Error("Error trying to close postgres connection", err)
	}
}

func (db *PostGresqlDB) SaveNewPlayer(name string, guid string, ipAddress string) (int, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	initialAliases, _ := json.Marshal([]string{name})
	p, err := db.queries.CreatePlayer(c, postgres_genererated.CreatePlayerParams{
		Name:       name,
		Guid:       guid,
		IpAddress:  ipAddress,
		TimeJoined: pgtype.Timestamp{Time: time.Now(), Valid: true},
		Aliases:    string(initialAliases),
	})
	if err != nil {
		return 0, err
	}
	logrus.Debugf("Player created: %v", p)
	return int(p.ID), nil
}

func (db *PostGresqlDB) UpdatePlayerOnJoin(guid, name, ip string, aliases []string) error {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	aliasesJSON, _ := json.Marshal(aliases)
	return db.queries.UpdatePlayerOnJoin(c, postgres_genererated.UpdatePlayerOnJoinParams{
		Guid:       guid,
		Name:       name,
		IpAddress:  ip,
		TimeJoined: pgtype.Timestamp{Time: time.Now(), Valid: true},
		Aliases:    string(aliasesJSON),
	})
}

func (db *PostGresqlDB) UpdatePlayer() error {
	return fmt.Errorf("To implement")
}

func (db *PostGresqlDB) SetPlayerRole(guid string, role int) error {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	return db.queries.SetPlayerRole(c, postgres_genererated.SetPlayerRoleParams{
		Guid: guid,
		Role: int32(role),
	})
}

func (db *PostGresqlDB) PenAdd(guid string, size float64) error {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	pen, err := db.queries.UpsertPen(c, postgres_genererated.UpsertPenParams{
		Guid: guid,
		Size: size,
		Date: pgtype.Date{Time: time.Now(), Valid: true},
	})
	logrus.Debugf("Pen upserted: %v", pen)
	return err
}

func (db *PostGresqlDB) PenGetAttempts(guid string) (int, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	count, err := db.queries.GetPenCounter(c, postgres_genererated.GetPenCounterParams{
		Guid: guid,
		Year: int32(time.Now().Year()),
	})
	if err != nil {
		return 0, nil
	}
	return int(count), nil
}

func (db *PostGresqlDB) PenIncrementAttempts(guid string) error {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	return db.queries.IncrementPenCounter(c, postgres_genererated.IncrementPenCounterParams{
		Guid: guid,
		Year: int32(time.Now().Year()),
	})
}

func (db *PostGresqlDB) PenDecrementAttempts(guid string) error {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	return db.queries.DecrementPenCounter(c, postgres_genererated.DecrementPenCounterParams{
		Guid: guid,
		Year: int32(time.Now().Year()),
	})
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

	logrus.Debugf("Potd: %v", potd)
	res := make([]mydb.PenData, 0, len(potd))
	for _, v := range potd {
		res = append(res, mydb.PenData{Name: sql.NullString{String: v.Name, Valid: true}, Size: v.Size})
	}
	return utils.GetTodayDateFormated(), res, nil
}
func (db *PostGresqlDB) PenPenHallOfFame() ([]mydb.PenData, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	phof, err := db.queries.GetPensOrderBySizeDesc(c, postgres_genererated.GetPensOrderBySizeDescParams{
		Column1: pgtype.Date{Time: time.Now(), Valid: true},
		Limit:   20,
	})

	if err != nil {
		return []mydb.PenData{}, nil
	}

	logrus.Debugf("Phof: %v", phof)
	res := make([]mydb.PenData, 0, len(phof))
	for _, v := range phof {
		res = append(res, mydb.PenData{Name: sql.NullString{String: v.Name, Valid: true}, Size: v.Size, Date: v.Date.Time})
	}
	return res, nil
}

func (db *PostGresqlDB) PenPenHallOfShame() ([]mydb.PenData, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	phos, err := db.queries.GetPensOrderBySizeAsc(c, postgres_genererated.GetPensOrderBySizeAscParams{
		Column1: pgtype.Date{Time: time.Now(), Valid: true},
		Limit:   20,
	})

	if err != nil {
		return []mydb.PenData{}, nil
	}

	logrus.Debugf("Phos: %v", phos)
	res := make([]mydb.PenData, 0, len(phos))
	for _, v := range phos {
		res = append(res, mydb.PenData{Name: sql.NullString{String: v.Name, Valid: true}, Size: v.Size, Date: v.Date.Time})
	}
	return res, nil
}

func (db *PostGresqlDB) PenPlayerGetDailySize(guid string) (float64, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()

	return db.queries.GetPlayerPenByDate(c, postgres_genererated.GetPlayerPenByDateParams{
		Guid: guid,
		Date: pgtype.Date{Time: time.Now(), Valid: true},
	})
}

func mustMarshalCheckpoints(checkpoints []int) string {
	b, _ := json.Marshal(checkpoints)
	return string(b)
}

func (db *PostGresqlDB) GetBestCheckpoints(guid, mapname, way string) ([]int, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	checkpointsStr, err := db.queries.GetBestCheckpointsByGuidMapWay(c, postgres_genererated.GetBestCheckpointsByGuidMapWayParams{
		Guid:    guid,
		Mapname: mapname,
		Way:     way,
	})
	if err != nil {
		return nil, err
	}
	var checkpoints []int
	err = json.Unmarshal([]byte(checkpointsStr), &checkpoints)
	return checkpoints, err
}

func (db *PostGresqlDB) HandleRun(info models.PlayerRunInfo, checkpoints []int) error {
	logrus.Debugf("HandleRun: %v | %v", info, checkpoints)
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
		logrus.Debugf("HandleRun: Time diff: %dms", timeDiff)

		if timeDiff > 0 {
			if err = db.queries.UpdateRunByGuidAndUTJ(c, postgres_genererated.UpdateRunByGuidAndUTJParams{
				Runtime:     int32(runtime),
				Checkpoints: mustMarshalCheckpoints(checkpoints),
				RunDate:     pgtype.Timestamp{Time: time.Now(), Valid: true},
				Guid:        info.Guid,
				Mapname:     info.Mapname,
				Way:         info.Way,
				Utj:         info.Utj,
			}); err != nil {
				return err
			}
			logrus.Debugf("HandleRun: Successful update time: %d for guid: %s", runtime, info.Guid)
		} else {
			logrus.Debugf("HandleRun: Not an improvement")
		}
	} else {
		logrus.Debugf("HandleRun: No run found. Create a new entry in db")
		if err = db.queries.CreateRun(c, postgres_genererated.CreateRunParams{
			Guid:        info.Guid,
			Utj:         info.Utj,
			Mapname:     info.Mapname,
			Way:         info.Way,
			Runtime:     int32(runtime),
			Checkpoints: mustMarshalCheckpoints(checkpoints),
			RunDate:     pgtype.Timestamp{Time: time.Now(), Valid: true},
			Demopath:    info.Demopath,
		}); err != nil {
			return err
		}
		logrus.Debugf("HandleRun: Created new entry in db")
	}
	return nil
}

func (db *PostGresqlDB) SetMapOptions(mapname, options string) error {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	return db.queries.UpsertMapOptions(c, postgres_genererated.UpsertMapOptionsParams{
		Mapname: mapname,
		Options: options,
	})
}

func (db *PostGresqlDB) GetMapOptions(mapname string) (string, bool) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	options, err := db.queries.GetMapOptions(c, mapname)
	if err != nil {
		return "", false
	}
	return options, true
}

func (db *PostGresqlDB) DeleteMapOptions(mapname string) (bool, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	rows, err := db.queries.DeleteMapOptions(c, mapname)
	return rows > 0, err
}

func (db *PostGresqlDB) SaveGoto(mapname, jumpname string, posX, posY, posZ, angleV, angleH float64) error {
	if mapname == "" || jumpname == "" {
		return fmt.Errorf("SaveGoto: mapname and jumpname must not be empty (mapname=%q, jumpname=%q)", mapname, jumpname)
	}
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	return db.queries.UpsertGoto(c, postgres_genererated.UpsertGotoParams{
		Mapname:  mapname,
		Jumpname: jumpname,
		PosX:     posX,
		PosY:     posY,
		PosZ:     posZ,
		AngleV:   angleV,
		AngleH:   angleH,
	})
}

func (db *PostGresqlDB) GetGoto(mapname, jumpname string) (mydb.GotoData, bool) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	row, err := db.queries.GetGotoByMapAndJump(c, postgres_genererated.GetGotoByMapAndJumpParams{
		Mapname:  mapname,
		Jumpname: jumpname,
	})
	if err != nil {
		return mydb.GotoData{}, false
	}
	return mydb.GotoData{
		PosX:   row.PosX,
		PosY:   row.PosY,
		PosZ:   row.PosZ,
		AngleV: row.AngleV,
		AngleH: row.AngleH,
	}, true
}

func (db *PostGresqlDB) GetGotoNames(mapname string) ([]string, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	return db.queries.GetGotoNamesByMap(c, mapname)
}

func (db *PostGresqlDB) DeleteGoto(mapname, jumpname string) (bool, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	rows, err := db.queries.DeleteGoto(c, postgres_genererated.DeleteGotoParams{
		Mapname:  mapname,
		Jumpname: jumpname,
	})
	return rows > 0, err
}

func (db *PostGresqlDB) DeleteAllGotos(mapname string) (int, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	rows, err := db.queries.DeleteAllGotosByMap(c, mapname)
	return int(rows), err
}

func (db *PostGresqlDB) AddIgnore(guid, ignoredGuid string) error {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	return db.queries.AddIgnore(c, guid, ignoredGuid)
}

func (db *PostGresqlDB) GetIgnoredGuids(guid string) ([]string, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	return db.queries.GetIgnoredGuids(c, guid)
}

func (db *PostGresqlDB) RemoveIgnore(guid, ignoredGuid string) error {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	return db.queries.RemoveIgnore(c, guid, ignoredGuid)
}

func (db *PostGresqlDB) GetIgnoredPlayers(guid string) ([]mydb.IgnoredPlayer, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	rows, err := db.queries.GetIgnoredPlayers(c, guid)
	if err != nil {
		return nil, err
	}
	result := make([]mydb.IgnoredPlayer, 0, len(rows))
	for _, r := range rows {
		result = append(result, mydb.IgnoredPlayer{Id: int(r.ID), Name: r.Name, Guid: r.IgnoredGuid})
	}
	return result, nil
}

func (db *PostGresqlDB) GetRandomQuote() (string, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	quote, err := db.queries.GetRandomQuote(c)
	if err != nil {
		return "", err
	}
	return quote.Text, nil
}

func (db *PostGresqlDB) SaveQuote(text string) error {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	_, err := db.queries.SaveQuote(c, text)
	return err
}

func (db *PostGresqlDB) GetPlayerById(id int) (mydb.LookupResult, bool) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	r, err := db.queries.GetPlayerById(c, int32(id))
	if err != nil {
		return mydb.LookupResult{}, false
	}
	return mydb.LookupResult{
		Id:      int(r.ID),
		Name:    r.Name,
		Aliases: r.Aliases,
		Role:    int(r.Role),
		Ip:      r.IpAddress,
		Guid:    r.Guid,
	}, true
}

func (db *PostGresqlDB) GetPlayersByIp(ip string) ([]mydb.LookupResult, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	rows, err := db.queries.GetPlayersByIp(c, ip)
	if err != nil {
		return nil, err
	}
	results := make([]mydb.LookupResult, 0, len(rows))
	for _, r := range rows {
		results = append(results, mydb.LookupResult{
			Id:      int(r.ID),
			Name:    r.Name,
			Aliases: r.Aliases,
			Role:    int(r.Role),
			Ip:      r.IpAddress,
			Guid:    r.Guid,
		})
	}
	return results, nil
}

func (db *PostGresqlDB) LookupPlayers(search string) ([]mydb.LookupResult, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	rows, err := db.queries.LookupPlayersByNameOrAlias(c, postgres_genererated.LookupPlayersByNameOrAliasParams{
		Column1: pgtype.Text{String: search, Valid: true},
		Limit:   10,
	})
	if err != nil {
		return nil, err
	}
	results := make([]mydb.LookupResult, 0, len(rows))
	for _, r := range rows {
		results = append(results, mydb.LookupResult{
			Id:      int(r.ID),
			Name:    r.Name,
			Aliases: r.Aliases,
			Role:    int(r.Role),
			Ip:      r.IpAddress,
			Guid:    r.Guid,
		})
	}
	return results, nil
}

func (db *PostGresqlDB) GetPlayerByGuid(guid string) (models.Player, bool) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()

	if playerDb, err := db.queries.GetPLayerByGuid(c, guid); err != nil {
		logrus.Errorf("[GetPlayerByGuid] Error: %v", err)
		return models.Player{}, false
	} else {
		logrus.Debugf("Player found in db: %+v", playerDb)
		var aliases []string
		json.Unmarshal([]byte(playerDb.Aliases), &aliases) //nolint: errcheck
		return models.Player{
			Role:    int(playerDb.Role),
			Name:    playerDb.Name,
			Guid:    guid,
			Id:      strconv.Itoa(int(playerDb.ID)),
			Ip:      playerDb.IpAddress,
			Aliases: aliases,
		}, true
	}
}

func (db *PostGresqlDB) AddBan(guid, ip, reason string) error {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	return db.queries.AddBan(c, guid, ip, reason)
}

func (db *PostGresqlDB) GetBan(guid string) (string, bool, error) {
	c, cancel := context.WithTimeout(db.ctx, dbTimeout*time.Second)
	defer cancel()
	reason, err := db.queries.GetBan(c, guid)
	if err != nil {
		// pgx returns an error when no row is found
		return "", false, nil
	}
	return reason, true, nil
}
