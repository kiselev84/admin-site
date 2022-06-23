package functions

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func CheckShh() {
	var (
		command = "who"
		str     string
		textMes = " подключен по shh к suse-pc\n"
	)

	sshLog, err := os.OpenFile("../shh.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println("Не удалось открыть файл sshLog", err)
	}
	defer sshLog.Close()

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

				var b bytes.Buffer
				b.WriteString(fmt.Sprintf(tm + " " + word[sOpenIndex+1:sCloseIndex] + textMes))
				_, err = sshLog.Write(b.Bytes())
				if err != nil {
					fmt.Println("Не смогли записать в файл", err)
					return
				}

			} else if word[1] == '1' && word[2] == '9' && word[3] == '2' && word[4] == '.' {
				exec.Command("kdialog", "--passivepopup", word[sOpenIndex+1:sCloseIndex]+textMes).Output()

				var b bytes.Buffer
				b.WriteString(fmt.Sprintf(tm + " " + word[sOpenIndex+1:sCloseIndex] + textMes))
				_, err = sshLog.Write(b.Bytes())
				if err != nil {
					fmt.Println("Не смогли записать в файл", err)
					return
				}
			}
			str = str[spaceIndex+1:]
			str = strings.Trim(str, space)

		}
		if str[1] == '1' && str[2] == '0' && str[3] == '.' {
			sOpenIndex := strings.Index(str, sOpen)
			sCloseIndex := strings.Index(str, sClose)
			exec.Command("kdialog", "--passivepopup", str[sOpenIndex+1:sCloseIndex]+textMes).Output()

			var b bytes.Buffer
			b.WriteString(fmt.Sprintf(tm + " " + str[sOpenIndex+1:sCloseIndex] + textMes))
			_, err = sshLog.Write(b.Bytes())
			if err != nil {
				fmt.Println("Не смогли записать в файл", err)
				return
			}
		} else if str[1] == '1' && str[2] == '9' && str[3] == '2' && str[4] == '.' {
			sOpenIndex := strings.Index(str, sOpen)
			sCloseIndex := strings.Index(str, sClose)
			exec.Command("kdialog", "--passivepopup", str[sOpenIndex+1:sCloseIndex]+textMes).Output()

			var b bytes.Buffer
			b.WriteString(fmt.Sprintf(tm + " " + str[sOpenIndex+1:sCloseIndex] + textMes))
			_, err = sshLog.Write(b.Bytes())
			if err != nil {
				fmt.Println("Не смогли записать в файл", err)
				return
			}
		}
		time.Sleep(20 * time.Second)
	}
}
