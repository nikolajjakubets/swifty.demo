package main

func Main(rq *Request) (interface{}, *Response) {
	if rq.Args["act"] == "panic" {
		panic(rq.Args["message"])
	} else {
		print(rq.Args["act"])
	}

	return "done", nil
}
