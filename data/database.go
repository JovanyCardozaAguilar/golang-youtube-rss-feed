package data

import (
	"context"
	"docker-go-youtube-feed/models"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pgInstance *models.Postgres
	pgOnce     sync.Once
)

// CreateDBPool creates a new database connection pool.
func CreateDBPool(ctx context.Context, connString string) (*models.Postgres, error) {
	pgOnce.Do(func() {
		dbpool, err := pgxpool.New(ctx, connString)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
			os.Exit(1)
		}

		pgInstance = &models.Postgres{Db: dbpool}
	})

	return pgInstance, nil
}

// QueryGreeting queries the greeting message from the database.
func QueryGreeting(ctx context.Context, pg *models.Postgres) (string, error) {
	var greeting string
	err := pg.Db.QueryRow(ctx, "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		return "", fmt.Errorf("QueryRow failed: %v", err)
	}
	return greeting, nil
}

// Test Single Query queries from the database.
func QuerySingleTest(ctx context.Context, pg *models.Postgres) (*models.ChannelProfile, error) {
	var channel models.ChannelProfile
	err := pg.Db.QueryRow(ctx, "SELECT * FROM CHANNEL WHERE channelId = '1'").Scan(&channel.ChannelId, &channel.Username, &channel.Avatar)
	if err != nil {
		return nil, fmt.Errorf("QueryRow failed: %v", err)
	}
	return &channel, nil
}

// Test Multiple Query queries from the database.
func QueryMultiTest(ctx context.Context, pg *models.Postgres) ([]models.ChannelProfile, error) {
	rows, err := pg.Db.Query(ctx, "SELECT * FROM CHANNEL")
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.ChannelProfile])
}

func GetChannel(pg *models.Postgres, ctx context.Context, acctId string) (*models.ChannelProfile, error) {
	query := `
	SELECT * FROM CHANNEL WHERE channelId = @acctId
	`
	var channel models.ChannelProfile
	// Define the named arguments for the query.
	args := pgx.NamedArgs{
		"acctId": acctId,
	}

	err := pg.Db.QueryRow(ctx, query, args).Scan(&channel.ChannelId, &channel.Username, &channel.Avatar)
	if err != nil {
		return nil, fmt.Errorf("QueryRow failed: %v", err)
	}
	return &channel, nil
}

func UpdateChannel(pg *models.Postgres, ctx context.Context, acctId string, accountDetails models.ChannelProfile) error {
	query := `
			UPDATE Account 
			SET id = @Id, FirstName = @FirstName, LastName = @LastName, Token = @Token
			WHERE id = @Id
			`
	args := pgx.NamedArgs{
		"ChannelId":	accountDetails.ChannelId,
		"Username":	accountDetails.Username,
		"Avatar":	accountDetails.Avatar,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update row: %w", err)
	}

	return nil
}

func InsertChannel(pg *models.Postgres, ctx context.Context, accountDetails models.ChannelProfile) error {
	query := `INSERT INTO Account (id, FirstName, LastName, Token) VALUES (@Id, @FirstName, @LastName, @Token)`
	args := pgx.NamedArgs{
		"ChannelId":	accountDetails.ChannelId,
		"Username":	accountDetails.Username,
		"Avatar":	accountDetails.Avatar,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func DeleteChannel(pg *models.Postgres, ctx context.Context, acctId string) error {
	query := `
	DELETE FROM Account WHERE id = @Id
	`
	args := pgx.NamedArgs{
		"Id": acctId,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("error deleting row: %w", err)
	}

	return nil
}
