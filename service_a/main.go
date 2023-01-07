package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
)

const (
	XOtSpanContext  = "X-Ot-Span-Context"
	XRequestId      = "X-Request-Id"
	XB3TraceId      = "X-B3-TraceId"
	XB3SpanId       = "X-B3-SpanId"
	XB3ParentSpanId = "X-B3-ParentSpanId"
	XB3Sampled      = "X-B3-Sampled"
	XB3Flags        = "X-B3-Flags"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Calling Service B")

	req, err := http.NewRequest("GET", "http://service_a_envoy:8788/", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}

	req.Header.Add(XRequestId, r.Header.Get(XRequestId))
	req.Header.Add(XB3TraceId, r.Header.Get(XB3TraceId))
	req.Header.Add(XB3SpanId, r.Header.Get(XB3SpanId))
	req.Header.Add(XB3ParentSpanId, r.Header.Get(XB3ParentSpanId))
	req.Header.Add(XB3Sampled, r.Header.Get(XB3Sampled))
	req.Header.Add(XB3Flags, r.Header.Get(XB3Flags))
	req.Header.Add(XOtSpanContext, r.Header.Get(XOtSpanContext))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Fprintf(w, string(body))
	fmt.Fprintf(w, "Hello from service A")

	req, err = http.NewRequest("GET", "http://service_a_envoy:8791/", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}

	req.Header.Add(XRequestId, r.Header.Get(XRequestId))
	req.Header.Add(XB3TraceId, r.Header.Get(XB3TraceId))
	req.Header.Add(XB3SpanId, r.Header.Get(XB3SpanId))
	req.Header.Add(XB3ParentSpanId, r.Header.Get(XB3ParentSpanId))
	req.Header.Add(XB3Sampled, r.Header.Get(XB3Sampled))
	req.Header.Add(XB3Flags, r.Header.Get(XB3Flags))
	req.Header.Add(XOtSpanContext, r.Header.Get(XOtSpanContext))

	client = &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Fprintf(w, string(body))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
