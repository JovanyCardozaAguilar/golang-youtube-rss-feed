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
	err := pg.Db.QueryRow(ctx, "SELECT * FROM VIDEO WHERE videoId = 'videoID2'").Scan(&video.VideoId, &video.Title, &video.Thumbnail, &video.Watched)
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
	err := pg.Db.QueryRow(ctx, "SELECT * FROM CATEGORY WHERE categoryId = 'cat2'").Scan(&category.CategoryId, &category.CatName)
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

	err := pg.Db.QueryRow(ctx, query, args).Scan(&video.VideoId, &video.VChannelId, &video.Title, &video.Thumbnail, &video.Watched)
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

	err := pg.Db.QueryRow(ctx, query, args).Scan(&category.CategoryId, &category.CatName)
	if err != nil {
		return nil, fmt.Errorf("QueryRow failed: %v", err)
	}
	return &category, nil
}

func GetChannelCategory(pg *models.Postgres, ctx context.Context, ccId string) ([]models.ChannelProfile, error) {
	query := `
	SELECT channelId, username, avatar 
	FROM Channel
	JOIN Channel_Category ON channelId = ccChannelId
	WHERE ccCategoryId = @category
	`
	// Define the named arguments for the query.
	args := pgx.NamedArgs{
		"category": ccId,
	}

	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("Query failed: %w", err)
	}
	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.ChannelProfile])
}

func GetVideoCategory(pg *models.Postgres, ctx context.Context, vcId string) ([]models.VideoProfile, error) {
	query := `
	SELECT videoId, vChannelId, title, thumbnail, watched
	FROM Video
	JOIN Video_Category ON videoId = vcVideoId
	WHERE vcCategoryId = @category
	`
	// Define the named arguments for the query.
	args := pgx.NamedArgs{
		"category": vcId,
	}

	rows, err := pg.Db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("QueryRow failed: %v", err)
	}
	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.VideoProfile])
}

func GetFeed(pg *models.Postgres, ctx context.Context) ([]models.FeedProfile, error) {
	query := `
	SELECT videoId, vChannelId, title, thumbnail, watched
	FROM Video
	`

	rows, err := pg.Db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("QueryRow failed: %v", err)
	}
	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.FeedProfile])
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
			SET videoId = @VideoId, title = @Title, thumbnail = @Thumbnail, watched = @Watched
			WHERE videoId = @VideoId
			`
	args := pgx.NamedArgs{
		"VideoId":	vidDetails.VideoId,
		"Title":	vidDetails.Title,
		"Thumbnail":	vidDetails.Thumbnail,
		"Watched":	vidDetails.Watched,
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
			SET categoryId = @CategoryId, catName = @CatName
			WHERE categoryId = @CategoryId
			`
	args := pgx.NamedArgs{
		"CategoryId":	catDetails.CategoryId,
		"CatName":	catDetails.CatName,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update row: %w", err)
	}

	return nil
}

func InsertChannel(pg *models.Postgres, ctx context.Context, accountDetails models.ChannelProfile) error {
	query := `INSERT INTO Channel (channelid, username, avatar) VALUES (@ChannelId, @Username, @Avatar)`
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
	query := `INSERT INTO Video (VideoId, VChannelId, Title, Thumbnail, Watched) VALUES (@VideoId, @VChannelId, @Title, @Thumbnail, @Watched)`
	args := pgx.NamedArgs{
		"VideoId":	vidDetails.VideoId,
		"VChannelId":	vidDetails.VChannelId,
		"Title":	vidDetails.Title,
		"Thumbnail":	vidDetails.Thumbnail,
		"Watched":	vidDetails.Watched,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func InsertCategory(pg *models.Postgres, ctx context.Context, catDetails models.CategoryProfile) error {
	query := `INSERT INTO Category (categoryId, catName) VALUES (@CategoryId, @CatName)`
	args := pgx.NamedArgs{
		"CategoryId":	catDetails.CategoryId,
		"CatName":	catDetails.CatName,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func InsertChannelCategory(pg *models.Postgres, ctx context.Context, catDetails models.ChannelCategoryProfile) error {
	query := `INSERT INTO Channel_Category (CcChannelId, CcCategoryId) VALUES (@channel, @category)`
	args := pgx.NamedArgs{
		"channel":	catDetails.CcChannelId,
		"category":	catDetails.CcCategoryId,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func InsertVideoCategory(pg *models.Postgres, ctx context.Context, catDetails models.VideoCategoryProfile) error {
	query := `INSERT INTO Video_Category (VcVideoId, VcCategoryId) VALUES (@video, @category)`
	args := pgx.NamedArgs{
		"video":	catDetails.VcVideoId,
		"category":	catDetails.VcCategoryId,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func DeleteChannel(pg *models.Postgres, ctx context.Context, chanId string) error {
	query := `
	DELETE FROM Channel WHERE channelId = @Id
	`
	args := pgx.NamedArgs{
		"Id": chanId,
	}
	_, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("error deleting row: %w", err)
	}

	return nil
}

func DeleteVideo(pg *models.Postgres, ctx context.Context, vidId string) error {
	query := `
	DELETE FROM Video WHERE VideoId = @vidId
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
	DELETE FROM Category WHERE CategoryId = @CategoryId
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

func DeleteChannelCategory(pg *models.Postgres, ctx context.Context, ccDetails models.ChannelCategoryProfile) error {
	query := `
	DELETE FROM Channel_category WHERE ccChannelId = @channel AND ccCategoryId = @category
	`
	args := pgx.NamedArgs{
		"channel": ccDetails.CcChannelId,
		"category": ccDetails.CcCategoryId,
	}
	check, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("error deleting row: %w", err)
	}
	if check.RowsAffected() == 0 {
		return fmt.Errorf("no matching entry to delete %w", err)
	}

	return nil
}

func DeleteVideoCategory(pg *models.Postgres, ctx context.Context, vcDetails models.VideoCategoryProfile) error {
	query := `
	DELETE FROM Video_category WHERE vcVideoId = @video AND vcCategoryId = @category
	`
	args := pgx.NamedArgs{
		"video": vcDetails.VcVideoId,
		"category": vcDetails.VcCategoryId,
	}
	check, err := pg.Db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("error deleting row: %w", err)
	}
	if check.RowsAffected() == 0 {
		return fmt.Errorf("no matching entry to delete %w", err)
	}

	return nil
}
