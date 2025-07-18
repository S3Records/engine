package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/medivue/adapters/db"
	"github.com/medivue/core/ports"
)

type Repositories struct {
	User         ports.UserRepository
	Schedule     ports.ScheduleRepository
	Diagnostic   ports.DiagnosticRepository
	Availability ports.AvailabilityRepository
	Record       ports.RecordRepository
	Payment      ports.PaymentRepository
	Appointment  ports.AppointmentRepository
	TestPrice    ports.TestPriceRepository
}

// InitializeRepositories creates and returns all repositories
func InitializeRepositories(store *db.Queries, conn *pgxpool.Pool) *Repositories {
	return &Repositories{
		User:         NewUserRepository(store),
		Record:       NewRecordRepository(store),
		Schedule:     NewScheduleRepository(store),
		TestPrice:    NewTestPriceRepository(store),
		Availability: NewAvailabilityRepository(store),
		Payment:      NewPaymentRepository(store, conn),
		Appointment:  NewAppointmentRepository(store, conn),
		Diagnostic:   NewDiagnosticCentreRepository(store, conn),
	}
}
