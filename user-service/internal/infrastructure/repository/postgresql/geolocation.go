package postgresql

import (
	"context"
	"fmt"
	"user-service/internal/entity"
	"user-service/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

const (
	geolocationsTableName      = "geolocations"
	geolocationsServiceName    = "geolocationService"
	geolocationsSpanRepoPrefix = "geolocationsRepo"
)

type geolocationsRepo struct {
	tableName string
	db        *postgres.PostgresDB
}

func NewGeolocationsRepo(db *postgres.PostgresDB) *geolocationsRepo {
	return &geolocationsRepo{
		tableName: geolocationsTableName,
		db:        db,
	}
}

func (p *geolocationsRepo) geolocationsSelectQueryPrefix() squirrel.SelectBuilder {
	return p.db.Sq.Builder.
		Select(
			"id",
			"latitude",
			"longitude",
			"owner_id",
		).From(p.tableName)
}

func (p geolocationsRepo) CreateGeolocation(ctx context.Context, geolocation *entity.Geolocation) error {
	data := map[string]any{
		"id":         geolocation.Id,
		"latitude":  geolocation.Latitude,
		"longitude":  geolocation.Longitude,
		"owner_id":   geolocation.OwnerId,
	}
	query, args, err := p.db.Sq.Builder.Insert(p.tableName).SetMap(data).ToSql()
	if err != nil {
		return p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "create"))
	}

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		return p.db.Error(err)
	}

	return nil
}

func (p geolocationsRepo) GetGeolocation(ctx context.Context, params map[string]int64) (*entity.Geolocation, error) {
	var (
		geolocation entity.Geolocation
	)

	queryBuilder := p.geolocationsSelectQueryPrefix()

	for key, value := range params {
		if key == "id" {
			queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value))
		}
	}
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "get"))
	}
	if err = p.db.QueryRow(ctx, query, args...).Scan(
		&geolocation.Id,
		&geolocation.Latitude,
		&geolocation.Longitude,
		&geolocation.OwnerId,
	); err != nil {
		return nil, p.db.Error(err)
	}

	return &geolocation, nil
}

func (p geolocationsRepo) UpdateGeolocation(ctx context.Context, geolocations *entity.Geolocation) error {
	clauses := map[string]any{
		"latitude":  geolocations.Latitude,	
		"longitude":  geolocations.Longitude,
		"owner_id":   geolocations.OwnerId,
	}
	sqlStr, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", geolocations.Id)).
		ToSql()
	if err != nil {
		return p.db.ErrSQLBuild(err, p.tableName+" update")
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return p.db.Error(err)
	}

	if commandTag.RowsAffected() == 0 {
		return p.db.Error(fmt.Errorf("no sql rows"))
	}

	return nil
}

func (p geolocationsRepo) DeleteGeolocation(ctx context.Context, guid int64) error {
	sqlStr, args, err := p.db.Sq.Builder.
		Delete(p.tableName).
		Where(p.db.Sq.Equal("id", guid)).
		ToSql()
	if err != nil {
		return p.db.ErrSQLBuild(err, p.tableName+" delete")
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return p.db.Error(err)
	}

	if commandTag.RowsAffected() == 0 {
		return p.db.Error(fmt.Errorf("no sql rows"))
	}

	return nil
}

func (p geolocationsRepo) ListGeolocation(ctx context.Context, filter map[string]string) ([]*entity.Geolocation, error) {
	var (
		geolocations []*entity.Geolocation
	)
	queryBuilder := p.geolocationsSelectQueryPrefix()

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "list"))
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, p.db.Error(err)
	}
	defer rows.Close()
	geolocations = make([]*entity.Geolocation, 0)
	for rows.Next() {
		var geolocation entity.Geolocation
		if err = rows.Scan(
			&geolocation.Id,
			&geolocation.Latitude,
			&geolocation.Longitude,
			&geolocation.OwnerId,
		); err != nil {
			return nil, p.db.Error(err)
		}
		geolocations = append(geolocations, &geolocation)
	}

	return geolocations, nil
}
