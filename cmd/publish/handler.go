package main

import (
	"MyDouyin/kitex_gen/publish"
	"context"
)

// PublishSrvImpl implements the last service interface defined in the IDL.
type PublishSrvImpl struct{}

// PublishAction implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}
