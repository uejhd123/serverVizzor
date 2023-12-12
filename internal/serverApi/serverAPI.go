package serverapi

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

//В этом файле собраны функции для взаимодействия с Серверами
//Получение данных о нагрузке процессора, занятой ОЗУ и т.д.

type SystemStat struct {
	Mem      []string `json:"memory"`
	HostName string   `json:"hostname"`
	UserName string   `json:"username"`
	Cpu      int      `json:"cpu"`
}

func GetSystemData() *SystemStat {
	cpu, _ := getCPUUsage()
	hostname, _ := getHostname()
	username, _ := getUsername()
	mem := memStat()
	return &SystemStat{
		Mem:      mem,
		HostName: hostname,
		UserName: username,
		Cpu:      cpu,
	}
}

func getCPUUsage() (int, error) {
	v, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0.0, err
	}

	return int(v[0]), nil
}

func getHostname() (string, error) {
	h, err := host.Info()
	if err != nil {
		return "", err
	}

	return h.Hostname, nil
}

func getUsername() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	return u.Username, nil
}

func memStat() []string {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		fmt.Println("fault")
		os.Exit(1)
	}
	buf := make([]byte, 84)
	n1, err := file.Read(buf)
	if err != nil {
		os.Exit(1)
	}
	res := strings.ReplaceAll(string(buf[:n1]), " ", "")
	return strings.Split(res, "\n")
}
