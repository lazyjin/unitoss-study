package main

import (
	"common"
	"common/clog"
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"path"
	"regexp"
	"strconv"
)

type Page struct {
	Title string
	Body  []byte
}

type Response struct {
	Success string      `json:"success"`
	Message string      `json:"message"`
	data    interface{} `json:"data"`
	Error   error       `json:error`
}

const PNAME = "udrgenreq"

var (
	log       = clog.GetLogger()
	rabbitMgr = common.NewRabbitManager()
)

var templates *template.Template //= template.Must(template.ParseFiles("/Users/lazyjin/Developer/unitoss-study/pilot/src/udrgenreq/udrgen.html", "/Users/lazyjin/Developer/unitoss-study/pilot/src/udrgenreq/view.html"))
var validPath = regexp.MustCompile("^/(udrgen|genreq)/$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

func udrGenHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "udrgen")
}

func genReqHandler(w http.ResponseWriter, r *http.Request) {
	errortype, _ := strconv.Atoi(r.FormValue("errortype"))
	count, _ := strconv.Atoi(r.FormValue("count"))
	log.Debugf("%v, %v", errortype, count)

	suc := "true"

	msg := common.UdrReqMsg{
		ErrorType: errortype,
		Count:     count,
	}
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		suc = "false"
		log.Errorf("Fail to Marshal msg : %v", err)
	}

	err = rabbitMgr.PublishToQueue(string(jsonMsg))
	if err != nil {
		suc = "false"
		log.Errorf("Publish queue fail: %v", err)
	}

	json.NewEncoder(w).Encode(Response{
		Success: suc,
		Message: "good",
		Error:   err,
	})

	return
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	templates.Execute(w, nil)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		log.Debugf("m is %v", m)
		if m == nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r)
	}
}

func main() {
	log.Info("Start UDR generating Request Web Server...")

	http.HandleFunc("/udrgen/", makeHandler(udrGenHandler))
	http.HandleFunc("/genreq/", makeHandler(genReqHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init() {
	common.ReadConfigFile(PNAME)
	conf := common.GetConfig()

	// init log
	clog.InitWith(PNAME, conf.Logname, conf.Logdir, conf.Loglevel)

	// rabbit publisher connect
	rabbitMgr.ConnectRabbit(
		conf.Rabbithost,
		conf.Rabbitport,
		conf.Rabbituser,
		conf.Rabbitpw)
	rabbitMgr.UdrSendQueueDeclare(conf.Reqreciever)

	templates = template.Must(template.ParseFiles(
		path.Join(conf.Templatedir, "udrgen.html"),
		path.Join(conf.Templatedir, "header.html")))

	http.Handle("/library/", http.StripPrefix("/library/", http.FileServer(http.Dir("../library"))))
}
