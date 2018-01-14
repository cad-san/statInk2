package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/cad-san/statink2"
	"github.com/pkg/errors"
)

var rankList = []string{
	"c-",
	"c",
	"c+",
	"b-",
	"b",
	"b+",
	"a-",
	"a",
	"a+",
	"s",
	"s+",
}

type config struct {
	APIKey string
}

type options struct {
	configPath string
}

type history struct {
	Rule       statink2.Rule
	Mode       statink2.Mode
	Weapon     statink2.Weapon
	Stage      statink2.Stage
	Rank       string
	GachiPower int
	Day        string
	Time       string
	Duration   string
}

type appHandler struct {
	name    string
	version string
	client  *statink2.Client
	prev    history
}

func parseFlag() options {
	var opt options
	flag.StringVar(&opt.configPath, "config", "config.json", "file path")
	flag.StringVar(&opt.configPath, "c", "config.json", "file path")
	flag.Parse()
	return opt
}

func getFormString(values url.Values, key string) (string, error) {
	val, ok := values[key]
	if !ok {
		return "", errors.New("no form value")
	}
	if len(val[0]) == 0 {
		return "", errors.New("no form value")
	}
	return val[0], nil
}

func getFormInt(values url.Values, key string) (int, error) {
	val, ok := values[key]
	if !ok {
		return -1, errors.New("no form value")
	}
	if len(val[0]) == 0 {
		return -1, errors.New("no form value")
	}
	return strconv.Atoi(val[0])

}

func getFormTime(values url.Values, dayKey string, timeKey string) (*time.Time, error) {
	dayStr, err := getFormString(values, dayKey)
	if err != nil {
		return nil, err
	}
	timeStr, err := getFormString(values, timeKey)
	if err != nil {
		return nil, err
	}
	dateStr := dayStr + " " + timeStr
	date, err := time.ParseInLocation("2006-01-02 15:04", dateStr, time.Local)
	return &date, err
}

func getFormDuration(values url.Values, durationKey string) (*time.Duration, error) {
	durStr, err := getFormString(values, durationKey)
	if err != nil {
		return nil, err
	}
	t, err := time.Parse("15:04:05", durStr)
	if err != nil {
		return nil, err
	}
	h, m, s := t.Clock()
	d := time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(s)*time.Second
	return &d, nil
}

func parseForm(r *http.Request) (*statink2.Battle, error) {
	var result statink2.Battle
	err := r.ParseForm()
	if err != nil {
		return nil, errors.Wrap(err, "parse form failed")
	}
	var str string
	// mode
	str, err = getFormString(r.Form, "mode")
	if err != nil {
		return nil, errors.Wrap(err, "form value[mode] is invald")
	}
	result.Mode = statink2.Mode(str)
	// rule
	str, err = getFormString(r.Form, "rule")
	if err != nil {
		return nil, errors.Wrap(err, "form value[rulemode] is invald")
	}
	result.Rule = statink2.Rule(str)
	// lobby
	str, err = getFormString(r.Form, "lobby")
	if err != nil {
		return nil, errors.Wrap(err, "form value[lobby] is invald")
	}
	result.Lobby = statink2.Lobby(str)
	// weapon
	str, err = getFormString(r.Form, "weapon")
	if err != nil {
		return nil, errors.Wrap(err, "form value[weapon] is invald")
	}
	result.Weapon = statink2.Weapon(str)
	// stage
	str, err = getFormString(r.Form, "stage")
	if err != nil {
		return nil, errors.Wrap(err, "form value[stage] is invald")
	}
	result.Stage = statink2.Stage(str)
	// startAt
	startTime, err := getFormTime(r.Form, "date", "start_time")
	if err != nil {
		return nil, errors.Wrap(err, "form value[date] is invald")
	}
	result.StartAt = startTime.Unix()
	// endAt
	duration, err := getFormDuration(r.Form, "duration")
	if err != nil {
		return nil, errors.Wrap(err, "form value[duration] is invald")
	}
	result.EndAt = startTime.Add(*duration).Unix()
	// gachi_power
	result.GachiPower, err = getFormInt(r.Form, "gachi_power")
	if err != nil {
		return nil, errors.Wrap(err, "form value[gachi_power] is invald")
	}
	// result
	result.Result, err = getFormString(r.Form, "result")
	if err != nil {
		return nil, errors.Wrap(err, "form value[result] is invald")
	}
	// knockout
	result.KnockOut, err = getFormString(r.Form, "knockout")
	if err != nil {
		return nil, errors.Wrap(err, "form value[knockout] is invald")
	}
	// kill_or_assist
	result.KillOrAssist, err = getFormInt(r.Form, "kill_or_assist")
	if err != nil {
		return nil, errors.Wrap(err, "form value[kill_or_assist] is invald")
	}
	// kill
	assist, err := getFormInt(r.Form, "assist")
	if err != nil {
		return nil, errors.Wrap(err, "form value[assist] is invald")
	}
	result.Kill = result.KillOrAssist - assist
	// death
	result.Death, err = getFormInt(r.Form, "death")
	if err != nil {
		return nil, errors.Wrap(err, "form value[death] is invald")
	}
	// special
	result.Special, err = getFormInt(r.Form, "special")
	if err != nil {
		return nil, errors.Wrap(err, "form value[special] is invald")
	}
	// my_team_count
	result.MyTeamCount, err = getFormInt(r.Form, "my_team_count")
	if err != nil {
		return nil, errors.Wrap(err, "form value[my_team_count] is invald")
	}
	// his_team_count
	result.HisTeamCount, err = getFormInt(r.Form, "his_team_count")
	if err != nil {
		return nil, errors.Wrap(err, "form value[his_team_count] is invald")
	}

	if result.KnockOut == statink2.Yes {
		if result.Result == statink2.ResultWin {
			result.MyTeamCount = 100
			result.HisTeamCount = 0
		} else {
			result.MyTeamCount = 0
			result.HisTeamCount = 100
		}
	}

	// rank_in_team
	result.RankInTeam, err = getFormInt(r.Form, "rank_in_team")
	if err != nil {
		return nil, errors.Wrap(err, "form value[rank_in_team] is invald")
	}
	result.Rank, err = getFormString(r.Form, "rank")
	if err != nil {
		return nil, errors.Wrap(err, "form value[rank] is invald")
	}
	result.RankAfter, err = getFormString(r.Form, "rank_after")
	if err != nil {
		return nil, errors.Wrap(err, "form value[rank_after] is invald")
	}
	return &result, nil
}

