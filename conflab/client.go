package conflab

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"log"
	"encoding/json"
	"os"
)

//config data local cache
var cache = make(map[string]map[string]string)

var Prefix = "/conflab/config/"

//register app to conflab
func Register(appId string) {
	//get ZK_ADDRESS from env
	if s, b := os.LookupEnv("ZK_ADDRESS");b{
		watch(s, appId)
	}else {
		log.Fatalln("can not find ZK_ADDRESS in ENV.")
	}
}

//watch app(zk node)
func watch(server, appId string) {
	c, _, err := zk.Connect([]string{server}, time.Second*10)
	if err != nil {
		panic(err)
	}
	path := Prefix + appId
	bytes, _, event, _ := c.GetW(path)
	//init cache
	refreshCache(appId, bytes)
	//refresh local cache, when node data changed
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

//refresh cache
func refreshCache(appId string, bytes []byte) {
	var app = make(map[string]string)
	log.Println(appId + ":" + string(bytes))
	json.Unmarshal(bytes, &app)
	cache[appId] = app
}
