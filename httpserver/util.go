package httpserver

import (
	"encoding/json"
	"inyuapp.com/golang/log"

	"io/ioutil"

	"net/http"
	"strconv"
)

func Getfilter(req *http.Request) (filter map[string]interface{}) {
	filter = make(map[string]interface{})
	for k, vs := range req.URL.Query() {
		if len(vs) > 1 || len(vs) < 1 {
			log.Warn(log.Fields{
				"req.URL.Query()": "len(vs)>1",
				"values":          vs,
			})
			continue
		}
		v := vs[0]
		intv, err := strconv.Atoi(v)
		if err == nil {
			filter[k] = intv
			continue
		}
		floatv, err := strconv.ParseFloat(v, 64)
		if err == nil {
			filter[k] = floatv
			continue
		}
		boolv, err := strconv.ParseBool(v)
		if err == nil {
			filter[k] = boolv
			continue
		}
		filter[k] = v
	}
	return
}

func Options(req *http.Request, resp http.ResponseWriter, contenttype, server, methods string) {
	Header(resp, contenttype, server, methods)
	resp.Write([]byte(""))
}

func Header(resp http.ResponseWriter, contenttype, server, methods string) {

	resp.Header().Set("Access-Control-Allow-Origin", "*")                   //允许访问所有域
	resp.Header().Add("Access-Control-Allow-Headers", "Content-Type,token") //header的类型
	resp.Header().Set("content-type", contenttype)
	resp.Header().Set("HttpServer", server)

	resp.Header().Set("Access-Control-Allow-Methods", methods)
	resp.Header().Set("Access-Control-Allow-Credentials", "true")
}
func To(req *http.Request, s interface{}) (err error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		return err
	}
	return
}
