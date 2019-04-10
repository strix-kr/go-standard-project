package notifications

type SMS struct {
    Type        string `json:"type"`
    PhoneNumber string `json:"phone_number"`
    Title       string `json:"title"`
    Content     string `json:"content"`
}

// SMS 를 전송합니다.
func SendSMS(sms SMS) (messageId uint64, err error) {
    
    // do something
    return
}
