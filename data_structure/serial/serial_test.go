package serial

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSerializationAndDeserialization(t *testing.T) {
	tests := []struct {
		name string
		want *KubeHpaDao
	}{
		// TODO: Add test cases.
		{
			name: "test",
			want: &KubeHpaDao{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SerializationAndDeserialization()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SerializationAndDeserialization() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}
