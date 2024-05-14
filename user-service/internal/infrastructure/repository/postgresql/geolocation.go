package postgresql

import (
	"context"
	"fmt"
	"user-service/internal/entity"
	"user-service/internal/pkg/otlp"
	"user-service/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/k0kubun/pp"
)

const (
	geolocationsTableName      = "geolocations"
	geolocationsServiceName    = "geolocationService"
	geolocationsSpanRepoPrefix = "geolocationRepo"
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

func (p geolocationsRepo) Create(ctx context.Context, geolocation *entity.Geolocation) (*entity.Geolocation, error) {
	ctx, span := otlp.Start(ctx, geolocationsServiceName, geolocationsSpanRepoPrefix+"Create")
	defer span.End()
	data := map[string]any{
		"latitude":  geolocation.Latitude,
		"longitude": geolocation.Longitude,
		"owner_id":  geolocation.OwnerId,
	}
	query, args, err := p.db.Sq.Builder.Insert(p.tableName).SetMap(data).ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "create"))
	}

	query += "RETURNING id, latitude, longitude, owner_id"

	row := p.db.QueryRow(ctx, query, args...)

	if err = row.Scan(&geolocation.Id, &geolocation.Latitude, &geolocation.Longitude, &geolocation.OwnerId); err != nil {
		return nil, err
	}

	pp.Println(geolocation.Id)

	return geolocation, nil
}

func (p geolocationsRepo) Get(ctx context.Context, params map[string]int64) (*entity.Geolocation, error) {
	ctx, span := otlp.Start(ctx, geolocationsServiceName, geolocationsSpanRepoPrefix+"Create")
	defer span.End()
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

func (p geolocationsRepo) Update(ctx context.Context, geolocations *entity.Geolocation) (*entity.Geolocation, error) {
	ctx, span := otlp.Start(ctx, geolocationsServiceName, geolocationsSpanRepoPrefix+"Create")
	defer span.End()
	clauses := map[string]any{
		"latitude":  geolocations.Latitude,
		"longitude": geolocations.Longitude,
		"owner_id":  geolocations.OwnerId,
	}
	sqlStr, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", geolocations.Id)).
		ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, p.tableName+" update")
	}

	sqlStr += " RETURNING id, latitude, longitude, owner_id"

	row := p.db.QueryRow(ctx, sqlStr, args...)
	var resclient entity.Geolocation
	if err = row.Scan(&resclient.Id, &resclient.Latitude, &resclient.Longitude, &resclient.OwnerId); err != nil {
		return nil, err
	}

	return &resclient, nil
}

func (p geolocationsRepo) Delete(ctx context.Context, guid int64) (*entity.CheckResponse, error) {
	ctx, span := otlp.Start(ctx, geolocationsServiceName, geolocationsSpanRepoPrefix+"Create")
	defer span.End()
	sqlStr, args, err := p.db.Sq.Builder.
		Delete(p.tableName).
		Where(p.db.Sq.Equal("id", guid)).
		ToSql()
	if err != nil {
		return &entity.CheckResponse{Check: false}, nil
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return &entity.CheckResponse{Check: false}, nil
	}

	if commandTag.RowsAffected() == 0 {
		return &entity.CheckResponse{Check: false}, nil
	}

	return &entity.CheckResponse{Check: true}, nil
}

func (p geolocationsRepo) List(ctx context.Context, id string, limit, offset uint64, filter map[string]string) (*entity.AllGeolocation, error) {
	ctx, span := otlp.Start(ctx, geolocationsServiceName, geolocationsSpanRepoPrefix+"Create")
	defer span.End()

	queryBuilder := p.geolocationsSelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset).Where(p.db.Sq.Equal("owner_id", id))
	}


	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "list"))
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, p.db.Error(err)
	}
	defer rows.Close()

	var count int
	query = `SELECT COUNT(*) FROM geolocations`
	err = p.db.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return nil, p.db.Error(err)
	}

	geolocations := []entity.Geolocation{}
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
		geolocations = append(geolocations, geolocation)
	}

	return &entity.AllGeolocation{
		Geolocations: geolocations,
		Count:        count,
	}, nil
}
