package main

import (
	"fmt"
)

type Server struct {
	Name  string
	Ports []string
	Ip    string
}

var servers []Server = []Server{
	Server{
		Name: "content212",
		Ports: []string{
			"8012", "9012", "9011", "8022", "9022", "9021", "9004", "17001",
			"11072", "11071", "11070",
		},
		Ip: "192.168.2.212",
	},
}

func main() {
	serverTpl := "server {\n"
	serverTpl += "    listen %s;\n"
	serverTpl += "    server_name %s.lexisnexis.com.cn;\n"
	serverTpl += "    proxy_cache off;\n"
	serverTpl += "    location / {\n"
	serverTpl += "        proxy_pass http://%s:%s;\n"
	serverTpl += "    }\n"
	serverTpl += "    location ~ .*\\.(js)$ {\n"
	serverTpl += "        proxy_pass http://%s:%s;\n"
	serverTpl += "    }\n"
	serverTpl += "    access_log  /data/logs/idol/%s.log  wwwlog;\n"
	serverTpl += "    error_log   /data/logs/idol/%s.error.log debug;\n"
	serverTpl += "}\n"

	for _, server := range servers {
		for _, port := range server.Ports {
			fmt.Printf(serverTpl, port, server.Name, server.Ip, port,
				server.Ip, port, server.Name, server.Name)
		}
	}
}
