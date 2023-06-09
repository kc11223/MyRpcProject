package test

import (
	"fmt"
	"service_discovery/model"
	"strings"
	"testing"
)

var req1 = &model.ArgRequest{Zone: "zone1", AppID: "com.xx.testapp", Env: "test", Hostname: "myhost1", Addrs: []string{"http://1.1.1.1/testapp"}, Status: 1}
var req2 = &model.ArgRequest{Zone: "zone1", AppID: "com.xx.testapp", Env: "test", Hostname: "myhost2", Addrs: []string{"http://2.2.2.2/testapp"}, Status: 1}

var r = model.NewRegistry(&model.ConnOption{})
var instance1 = model.NewInstance(req1)
var instance2 = model.NewInstance(req2)

func TestRegister(t *testing.T) {
	app, _ := r.Register(instance1, req1.LatestTimestamp)
	t.Log(app)
}

func TestFetch(t *testing.T) {
	r.Register(instance1, req1.LatestTimestamp)
	r.Register(instance2, req2.LatestTimestamp)
	rs, err := r.Fetch(req1.Zone, req1.Env, req1.AppID, req1.Status, 0)
	if err != nil {
		t.Error(err)
		return
	}
	for _, instance := range rs.Instances {
		t.Log(fmt.Sprintf("AppID:%s,env:%s,hostname:%s,addrs:%s\n",
			instance.AppID,
			instance.Env,
			instance.Hostname,
			strings.Join(instance.Addrs, " ")))
	}
}

func TestCancel(t *testing.T) {
	r.Register(instance1, req1.LatestTimestamp)
	_, err := r.Cancel(req1.Env, req1.AppID, req1.Hostname, 0)
	if err != nil {
		t.Error(err)
		return
	}
	rs, err := r.Fetch(req1.Zone, req1.Env, req1.AppID, req1.Status, 0)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rs)
}

func TestRenew(t *testing.T) {
	r.Register(instance1, req1.LatestTimestamp)
	instance, err := r.Renew(req1.Env, req1.AppID, req1.Hostname)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(instance)
}
