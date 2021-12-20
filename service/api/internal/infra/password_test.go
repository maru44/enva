package infra

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CheckPassword(t *testing.T) {
	tests := []struct {
		testName  string
		raw       string
		hashed    string
		isWantErr bool
	}{
		{
			"success",
			"MoFXYn3j3oHas3vGZRkZ9Y_vYl1aMzx5atmKHJ851w7lrNGU",
			"$2a$08$cG4HVxmut6wiO2jPf3GGEu2ljfX6GliJpuGkqu0eX.SIEpBw9rwxe",
			false,
		},
		{
			"failed",
			"MoFXYn3j3oHas3vGZRkZ9Y_vYl1aMzx5atmKHJ851w7lrNGu",
			"$2a$08$cG4HVxmut6wiO2jPf3GGEu2ljfX6GliJpuGkqu0eX.SIEpBw9rwxe",
			true,
		},
	}

	pass := Password{}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			err := pass.Check(tt.hashed, tt.raw)

			if !tt.isWantErr {
				require.NoError(t, err)
			} else {
				require.NotNil(t, err)
			}
		})
	}
}
