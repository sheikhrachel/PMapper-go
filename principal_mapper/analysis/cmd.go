package analysis

type Analysis struct {
	OutputType string `kong:"name='output-type',default='text',enum='text,json',help='The type of output for identified issues.'"`
}
