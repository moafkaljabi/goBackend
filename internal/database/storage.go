/*

Create table for the Device model.

Create CRUD operations for the Device model:
func CreateDevice
func UpdateDevice
func DeleteDevice
func GetDeviceByID


*/

package database

import (
	"database/sql"
	"fmt"
	"goBackend/internal/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Storage interface
type Storage interface {

	// Account methods
	CreateAccount(*models.Account) error
	DeleteAccount(int) error
	UpdateAccount(*models.Account) error
	GetAccountByID(int) (*models.Account, error)

	// Device methods
	CreateDevice(*models.Device) error
	UpdateDevice(*models.Device) error
	DeleteDevice(int) error
	GetDeviceByID(int) (*models.Device, error)
	GetDevicesByUserID(int) ([]*models.Device, error)
}

// Postgres storage implementation
type PostgresStore struct {
	db *sql.DB
}

// Initialize a new database connection
func NewPostgresStore() (*PostgresStore, error) {
	// Load environment variables from .env
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	connStr := fmt.Sprintf(
		"host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	println("Connected to database")

	return &PostgresStore{db: db}, nil
}

// Initialize the database (Create tables if needed)
func (s *PostgresStore) Init() error {
	if err := s.CreateAccountTable(); err != nil {
		return err
	}

	if err := s.CreateDeviceTable(); err != nil {
		return err
	}
	return nil
}

// Create the accounts table if it doesn't exist
func (s *PostgresStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		number SERIAL, 
		balance NUMERIC(12, 2),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := s.db.Exec(query)
	return err
}

// Insert a new account into the database
func (s *PostgresStore) CreateAccount(acc *models.Account) error {
	query := `INSERT INTO account (first_name, last_name, number, balance) VALUES ($1, $2, $3, $4) RETURNING id`
	return s.db.QueryRow(query, acc.FirstName, acc.LastName, acc.Number, acc.Balance).Scan(&acc.ID)
}

// Update an existing account
func (s *PostgresStore) UpdateAccount(acc *models.Account) error {
	query := `UPDATE account SET first_name=$1, last_name=$2, balance=$3 WHERE id=$4`
	_, err := s.db.Exec(query, acc.FirstName, acc.LastName, acc.Balance, acc.ID)
	return err
}

// Delete an account
func (s *PostgresStore) DeleteAccount(id int) error {
	query := `DELETE FROM account WHERE id=$1`
	_, err := s.db.Exec(query, id)
	return err
}

// Get an account by ID
func (s *PostgresStore) GetAccountByID(id int) (*models.Account, error) {
	query := `SELECT id, first_name, last_name, number, balance FROM account WHERE id=$1`
	acc := &models.Account{}
	err := s.db.QueryRow(query, id).Scan(&acc.ID, &acc.FirstName, &acc.LastName, &acc.Number, &acc.Balance)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

// Device methods

func (s *PostgresStore) CreateDeviceTable() error {
	query := `CREATE TABLE IF NOT EXISTS device (
		device_id SERIAL PRIMARY KEY,
		name VARCHAR(50),
		status VARCHAR(50),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`
	_, err := s.db.Exec((query))
	return err
}

func (s *PostgresStore) CreateDevice(device *models.Device) error {
	query := `INSERT INTO device (name, status) VALUES ($1, $2) RETURNINIG device_id`
	return s.db.QueryRow(query, device.Name, device.Status).Scan(&device.DeviceID)
}

func (s *PostgresStore) UpdateDevice(device *models.Device) error {
	query := `UPDATE device SET name=$1, status=$2 WHERE device_id=$3`
	_, err := s.db.Exec(query, device.Name, device.Status, device.DeviceID)
	return err
}

func (s *PostgresStore) DeleteDevice(deviceID int) error {
	query := `DELETE FROM device WHERE device_id=$1`
	_, err := s.db.Exec(query, deviceID)
	return err
}

func (s *PostgresStore) GetDeviceByID(deviceID int) (*models.Device, error) {
	query := `SELECT device_id, name, status FROM device WHERE device_id=$1`
	device := &models.Device{}
	err := s.db.QueryRow(query, deviceID).Scan(&device.DeviceID, &device.Name, &device.Status)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (s *PostgresStore) GetDevicesByUserID(userID int) ([]*models.Device, error) {
	query := `SELECT device_id, name, status FROM device WHERE user_id=$1`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	devices := make([]*models.Device, 0)
	for rows.Next() {
		device := &models.Device{}
		err := rows.Scan(&device.DeviceID, &device.Name, &device.Status)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}

	return devices, nil
}