func (a *appHandler) battleForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("battle.html")
	if err != nil {
		a.client.Logger.Printf("template parse failed. err:%v", err)
		return
	}
	var params = map[string]interface{}{
		"Stages":  statink2.StageList,
		"Weapons": statink2.WeaponList,
		"Ranks":   rankList,
		"History": a.prev,
	}
	err = tmpl.Execute(w, params)
	if err != nil {
		a.client.Logger.Printf("template execute failed. err:%v", err)
		return
	}
}

func (a *appHandler) postBattle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("url : %v", r.Form)
	result, err := parseForm(r)
	if err != nil {
		a.client.Logger.Printf("parse form failed. err:%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result.Agent = a.name
	result.AgentVersion = a.version
	a.prev = history{
		Mode:       result.Mode,
		Rule:       result.Rule,
		Stage:      result.Stage,
		Weapon:     result.Weapon,
		GachiPower: result.GachiPower,
		Rank:       result.Rank,
	}
	a.prev.Day, _ = getFormString(r.Form, "date")
	a.prev.Time, _ = getFormString(r.Form, "start_time")
	a.prev.Duration, _ = getFormString(r.Form, "duration")
	err = a.client.PostBattle(nil, *result)
	if err != nil {
		a.client.Logger.Printf("post battle failed. err:%v", err)
	}
	http.Redirect(w, r, "/", 301)
}

func (a *appHandler) index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		a.battleForm(w, r)
	} else if r.Method == "POST" {
		a.postBattle(w, r)
	}
}

func main() {
	opt := parseFlag()
	if len(opt.configPath) == 0 {
		opt.configPath = "config.json"
	}
	file, err := os.Open(opt.configPath)
	if err != nil {
		fmt.Printf("cannot read config file:%v", err)
		os.Exit(1)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var conf config
	if err = decoder.Decode(&conf); err != nil {
		fmt.Printf("connt read config file:%v", err)
		os.Exit(1)
	}

	logger := log.New(os.Stderr, "", log.LstdFlags|log.Llongfile)
	client, err := statink2.NewClient(conf.APIKey, logger)
	if err != nil {
		os.Exit(1)
	}
	now := time.Now()
	app := appHandler{
		name:    "go-statink2",
		version: "0.1.0",
		client:  client,
		prev: history{
			Day:      now.Format("2006-01-02"),
			Time:     now.Format("15:04"),
			Duration: "00:05:00",
		},
	}

	http.HandleFunc("/", app.index)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("server failed. err:%v", err)
		os.Exit(1)
	}
}
