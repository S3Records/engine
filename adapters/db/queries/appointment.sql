-- name: CreateAppointment :one
INSERT INTO appointments (
    patient_id,
    schedule_id,
    diagnostic_centre_id,
    appointment_date,
    time_slot,
    status,
    notes,
    reminder_sent
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, false
) RETURNING *;

-- name: GetAppointment :one
SELECT * FROM appointments WHERE id = $1;
-- name: GetPatientAppointments :many
SELECT 
    a.*,
    dc.diagnostic_centre_name as diagnostic_centre_name,
    dc.address as diagnostic_centre_address
FROM appointments a
JOIN diagnostic_centres dc ON a.diagnostic_centre_id = dc.id
JOIN diagnostic_schedules s ON a.schedule_id = s.id
WHERE a.patient_id = $1
    AND a.status = ANY($2::appointment_status[])
    AND a.appointment_date BETWEEN $3 AND $4
ORDER BY a.appointment_date ASC, a.time_slot ASC
LIMIT $5 OFFSET $6;

-- name: GetCentreAppointments :many
SELECT /*+ INDEX(appointments idx_appointments_diagnostic_centre) */ * FROM appointments 
WHERE diagnostic_centre_id = $1
AND status = ANY($2::appointment_status[])
AND appointment_date BETWEEN $3 AND $4
ORDER BY appointment_date ASC
LIMIT $5 OFFSET $6;

-- name: UpdateAppointmentStatus :one
UPDATE appointments 
SET 
    status = COALESCE($2, status),
    updated_at = CURRENT_TIMESTAMP,
    check_in_time = CASE 
        WHEN $2 = 'in_progress' AND check_in_time IS NULL THEN CURRENT_TIMESTAMP 
        ELSE check_in_time 
    END,
    completion_time = CASE 
        WHEN $2 = 'completed' AND completion_time IS NULL THEN CURRENT_TIMESTAMP 
        ELSE completion_time 
    END
WHERE id = $1 
RETURNING *;

-- name: CancelAppointment :one
UPDATE appointments 
SET 
    status = 'cancelled',
    cancellation_reason = $2,
    cancelled_by = $3,
    cancellation_time = CURRENT_TIMESTAMP,
    cancellation_fee = COALESCE($4, 0),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 
RETURNING *;

-- name: RescheduleAppointment :one
WITH old_appointment AS (
    UPDATE appointments 
    SET 
        status = 'rescheduled',
        rescheduling_reason = $2,
        rescheduled_by = $3,
        rescheduling_time = CURRENT_TIMESTAMP,
        rescheduling_fee = COALESCE($4, 0),
        updated_at = CURRENT_TIMESTAMP
    WHERE appointments.id = $1 
    RETURNING 
        appointments.id as original_appointment_id,
        appointments.patient_id,
        appointments.diagnostic_centre_id,
        appointments.payment_amount,
        appointments.notes
)
INSERT INTO appointments (
    patient_id,
    schedule_id,
    diagnostic_centre_id,
    appointment_date,
    time_slot,
    status,
    payment_amount,
    notes,
    original_appointment_id
) 
SELECT 
    old_appointment.patient_id,
    $5,  -- new_schedule_id
    old_appointment.diagnostic_centre_id,
    $6,  -- new_appointment_date
    $7,  -- new_time_slot
    'confirmed'::appointment_status,
    old_appointment.payment_amount,
    old_appointment.notes,
    old_appointment.original_appointment_id
FROM old_appointment
RETURNING *;

-- name: UpdateAppointmentPayment :one
UPDATE appointments 
SET 
    payment_id = COALESCE($2, payment_id),
    payment_status = COALESCE($3, payment_status),
    payment_amount = COALESCE($4, payment_amount),
    status = COALESCE($5, status),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 
RETURNING *;

-- name: DeleteAppointment :exec
DELETE FROM appointments WHERE id = $1;

-- name: UpdateReminderSent :one
UPDATE appointments 
SET reminder_sent = true,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 
RETURNING *;
