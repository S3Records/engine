// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: schedules.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const create_Diagnostic_Schedule = `-- name: Create_Diagnostic_Schedule :one
INSERT INTO diagnostic_schedules (
  user_id,
  diagnostic_centre_id,
  schedule_time,
  test_type,
  -- schedule_status,
  doctor,
  acceptance_status,
  notes
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING id, user_id, diagnostic_centre_id, schedule_time, test_type, schedule_status, doctor, acceptance_status, notes, rejection_note, created_at, updated_at
`

type Create_Diagnostic_ScheduleParams struct {
	UserID             string                   `db:"user_id" json:"user_id"`
	DiagnosticCentreID string                   `db:"diagnostic_centre_id" json:"diagnostic_centre_id"`
	ScheduleTime       pgtype.Timestamptz       `db:"schedule_time" json:"schedule_time"`
	TestType           string                   `db:"test_type" json:"test_type"`
	Doctor             string                   `db:"doctor" json:"doctor"`
	AcceptanceStatus   ScheduleAcceptanceStatus `db:"acceptance_status" json:"acceptance_status"`
	Notes              pgtype.Text              `db:"notes" json:"notes"`
}

// Create a diagnostic schedule
func (q *Queries) Create_Diagnostic_Schedule(ctx context.Context, arg Create_Diagnostic_ScheduleParams) (*DiagnosticSchedule, error) {
	row := q.db.QueryRow(ctx, create_Diagnostic_Schedule,
		arg.UserID,
		arg.DiagnosticCentreID,
		arg.ScheduleTime,
		arg.TestType,
		arg.Doctor,
		arg.AcceptanceStatus,
		arg.Notes,
	)
	var i DiagnosticSchedule
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DiagnosticCentreID,
		&i.ScheduleTime,
		&i.TestType,
		&i.ScheduleStatus,
		&i.Doctor,
		&i.AcceptanceStatus,
		&i.Notes,
		&i.RejectionNote,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const delete_Diagnostic_Schedule = `-- name: Delete_Diagnostic_Schedule :one
DELETE FROM diagnostic_schedules
WHERE id = $1 AND user_id = $2
RETURNING id, user_id, diagnostic_centre_id, schedule_time, test_type, schedule_status, doctor, acceptance_status, notes, rejection_note, created_at, updated_at
`

type Delete_Diagnostic_ScheduleParams struct {
	ID     string `db:"id" json:"id"`
	UserID string `db:"user_id" json:"user_id"`
}

func (q *Queries) Delete_Diagnostic_Schedule(ctx context.Context, arg Delete_Diagnostic_ScheduleParams) (*DiagnosticSchedule, error) {
	row := q.db.QueryRow(ctx, delete_Diagnostic_Schedule, arg.ID, arg.UserID)
	var i DiagnosticSchedule
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DiagnosticCentreID,
		&i.ScheduleTime,
		&i.TestType,
		&i.ScheduleStatus,
		&i.Doctor,
		&i.AcceptanceStatus,
		&i.Notes,
		&i.RejectionNote,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const get_Diagnostic_Schedule = `-- name: Get_Diagnostic_Schedule :one
SELECT id, user_id, diagnostic_centre_id, schedule_time, test_type, schedule_status, doctor, acceptance_status, notes, rejection_note, created_at, updated_at FROM diagnostic_schedules WHERE id = $1 AND user_id = $2
`

type Get_Diagnostic_ScheduleParams struct {
	ID     string `db:"id" json:"id"`
	UserID string `db:"user_id" json:"user_id"`
}

// Get Diagnostic Schedule
func (q *Queries) Get_Diagnostic_Schedule(ctx context.Context, arg Get_Diagnostic_ScheduleParams) (*DiagnosticSchedule, error) {
	row := q.db.QueryRow(ctx, get_Diagnostic_Schedule, arg.ID, arg.UserID)
	var i DiagnosticSchedule
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DiagnosticCentreID,
		&i.ScheduleTime,
		&i.TestType,
		&i.ScheduleStatus,
		&i.Doctor,
		&i.AcceptanceStatus,
		&i.Notes,
		&i.RejectionNote,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const get_Diagnostic_Schedules = `-- name: Get_Diagnostic_Schedules :many
SELECT id, user_id, diagnostic_centre_id, schedule_time, test_type, schedule_status, doctor, acceptance_status, notes, rejection_note, created_at, updated_at FROM diagnostic_schedules
WHERE user_id = $1
ORDER BY schedule_time DESC
LIMIT $2 OFFSET $3
`

type Get_Diagnostic_SchedulesParams struct {
	UserID string `db:"user_id" json:"user_id"`
	Limit  int32  `db:"limit" json:"limit"`
	Offset int32  `db:"offset" json:"offset"`
}

// Get Diagnostic Schedules
func (q *Queries) Get_Diagnostic_Schedules(ctx context.Context, arg Get_Diagnostic_SchedulesParams) ([]*DiagnosticSchedule, error) {
	rows, err := q.db.Query(ctx, get_Diagnostic_Schedules, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*DiagnosticSchedule
	for rows.Next() {
		var i DiagnosticSchedule
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.DiagnosticCentreID,
			&i.ScheduleTime,
			&i.TestType,
			&i.ScheduleStatus,
			&i.Doctor,
			&i.AcceptanceStatus,
			&i.Notes,
			&i.RejectionNote,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const get_Diagnsotic_Schedule_By_Centre = `-- name: Get_Diagnsotic_Schedule_By_Centre :one
SELECT id, user_id, diagnostic_centre_id, schedule_time, test_type, schedule_status, doctor, acceptance_status, notes, rejection_note, created_at, updated_at FROM diagnostic_schedules
WHERE id = $1 AND diagnostic_centre_id = $2
`

type Get_Diagnsotic_Schedule_By_CentreParams struct {
	ID                 string `db:"id" json:"id"`
	DiagnosticCentreID string `db:"diagnostic_centre_id" json:"diagnostic_centre_id"`
}

func (q *Queries) Get_Diagnsotic_Schedule_By_Centre(ctx context.Context, arg Get_Diagnsotic_Schedule_By_CentreParams) (*DiagnosticSchedule, error) {
	row := q.db.QueryRow(ctx, get_Diagnsotic_Schedule_By_Centre, arg.ID, arg.DiagnosticCentreID)
	var i DiagnosticSchedule
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DiagnosticCentreID,
		&i.ScheduleTime,
		&i.TestType,
		&i.ScheduleStatus,
		&i.Doctor,
		&i.AcceptanceStatus,
		&i.Notes,
		&i.RejectionNote,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const get_Diagnsotic_Schedules_By_Centre = `-- name: Get_Diagnsotic_Schedules_By_Centre :many
SELECT id, user_id, diagnostic_centre_id, schedule_time, test_type, schedule_status, doctor, acceptance_status, notes, rejection_note, created_at, updated_at FROM diagnostic_schedules
WHERE diagnostic_centre_id = $1
ORDER BY schedule_time DESC
LIMIT $2 OFFSET $3
`

type Get_Diagnsotic_Schedules_By_CentreParams struct {
	DiagnosticCentreID string `db:"diagnostic_centre_id" json:"diagnostic_centre_id"`
	Limit              int32  `db:"limit" json:"limit"`
	Offset             int32  `db:"offset" json:"offset"`
}

func (q *Queries) Get_Diagnsotic_Schedules_By_Centre(ctx context.Context, arg Get_Diagnsotic_Schedules_By_CentreParams) ([]*DiagnosticSchedule, error) {
	rows, err := q.db.Query(ctx, get_Diagnsotic_Schedules_By_Centre, arg.DiagnosticCentreID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*DiagnosticSchedule
	for rows.Next() {
		var i DiagnosticSchedule
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.DiagnosticCentreID,
			&i.ScheduleTime,
			&i.TestType,
			&i.ScheduleStatus,
			&i.Doctor,
			&i.AcceptanceStatus,
			&i.Notes,
			&i.RejectionNote,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const update_Diagnostic_Schedule = `-- name: Update_Diagnostic_Schedule :one
UPDATE diagnostic_schedules
SET
  schedule_time = COALESCE($1, schedule_time),
  test_type = COALESCE($2, test_type),
  schedule_status = COALESCE($3, schedule_status),
  notes = COALESCE($4, notes),
  doctor = COALESCE($5, doctor),
  updated_at = NOW()
WHERE id = $6 AND user_id = $7
RETURNING id, user_id, diagnostic_centre_id, schedule_time, test_type, schedule_status, doctor, acceptance_status, notes, rejection_note, created_at, updated_at
`

type Update_Diagnostic_ScheduleParams struct {
	ScheduleTime   pgtype.Timestamptz `db:"schedule_time" json:"schedule_time"`
	TestType       string             `db:"test_type" json:"test_type"`
	ScheduleStatus ScheduleStatus     `db:"schedule_status" json:"schedule_status"`
	Notes          pgtype.Text        `db:"notes" json:"notes"`
	Doctor         string             `db:"doctor" json:"doctor"`
	ID             string             `db:"id" json:"id"`
	UserID         string             `db:"user_id" json:"user_id"`
}

// Update a diagnostic schedule
func (q *Queries) Update_Diagnostic_Schedule(ctx context.Context, arg Update_Diagnostic_ScheduleParams) (*DiagnosticSchedule, error) {
	row := q.db.QueryRow(ctx, update_Diagnostic_Schedule,
		arg.ScheduleTime,
		arg.TestType,
		arg.ScheduleStatus,
		arg.Notes,
		arg.Doctor,
		arg.ID,
		arg.UserID,
	)
	var i DiagnosticSchedule
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DiagnosticCentreID,
		&i.ScheduleTime,
		&i.TestType,
		&i.ScheduleStatus,
		&i.Doctor,
		&i.AcceptanceStatus,
		&i.Notes,
		&i.RejectionNote,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const update_Diagnostic_Schedule_By_Centre = `-- name: Update_Diagnostic_Schedule_By_Centre :one
UPDATE diagnostic_schedules
SET
  acceptance_status = COALESCE($1, acceptance_status),
  updated_at = NOW()
WHERE id = $2 AND diagnostic_centre_id = $3
RETURNING id, user_id, diagnostic_centre_id, schedule_time, test_type, schedule_status, doctor, acceptance_status, notes, rejection_note, created_at, updated_at
`

type Update_Diagnostic_Schedule_By_CentreParams struct {
	AcceptanceStatus   ScheduleAcceptanceStatus `db:"acceptance_status" json:"acceptance_status"`
	ID                 string                   `db:"id" json:"id"`
	DiagnosticCentreID string                   `db:"diagnostic_centre_id" json:"diagnostic_centre_id"`
}

func (q *Queries) Update_Diagnostic_Schedule_By_Centre(ctx context.Context, arg Update_Diagnostic_Schedule_By_CentreParams) (*DiagnosticSchedule, error) {
	row := q.db.QueryRow(ctx, update_Diagnostic_Schedule_By_Centre, arg.AcceptanceStatus, arg.ID, arg.DiagnosticCentreID)
	var i DiagnosticSchedule
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DiagnosticCentreID,
		&i.ScheduleTime,
		&i.TestType,
		&i.ScheduleStatus,
		&i.Doctor,
		&i.AcceptanceStatus,
		&i.Notes,
		&i.RejectionNote,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
