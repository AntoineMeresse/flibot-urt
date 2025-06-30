package postgres_impl

import (
	"context"
	"fmt"
	"time"

	mydb "github.com/AntoineMeresse/flibot-urt/src/db"
	postgres_genererated "github.com/AntoineMeresse/flibot-urt/src/db/postgres_impl/generated"
	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type PostGresqlDB struct {
	ctx     context.Context
	conn    *pgx.Conn
	queries *postgres_genererated.Queries
}

func InitPostGresDb(ctx context.Context, uri string) (*PostGresqlDB, error) {
	conn, err := pgx.Connect(ctx, uri)
	if err != nil {
		return nil, err
	}

	queries := postgres_genererated.New(conn)
	schema := mydb.ReadSchema("/home/antoine/dev/go/flibot-urt/sqlc/postgres/schema.sql")
	_, err = conn.Exec(ctx, schema)

	if err != nil {
		return nil, err
	}

	logrus.Debugf("Schema: %s", schema)

	return &PostGresqlDB{conn: conn, queries: queries, ctx: ctx}, nil
}

func (db *PostGresqlDB) Close() {
	db.conn.Close(db.ctx)
}

func (db *PostGresqlDB) SaveNewPlayer(name string, guid string, ipAddress string) (int, error) {
	c, cancel := context.WithTimeout(db.ctx, 2*time.Second)
	defer cancel()
	p, err := db.queries.CreatePlayer(c, postgres_genererated.CreatePlayerParams{
		Name:       name,
		Guid:       guid,
		IpAddress:  ipAddress,
		TimeJoined: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	logrus.Debugf("Player created: %v", p)
	return int(p.ID), err
}

func (db *PostGresqlDB) UpdatePlayer() error {
	return fmt.Errorf("To implement")
}

func (db *PostGresqlDB) PenAdd(guid string, size float64) error {
	return fmt.Errorf("To implement")
}
func (db *PostGresqlDB) PenPenOfTheDay() (string, []mydb.PenData, error) {
	return "", []mydb.PenData{}, fmt.Errorf("To implement")
}
func (db *PostGresqlDB) PenPenHallOfFame() ([]mydb.PenData, error) {
	return []mydb.PenData{}, fmt.Errorf("To implement")
}
func (db *PostGresqlDB) PenPenHallOfShame() ([]mydb.PenData, error) {
	return []mydb.PenData{}, fmt.Errorf("To implement")
}

func (db *PostGresqlDB) HandleRun(info models.PlayerRunInfo, checkpoints []int) error {
	return fmt.Errorf("To implement")
}
