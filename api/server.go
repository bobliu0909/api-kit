package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rl5c/api-gin/conf"
	"github.com/rl5c/api-gin/logger"
	"github.com/rl5c/api-gin/pkg/controllers"
)

type APIServer struct {
	ctx context.Context
	svc *http.Server
	tlsConfig *conf.TLSConfig
}

func NewServer(ctx context.Context, controller controllers.IController, config *conf.API) (*APIServer, error) {
	var tlsConfig *tls.Config
	if config.TLSConfig != nil {
		certPool := x509.NewCertPool()
		caCert, err := ioutil.ReadFile(config.TLSConfig.CaCert)
		if err != nil {
			return nil, err
		}
		certPool.AppendCertsFromPEM(caCert)
		tlsConfig = &tls.Config{
			ClientCAs: certPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
			NextProtos: []string{"http/1.1"},
		}
	}
	router := NewRouter(controller, config)
	server := &APIServer{
		ctx: ctx,
		svc: &http.Server{
			Addr: config.Bind,
			Handler: router,
			WriteTimeout: 15 * time.Second,
			ReadTimeout: 15 * time.Second,
			IdleTimeout: 60 * time.Second,
			TLSConfig: tlsConfig,
		},
		tlsConfig: config.TLSConfig,
	}
	return server, nil
}

func (server *APIServer) Startup() error {
	var err error
	if server.tlsConfig != nil {
		logger.INFO("[#api#] server TLS enabled and listening [%s]...", server.svc.Addr)
		err = server.svc.ListenAndServeTLS(server.tlsConfig.ServerCert, server.tlsConfig.ServerKey)
	} else {
		logger.INFO("[#api#] server listening [%s]...", server.svc.Addr)
		err = server.svc.ListenAndServe()
	}
	if err != nil {
		if err != http.ErrServerClosed {
			return err
		}
	}
	return nil
}

func (server *APIServer) Stop() error {
	logger.INFO("[#api#] server stopping...")
	shutdownCtx, cancel := context.WithTimeout(server.ctx, time.Second * 15)
	defer cancel()
	return server.svc.Shutdown(shutdownCtx)
}
