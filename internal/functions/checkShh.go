package functions

import (
	"database/sql"
	"fmt"
	"os/exec"
	"project/test_site/internal/entity"
	"strings"
	"time"
)

func CheckShh() {
	var (
		command = "who"
		str     string
		textMes = " подключен по shh к suse-pc\n"
	)

	for {
		out, _ := exec.Command(command).Output()
		str = string(out)
		space := " "
		sOpen := "("
		sClose := ")"
		tm := time.Now().Format("2006-01-02 15:04:05")

		for strings.Contains(str, space) {
			spaceIndex := strings.Index(str, space)
			sOpenIndex := strings.Index(str, sOpen)
			sCloseIndex := strings.Index(str, sClose)
			word := str[0:spaceIndex]

			if word[1] == '1' && word[2] == '0' && word[3] == '.' {

				exec.Command("kdialog", "--passivepopup", word[sOpenIndex+1:sCloseIndex]+textMes).Output()
				addSql(tm, word, sOpenIndex, sCloseIndex, textMes)

			} else if word[1] == '1' && word[2] == '9' && word[3] == '2' && word[4] == '.' {

				exec.Command("kdialog", "--passivepopup", word[sOpenIndex+1:sCloseIndex]+textMes).Output()
				addSql(tm, word, sOpenIndex, sCloseIndex, textMes)
			}
			str = str[spaceIndex+1:]
			str = strings.Trim(str, space)

		}
		if str[1] == '1' && str[2] == '0' && str[3] == '.' {
			sOpenIndex := strings.Index(str, sOpen)
			sCloseIndex := strings.Index(str, sClose)

			exec.Command("kdialog", "--passivepopup", str[sOpenIndex+1:sCloseIndex]+textMes).Output()
			addSql(tm, str, sOpenIndex, sCloseIndex, textMes)

		} else if str[1] == '1' && str[2] == '9' && str[3] == '2' && str[4] == '.' {
			sOpenIndex := strings.Index(str, sOpen)
			sCloseIndex := strings.Index(str, sClose)

			exec.Command("kdialog", "--passivepopup", str[sOpenIndex+1:sCloseIndex]+textMes).Output()
			addSql(tm, str, sOpenIndex, sCloseIndex, textMes)
		}
		time.Sleep(20 * time.Second)
	}
}

func addSql(tm string, word string, sOpenIndex int, sCloseIndex int, textMes string) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/Check",
		entity.UserSql, entity.PassSql, entity.HostSql))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("insert into Check.log_ssh (time, ip, text) values (?, ?,?)",
		tm, word[sOpenIndex+1:sCloseIndex], textMes)
	if err != nil {
		panic(err)
	}
}
