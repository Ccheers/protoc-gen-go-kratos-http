package audit

import (
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestExtractFromRequest1(t *testing.T) {
	type args struct {
		r     *http.Request
		audit *Audit
	}
	tests := []struct {
		name    string
		args    args
		want    *Audit
		wantErr bool
	}{
		{
			name: "TestExtractFromRequest1",
			args: args{
				r: &http.Request{
					Method: http.MethodGet,
					URL:    &url.URL{RawQuery: "client_ip=127.0.0.1&cluster_name=cluster1&cluster_id=123456"},
				},
				audit: &Audit{
					Module: "module",
					Action: "action",
					Metas: []Meta{
						{
							Key: "client_ip",
							Value: MetaValue{
								Extract: "client_ip",
							},
						},
						{
							Key: "cluster_name",
							Value: MetaValue{
								Extract: "cluster_name",
							},
						},
						{
							Key: "cluster_id2",
							Value: MetaValue{
								Const: "2222",
							},
						},
					},
				},
			},
			want: &Audit{
				Module: "module",
				Action: "action",
				Metas: []Meta{
					{
						Key: "client_ip",
						Value: MetaValue{
							Extract: "client_ip",
						},
					},
					{
						Key: "cluster_name",
						Value: MetaValue{
							Extract: "cluster_name",
						},
					},
					{
						Key: "cluster_id2",
						Value: MetaValue{
							Const: "2222",
						},
					},
				},
				ExtractedMetas: []ExtractedMeta{
					{
						Key:   "client_ip",
						Value: "127.0.0.1",
					},
					{
						Key:   "cluster_name",
						Value: "cluster1",
					},
					{
						Key:   "cluster_id2",
						Value: "2222",
					},
				},
			},
		},
		{
			name: "TestExtractFromRequest2",
			args: args{
				r: &http.Request{
					Method: http.MethodPost,
					Body:   nil,
				},
				audit: &Audit{
					Module: "module",
					Action: "action",
					Metas: []Meta{
						{
							Key: "client_ip",
							Value: MetaValue{
								Extract: "client_ip",
							},
						},
					},
				},
			},
			want: &Audit{
				Module: "module",
				Action: "action",
				Metas: []Meta{
					{
						Key: "client_ip",
						Value: MetaValue{
							Extract: "client_ip",
						},
					},
				},
			},
			// wantErr: true,
		},
		{
			name: "TestExtractFromRequest_WithBody",
			args: args{
				r: &http.Request{
					Method: http.MethodPost,
					Body:   io.NopCloser(strings.NewReader(`{"client_ip":"127.0.0.1"}`)),
				},
				audit: &Audit{
					Module: "module",
					Action: "action",
					Metas: []Meta{
						{
							Key: "client_ip",
							Value: MetaValue{
								Extract: "client_ip",
							},
						},
					},
				},
			},
			want: &Audit{
				Module: "module",
				Action: "action",
				Metas: []Meta{
					{
						Key: "client_ip",
						Value: MetaValue{
							Extract: "client_ip",
						},
					},
				},
				ExtractedMetas: []ExtractedMeta{
					{
						Key:   "client_ip",
						Value: "127.0.0.1",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractFromRequest(tt.args.r, tt.args.audit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractFromRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractFromRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAudit(t *testing.T) {
	type args struct {
		module string
		action string
		metas  []Meta
	}
	tests := []struct {
		name string
		args args
		want *Audit
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAudit(tt.want.Module, tt.want.Action, tt.want.Metas); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAudit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractFromBody(t *testing.T) {
	type args struct {
		r     *http.Request
		metas []Meta
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractFromBody(tt.args.r, tt.args.metas)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractFromBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractFromBody() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractFromQuery(t *testing.T) {
	type args struct {
		r     *http.Request
		metas []Meta
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractFromQuery(tt.args.r, tt.args.metas); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractFromQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
