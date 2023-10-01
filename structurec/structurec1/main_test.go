//go:generate mockgen -source ${GOFILE} -destination mocks_test.go -package ${GOPACKAGE}_test

package main

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestPlayer_Timelen(t *testing.T) {
	ctrl:=gomock.NewController(t)
	mock:=NewMockItem(ctrl)
	mock.EXPECT().Getlen().Return(10)
	tests := []struct {
		name string
		Playlist map[string]Item
		want int
	}{
	//	{
	//		name: "sucsess",
	//		Playlist: map[string]Item{
	//			"one": mock,
	//		},
	//		want: 10,
	//	},

		{
			name: "sucsess",
			Playlist: map[string]Item{
				"one": mock,
				"second" : &Song{
					Time: 23,
				},
			},
			want: 33,
		},

	}



	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player:=Player{
				Playlist: tt.Playlist,
			}
			if got := player.Timelen(); got != tt.want {
				t.Errorf("Player.Timelen() = %v, want %v", got, tt.want)
			}
		})
	}
}
