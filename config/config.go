package config

var JwtSecret = []byte("supersecretkey") // ใส่ใน ENV จริงๆ
var AllowedOrigins = []string{
	"http://localhost:3000",
	"https://frontend1.com",
	"https://app.example.com",
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