package reamaze

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClient_GetArticles(t *testing.T) {
	type fields struct {
		baseURL    string
		auth       string
		httpClient *http.Client
	}
	type args struct {
		o []ArticlesOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetArticlesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				auth:       tt.fields.auth,
				httpClient: tt.fields.httpClient,
			}
			got, err := c.GetArticles(tt.args.o...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetArticles() = %v, want %v", got, tt.want)
			}
		})
	}
}
