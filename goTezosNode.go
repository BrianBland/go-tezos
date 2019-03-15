package goTezos

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type TezosRPCClient struct {
	Host        string
	Port        string
	logfunction func(level, msg string)
	logger      *log.Logger
	isWebClient bool
	httpClient  *http.Client
}

// Create a new RPC client using the specified hostname and port.
// Also acceptable is the hostname of a web-endpoint that supports https
func NewTezosRPCClient(hostname string, port string) *TezosRPCClient {
	t := TezosRPCClient{}

	// Strip off posible trailing '/'
	hLen := len(hostname)
	if hostname[hLen-1] == '/' {
		hostname = hostname[:hLen-1]
	}

	// Strip off URI scheme
	if hostname[:8] == "https://" {
		hostname = hostname[8:]
		t.isWebClient = true
	} else if hostname[:7] == "http://" {
		hostname = hostname[7:]
	}

	t.Host = hostname
	t.Port = port
	t.logfunction = func(level, msg string) {
		fmt.Println(level + ": " + msg)
	}
	t.SetLogger(log.New(os.Stdout, hostname, 0))

	var netTransport = &http.Transport{ // TODO make this as config option, but with defaults like this
		Dial: (&net.Dialer{
			Timeout: 3 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 3 * time.Second,
	}

	t.httpClient = &http.Client{
		Timeout:   time.Second * 3,
		Transport: netTransport,
	}

	return &t
}

//Set the logger for the RPC Client
func (this *TezosRPCClient) SetLogger(log *log.Logger) {
	this.logger = log
}

func (this *TezosRPCClient) IsWebClient(b bool) {
	this.isWebClient = b
}

func (this *TezosRPCClient) GetResponse(method string, path string, args string) (ResponseRaw, error) {

	var url string

	if this.isWebClient {
		url = fmt.Sprintf("https://%s:%s%s", this.Host, this.Port, path)
	} else {
		url = fmt.Sprintf("http://%s:%s%s", this.Host, this.Port, path)
	}

	var body io.Reader

	if method != "GET" {
		var jsonStr = []byte(args)
		body = bytes.NewReader(jsonStr)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		this.logger.Println("Error in GetResponse: " + err.Error())
		return ResponseRaw{}, err
	}

	resp, err := this.httpClient.Do(req)
	if err != nil {
		this.logger.Println("Error in GetResponse: " + err.Error())
		return ResponseRaw{}, err
	}
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		this.logger.Println("Error in GetResponse - readAll bytes: " + err.Error())
		return ResponseRaw{}, err
	}
	defer resp.Body.Close()
	return ResponseRaw{b}, nil
}

//A function just to perform a query to see if an RPC Client's endpoint is alive (heartbeat)
func (this *TezosRPCClient) Healthcheck() bool {
	_, err := this.GetResponse("GET", "/chains/main/blocks", "")
	if err == nil {
		return true // healthy
	}
	return false // unhelaty
}
