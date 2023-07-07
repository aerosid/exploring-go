package city

import (
	"testing"
)

func TestInsert(t *testing.T) {
	jaipur := &City{
		Name:  "Jaipur",
		State: "Rajasthan"}
	jaipur.Save()
	id := jaipur.CityId
	jaipur.Delete()
	t.Log("Jaipur", id)
}

func TestUpdate(t *testing.T) {
	chandigarh := &City{
		Name:  "Chandigarh",
		State: "Punjab"}
	chandigarh.Save()
	id := chandigarh.CityId
	chandigarh.State = "Haryana"
	chandigarh.Save()
	t.Log("Chandigarh", id)
	chandigarh.Delete()
}

func TestFindByState(t *testing.T) {
	list := FindByState("Karnataka")
	chd := FindById(7)
	chd.Delete()
	t.Log("FindByState:", list)
}
