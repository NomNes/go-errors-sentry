package errors

import (
	"net/http"

	"github.com/getsentry/sentry-go"
)

func SentryBeforeSend(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
	if ex, ok := hint.OriginalException.(*Error); ok {
		for key, val := range ex.Extra() {
			event.Extra[key] = val
		}
	}
	return event
}

func SentryHttpCaptureException(exception error, r *http.Request) *sentry.EventID {
	if r != nil {
		hub := sentry.CurrentHub().Clone()
		hub.Scope().SetRequest(r)
		return hub.CaptureException(exception)
	}
	return sentry.CaptureException(exception)
}
