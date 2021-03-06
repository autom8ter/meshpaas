package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/autom8ter/helmgate/gen/gql/go/generated"
	"github.com/autom8ter/helmgate/gen/gql/go/model"
	helmgatepb "github.com/autom8ter/helmgate/gen/grpc/go"
	"github.com/autom8ter/helmgate/internal/helpers"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) InstallApp(ctx context.Context, input model.AppInput) (*model.App, error) {
	resp, err := r.client.InstallApp(ctx, toAppInput(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return gqlApp(resp), nil
}

func (r *mutationResolver) UpdateApp(ctx context.Context, input model.AppInput) (*model.App, error) {
	resp, err := r.client.UpdateApp(ctx, toAppInput(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return gqlApp(resp), nil
}

func (r *mutationResolver) RollbackApp(ctx context.Context, input model.AppRef) (*model.App, error) {
	_, err := r.client.RollbackApp(ctx, toAppRef(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *mutationResolver) UninstallApp(ctx context.Context, input model.AppRef) (*string, error) {
	_, err := r.client.UninstallApp(ctx, toAppRef(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *queryResolver) GetApp(ctx context.Context, input model.AppRef) (*model.App, error) {
	resp, err := r.client.GetApp(ctx, toAppRef(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return gqlApp(resp), nil
}

func (r *queryResolver) GetHistory(ctx context.Context, input model.HistoryFilter) ([]*model.App, error) {
	resp, err := r.client.GetHistory(ctx, &helmgatepb.HistoryFilter{
		Ref:   toAppRef(*input.Ref),
		Limit: uint32(helpers.FromIntPointer(input.Limit)),
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	var apps []*model.App
	for _, a := range resp.Apps {
		apps = append(apps, gqlApp(a))
	}
	return apps, nil
}

func (r *queryResolver) SearchApps(ctx context.Context, input model.AppFilter) ([]*model.App, error) {
	resp, err := r.client.SearchApps(ctx, &helmgatepb.AppFilter{
		Namespace: helpers.FromStringPointer(input.Namespace),
		Selector:  helpers.FromStringPointer(input.Selector),
		Limit:     uint32(helpers.FromIntPointer(input.Limit)),
		Offset:    uint32(helpers.FromIntPointer(input.Offset)),
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	var apps []*model.App
	for _, a := range resp.Apps {
		apps = append(apps, gqlApp(a))
	}
	return apps, nil
}

func (r *queryResolver) SearchCharts(ctx context.Context, input model.ChartFilter) ([]*model.Chart, error) {
	resp, err := r.client.SearchCharts(ctx, &helmgatepb.ChartFilter{
		Term:  input.Term,
		Regex: helpers.FromBoolPointer(input.Regex),
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	var templates []*model.Chart
	for _, t := range resp.GetCharts() {
		templates = append(templates, gqlChart(t))
	}
	return templates, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
