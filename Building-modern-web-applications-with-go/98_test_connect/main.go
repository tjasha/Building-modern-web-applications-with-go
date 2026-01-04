package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	// connect to DB
	// conn=connection pool
	conn, err := sql.Open("pgx", "host=localhost port =xxx dbname=test_connect user=xxx password=")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect : %v\n", err))
	}	
	//we need to close it
	defer conn.Close()

	log.Println("Connected to database")


	// test connection 
	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot ping database!")
	}

	log.Println("Pinged database!")


	// get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}


	// insert a row
	//creating a query with placeholders
	query := `insert into users (first_name, last_name) values ($1, $2)`
	//executing query with placeholder values. it prevents malicious code to be executed!
	_, err = conn.Exec(query, "Jack", "Brown")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Insert the row!")


	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}


	// update a row
	stmt := `update users set first_name = $1 where id = $2`
	//executing query with placeholder values. it prevents malicious code to be executed!
	_, err = conn.Exec(stmt, "Jackie", 5)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update the row!")


	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}


	// get one row by id 
	query = `select id, first_name, last_name from users where id =$1`
	var firstName, lastName string
	var id int 

	//only return one row. in this case what return is a row OR an error
	row := conn.QueryRow(query, 1)
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("QueryRow returns ", id, firstName, lastName)


	// delete a row
	query = `delete from users where id = $1`
	_, err = conn.Exec(query, 6)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Deleted a  row!")


	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

}

func getAllRows(conn *sql.DB) error {

	rows, err := conn.Query("select id, first_name, last_name from users")
	if err != nil {
		log.Println(err)
		return err
	}

	//we need to close the rows!
	defer rows.Close()

	var firstName, lastName string
	var id int

	//looping over rows
	for rows.Next() {
		//we want to scan all rows and pass id into id, first_name into firstName...
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}

		fmt.Println("Record is:", id, firstName, lastName)
	
	}

	//we should always check for the errors just before exiting
	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	fmt.Println("----------------------------------------")

	return nil
}
