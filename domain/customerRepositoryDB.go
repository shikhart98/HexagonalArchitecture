package domain

import (
	"HexagonalArch/errs"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppErrors) {
	findAllSQLQuery := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers;"

	rows, err := d.client.Query(findAllSQLQuery)
	if err != nil {
		errStr := "Error while querying customer table " + err.Error()
		log.Println(errStr)
		return nil, errs.NewUnexpectedError(errStr)
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			errStr := "Error while scanning customer table " + err.Error()
			log.Println(errStr)
			return nil, errs.NewUnexpectedError(errStr)
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDB) ByID(id string) (*Customer, *errs.AppErrors) {
	customerSQLQuery := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id=?;"

	row := d.client.QueryRow(customerSQLQuery, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		log.Println("Error while scanning customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:@tcp(localhost:3308)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{client: client}
}
