package pipe

import (
	"net/url"

	"github.com/rishiloyola/zoop-proxy/Godeps/_workspace/src/github.com/rishiloyola/zoop-client"
	"github.com/rishiloyola/zoop-proxy/Godeps/_workspace/src/stathat.com/c/consistent"
)

type PipeClient struct{}

var zkStreamClient *zkClient.Client
var zkHttpClient *zkClient.Client

func (p *PipeClient) Init(zkHTTPpath string, zkStreampath string, zkIP string) error {
	var err error
	zkStreamClient = zkClient.New()
	zkHttpClient = zkClient.New()

	err = zkStreamClient.Connect(zkIP)
	if err != nil {
		panic(err)
		return err
	}
	err = zkHttpClient.Connect(zkIP)
	if err != nil {
		panic(err)
		return err
	}

	zkHttpClient.GetWatch(zkHTTPpath)
	zkStreamClient.GetWatch(zkStreampath)

	return nil
}

func (p *PipeClient) GetServers(path *url.URL) *consistent.Consistent {
	streamReq := identifyReqType(path)
	if streamReq {
		return zkStreamClient.GetHash()
	} else {
		return zkHttpClient.GetHash()
	}

}

func identifyReqType(path *url.URL) bool {
	query := path.Query()
	if len(query) == 0 {
		return false
	} else {
		if _, ok := query["stream"]; ok {
			return true
		} else {
			return false
		}
	}
}
