package main

import (
	"github.com/resurtm/boomak-server/config"
	"fmt"
)

func main() {
	fmt.Println("%+v\n", config.Config())

	//router := hs.NewRouter()

	//LoadConfig()
	//ConnectToDb()
	//InitMailing()
	//InitHttp()
}

/*func InitHttp() {




	listenAddr := fmt.Sprintf("%s:%d", Config.Server.Hostname, Config.Server.Port)
	fmt.Printf("Listening at \"%s\"...\n", listenAddr)

	h1 := SetupCors(r)
	h2 := handlers.LoggingHandler(os.Stdout, h1)
	http.ListenAndServe(listenAddr, h2)
}
*/
