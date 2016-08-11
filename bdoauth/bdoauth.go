package bdoauth

import "net/http"
import "net/url"
import "encoding/json"
import "io/ioutil"

func RequestClientCredentials(id,secret string) map[string]interface{} {
	u, _ := url.Parse("https://openapi.baidu.com/oauth/2.0/token")
	q := u.Query()
	q.Set("grant_type","client_credentials")
	q.Set("client_id",id)
	q.Set("client_secret",secret)
	u.RawQuery = q.Encode()
	res,err := http.Get(u.String())
	if err!=nil {
		return nil
	}
	ret,err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err!=nil {
		return nil
	}
	values := make(map[string]interface{})
	err = json.Unmarshal(ret,&values)
	return values
}
