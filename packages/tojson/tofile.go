// save text to file
package tojson

import (
	"FakeDataSender/packages/atm"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const outPath = "./out/"

func SaveJSON(a *atm.ATM, path string) {
	if len(path) == 0 {
		path = outPath
	}

	path = outPath + "/JSON/"

	body := PackToJSON(a)
	now := time.Now()
	// timestamp := now.Unix()
	timestampNano := now.UnixNano()

	fileBody := "[" + body + "]"

	if len(body) > 0 {

		fileName := fmt.Sprintf("ATM_%d_%d.json", a.AtmNumber, timestampNano)

		dir, _ := filepath.Split(path)

		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.MkdirAll(dir, 0777)
		}
		out, err := os.Create(path + fileName)
		if err != nil {
			out.Close()
		}

		// Write the body to file

		r := strings.NewReader(fileBody)

		if _, err := io.Copy(out, r); err != nil {
			log.Println(err)
		}

		// Check(err)
		out.Close()
	}

}

func SaveToZabbixSenderData(a *atm.ATM, path string) {
	if len(path) == 0 {
		path = outPath
	}

	path = outPath + "/to_zabbix/"

	body := PackToJSON(a)
	now := time.Now()
	timestamp := now.Unix()
	timestampNano := now.UnixNano()

	fileBody := "ATM" + strconv.Itoa(a.AtmNumber) + "    " + "rds.raw.data" + "    " + strconv.Itoa(int(timestamp)) + "    " + "[" + body + "]"

	if len(body) > 0 {

		fileName := fmt.Sprintf("ATM_%d_%d.zabbix", a.AtmNumber, timestampNano)

		dir, _ := filepath.Split(path)

		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.MkdirAll(dir, 0777)
		}
		out, err := os.Create(path + fileName)
		if err != nil {
			out.Close()
		}

		// Write the body to file

		r := strings.NewReader(fileBody)

		if _, err := io.Copy(out, r); err != nil {
			log.Println(err)
		}

		// Check(err)
		out.Close()
	}

}
