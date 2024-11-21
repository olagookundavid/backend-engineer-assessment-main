package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/masena-dev/bookstore-api/internal/jsonlog"
)

type envelope map[string]any

func WriteJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// Encode the data to JSON, returning the error if there was one.
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "Helper/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}
func ReadJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	//disallow unknown fields, doesn't ignore fields that are not meant to be in json request body
	dec.DisallowUnknownFields()
	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError *http.MaxBytesError
		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}
	return nil
}

func ReadIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}
func ReadIntParam(r *http.Request, intName string) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName(intName), 10, 64)
	if err != nil || id < 1 {
		return 0, fmt.Errorf("invalid %s parameter", intName)
	}
	return id, nil
}

func ReadStringParam(r *http.Request, paramName string) (string, error) {
	params := httprouter.ParamsFromContext(r.Context())
	value := params.ByName(paramName)
	if value == "" {
		return "", fmt.Errorf("missing %s parameter", paramName)
	}
	return value, nil
}

func GetDate(r *http.Request) (time.Time, error) {
	dateString, err := ReadStringParam(r, "date")
	if err != nil {
		return time.Now(), err
	}
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return time.Now(), errors.New("invalid date format")
	}
	return date, nil
}

func ReadBoolParam(r *http.Request, paramName string) (bool, error) {
	params := httprouter.ParamsFromContext(r.Context())
	value := params.ByName(paramName)
	if value == "" {
		return false, fmt.Errorf("missing %s parameter", paramName)
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("invalid boolean value for %s parameter", paramName)
	}
	return boolValue, nil
}

// Errors helpers
func logError(r *http.Request, err error) {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	logger.PrintError(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String()})
}

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}
	err := WriteJSON(w, status, env, nil)
	if err != nil {
		logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, err)
	message := "the server encountered a problem and could not process your request"
	env := envelope{
		"error":    message,
		"devError": err.Error(),
	}
	err = WriteJSON(w, http.StatusInternalServerError, env, nil)
	if err != nil {
		logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//////

//For flexible update querys

// func  readString(qs url.Values, key string, defaultValue string) string {

// 	s := qs.Get(key)
// 	if s == "" {
// 		return defaultValue
// 	}
// 	// Otherwise return the string.
// 	return s
// }

// func  readCSV(qs url.Values, key string, defaultValue []string) []string {
// 	csv := qs.Get(key)
// 	if csv == "" {
// 		return defaultValue
// 	}
// 	return strings.Split(csv, ",")
// }