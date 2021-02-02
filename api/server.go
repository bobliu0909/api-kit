package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rl5c/api-server/conf"
	"github.com/rl5c/api-server/pkg/controllers"
	"github.com/rl5c/api-server/pkg/logger"
)

type APIServer struct {
	ctx context.Context
	svc *http.Server
	tlsConfig *conf.TLSConfig
	shutdown bool
}

func NewServer(ctx context.Context, cluster string, controller controllers.BaseController, config *conf.APIConfig) (*APIServer, error) {
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
	router := NewRouter(cluster, controller, config)
	server := &APIServer{
		ctx: ctx,
		svc: &http.Server{
			Addr: config.Bind,
			Handler: router,
			WriteTimeout: 30 * time.Second,
			ReadTimeout: 30 * time.Second,
			IdleTimeout: 60 * time.Second,
			TLSConfig: tlsConfig,
		},
		tlsConfig: config.TLSConfig,
		shutdown: false,
	}
	return server, nil
}

func (server *APIServer) Startup() error {
	server.shutdown = false
	startCtx, cancel := context.WithTimeout(server.ctx, time.Second * 3)
	defer cancel()
	var	err error
	errCh := make(chan error)
	go func(errCh chan<- error) {
		var e error
		if server.tlsConfig != nil {
			logger.INFO("[#api#] server https TLS enabled.", server.svc.Addr)
			e = server.svc.ListenAndServeTLS(server.tlsConfig.ServerCert, server.tlsConfig.ServerKey)
		} else {
			e = server.svc.ListenAndServe()
		}
		if !server.shutdown {
			errCh <- e
		}
	}(errCh)
	select {
		case <-startCtx.Done():
			logger.INFO("[#api#] server listening on [%s]...", server.svc.Addr)
		case err = <-errCh:
			if err != nil && !server.shutdown {
				logger.ERROR("[#api#] server listening failed.")
			}
	}
	close(errCh)
	return err
}

func (server *APIServer) Stop() error {
	logger.INFO("[#api#] server stopping...")
	shutdownCtx, cancel := context.WithTimeout(server.ctx, time.Second * 15)
	defer cancel()
	server.shutdown = true
	return server.svc.Shutdown(shutdownCtx)
}
