package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mdouchement/standardfile/internal/database"
	"github.com/mdouchement/standardfile/internal/server/service"
	"github.com/mdouchement/standardfile/internal/sferror"
)

type (
	sub struct {
		db           database.Client
		subscription string
	}
)

type ResponseRecorder struct {
	body       []byte
	statusCode int
	origin     *echo.Response
}

func (w *ResponseRecorder) Header() http.Header {
	return w.origin.Header()
}

func (w *ResponseRecorder) Write(b []byte) (int, error) {
	if w.body == nil {
		w.body = make([]byte, 0)
	}
	w.body = append(w.body, b...)
	return len(b), nil
}

func (w *ResponseRecorder) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (s *sub) DecorateJsonResponse(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := currentUser(c)
		recorder := &ResponseRecorder{
			origin: c.Response(),
		}
		c.SetResponse(echo.NewResponse(recorder, c.Echo()))
		if err := next(c); err != nil {
			return err
		}
		session := currentSession(c)
		c.SetResponse(recorder.origin)
		if !strings.HasPrefix(recorder.Header().Get(echo.HeaderContentType), "application/json") {
			// replay recorded response if content is not json
			return c.Blob(recorder.statusCode, recorder.Header().Get(echo.HeaderContentType), recorder.body)
		}
		var body interface{}
		err := json.Unmarshal(recorder.body, &body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, sferror.New("Could not parse Json body"))
		}
		sub := service.NewSubscriptionService(s.db, s.subscription, service.Params{
			APIVersion: session.APIVersion,
			UserAgent:  session.UserAgent,
			Session:    session,
		})
		return c.JSON(recorder.statusCode, echo.Map{
			"meta": echo.Map{
				"auth": echo.Map{
					"userUuid": user.ID,
					"roles":    sub.UserRoles(user),
				},
			}, "data": body,
		})
	}
}

func (s *sub) Subscriptions(c echo.Context) error {
	return c.HTML(http.StatusInternalServerError, "getaddrinfo ENOTFOUND payments")
}

func (s *sub) Subscription(c echo.Context) error {
	user := currentUser(c)

	var params service.SyncParams
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, sferror.New("Could not get syncing params."))
	}
	params.UserAgent = c.Request().UserAgent()
	params.Session = currentSession(c)

	sub := service.NewSubscriptionService(s.db, s.subscription, params.Params)
	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"user": echo.Map{
			"uuid":  user.ID,
			"email": user.Email,
		},
		"subscription": sub.GetSubscription(user),
	})
}

func (s *sub) Features(c echo.Context) error {
	user := currentUser(c)

	var params service.SyncParams
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, sferror.New("Could not get syncing params."))
	}
	params.UserAgent = c.Request().UserAgent()
	params.Session = currentSession(c)

	sub := service.NewSubscriptionService(s.db, s.subscription, params.Params)

	return c.JSON(http.StatusOK, echo.Map{
		"success":  true,
		"userUuid": user.ID,
		"features": sub.Features(user, s.subscription),
	})
}
