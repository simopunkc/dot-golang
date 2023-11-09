package service

import (
	"dot-golang/internal/abstraction"
	"dot-golang/internal/domain"
	"dot-golang/internal/stub"
	"errors"
	"testing"
)

func TestGetNews(t *testing.T) {
	type fields struct {
		blogCache abstraction.BlogCache
		blogRepo  abstraction.BlogRepository
		blogEvent abstraction.BlogEvent
		blogUtil  abstraction.BlogUtil
	}
	type args struct {
		page int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "success get list news from database",
			fields: fields{
				blogCache: &BlogCacheMock{
					ExistsFunc: func(key string) bool {
						return false
					},
					GetFunc: func(key string) (string, error) {
						return "", errors.New("cache not found")
					},
					SetFunc: func(key string, value string) error {
						return nil
					},
				},
				blogRepo: &BlogRepositoryMock{
					GetNewsFunc: func(limit int, offset int) ([]domain.News, error) {
						return stub.StubNews, nil
					},
				},
				blogEvent: &BlogEventMock{},
				blogUtil: &BlogUtilMock{
					GetLimitAndOffsetFunc: func(page int) (int, int) {
						return 0, 10
					},
					IntToStringFunc: func(raw int) string {
						return "0"
					},
					ToJsonFunc: func(raw interface{}) []byte {
						return []byte{}
					},
					ByteToStrFunc: func(raw []byte) string {
						return "[]"
					},
				},
			},
			args: args{
				page: 1,
			},
			want: 2,
		},
		{
			name: "success get list news from redis",
			fields: fields{
				blogCache: &BlogCacheMock{
					ExistsFunc: func(key string) bool {
						return true
					},
					GetFunc: func(key string) (string, error) {
						return stub.StubCacheNews, nil
					},
					SetFunc: func(key string, value string) error {
						return nil
					},
				},
				blogRepo:  &BlogRepositoryMock{},
				blogEvent: &BlogEventMock{},
				blogUtil: &BlogUtilMock{
					GetLimitAndOffsetFunc: func(page int) (int, int) {
						return 0, 10
					},
					IntToStringFunc: func(raw int) string {
						return "0"
					},
					ToJsonFunc: func(raw interface{}) []byte {
						return []byte{}
					},
					ByteToStrFunc: func(raw []byte) string {
						return "[]"
					},
					StringToArrayNewsFunc: func(jsonStr string) []domain.News {
						return stub.StubNews
					},
				},
			},
			args: args{
				page: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := NewBlogService(tt.fields.blogCache, tt.fields.blogRepo, tt.fields.blogEvent, tt.fields.blogUtil)
			if resp, _ := ns.GetNews(tt.args.page); len(resp) != tt.want {
				t.Errorf("GetNews() = %v, want %v", len(resp), tt.want)
			}
		})
	}
}
