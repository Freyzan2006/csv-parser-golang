package config

type Config struct {
	FilePath     string
	Required     []string
	ValidateType map[string]string
	Range        map[string][2]float64
	Verbose 	 bool
	Filter		 string
	Sort		 string
	Header 		 bool
	Export       string
}