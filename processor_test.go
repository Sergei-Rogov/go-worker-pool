package main

import "testing"

func TestProcess(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{"positive", 5, 10, false},
		{"negative", -3, -3, true},
		{"zero", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Process(tt.input)

			if got.Value != tt.want {
				t.Errorf("got %d, want %d", got.Value, tt.want)
			}

			if (got.Err != nil) != tt.wantErr {
				t.Errorf("got error %v, wantErr %v", got.Err, tt.wantErr)
			}
		})
	}
}
