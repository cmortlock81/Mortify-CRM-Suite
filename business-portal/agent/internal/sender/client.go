package sender
import("bytes";"crypto/tls";"encoding/json";"fmt";"net/http";"strings";"time")
type Client struct{APIURL,Token string; HTTP *http.Client}
func New(api,token string,verify bool)*Client{return &Client{strings.TrimRight(api,"/"),token,&http.Client{Timeout:30*time.Second,Transport:&http.Transport{TLSClientConfig:&tls.Config{InsecureSkipVerify:!verify}}}}}
func (c *Client) Post(path string,payload any) error{b,_:=json.Marshal(payload); req,_:=http.NewRequest("POST",c.APIURL+"/api"+path,bytes.NewReader(b)); req.Header.Set("Content-Type","application/json"); req.Header.Set("Authorization","Bearer "+c.Token); resp,err:=c.HTTP.Do(req); if err!=nil{return err}; defer resp.Body.Close(); if resp.StatusCode>=300{return fmt.Errorf("api status %d",resp.StatusCode)}; return nil}
