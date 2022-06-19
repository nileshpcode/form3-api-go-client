package main

import (
	uuid "github.com/satori/go.uuid"
	"reflect"
	"testing"
)

func Test_accountService_Create(t *testing.T) {
	client := NewClient("localhost:8080")

	type args struct {
		data AccountRequestDTO
	}
	tests := []struct {
		name    string
		args    args
		want    *AccountResponseDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.AccountInterface.Create(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountService_Delete(t *testing.T) {
	client := NewClient("localhost:8080")

	type args struct {
		id      uuid.UUID
		version int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := client.AccountInterface.Delete(tt.args.id, tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_accountService_Fetch(t *testing.T) {
	client := NewClient("http://127.0.0.1:8080")

	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    *AccountResponseDTO
		wantErr bool
	}{
		{
			name: "ShouldGetAccount",
			args: args{
				id: uuid.NewV4(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.AccountInterface.Fetch(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fetch() got = %v, want %v", got, tt.want)
			}
		})
	}
}
