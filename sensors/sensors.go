package sensors

type Sensors struct{}

type SensorValues struct{
	A int
}

func Start() *Sensors {
	return &Sensors{}
}

func (s *Sensors) Read() (*SensorValues, error) {
	return &SensorValues{}, nil
}
