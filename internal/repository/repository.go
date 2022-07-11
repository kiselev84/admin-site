package repository

import (
	"database/sql"
	"fmt"
	"project/test_site/internal/entity"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type repository struct {
	sync.Mutex
	usersById map[uint8]*entity.Ipcheck
}

func NewRepository() *repository {
	return &repository{
		usersById: make(map[uint8]*entity.Ipcheck),
	}
}

//Добавление проверяемого ip в хранилище
func (r *repository) AddNewIp(ipcheck *entity.Ipcheck) (uint8, error) {
	r.Lock()
	defer r.Unlock()

	db, err := sql.Open("mysql", "usersql:Nomu8@RAmBat@tcp(10.101.2.194:3306)/Check")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into Check.ipcheck (office, ip, city, server) values (?, ?,?,?)",
		ipcheck.Office, ipcheck.Ip, ipcheck.City, ipcheck.Server)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()

	return uint8(id), nil
}

//Получение всех ip в хранилище
func (r *repository) GetAll() []*entity.Ipcheck {
	db, err := sql.Open("mysql", "usersql:Nomu8@RAmBat@tcp(10.101.2.194:3306)/Check")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from Check.ipcheck")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := make([]*entity.Ipcheck, 0)

	for rows.Next() {
		p := &entity.Ipcheck{}
		err = rows.Scan(&p.Id, &p.Office, &p.Ip, &p.City, &p.Server)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}
	return users
}

//Получение log_ssh в хранилище
func (r *repository) GetLogSsh() []*entity.SshLog {
	db, err := sql.Open("mysql", "usersql:Nomu8@RAmBat@tcp(10.101.2.194:3306)/Check")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from Check.log_ssh ORDER BY ID DESC")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := make([]*entity.SshLog, 0)

	for rows.Next() {
		p := &entity.SshLog{}
		err = rows.Scan(&p.Id, &p.Time, &p.Ip, &p.Text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}
	return users
}

//Получение log_check_net в хранилище
func (r *repository) GetLogCheckNet() []*entity.CheckNetLog {
	db, err := sql.Open("mysql", "usersql:Nomu8@RAmBat@tcp(10.101.2.194:3306)/Check")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from Check.log_check_net ORDER BY ID DESC")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := make([]*entity.CheckNetLog, 0)

	for rows.Next() {
		p := &entity.CheckNetLog{}
		err = rows.Scan(&p.Id, &p.Time, &p.Office, &p.Ip, &p.City, &p.Server, &p.Text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}
	return users
}

//Получение log_check_net в хранилище по City
func (r *repository) GetLogCheckNetCity(city string) []*entity.CheckNetLog {
	db, err := sql.Open("mysql", "usersql:Nomu8@RAmBat@tcp(10.101.2.194:3306)/Check")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from Check.log_check_net WHERE city=? ORDER BY ID DESC", city)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := make([]*entity.CheckNetLog, 0)

	for rows.Next() {
		p := &entity.CheckNetLog{}
		err = rows.Scan(&p.Id, &p.Time, &p.Office, &p.Ip, &p.City, &p.Server, &p.Text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}
	return users
}
