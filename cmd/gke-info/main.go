package main

import (
	"net/http"
	"os"

	"context"

	"github.com/go-kit/kit/log"
)

var (
	version = os.Getenv("VERSION")
)

func main() {
	// Create Local logger
	localLogger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	projectID := "vic-goog"
	serviceName := "gke-info"
	serviceComponent := os.Getenv("COMPONENT")
	sdc, err := NewStackDriverClient(ctx, projectID, serviceName+"-"+serviceComponent, version)
	if err != nil {
		panic("Unable to create stackdriver clients: " + err.Error())
	}

	var common CommonService
	common = commonService{backendURL: "http://info-backend:8080/metadata", sdc: sdc}
	common = stackDriverMiddleware{ctx, sdc, localLogger, common.(commonService)}

	createCommonEndpoints(common, sdc)
	if serviceComponent == "frontend" {
		createFrontendEndpoints(common, sdc)
	} else if serviceComponent == "backend" {
		createBackendEndpoints(common, sdc)
	} else {
		panic("Unknown component: " + serviceComponent)
	}

	localLogger.Log("msg", "HTTP", "addr", ":8080")
	localLogger.Log("err", http.ListenAndServe(":8080", nil))
}
