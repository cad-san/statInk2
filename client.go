package statink2

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

// Client is a http client for stat.ink (Splatoon2 version)
type Client struct {
	hc     *http.Client
	URL    *url.URL
	APIKey string
	Logger *log.Logger
}

// Battle is battle infomation
type Battle struct {
	// rule info
	Lobby Lobby `json:"lobby"`
	Mode  Mode  `json:"mode"`
	Rule  Rule  `json:"rule"`

	// match info
	Stage      Stage `json:"stage"`
	GachiPower int   `json:"estimate_gachi_power"`
	StartAt    int64 `json:"start_at"`
	EndAt      int64 `json:"end_at"`

	// result
	Result       string `json:"result"`
	KnockOut     string `json:"knock_out"`
	MyTeamCount  int    `json:"my_team_count"`
	HisTeamCount int    `json:"his_team_count"`

	// player info
	Weapon       Weapon `json:"weapon"`
	RankInTeam   int    `json:"rank_in_team"`
	KillOrAssist int    `json:"kill_or_assist"`
	Kill         int    `json:"kill"`
	Death        int    `json:"death"`
	Special      int    `json:"special"`
	Rank         string `json:"rank"`
	RankExp      int    `json:"rank_exp"`
	RankAfter    string `json:"rank_after"`
	RankAfterExp int    `json:"rank_exp_after"`
}

const (
	statInkDomainURL = "https://stat.ink"
	battleAPI        = "api/v2/battle"
	version          = "0.1"
)

var userAgent = fmt.Sprintf("stat.ink GoClient/%s (%s", version, runtime.Version())

// NewClient generates Client, http client object for stat.ink
// It provides a http client with empty cookiejar.
func NewClient(key string, logger *log.Logger) (*Client, error) {
	return newClient(statInkDomainURL, key, logger)
}

func newClient(domain, key string, logger *log.Logger) (*Client, error) {
	client := &Client{
		hc: &http.Client{},
	}
	var err error
	client.URL, err = url.ParseRequestURI(domain)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url : %s", domain)
	}

	if len(key) == 0 {
		return nil, errors.New("missing api key")
	}
	client.APIKey = key

	if logger == nil {
		client.Logger = log.New(ioutil.Discard, "", log.LstdFlags)
	} else {
		client.Logger = logger
	}

	return client, nil
}

// PostBattle send HTTP POST request to stat.ink
func (c *Client) PostBattle(ctx context.Context, battle Battle) error {
	json, err := json.Marshal(battle)
	c.Logger.Printf("request body : %v", string(json))
	if err != nil {
		return errors.Wrap(err, "cannot marshal battle struct")
	}
	req, err := c.newRequest(ctx, "POST", battleAPI, strings.NewReader(string(json)))
	if err != nil {
		return errors.Wrap(err, "cannot create Request for POST")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.hc.Do(req)
	if err != nil {
		return errors.Wrap(err, "http request failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 > 2 {
		body, _ := ioutil.ReadAll(resp.Body)
		c.Logger.Printf("error : %s\n", body)
		return errors.Errorf("request failed. URL : %v, status code:%d)", req.URL, resp.StatusCode)
	}
	return nil
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}, f *os.File) error {
	if f != nil {
		resp.Body = ioutil.NopCloser(io.TeeReader(resp.Body, f))
		defer f.Close()
	}

	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
