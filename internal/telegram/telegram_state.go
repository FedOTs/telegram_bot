package telegram

import (
	"errors"
	"fmt"
	statemachine "telegram_bot/pkg/statemashine"
)

const (
	Hello     statemachine.StateType = "UserHello"
	Phone     statemachine.StateType = "UserPhone"
	PhoneFail statemachine.StateType = "UserPhoneFailed"
	IdAdress  statemachine.StateType = "UserIdAdress"

	SayHello     statemachine.EventType = "SayHello"
	GetPhone     statemachine.EventType = "GetPhone"
	FailGetPhone statemachine.EventType = "FailGetPhone"
	GetIdAdres   statemachine.EventType = "GetIdAdres"
)

type ContextSendMessage interface {
	Answer() string
}

type HelloContext struct {
	message string
	err     error
}

func (c *HelloContext) Answer() string {
	return `Добрый день. Для внесения показаний нам необходимо знать ваш номер телефона для обратной связи.`
}

type HelloAction struct{}

func (a *HelloAction) Execute(eventCtx statemachine.EventContext) statemachine.EventType {
	fmt.Println("Hello new user")
	return GetPhone
}

type PhoneContext struct {
	number string
	err    error
}

func (c *PhoneContext) Answer() string {
	if c.err != nil {
		return `Вы не ввели данные телефона. \nК сожалению заносить показания можно только после ввода телефонного номера`
	}
	return `Пожалуйста введите № лицевого счета.\nНомер лицевого счета должен состоять только из *ЦИФР*, не более 10 символов.`
}

func (c *PhoneContext) String() string {
	return fmt.Sprintf("Phone context number: %s err: %v", c.number, c.err)
}

type PhoneAction struct{}

func (a *PhoneAction) Execute(eventCtx statemachine.EventContext) statemachine.EventType {
	number := eventCtx.(*PhoneContext)
	if len(number.number) == 0 {
		number.err = errors.New("Не ввели номер")
		return FailGetPhone
	}
	return GetIdAdres
}

type PhoneFailedAction struct{}

func (a *PhoneFailedAction) Execute(eventCtx statemachine.EventContext) statemachine.EventType {
	return statemachine.NoOp
}

type IdAdressAction struct{}

func (a *IdAdressAction) Execute(eventCtx statemachine.EventContext) statemachine.EventType {
	return statemachine.NoOp
}

func newTelegramFSM() *statemachine.StateMachine {
	return &statemachine.StateMachine{
		States: statemachine.States{
			statemachine.Default: statemachine.State{
				Events: statemachine.Events{
					SayHello: Hello,
				},
			},
			Hello: statemachine.State{
				Action: &HelloAction{},
				Events: statemachine.Events{
					GetPhone: Phone,
				},
			},
			Phone: statemachine.State{
				Action: &PhoneAction{},
				Events: statemachine.Events{
					GetIdAdres:   IdAdress,
					FailGetPhone: PhoneFail,
				},
			},
			PhoneFail: statemachine.State{
				Action: &PhoneFailedAction{},
				Events: statemachine.Events{
					GetPhone: Phone,
				},
			},
			IdAdress: statemachine.State{
				Action: &IdAdressAction{},
			},
		},
	}
}
