package main

import (
	"net/http"
	"net/url"
	"testing"
)

func TestAgentLoginFunc(t *testing.T) {
	form := url.Values{}
	form.Add("name", "tester")
	form.Add("token", "123")
	resp, err := http.PostForm("http://localhost:9090/api_agent/login", form)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Error(resp.Status)
	}
}

func TestAgentQiangZhuangFunc(t *testing.T) {
	resp, err := http.PostForm("http://localhost:9090/api_agent/qiangzhuang", url.Values{})
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Error(resp.Status)
	}
}
