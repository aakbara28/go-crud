package logging

import (
	"encoding/json"
	"log"
	"time"
)

type LogEntry struct {
	RequestId   string      `json:"requestId"`
	RequestKey  string      `json:"requestKey"`
	Code        string      `json:"code"`
	Type        string      `json:"type"`
	Method      string      `json:"method"`
	Function    string      `json:"function"`
	Data        interface{} `json:"data"`
	LogType     string      `json:"logType"`
	Timestamp   string      `json:"timestamp"`
}

func LoggingDb(requestId, requestKey, code, logType, method, function string, data interface{}, logCategory string) {
	entry := LogEntry{
		RequestId:  requestId,
		RequestKey: requestKey,
		Code:       code,
		Type:       logType,
		Method:     method,
		Function:   function,
		Data:       data,
		LogType:    logCategory,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
	}
	
	jsonData, _ := json.Marshal(entry)
	log.Printf("[%s] %s", logCategory, string(jsonData))
}