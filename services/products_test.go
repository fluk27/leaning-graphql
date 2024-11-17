package services

import (
	"leaning-graphql/models"
	"leaning-graphql/repository"
	"reflect"
	"testing"
)

func Test_productsServices_GetProductsService(t *testing.T) {
	type fields struct {
		productRepo repository.ProductRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := productsServices{
				productRepo: tt.fields.productRepo,
			}
			got, err := p.GetProductsService()
			if (err != nil) != tt.wantErr {
				t.Errorf("productsServices.GetProductsService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productsServices.GetProductsService() = %v, want %v", got, tt.want)
			}
		})
	}
}
