package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	hosts = make([]Hosts, 0, 10)
}

var hosts []Hosts

type Hosts struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
}

// UpdateHosts update hosts data
func UpdateHosts(c *gin.Context) {
	var host Hosts
	if err := c.ShouldBindJSON(&host); err != nil {
		c.Data(http.StatusBadRequest, "text/plain", []byte("Wrong data"))
		return
	}
	for i, v := range hosts {
		if v.Hostname == host.Hostname {
			hosts[i].IP = host.IP
			c.Data(http.StatusOK, "text/plain", []byte("Success!"))
			return
		}
	}
	hosts = append(hosts, host)
	c.Data(http.StatusOK, "text/plain", []byte("Success!"))
}

// ResetHosts remove all hosts data
func ResetHosts(c *gin.Context) {
	hosts = make([]Hosts, 0, 10)
	c.Data(http.StatusOK, "text/plain", []byte("Hosts Reset!"))
}

// ListHosts list all hosts data
func ListHosts(c *gin.Context) {
	var hostValue string
	for _, v := range hosts {
		hostValue += fmt.Sprintf("%s\t%s\n", v.IP, v.Hostname)
	}
	c.Data(http.StatusOK, "text/plain", []byte(hostValue))
}
