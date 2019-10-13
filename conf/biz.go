package conf

var condData = map[string]interface{}{
	"EsHost": "127.0.0.1:9200",
}


func GetString(key string) string {
	v, ok := condData[key]
	if !ok {
		return ""
	}
	return v.(string)
}
