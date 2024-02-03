package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

type Department struct {
	DepartmentID   int64  `db:"department_id"`
	DepartmentName string `db:"department_name"`
}

type Employee struct {
	EmployeeID   int64  `db:"employee_id"`
	EmployeeName string `db:"employee_name"`
	DepartmentID int64  `db:"department_id"`
}

func insertDepartmentsBatch(db *sql.DB, departments []Department) error {
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(context.Background(), "INSERT INTO departements (department_name) VALUES (?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, d := range departments {
		_, err := stmt.ExecContext(context.Background(), d.DepartmentName)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func insertEmployeesBatch(db *sql.DB, employees []Employee) error {
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(context.Background(), "INSERT INTO employee (employee_name, department_id) VALUES (?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, e := range employees {
		_, err := stmt.ExecContext(context.Background(), e.EmployeeName, e.DepartmentID)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	departments := []Department{
		{DepartmentName: "Marketing"},
		{DepartmentName: "Sales"},
		{DepartmentName: "Engineering"},
	}

	employees := []Employee{
		{EmployeeName: "John Doe", DepartmentID: 1},
		{EmployeeName: "Jane Doe", DepartmentID: 2},
		{EmployeeName: "Peter Smith", DepartmentID: 3},
	}

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/office")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = insertDepartmentsBatch(db, departments)
	if err != nil {
		log.Fatal(err)
	}

	err = insertEmployeesBatch(db, employees)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Batch insert berhasil!")
}
