package conflab

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"log"
	"encoding/json"
)

var cache = make(map[string]map[string]string)

var Server = "172.30.10.161"
var Prefix = "/conflab/config/"

func Register(appId string) {
	watch(Server, appId)
}
func watch(server, appId string) {
	c, _, err := zk.Connect([]string{server}, time.Second*10)
	if err != nil {
		panic(err)
	}
	path := Prefix + appId
	bytes, _, event, _ := c.GetW(path)
	refreshCache(appId, bytes)
	go func() {
		for {
			ev := <-event
			if ev.Type == zk.EventNodeDataChanged {
				if bt, _, ev, e := c.GetW(path); e == nil {
					event = ev
					refreshCache(appId, bt)
				} else {
					panic(e)
				}
			}
			if ev.Err != nil {
				log.Fatalf("zk watcher error %+v", ev.Err)
			}
		}
	}()
	log.Println("watch " + server + path)
}

func refreshCache(appId string, bytes []byte) {
	var app = make(map[string]string)
	log.Println(appId + ":" + string(bytes))
	json.Unmarshal(bytes, &app)
	cache[appId] = app
}
