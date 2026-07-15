package metrics

import (
	"github.com/bytedance/gopkg/util/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

const serviceLabel = "nofelet-web"

var (
	requestsTotal   metric.Int64Counter
	responseTotal   metric.Int64Counter
	requestDuration metric.Float64Histogram
	activeRequests  metric.Int64UpDownCounter

	AuthFail            metric.Int64Counter
	AuthSuccess         metric.Int64Counter
	RegisterFail        metric.Int64Counter
	RegisterSuccess     metric.Int64Counter
	UserCreationFail    metric.Int64Counter
	UserCreationSuccess metric.Int64Counter
)

func Init() {
	var err error

	// Создаем Meter (пространство имен для метрик вашего приложения)
	meter := otel.Meter(serviceLabel)

	// Счетчик всех HTTP-запросов
	requestsTotal, err = meter.Int64Counter(
		"http.server.requests.total",
		metric.WithDescription("Total number of HTTP requests."),
		metric.WithUnit("{request}"),
	)
	if err != nil {
		logger.Warn("сбой инициализации метрики requestsTotal", err)
	}

	// Счетчик всех HTTP-ответов
	responseTotal, err = meter.Int64Counter(
		"http.server.responses.total",
		metric.WithDescription("Total number of HTTP responses."),
		metric.WithUnit("{response}"),
	)
	if err != nil {
		logger.Warn("сбой инициализации метрики responseTotal", err)
	}

	// Гистограмма времени выполнения запроса
	requestDuration, err = meter.Float64Histogram(
		"http.server.request.duration",
		metric.WithDescription("Duration of HTTP requests."),
		metric.WithUnit("ms"),
	)
	if err != nil {
		logger.Warn("сбой инициализации метрики requestDuration", err)
	}

	// Активные запросы
	activeRequests, err = meter.Int64UpDownCounter(
		"http.server.requests.active",
		metric.WithDescription("Current number of active HTTP requests."),
		metric.WithUnit("{request}"),
	)
	if err != nil {
		logger.Warn("сбой инициализации метрики activeRequests", err)
	}

	// Счетчик ошибок авторизации
	AuthFail, err = meter.Int64Counter(
		"auth.fail.total",
		metric.WithDescription("Failed auth"),
		metric.WithUnit("{auth}"),
	)
	if err != nil {
		logger.Warn("сбой инициализации метрики authFail", err)
	}

	// Счетчик успешных логинов
	AuthSuccess, err = meter.Int64Counter(
		"auth.success.total",
		metric.WithDescription("Total successful auths"),
		metric.WithUnit("{auth}"),
	)
	if err != nil {
		logger.Warn("сбой инициализации метрики authSuccess", err)
	}

	// Счетчик успешных регистраций
	RegisterSuccess, err = meter.Int64Counter(
		"register.success.total",
		metric.WithDescription("Total successful user registrations"),
		metric.WithUnit("{registration}"),
	)
	if err != nil {
		logger.Warn("сбой инициализации метрики registerSuccess", err)
	}

	// Счетчик ошибок регистраций
	RegisterFail, err = meter.Int64Counter(
		"register.fail.total",
		metric.WithDescription("Total fail user registrations"),
		metric.WithUnit("{registration}"),
	)
	if err != nil {
		logger.Warn("сбой инициализации метрики registerFail", err)
	}

	// Счетчик успешного создания пользователя
	UserCreationSuccess, err = meter.Int64Counter(
		"user.create.success.total",
		metric.WithDescription("Total successful users created"),
		metric.WithUnit("{createdUsers}"),
	)
	if err != nil {
		logger.Warn("сбой инициализации метрики userCreationSuccess", err)
	}

	// Счетчик ошибок создания пользователя
	UserCreationFail, err = meter.Int64Counter(
		"user.create.fail.total",
		metric.WithDescription("Total fail user created"),
		metric.WithUnit("{createdUsers}"),
	)
	if err != nil {
		logger.Warn("сбой инициализации метрики userCreationFail", err)
	}
}
