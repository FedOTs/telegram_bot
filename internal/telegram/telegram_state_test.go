package telegram

import (
	"testing"
)

func TestOrderFSM(t *testing.T) {
	telegramFSM := newTelegramFSM()

	phoneCtx := &PhoneContext{
		number: "",
	}

	err := telegramFSM.SendEvent(SayHello, phoneCtx)
	if err != nil {
		t.Errorf("Failed to send SayHello, err: %v", err)
	}

	/*
		err = telegramFSM.SendEvent(GetPhone, phoneCtx)
		if err != nil {
			t.Errorf("Failed to send GetPhone, err: %v", err)
		}

		if telegramFSM.Current != PhoneFail {
			t.Errorf("Expected the FSM to be in the OrderFailed state, actual: %s",
				telegramFSM.Current)
		}

		phoneCtx = &PhoneContext{
			number: "+79999999999",
		}

		err = telegramFSM.SendEvent(GetPhone, phoneCtx)
		if err != nil {
			t.Errorf("Failed to send GetPhone, err: %v", err)
		}
	*/
}
