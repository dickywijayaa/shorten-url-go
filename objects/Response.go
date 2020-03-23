package objects

type Response struct {
	Code		int
	Data		interface{}
	Message		string
	Status		string
}