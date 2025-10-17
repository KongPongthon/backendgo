package config

var JwtSecret = []byte("supersecretkey") // ใส่ใน ENV จริงๆ
var AllowedOrigins = []string{
	"http://localhost:3000",
	"*",
}

var AllowedMethods = []string{
	"GET",
	"POST",
	"PUT",
	"PATCH",
	"DELETE",
	"OPTIONS",
}