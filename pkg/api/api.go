// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/ethersphere/bee/pkg/logging"
	"github.com/ethersphere/bee/pkg/pingpong"
	"github.com/prometheus/client_golang/prometheus"
)

type Service interface {
	http.Handler
	Metrics() (cs []prometheus.Collector)
}

type server struct {
	Options
	http.Handler
	metrics metrics
}

type Options struct {
	Pingpong pingpong.Interface
	Logger   logging.Logger
}

func New(o Options) Service {
	s := &server{
		Options: o,
		metrics: newMetrics(),
	}

	s.setupRouting()

	return s
}
