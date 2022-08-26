package callbacks

import (
	"context"
	"reflect"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/anuramat/homework-1/internal/models"
)

func Test_upvoteCallback(t *testing.T) {
	type args struct {
		ctx   context.Context
		query *tgbotapi.CallbackQuery
		data  *models.BotData
	}
	tests := []struct {
		name string
		args args
		want models.ChattableSlice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := upvoteCallback(tt.args.ctx, tt.args.query, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("upvoteCallback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_downvoteCallback(t *testing.T) {
	type args struct {
		ctx   context.Context
		query *tgbotapi.CallbackQuery
		data  *models.BotData
	}
	tests := []struct {
		name string
		args args
		want models.ChattableSlice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := downvoteCallback(tt.args.ctx, tt.args.query, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("downvoteCallback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_editDescriptionCallback(t *testing.T) {
	type args struct {
		in0   context.Context
		query *tgbotapi.CallbackQuery
		data  *models.BotData
	}
	tests := []struct {
		name string
		args args
		want models.ChattableSlice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := editDescriptionCallback(tt.args.in0, tt.args.query, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("editDescriptionCallback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextImageCallback(t *testing.T) {
	type args struct {
		ctx   context.Context
		query *tgbotapi.CallbackQuery
		data  *models.BotData
	}
	tests := []struct {
		name string
		args args
		want models.ChattableSlice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextImageCallback(tt.args.ctx, tt.args.query, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextImageCallback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_previousImageCallback(t *testing.T) {
	type args struct {
		ctx   context.Context
		query *tgbotapi.CallbackQuery
		data  *models.BotData
	}
	tests := []struct {
		name string
		args args
		want models.ChattableSlice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := previousImageCallback(tt.args.ctx, tt.args.query, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("previousImageCallback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deltaIndexImage(t *testing.T) {
	type args struct {
		delta_index int
		ctx         context.Context
		query       *tgbotapi.CallbackQuery
		data        *models.BotData
	}
	tests := []struct {
		name string
		args args
		want models.ChattableSlice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deltaIndexImage(tt.args.delta_index, tt.args.ctx, tt.args.query, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deltaIndexImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteImageCallback(t *testing.T) {
	type args struct {
		ctx   context.Context
		query *tgbotapi.CallbackQuery
		data  *models.BotData
	}
	tests := []struct {
		name string
		args args
		want models.ChattableSlice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteImageCallback(tt.args.ctx, tt.args.query, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteImageCallback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randomImageCallback(t *testing.T) {
	type args struct {
		ctx   context.Context
		query *tgbotapi.CallbackQuery
		data  *models.BotData
	}
	tests := []struct {
		name string
		args args
		want models.ChattableSlice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomImageCallback(tt.args.ctx, tt.args.query, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randomImageCallback() = %v, want %v", got, tt.want)
			}
		})
	}
}
