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
func QuerySingleTestChannel(ctx context.Context, pg *models.Postgres) (*models.ChannelProfile, error) {
	var channel models.ChannelProfile
	err := pg.Db.QueryRow(ctx, "SELECT * FROM CHANNEL WHERE channelId = '1'").Scan(&channel.ChannelId, &channel.Username, &channel.Avatar)
	if err != nil {
		return nil, fmt.Errorf("QueryRow failed: %v", err)
	}
	return &channel, nil
}

// Test Multiple Query queries from the database.
func QueryMultiTestChannel(ctx context.Context, pg *models.Postgres) ([]models.ChannelProfile, error) {
	rows, err := pg.Db.Query(ctx, "SELECT * FROM CHANNEL")
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.ChannelProfile])
}

func QuerySingleTestVideo(ctx context.Context, pg *models.Postgres) (*models.VideoProfile, error) {
	var video models.VideoProfile
	err := pg.Db.QueryRow(ctx, "SELECT * FROM VIDEO WHERE videoId = 'videoID2'").Scan(&video.VideoId, &video.Title, &video.Thumbnail, &video.Watched, &video.VideoChannel)
	if err != nil {
		return nil, fmt.Errorf("QueryRow failed: %v", err)
	}
	return &video, nil
}

func QueryMultiTestVideo(ctx context.Context, pg *models.Postgres) ([]models.VideoProfile, error) {
	rows, err := pg.Db.Query(ctx, "SELECT * FROM VIDEO")
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.VideoProfile])
}

func QuerySingleTestCategory(ctx context.Context, pg *models.Postgres) (*models.CategoryProfile, error) {
	var category models.CategoryProfile
	err := pg.Db.QueryRow(ctx, "SELECT * FROM CATEGORY WHERE categoryId = 'cat2'").Scan(&category.CategoryId, &category.CatName, &category.CatChannel)
	if err != nil {
		return nil, fmt.Errorf("QueryRow failed: %v", err)
	}
	return &category, nil
}

func QueryMultiTestCategory(ctx context.Context, pg *models.Postgres) ([]models.CategoryProfile, error) {
	rows, err := pg.Db.Query(ctx, "SELECT * FROM CATEGORY")
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}

	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.CategoryProfile])
}

func GetChannel(pg *models.Postgres, ctx context.Context, chanId string) (*models.ChannelProfile, error) {
	query := `
	SELECT * FROM CHANNEL WHERE channelId = @chanId
	`
	var channel models.ChannelProfile
	// Define the named arguments for the query.
	args := pgx.NamedArgs{
		"chanId": chanId,
	}

	err := pg.Db.QueryRow(ctx, query, args).Scan(&channel.ChannelId, &channel.Username, &channel.Avatar)
	if err != nil {
		return nil, fmt.Errorf("QueryRow failed: %v", err)
	}
	return &channel, nil
}

func GetVideo(pg *models.Postgres, ctx context.Context, vidId string) (*models.VideoProfile, error) {
	query := `
	SELECT * FROM VIDEO WHERE videoId = @vidId
	`
	var video models.VideoProfile
	// Define the named arguments for the query.
	args := pgx.NamedArgs{
		"vidId": vidId,
	}

	err := pg.Db.QueryRow(ctx, query, args).Scan(&video.VideoId, &video.Title, &video.Thumbnail, &video.Watched, &video.VideoChannel)
	if err != nil {
		return nil, fmt.Errorf("QueryRow failed: %v", err)
	}
	return &video, nil
}

func GetCategory(pg *models.Postgres, ctx context.Context, catId string) (*models.CategoryProfile, error) {
	query := `
	SELECT * FROM Category WHERE categoryId = @catId
	`
	var category models.CategoryProfile
	// Define the named arguments for the query.
	args := pgx.NamedArgs{
		"catId": catId,
	}

	err := pg.Db.QueryRow(ctx, query, args).Scan(&category.CategoryId, &category.CatName, &category.CatChannel)
	if err != nil {
		return nil, fmt.Errorf("QueryRow failed: %v", err)
	}
	return &category, nil
}

func UpdateChannel(pg *models.Postgres, ctx context.Context, chanId string, chanDetails models.ChannelProfile) error {
	query := `
			UPDATE Channel
			SET channelId = @ChannelId, username = @Username, avatar = @Avatar
			WHERE channelId = @ChannelId
			`
	args := pgx.NamedArgs{
		"ChannelId":	chanDetails.ChannelId,
		"Username":	chanDetails.Username,
		"Avatar":	chanDetails.Avatar,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update row: %w", err)
	}

	return nil
}

func UpdateVideo(pg *models.Postgres, ctx context.Context, vidId string, vidDetails models.VideoProfile) error {
	query := `
			UPDATE Video
			SET videoId = @VideoId, title = @Title, thumbnail = @Thumbnail, watched = @Watched, videoChannel = @VideoChannel
			WHERE videoId = @VideoId
			`
	args := pgx.NamedArgs{
		"VideoId":	vidDetails.VideoId,
		"Title":	vidDetails.Title,
		"Thumbnail":	vidDetails.Thumbnail,
		"Watched":	vidDetails.Watched,
		"VideoChannel":	vidDetails.VideoChannel,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update row: %w", err)
	}

	return nil
}

func UpdateCategory(pg *models.Postgres, ctx context.Context, catId string, catDetails models.CategoryProfile) error {
	query := `
			UPDATE Category
			SET categoryId = @CategoryId, catName = @CatName, catChannel = @catChannel
			WHERE categoryId = @CategoryId
			`
	args := pgx.NamedArgs{
		"CategoryId":	catDetails.CategoryId,
		"CatName":	catDetails.CatName,
		"CatChannel":	catDetails.CatChannel,
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

func InsertVideo(pg *models.Postgres, ctx context.Context, vidDetails models.VideoProfile) error {
	query := `INSERT INTO Video (VideoId, Title, Thumbnail, Watched, VideoChannel) VALUES (@VideoId, @Title, @Thumbnail, @Watched, @VideoChannel)`
	args := pgx.NamedArgs{
		"VideoId":	vidDetails.VideoId,
		"Title":	vidDetails.Title,
		"Thumbnail":	vidDetails.Thumbnail,
		"Watched":	vidDetails.Watched,
		"VideoChannel":	vidDetails.VideoChannel,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func InsertCategory(pg *models.Postgres, ctx context.Context, catDetails models.CategoryProfile) error {
	query := `INSERT INTO Category (categoryId, catName, catChannel) VALUES (@CategoryId, @CatName, @CatChannel)`
	args := pgx.NamedArgs{
		"CategoryId":	catDetails.CategoryId,
		"CatName":	catDetails.CatName,
		"CatChannel":	catDetails.CatChannel,
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

func DeleteVideo(pg *models.Postgres, ctx context.Context, vidId string) error {
	query := `
	DELETE FROM Account WHERE VideoId = @vidId
	`
	args := pgx.NamedArgs{
		"vidId": vidId,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("error deleting row: %w", err)
	}

	return nil
}

func DeleteCategory(pg *models.Postgres, ctx context.Context, catId string) error {
	query := `
	DELETE FROM Category WHERE CategoryId = @catId
	`
	args := pgx.NamedArgs{
		"CategoryId":	catId,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("error deleting row: %w", err)
	}

	return nil
}
