package main

import (
	"blog_kitex/bapi"
	api "blog_kitex/kitex_gen/api"
	"context"
	"encoding/json"
)

// TagServiceImpl implements the last service interface defined in the IDL.
type TagServiceImpl struct{}

// GetTagList implements the TagServiceImpl interface.
func (s *TagServiceImpl) GetTagList(ctx context.Context, req *api.GetTagListRequest) (resp *api.TagListResponse, err error) {
	// TODO: Your code here...
	sdk := bapi.NewAPI("http://127.0.0.1:8888")
	body, err := sdk.GetTagList(ctx, req.Name, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	tagList := api.TagListResponse{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, err
	}
	return &tagList, nil
}
