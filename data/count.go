package data

type Counter struct {
	Count int `json:"Count"`
}

func GetCounterUser(id string) (Counter, error) {

	return Counter{}, nil
}

func UpdateCounterUser(id string) error {

	return nil
}
