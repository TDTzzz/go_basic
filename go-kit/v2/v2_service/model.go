package v2_service

type TestRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type TestResponse struct {
	Res int `json:"res"`
}
