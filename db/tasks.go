package db

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

var (
	dbString   string
	bucketName []byte
)

// Task defines a TODO task
type Task struct {
	ID   int
	Name string
}

// Initialise sorts out the setup of the database
func Initialise(dbName string) {
	userHomeDir := os.Getenv("HOME")
	if userHomeDir == "" {
		fmt.Fprintln(os.Stderr, "$HOME Environment Variable not set")
		os.Exit(2)
	}

	// set directory and create if necessary
	dbDirString := fmt.Sprintf("%s/.tasks", userHomeDir)
	os.Mkdir(dbDirString, 0700)

	// Update package variables
	dbString = fmt.Sprintf("%s/%s", dbDirString, dbName)
	bucketName = []byte("tasks")
}

// AllTasks lists out the open tasks
func AllTasks() ([]Task, error) {
	var tasks []Task
	// connect to database
	db, err := connect(dbString)
	if err != nil {
		return nil, errors.Wrap(err, "Could Not Connect to db")
	}
	defer db.Close()

	// read-only transaction
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found", bucketName)
		}

		// print out all open tasks
		bucket.ForEach(func(k, v []byte) error {
			// convert byte slice to integer
			tasks = append(tasks, Task{
				ID:   btoi(k),
				Name: string(v),
			})
			return nil
		})

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "Could not get all tasks")
	}

	return tasks, nil
}

// CreateTask takes the arguments after the add command and creates a new entry in the database
func CreateTask(taskName string) error {
	// connect to database
	db, err := connect(dbString)
	if err != nil {
		return errors.Wrap(err, "Could Not Connect to db")
	}
	defer db.Close()

	// read/write transaction
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return errors.Wrapf(err, "Failed to create non existant bucket: %s", string(bucketName))
		}

		// get next available key
		id, _ := bucket.NextSequence()

		// add task to database
		err = bucket.Put(itob(int(id)), []byte(taskName))
		if err != nil {
			return errors.Wrapf(err, "Could Not Create Task: %s", taskName)
		}
		return nil
	})
	return nil
}

// DeleteTask takes the ID of a task and updates its status to complete
func DeleteTask(taskKey int) error {
	// connect to database
	db, err := connect(dbString)
	if err != nil {
		return errors.Wrap(err, "Could Not Connect to db")
	}
	defer db.Close()

	// read/write transaction
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return errors.Wrapf(err, "Failed to create non existant bucket: %s", string(bucketName))
		}

		// delete task from database
		err = bucket.Delete(itob(taskKey))
		if err != nil {
			return errors.Wrapf(err, "Could not Delete item %d", taskKey)
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "Could not complete transaction")
	}

	return nil
}

// connect Opens a bolt DB instance and returns a pointer to it
func connect(dbPath string) (*bolt.DB, error) {
	var (
		db  *bolt.DB
		err error
	)

	// setup database configuration
	dbCfg := &bolt.Options{Timeout: 1 * time.Second}

	// get database object
	db, err = bolt.Open(dbPath, 0600, dbCfg)
	if err != nil {
		return db, errors.Wrap(err, "Could not open bolt DB")
	}

	// Return db object and error (if any)
	// The closure checks if bucketName exists and creates if it doesn't
	return db, db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(bucketName); err != nil {
			return errors.Wrapf(err, "Could not create bucket: %s", string(bucketName))
		}
		return nil
	})
}

// itob creates a byte slice from an integer
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi creates an integer from a byte slice
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
