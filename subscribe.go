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
	subStruct := new(comm.SubStruct)
	subStruct.Id = strconv.Itoa(time.Now().Nanosecond())
	subStruct.Sub = conf.LtcTopic.KLineTopicDesc
	reqBytes , err := json.Marshal(subStruct)
	if err!=nil {
		return
	}
	if err := ws.WriteMessage(websocket.TextMessage,reqBytes); err == nil {

	}else{
		fmt.Errorf("send req response error %s",err.Error())
	}
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

				}else{
					fmt.Errorf("huobi server ping client error %s",err.Error())
					continue
				}
			}
			if  _, ok := resMap["subbed"];ok  {
				var resStruct comm.ResStruct
				json.Unmarshal(res,&resStruct)
				//resStruct.Status
				fmt.Println(resStruct)
			}

			if  _, ok := resMap["ch"];ok  {
				var resStruct comm.ResStruct
				json.Unmarshal(res,&resStruct)
				//resStruct.Status
				fmt.Println(resStruct)
				var ddRobot comm.DDRobotStruct;
				//append(ddRobot.At.Atmobiles, "18210048936")
				ddRobot.Msgtype = "text"
				if resStruct.Tick.Open < 290 {
					ddRobot.Text.Content = "莱特币已经" + strconv.FormatFloat(resStruct.Tick.Open,'f',30,32)+"了！"
					comm.SendDDRobot(ddRobot)
				}
				//if resStruct.Tick.Open > 320 {
				//	ddRobot.Text.Content = "ltc已经大于320了"
				//	comm.SendDDRobot(ddRobot)
				//}
			}

		}
	}
}

