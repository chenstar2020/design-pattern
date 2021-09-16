package bridge

import "testing"

func TestNewBridge(t *testing.T) {
	sms := ViaSMS()
	commonSms := NewCommonMessage(sms)
	commonSms.SendMessage("下午一起吃饭", "tom")
	email := ViaEmail()
	commonEmail := NewCommonMessage(email)
	commonEmail.SendMessage("下午一起打球", "jack")

	urgencySms := NewUrgencyMessage(sms)
	urgencySms.SendMessage("发生地震了，快跑", "chen")
	urgencyEmail := NewUrgencyMessage(email)
	urgencyEmail.SendMessage("发生火灾了，快跑", "star")
}
