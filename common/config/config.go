package config

type envType struct {
	Dev    int
	Test   int
	Online int
}

var EnvironmentType = envType{
	Dev:    1,
	Test:   2,
	Online: 3,
}

var Environment = EnvironmentType.Dev
