package notifier

import (
	"context"
	"net/url"

	"github.com/signal-ai/alertmanager-to-github/pkg/types"
)

type Notifier interface {
	Notify(context.Context, *types.WebhookPayload, url.Values) error
}
