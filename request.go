package request

import (
	"fmt"
	"io/ioutil"
	"time"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

var (
	sessionCookies []*http.Cookie
	sessionJar     *cookiejar.Jar
)

func init() {
	sessionCookies = nil
	sessionJar, _ = cookiejar.New(nil)
}

func Get(url string, headers map[string]string) (*Response, error) {

	client := &http.Client{
		Jar:           sessionJar,
		Timeout: 30 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("Erro ao iniciar requisição: %w", err)
	}

	for v, k := range headers {
		req.Header.Set(v, k)
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Erro ao finalizar requisição: %w", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("Erro ao decodificar resosta da requisição: %w", err)
	}

	if res.StatusCode == http.StatusFound {
		return &Response{res.Status, res.StatusCode, string(body), res.Header}, nil
	}

	sessionCookies = sessionJar.Cookies(req.URL)

	return &Response{res.Status, res.StatusCode, string(body), res.Header}, nil

}

func Post(url, postdata string, headers map[string]string) (*Response, error) {

	client := &http.Client{
		Jar:           sessionJar,
		Timeout: 30 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(postdata))

	if err != nil {
		return nil, fmt.Errorf("Erro ao iniciar requisição: %w", err)
	}

	for v, k := range headers {
		req.Header.Set(v, k)
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Erro ao finalizar requisição: %w", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("Erro ao decodificar resosta da requisição: %w", err)
	}

	if res.StatusCode == http.StatusFound {
		return &Response{res.Status, res.StatusCode, string(body), res.Header}, nil
	}

	sessionCookies = sessionJar.Cookies(req.URL)

	return &Response{res.Status, res.StatusCode, string(body), res.Header}, nil

}

func Delete(url string, headers map[string]string) (*Response, error) {

	client := &http.Client{
		Jar:           sessionJar,
		Timeout: 30 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return nil, fmt.Errorf("Erro ao iniciar requisição: %w", err)
	}

	for v, k := range headers {
		req.Header.Set(v, k)
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Erro ao finalizar requisição: %w", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("Erro ao decodificar resosta da requisição: %w", err)
	}

	if res.StatusCode == http.StatusFound {
		return &Response{res.Status, res.StatusCode, string(body), res.Header}, nil
	}

	sessionCookies = sessionJar.Cookies(req.URL)

	return &Response{res.Status, res.StatusCode, string(body), res.Header}, nil

}

func Put(url string, headers map[string]string) (*Response, error) {

	client := &http.Client{
		Jar:           sessionJar,
		Timeout: 30 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	req, err := http.NewRequest("PUT", url, nil)

	if err != nil {
		return nil, fmt.Errorf("Erro ao iniciar requisição: %w", err)
	}

	for v, k := range headers {
		req.Header.Set(v, k)
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Erro ao finalizar requisição: %w", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("Erro ao decodificar resosta da requisição: %w", err)
	}

	if res.StatusCode == http.StatusFound {
		return &Response{res.Status, res.StatusCode, string(body), res.Header}, nil
	}

	sessionCookies = sessionJar.Cookies(req.URL)

	return &Response{res.Status, res.StatusCode, string(body), res.Header}, nil

}

func Options(url string, headers map[string]string) (*Response, error) {

	client := &http.Client{
		Jar:           sessionJar,
		Timeout: 30 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	req, err := http.NewRequest("OPTIONS", url, nil)

	if err != nil {
		return nil, fmt.Errorf("Erro ao iniciar requisição: %w", err)
	}

	for v, k := range headers {
		req.Header.Set(v, k)
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Erro ao finalizar requisição: %w", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("Erro ao decodificar resosta da requisição: %w", err)
	}

	if res.StatusCode == http.StatusFound {
		return &Response{res.Status, res.StatusCode, string(body), res.Header}, nil
	}

	sessionCookies = sessionJar.Cookies(req.URL)

	return &Response{res.Status, res.StatusCode, string(body), res.Header}, nil

}

func Do(method, urli, postdata string, headers map[string]string, proxy string) (*Response, error) {

	transport := &http.Transport{}

	if proxy != "" {
		parse, _ := url.Parse(proxy)
		transport = &http.Transport{Proxy: http.ProxyURL(parse)}
	}

	client := &http.Client{
		Jar:           sessionJar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Transport:     transport,
	}

	req, err := http.NewRequest(strings.ToUpper(method), urli, strings.NewReader(postdata))

	if err != nil {
		return nil, fmt.Errorf("Erro ao iniciar requisição: %w", err)
	}

	for v, k := range headers {
		req.Header.Set(v, k)
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Erro ao finalizar requisição: %w", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("Erro ao decodificar resosta da requisição: %w", err)
	}

	if res.StatusCode == http.StatusFound {
		return &Response{res.Status, res.StatusCode, string(body), res.Header}, nil
	}

	sessionCookies = sessionJar.Cookies(req.URL)

	return &Response{res.Status, res.StatusCode, string(body), res.Header}, nil

}

func DeleteCookies() {

	sessionJar, _ = cookiejar.New(nil)

}
