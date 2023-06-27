package main

import (
	"fmt"

	"github.com/lukechampine/freeze"
)


type MessageQueName struct{
	IncomeRequest string
	IncomeResponse string
	Outcome string
    OutcomeRequest string
    OutcomeResponse string
}

type MessageQueueRouteKey struct{
	IncomeRequest string
	IncomeResponse string
	OutcomeRequest string
	OutcomeResponse string
}

type MessageQueueExchange struct{
    Incomes string
	Outcomes string
}

func insertOcppHistories(isRemote bool, code string, message [] interface {})

func main() {

	messageQueue := &MessageQueName{"income-req","income-res","outcome","outcome-req","outcome-res"}
	messageQueue = freeze.Object(messageQueue).(*MessageQueName)

	routeKey := &MessageQueueRouteKey{"imcome-req","imcome-res","outcome-req","outcome-res"}
	routeKey = freeze.Object(routeKey).(*MessageQueueRouteKey)

	messageExchange := &MessageQueueExchange{"imcomes","outcomes"}
	messageExchange = freeze.Object(messageExchange).(*MessageQueueExchange)
	
	fmt.Println(messageQueue.Outcome)
	fmt.Println("ROUTE", routeKey.IncomeRequest)
	fmt.Println("EXCHANGE", messageExchange.Incomes)
	

}
