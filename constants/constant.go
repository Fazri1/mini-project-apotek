package constants

import "os"

const JWT_SECRET_KEY = "miniProject"
// raja ongkir API
var RO_API_KEY = os.Getenv("RO_API_KEY")

// midtrans API
var MT_MERCHANT_ID = os.Getenv("MT_MERCHANT_ID")
var MT_CLIENT_KEY = os.Getenv("MT_CLIENT_KEY")
var MT_SERVER_KEY = os.Getenv("MT_SERVER_KEY")

// AWS
var AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
var AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
var AWS_REGION = os.Getenv("AWS_REGION")
