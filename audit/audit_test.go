package audit

import (
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestExtractFromRequest(t *testing.T) {
	type args struct {
		request *http.Request
		mm      map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "TestExtractFromRequest",
			args: args{
				request: &http.Request{
					Method: http.MethodGet,
					URL:    &url.URL{RawQuery: "client_ip=192.168.1.1&cluster_name=cluster1&cluster_id=123456"},
				},
				mm: map[string]string{
					"client_ip":    "cluster_id",
					"cluster_name": "cluster_name",
				},
			},
			want: map[string]string{
				"client_ip":    "123456",
				"cluster_name": "cluster1",
			},
		},
		{
			name: "TestExtractFromRequest",
			args: args{
				request: &http.Request{
					Method: http.MethodPost,
					URL:    &url.URL{RawQuery: "client_ip=192.168.1.1&cluster_name=cluster1&extra_param=extra_value"},
					Body:   io.NopCloser(strings.NewReader(`{"client_ip":"192.168.1.1","cluster_name":"cluster1","cluster_id":123456}`)),
				},
				mm: map[string]string{
					"client_ip":    "cluster_id",
					"cluster_name": "cluster_name",
				},
			},
			want: map[string]string{
				"client_ip":    "123456",
				"cluster_name": "cluster1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractFromRequest(tt.args.request, tt.args.mm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractFromRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
