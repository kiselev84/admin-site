package functions

import (
	"database/sql"
	"fmt"
	"os/exec"
	"project/test_site/internal/entity"
	"strings"
	"time"
)

func CheckNet() {

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

	for {
		for _, str := range users {
			checkIp(str)
			//fmt.Println(str)
		}

	}

}

func checkIp(p *entity.Ipcheck) {

	out, _ := exec.Command("ping", p.Ip, "-c 5", "-i 1", "-w 30").Output()
	tm := time.Now().Format("2006-01-02 15:04:05")
	textCheck := "100% Потеря канала"

	// Потеря канала
	if strings.Contains(string(out), "100% packet loss") {

		exec.Command("kdialog", "--passivepopup", "100% Потеря канала "+p.Ip+" "+p.City+" "+p.Office).Output()

		db, err := sql.Open("mysql", "usersql:Nomu8@RAmBat@tcp(10.101.2.194:3306)/Check")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		_, err = db.Exec("insert into Check.log_check_net (time, office, ip, city, server, text) values (?, ?,?,?,?,?)",
			tm, p.Office, p.Ip, p.City, p.Server, textCheck)
		if err != nil {
			panic(err)
		}
	}
}
