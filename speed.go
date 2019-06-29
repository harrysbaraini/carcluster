package main

type SpeedData struct {
	Value int `default:0 json:"value, omitempty"`
}

func OnSpeedUpdate() Message {
	data := SpeedData{53}

	m := Message{"speed:update", "", nil}

	return m.ToResponse(data)
}