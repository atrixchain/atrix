// Copyright 2022 Atrix Foundation
// This file is part of the Atrix Network packages.
//
// Atrix is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Atrix packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Atrix packages. If not, see https://github.com/Atrix/Atrix/blob/main/LICENSE

package config

func EnableObservability() error {
	return nil
	// if true {
	// Temporarily disabling this until we can configure out port reuse
	// fast enough or enabling observability through the config.
	// Please see https://github.com/Atrix/Atrix/v9/issues/84
	// return nil
	// }

	// pe, err := prometheus.NewExporter(prometheus.Options{
	// 	Namespace: "Atrixd",
	// })
	// if err != nil {
	// 	return fmt.Errorf("cmd/config: failed to create the OpenCensus Prometheus exporter: %w", err)
	// }

	// views := app.ObservabilityViews()
	// if err := view.Register(views...); err != nil {
	// 	return fmt.Errorf("cmd/config: failed to register OpenCensus views: %w", err)
	// }
	// view.RegisterExporter(pe)

	// mux := http.NewServeMux()
	// mux.Handle("/metrics", pe)

	// // TODO: Derive the Prometheus observability exporter from the Atrix config.
	// addr := ":8877"
	// go func() {
	// 	println("Serving the Prometheus observability exporter at", addr)
	// 	if err := http.ListenAndServe(addr, mux); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// return nil
}
