package examples

import (
	"context"
	cxsdk "coralogix-management-sdk/go"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestActions(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	t.Log(region)
	assert.Nil(t, err)
	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, apiKey)
	c := cxsdk.NewActionsClient(creator)

	action, e := c.CreateAction(context.Background(), &cxsdk.CreateActionRequest{
		Name:             wrapperspb.String("google search action"),
		Url:              wrapperspb.String("https://www.google.com/search?q={{$p.selected_value}}"),
		IsPrivate:        wrapperspb.Bool(false),
		SourceType:       cxsdk.SourceTypeSourceTypeLog,
		ApplicationNames: []*wrapperspb.StringValue{},
		SubsystemNames:   []*wrapperspb.StringValue{},
	})
	assert.Nil(t, e)

	_, e = c.UpdateAction(context.Background(), &cxsdk.ReplaceActionRequest{
		Action: &cxsdk.Action{
			Id:               action.Action.Id,
			Name:             wrapperspb.String("bing search action"),
			Url:              wrapperspb.String("https://www.bing.com/search?q={{$p.selected_value}}"),
			IsPrivate:        wrapperspb.Bool(false),
			IsHidden:         wrapperspb.Bool(false),
			SourceType:       cxsdk.SourceTypeSourceTypeLog,
			ApplicationNames: []*wrapperspb.StringValue{},
			SubsystemNames:   []*wrapperspb.StringValue{},
		},
	})

	assert.Nil(t, e)

	updated, _ := c.GetAction(context.Background(), &cxsdk.GetActionRequest{
		Id: action.Action.Id,
	})

	assert.Equal(t, updated.Action.Url.Value, "https://www.bing.com/search?q={{$p.selected_value}}")

	c.DeleteAction(context.Background(), &cxsdk.DeleteActionRequest{
		Id: action.Action.Id,
	})
}
