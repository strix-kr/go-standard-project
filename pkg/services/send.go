package services

import (
    "github.com/strix-kr/go-standard-project/internal/notifications"
    "log"
)

//
func SendSMS() {
    
    sms := notifications.SMS{}
   
    // 내부 서비스를 호출합니다.
    messageId, err := notifications.SendSMS(sms)
    if err != nil {
        log.Fatal(err)
        return
    }
    
    log.Printf("message id : %d", messageId)
}
