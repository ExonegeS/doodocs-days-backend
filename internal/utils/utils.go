package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/exoneges/doodocs-days-backend/internal/config"
)

func LogRequest(r *http.Request, addInfo string) {
	if config.PROTECTED {
		config.LOGGER.Info("Request",
			"From", "x.x.x.x",
			"Method", r.Method,
			"URL", r.URL,
			"Additional", addInfo)
	} else {
		config.LOGGER.Info("Request",
			"From", r.RemoteAddr,
			"Method", r.Method,
			"URL", r.URL,
			"Additional", addInfo)
	}
}

func SendJSONError(w http.ResponseWriter, code int, err error, message string) {
	config.LOGGER.Error(err.Error())
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func ReadJSONFile(filepath string, v interface{}) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	return json.Unmarshal(byteValue, v)
}

func ConvertToBinary(item interface{}) ([]byte, error) {
	return json.Marshal(item)
}

func Itoa(num int) string {
	return strconv.Itoa(num)
}

func Atoi(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1, err
	}
	return i, nil
}
