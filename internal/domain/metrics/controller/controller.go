package controller

import (
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

type Controller struct {
	reader *sdkmetric.ManualReader
}

func New(reader *sdkmetric.ManualReader) *Controller {
	return &Controller{
		reader: reader,
	}
}
