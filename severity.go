package zabbix

func SeverityDict(key int) string {
	severity := map[int]string{
		0: "not classified",
		1: "information",
		2: "warning",
		3: "average",
		4: "high",
		5: "disaster",
	}
	return severity[key]
}
