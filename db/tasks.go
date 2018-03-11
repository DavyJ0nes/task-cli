package db

import (
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

var (
	dbString   string
	bucketName []byte
)

func init() {
	userHomeDir := os.Getenv("HOME")
	if userHomeDir == "" {
		fmt.Fprintln(os.Stderr, "$HOME Environment Variable not set")
		os.Exit(2)
	}

	dbDirString := fmt.Sprintf("%s/.tasks", userHomeDir)
	os.Mkdir(dbDirString, 0700)

	dbString = fmt.Sprintf("%s/tasks.db", dbDirString)
	bucketName = []byte("tasks")
}

// connect Opens a bolt DB instance and returns a pointer to it
func connect(dbPath string) (*bolt.DB, error) {
	var (
		db  *bolt.DB
		err error
	)

	// setup database configuration
	dbCfg := &bolt.Options{Timeout: 1 * time.Second}

	// open database
	db, err = bolt.Open(dbPath, 0600, dbCfg)
	if err != nil {
		return db, errors.Wrap(err, "Could not open bolt DB")
	}

	// Start writable transaction.
	tx, err := db.Begin(true)
	if err != nil {
		return db, errors.Wrap(err, "Start Writable Transaction")
	}
	defer tx.Rollback()

	// Initialize top-level buckets.
	if _, err := tx.CreateBucketIfNotExists(bucketName); err != nil {
		return db, errors.Wrap(err, "Could not create bucket")
	}

	// Save transaction to disk.
	return db, errors.Wrap(tx.Commit(), "Error commiting transaction")
}

// ListTasks lists out the open tasks
func ListTasks() error {
	db, err := connect(dbString)
	if err != nil {
		return errors.Wrap(err, "Could Not Connect to db")
	}
	defer db.Close()

	// read-only transaction
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found", bucketName)
		}

		// check if there are any open tasks
		if bucket.Stats().KeyN == 0 {
			fmt.Println("No Tasks. You're all clear")
			return nil
		}

		// print out all open tasks
		bucket.ForEach(func(k, v []byte) error {
			// convert byte slice to integer
			key := binary.BigEndian.Uint64(k)
			fmt.Printf("%d. | %s\n", key, string(v))
			return nil
		})

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "Could not list tasks")
	}

	return nil
}

// AddTask takes the arguments after the add command and creates a new entry in the database
func AddTask(args []string) error {
	db, err := connect(dbString)
	if err != nil {
		return errors.Wrap(err, "Could Not Connect to db")
	}
	defer db.Close()

	taskName := strings.Join(args, " ")

	// read/write transaction
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}

		// get next available key
		id, _ := bucket.NextSequence()

		// add task to database
		err = bucket.Put(itob(id), []byte(taskName))
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}

// CompleteTask takes the ID of a task and updates its status to complete
func CompleteTask(args []string) error {
	var (
		ids []int
		err error
	)

	db, err := connect(dbString)
	if err != nil {
		return errors.Wrap(err, "Could Not Connect to db")
	}
	defer db.Close()

	// convert arguments into integers
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			return errors.Errorf("Could not parse Argument: %s", arg)
		}
		ids = append(ids, id)
	}

	// delete tasks from database
	for _, id := range ids {
		// read/write transaction
		err = db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists(bucketName)
			if err != nil {
				return err
			}

			// get name of task
			taskVal := bucket.Get(itob(uint64(id)))

			// delete task from database
			err = bucket.Delete(itob(uint64(id)))
			if err != nil {
				return errors.Wrapf(err, "Could not Delete item %d", id)
			}

			// print out for user
			fmt.Printf("Completed: %s\n", taskVal)

			return nil
		})

		if err != nil {
			return errors.Wrap(err, "Could not complete transaction")
		}

	}
	return nil
}

// itob creates a byte slice from an integer
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
