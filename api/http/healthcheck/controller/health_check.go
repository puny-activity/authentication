package controller

import "net/http"

func (c *Controller) HealthCheck(w http.ResponseWriter, r *http.Request) error {
	return c.responseWriter.Write(w, http.StatusOK, nil)
}
