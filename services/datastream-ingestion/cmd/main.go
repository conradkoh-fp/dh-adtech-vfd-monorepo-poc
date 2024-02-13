package main

import (
	deal_entities "common/domain/entities/deal"
	"fmt"
)

func main() {
	deal := deal_entities.Deal{
		DealID: "deal-id",
	}
	fmt.Println(deal.Format()) //POC: consume the entities from the domain deal entities package
}
