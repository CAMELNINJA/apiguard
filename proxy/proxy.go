package proxy

import (
    "net/http"
    "net/http/httputil"
    "net/url"
)

func NewReverseProxy(target string) (*httputil.ReverseProxy, error) {
    parsedURL, err := url.Parse(target)
    if err != nil {
        return nil, err
    }

    proxy := httputil.NewSingleHostReverseProxy(parsedURL)

    originalDirector := proxy.Director
    proxy.Director = func(req *http.Request) {
        originalDirector(req)
        req.Host = parsedURL.Host
    }

    return proxy, nil
}
