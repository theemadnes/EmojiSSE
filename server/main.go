package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/enescakir/emoji"
	"golang.org/x/exp/maps"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// set default port number if env var $PORT isn't set
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// get a random Emoji
func getRandomEmoji() string {
	// set random seed
	//rand.Seed(time.Now().UnixNano())
	emojiMap := emoji.Map()
	var emojiValues []string = maps.Values(emojiMap)
	//emojiValues[1]
	//randomEmoji := emojiValues[0]
	return emojiValues[0]
}

func main() {
	H2CServerUpgrade()
	//H2CServerPrior()
}

func checkErr(err error, msg string) {
	if err == nil {
		return
	}
	fmt.Printf("ERROR: %s: %s\n", msg, err)
	os.Exit(1)
}

func H2CServerUpgrade() {
	port := getEnv("PORT", "8080")
	h2s := &http2.Server{}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Got request\n")
		//fmt.Fprintf(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
		fmt.Fprintf(w, "%v", getRandomEmoji())
	})

	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: h2c.NewHandler(handler, h2s),
	}

	fmt.Printf("Listening [0.0.0.0:%s]...\n", port)
	checkErr(server.ListenAndServe(), "while listening")
}
