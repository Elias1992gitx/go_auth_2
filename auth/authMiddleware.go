package main

import (
        "context"
        "errors"

        "go-micro.dev/v4/metadata"
        "go-micro.dev/v4/server"


        
)

func AuthMiddleware() server.HandlerWrapper {
        return func(h server.HandlerFunc) server.HandlerFunc {
                return func(ctx context.Context, req server.Request, rsp interface{}) error {
                        // Extract the metadata from the context
                        meta, ok := metadata.FromContext(ctx)
                        if !ok {
                                return errors.New("no auth meta-data found in request")
                        }

                        token, ok := meta["Token"]
                        if !ok || token != "valid-token" {
                                return errors.New("invalid or missing auth token")
                        }

                        // Call the next handler
                        return h(ctx, req, rsp)
                }
        }
}