// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: recipe.sql

package store

import (
	"context"
	"database/sql"
)

const createRecipe = `-- name: CreateRecipe :one
INSERT INTO recipe (id, name, description, instructions) VALUES (?, ?, ?, ?) RETURNING id, created_at, name, description, instructions
`

type CreateRecipeParams struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
}

func (q *Queries) CreateRecipe(ctx context.Context, arg CreateRecipeParams) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, createRecipe,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Instructions,
	)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.Instructions,
	)
	return i, err
}

const deleteRecipe = `-- name: DeleteRecipe :exec
DELETE FROM recipe WHERE id = ?
`

func (q *Queries) DeleteRecipe(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteRecipe, id)
	return err
}

const getRandomRecipes = `-- name: GetRandomRecipes :many
SELECT id, created_at, name, description, instructions FROM recipe ORDER BY RANDOM() DESC LIMIT 10
`

func (q *Queries) GetRandomRecipes(ctx context.Context) ([]Recipe, error) {
	rows, err := q.db.QueryContext(ctx, getRandomRecipes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipe
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Description,
			&i.Instructions,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRecipe = `-- name: GetRecipe :one
SELECT id, created_at, name, description, instructions FROM recipe WHERE id = ?
`

func (q *Queries) GetRecipe(ctx context.Context, id string) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, getRecipe, id)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.Instructions,
	)
	return i, err
}

const getRecipes = `-- name: GetRecipes :many
SELECT id, created_at, name, description, instructions FROM recipe
`

func (q *Queries) GetRecipes(ctx context.Context) ([]Recipe, error) {
	rows, err := q.db.QueryContext(ctx, getRecipes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipe
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Description,
			&i.Instructions,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchRecipes = `-- name: SearchRecipes :many
SELECT id, created_at, name, description, instructions FROM recipe WHERE name LIKE ?
`

// Saerch anything that contains the given string
func (q *Queries) SearchRecipes(ctx context.Context, name string) ([]Recipe, error) {
	rows, err := q.db.QueryContext(ctx, searchRecipes, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Recipe
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Description,
			&i.Instructions,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRecipe = `-- name: UpdateRecipe :one
UPDATE recipe SET 
  name=COALESCE(?1, name),
  description=COALESCE(?2, description),
  instructions=COALESCE(?3, instructions)
WHERE id = ?4
RETURNING id, created_at, name, description, instructions
`

type UpdateRecipeParams struct {
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	Instructions sql.NullString `json:"instructions"`
	ID           string         `json:"id"`
}

func (q *Queries) UpdateRecipe(ctx context.Context, arg UpdateRecipeParams) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, updateRecipe,
		arg.Name,
		arg.Description,
		arg.Instructions,
		arg.ID,
	)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.Instructions,
	)
	return i, err
}
