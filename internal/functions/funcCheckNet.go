package functions

import (
	"os/exec"
	"project/test_site/internal/entity"
	"project/test_site/internal/repository"
	"strings"
	"time"
)

var repo = repository.NewRepository()

func CheckNet() {
	for {
		users := repo.GetAll()
		for _, str := range users {
			checkIp(str)
			//fmt.Println(str)
		}
	}
}

func checkIp(u *entity.Ipcheck) {

	out, _ := exec.Command("ping", u.Ip, "-c 5", "-i 1", "-w 30").Output()
	tm := time.Now().Format("2006-01-02 15:04:05")
	message := "100% Потеря канала"
	// Потеря канала
	if strings.Contains(string(out), "100% packet loss") {
		exec.Command("kdialog", "--passivepopup", message+"  "+u.City+" "+u.Office+" "+u.Ip).Output()
		repo.AddLogNet(u, message, tm)
	}
}
