package objects

type Response struct {
	Code		int				`json:"code"`
	Data		interface{}		`json:"data"`
	Message		string			`json:"message"`
	Status		string			`json:"status"`
}