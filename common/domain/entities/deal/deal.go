package deal_entities

import "fmt"

type Deal struct {
	DealID string
}

func (d *Deal) Format() string {
	return fmt.Sprintf("Deal ID: %v", d.DealID)
}
