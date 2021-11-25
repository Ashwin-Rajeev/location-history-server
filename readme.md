# Location History Server

toy in-memory location history server.

## Steps to Execute 

#### 1. export envs 
```bash
export HISTORY_SERVER_LISTEN_ADDR=8080
export LOCATION_HISTORY_TTL_SECONDS=50s
```
#### 2. start the application
```bash
go run main.go
```


### Add Location History [POST]

> http://localhost:8080/location/def457/now

Sample request body
```bash
{
	"lat": 19.34,
	"lng": 58.78
}
```

Sample response body
```bash
{
    "status": "ok",
    "result": "success"
}
```
### Get Location History [GET]

> http://localhost:8080/location/def456?max=2

Sample response body
```bash
{
	"order_id": "abc123",
	"history": [
		{"lat": 12.34, "lng": 56.78},
		{"lat": 12.34, "lng": 56.79}
	]
}
```

### Delete Location History [DELETE]

> http://localhost:8080/location/def457

Sample response body
```bash
{
    "status": "ok",
    "result": "success"
}
```
