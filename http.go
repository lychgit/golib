package lib

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) ([]byte, error) {
	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// 处理错误
		return nil, err
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	resp, err := client.Do(req)
	if err != nil {
		// 处理错误
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func HttpPost() (error) {

	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	return nil
}


func HttpsGet(url string) ([]byte, error) {
	//要管理代理、TLS配置、keep-alive、压缩和其他设置，创建一个Transport
	//Client和Transport类型都可以安全的被多个go程同时使用。出于效率考虑，应该一次建立、尽量重用。
	tr := &http.Transport{
		//InsecureSkipVerify用来控制客户端是否证书和服务器主机名。如果设置为true,
		//则不会校验证书以及证书中的主机名和服务器主机名是否一致。
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// 处理错误
		return nil, err
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	resp, err := client.Do(req)
	if err != nil {
		// 处理错误
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
