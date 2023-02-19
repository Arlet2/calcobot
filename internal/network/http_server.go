package network

import (
	"calcobot/internal/database"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type httpServer struct{
	db database.Database
	port int
}

func StartHttpServer(path string, db database.Database) {
	httpServer := httpServer{db: db, port: 8000}

	http.HandleFunc(path, httpServer.handler)

	log.Fatal(http.ListenAndServe("localhost:"+strconv.FormatInt(int64(httpServer.port), 10), nil))
}

func (httpServer httpServer) handler(writer http.ResponseWriter, req *http.Request) {

	if req.URL.Query().Has("username") {
		logs, err := httpServer.db.GetLogsByUsername(req.URL.Query().Get("username"))

		if err != nil {
			fmt.Fprint(writer, err)
			return
		}

		fmt.Fprintln(writer, "Data from "+req.URL.Query().Get("username")+":")

		for _, value := range logs {
			fmt.Fprintf(writer, "id: %d, time: %s, request: %s, answer: %f\n",
								value.Id, value.Time, value.Request, value.Answer)
		}

	} else {
		fmt.Fprint(writer, "Need query variable username")
	}

	
}