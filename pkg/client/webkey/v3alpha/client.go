package v2

import (
	"context"

	"github.com/zitadel/zitadel-go/v3/pkg/client/zitadel"
	webkey "github.com/zitadel/zitadel-go/v3/pkg/client/zitadel/resources/webkey/v3alpha"
)

type Client struct {
	Connection *zitadel.Connection
	webkey.WebkeyServiceClient
}

func NewClient(ctx context.Context, issuer, api string, scopes []string, options ...zitadel.Option) (*Client, error) {
	conn, err := zitadel.NewConnection(ctx, issuer, api, scopes, options...)
	if err != nil {
		return nil, err
	}

	return &Client{
		Connection:        conn,
		WebkeyServiceClient: webkey.NewWebkeyServiceClient(conn.ClientConn),
	}, nil
}
