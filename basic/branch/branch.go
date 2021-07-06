package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	//const filename = "abc.txt"
	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}

	//ctx0 := context.WithValue(context.Background(), "age", 18)
	//ctx1 := context.WithValue(ctx0, "age", 20)
	//ctx2 := context.WithValue(ctx0, "name", "gopher")
	//
	//v0 := ctx0.Value("age")
	//v1 := ctx1.Value("age")
	//v2 := ctx2.Value("age")

	//fmt.Println(v0, v1, v2)

	//var t interface{}
	//t = make(chan string, 1)
	//a, ok := t.(int)
	//if !ok {
	//	fmt.Printf("Type assersion failed: t is %T, but not int", t)
	//} else {
	//	fmt.Println(a)
	//}

	errStr := "cds_key=EGRESS|ecom.power.power_schedule:default:boe:|ecom.power.power_gateway:default::|Invoke|prod| reason=request timeout connect_timeout=50ms request_timeout=1000ms real_time=1003037us"
	errStr = "_cds_key=EGRESS|douyin.search.api:default:boe:|data.search.federation:ies:|search|prod|a b"
	//errStr = "abbA-/:..."
	reg := regexp.MustCompile(`[^a-zA-Z0-9_\-./:]+`)
	fmt.Println(reg.MatchString(errStr))
	fmt.Println(reg.FindAllString(errStr, len(errStr)))

	unValidChars := reg.FindAllString(errStr, len(errStr))
	fmt.Println(len(unValidChars))
	seen := map[string]struct{}{}
	for _, uc := range unValidChars {
		if _, ok := seen[uc]; !ok {
			errStr = strings.Replace(errStr, uc, "_", -1)
			seen[uc] = struct{}{}
		}
	}
	fmt.Println(errStr)
	fmt.Println(reg.MatchString(errStr))

	var isRetry bool
	fmt.Println(isRetry)

	var err error
	var state string
	isRetry = true
	err = fmt.Errorf("123")

	switch {
	case isRetry && err == nil:
		state = "mq重试消息任务执行成功但不投递下一轮"
		//_ = metrics.EmitCounter("HandleMQRetryMsgSuccess", 1)
	case isRetry && err != nil:
		state = "mq重试消息任务执行超时"
	default:
		state = "任务已执行但不投递下一轮"
	}
	fmt.Println(state)
	//if strings.Contains(errStr, "cds_key") {
	//	errStr = strings.Split(errStr, "|")[5]
	//	if strings.Contains(errStr, "request timeout") {
	//		errStr = "cds_key=EGRESS/request timeout"
	//	} else {
	//		errStr = "cds_key=EGRESS/other error"
	//	}
	//}
	//errVal:= strings.Split(errStr, ":")[0]
	//if errVal == "" {
	//	errVal = "nil"
	//}
	//errValQualified := strings.Join(strings.Split(errVal, " "), "_")
	//fmt.Println(errValQualified)
}
