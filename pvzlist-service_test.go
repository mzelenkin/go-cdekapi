package cdekapi

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_GetOffices(t *testing.T) {
	filters := OfficesFilter{}
	filters.AddParam(OfficesFilterCityID, "44")
	c := New(APITestBaseURL).SetAuth("z9GRRu7FxmO53CQ9cFfI6qiy32wpfTkd", "w24JTCv4MnAcuRTx0oHjHLDtyt3I6IBq")
	result, err := c.GetOffices(context.Background(), filters)
	fmt.Printf("ERROR: %v\n", err)
	fmt.Printf("%#v\n", result)
}
