package netconf

import (
	"log"
)

func main(){
	wait := make(chan bool,1)
	var nc CallhomeListener
	nc.Initialize("admin", "admin")
	log.Println("si")
	<-wait
}
