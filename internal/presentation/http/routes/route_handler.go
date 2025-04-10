package routes

import (
	"decard/internal/domain"
	"decard/internal/domain/valueobject"
	"decard/internal/presentation/http/common"
	"decard/internal/presentation/http/middleware"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"reflect"
)

// Изменить так, чтобы можно было прокидывать структуру кастомного реквеста, он сам
// валидируется и прокидывается в функцию + настройка логгера тоже на этом уровне
type publicAPIFunc func(w http.ResponseWriter, r any) error
type protectedAPIFunc func(w http.ResponseWriter, r any, profile valueobject.UUID) error

func HandlePublic(l *zerolog.Logger, h publicAPIFunc, req common.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		*l = l.With().Str("method", r.Method).Str("url", r.URL.String()).Logger()
		targetRequest, err := decodeRequest(r, req, l)

		if err != nil {
			http.Error(w, "invalid request body", http.StatusUnprocessableEntity)

			return
		}

		if err := h(w, targetRequest); err != nil {
			e := common.JSONErrorResponse(w, err)

			if e != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}

			l.Error().Err(err).Msg("http error")
		}
	}
}

func HandleProtected(l *zerolog.Logger, h protectedAPIFunc, req common.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		*l = l.With().Str("method", r.Method).Str("url", r.URL.String()).Logger()
		profileUUID, ok := r.Context().Value(middleware.ProfileUUIDKey).(valueobject.UUID)

		if !ok {
			l.Error().Err(domain.ErrInvalidUser).Msg("failed to assert user UUID")

			e := common.JSONErrorResponse(w, domain.ErrUnauthorized)
			if e != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}

			return
		}

		targetRequest, err := decodeRequest(r, req, l)

		if err != nil {
			http.Error(w, "invalid request body", http.StatusUnprocessableEntity)

			return
		}

		if err = h(w, targetRequest, profileUUID); err != nil {
			e := common.JSONErrorResponse(w, err)

			if e != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}

			l.Error().Err(err).Msg("http error")
		}
	}
}

func decodeRequest(r *http.Request, req any, l *zerolog.Logger) (any, error) {
	if req == nil {
		return nil, nil
	}

	var targetRequest any = nil
	var err error = nil

	targetRequest = reflect.New(reflect.TypeOf(req)).Interface()

	if r.ContentLength > 0 {
		if err = json.NewDecoder(r.Body).Decode(&targetRequest); err != nil {
			l.Error().Err(err).Msg("failed to parse request body")

			return nil, err
		}
	}

	if vars := mux.Vars(r); len(vars) > 0 {
		val := reflect.ValueOf(targetRequest).Elem()
		typ := val.Type()

		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			structField := typ.Field(i)

			if uriTag := structField.Tag.Get("uri"); uriTag != "" {
				if paramValue, exists := vars[uriTag]; exists && field.CanSet() {
					switch field.Kind() {
					case reflect.String:
						field.SetString(paramValue)
					default:
						return nil, fmt.Errorf("invalid uri tag type: %s", uriTag)
					}
				}
			}
		}
	}

	return targetRequest, nil
}
