package conflab

import (
	"testing"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func TestGet(t *testing.T) {
	Server = "172.30.10.161"
	Prefix = "/conflab/config/"
	appId := "a"
	Register(appId)
	go func() {
		for {
			time.Sleep(time.Second)
			println("---" + Get("a"))
		}
	}()
	time.Sleep(time.Second * 30)
}

func TestGetBoolean(t *testing.T) {
	Server = "172.30.10.161"
	Prefix = "/conflab/config/"
	appId := "a"
	Register(appId)
	go func() {
		for {
			time.Sleep(time.Second)
			b := GetBooleanDefault("b", false)
			println(b)
		}
	}()
	time.Sleep(time.Second * 5)
}

func TestGetInt(t *testing.T) {
	Server = "172.30.10.161"
	Prefix = "/conflab/config/"
	appId := "a"
	Register(appId)
	go func() {
		for {
			time.Sleep(time.Second)
			i := GetIntDefault("i",-1)
			println(i)
		}
	}()
	time.Sleep(time.Second * 30)
}

func WorldACL(perms int32) []zk.ACL {
	return []zk.ACL{{perms, "world", "anyone"}}
}
