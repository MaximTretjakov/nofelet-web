package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

func (c *Controller) GetMetrics(ctx *gin.Context) {
	var rm metricdata.ResourceMetrics

	if err := c.reader.Collect(ctx.Request.Context(), &rm); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to collect metrics")
		return
	}

	ctx.Header("Content-Type", "text/plain; version=0.0.4; charset=utf-8")

	// Парсинг OpenTelemetry:
	for _, sm := range rm.ScopeMetrics {
		for _, m := range sm.Metrics {
			fmt.Fprintf(ctx.Writer, "Метрика: %s\n", m.Name)
			// Логика обработки типов данных...
			switch data := m.Data.(type) {
			case metricdata.Sum[int64]:
				for _, dp := range data.DataPoints {
					attrs := dp.Attributes.Encoded(attribute.DefaultEncoder())
					fmt.Fprintf(ctx.Writer, "  Значение: %d | Теги: %s\n", dp.Value, attrs)
				}
			}
		}
	}
}
