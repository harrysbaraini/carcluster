package main

type SpeedData struct {
	Value int `default:0 json:"value, omitempty"`
}

func OnSpeedUpdate(m Message) Message {
	data := SpeedData{}
	err := m.FromResponse(&data)

	if err != nil {
		m.Data = ""
		return m
	}

	return m.ToResponse(data)
}
