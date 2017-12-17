package main

import (
	"github.com/gorilla/websocket"
	"fmt"
	"huobiapi/comm"
	//"golang.org/x/net/websocket"
	"time"
	"strconv"
	"huobiapi/conf"
	"encoding/json"
)

func main()  {
	dialer := websocket.Dialer{ /* set fields as needed */ }
	ws, _, err := dialer.Dial("wss://api.huobi.pro/ws", nil)
	if err != nil {
		// handle error
	}
	//if err := ws.WriteMessage(websocket.TextMessage, []byte("{\"ping\":112112121200000}")); err == nil {
	//	//_, p, err := ws.ReadMessage()
	//	//if err == nil {
	//	//	res := UnGzip(p)
	//	//	fmt.Println("the message is %s", string(res))
	//	//}
	//}
	for {
		if _,p,err := ws.ReadMessage();err == nil {
			res := comm.UnGzip(p)
			fmt.Println(string(res))
			resMap := comm.JsonDecodeByte(res)
			if  v, ok := resMap["ping"];ok  {
				pingMap := make(map[string]interface{})
				pingMap["pong"] = v
				pingParams := comm.JsonEncodeMapToByte(pingMap)
				if err := ws.WriteMessage(websocket.TextMessage, pingParams); err == nil {

					reqMap := new(comm.ReqStruct)
					reqMap.Id = strconv.Itoa(time.Now().Nanosecond())
					reqMap.Req = conf.LtcTopic.KLineTopicDesc
					reqBytes , err := json.Marshal(reqMap)
					if err!=nil {
						continue
					}
					if err := ws.WriteMessage(websocket.TextMessage,reqBytes); err == nil {

					}else{
						fmt.Errorf("send req response error %s",err.Error())
					}
				}else{
					fmt.Errorf("huobi server ping client error %s",err.Error())
					continue
				}
			}
			if  _, ok := resMap["rep"];ok  {
				var resStruct comm.ResStruct
				json.Unmarshal(res,&resStruct)
				//resStruct.Status
				fmt.Println(resStruct)
			}
			//if  v, ok := resMap["pong"];ok  {
			//	pingMap := make(map[string]interface{})
			//	pingMap["ping"] = v
			//	pingParams := comm.JsonEncodeToByte(pingMap)
			//	if err := ws.WriteMessage(websocket.TextMessage, pingParams); err == nil {
			//		if err != nil {
			//			fmt.Errorf("huobi server ping client error %s",err.Error())
			//		}
			//	}
			//}

		}
	}
}

