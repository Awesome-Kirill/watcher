package cache

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"io"
	"testing"
	"time"
	mock_cache "watcher/internal/cache/mock"
	"watcher/internal/dto"
	"watcher/internal/sorted"

	"github.com/golang/mock/gomock"
)

func TestMockCache_GetUrl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	aliver := mock_cache.NewMockaliver(ctrl)
	hoster := mock_cache.NewMockhoster(ctrl)

	ctx, cancel := context.WithCancel(context.Background())
	aliver.EXPECT().Alive(ctx, "ya.ru").Return(true, time.Duration(888)).AnyTimes()
	aliver.EXPECT().Alive(ctx, "habr.ru").Return(true, time.Duration(222)).AnyTimes()
	aliver.EXPECT().Alive(ctx, "xxx.ru").Return(true, time.Duration(7878)).AnyTimes()
	aliver.EXPECT().Alive(ctx, "slow.ru").Return(false, time.Duration(-1)).AnyTimes()

	hoster.EXPECT().Hosts(ctx).Return([]string{"ya.ru", "habr.ru", "xxx.ru", "slow.ru"}, nil).AnyTimes()
	var s io.Writer
	l := zerolog.New(s)
	cache := New(new(sorted.Sort), aliver, hoster, &l, 1)

	go func() {
		cache.Watch(ctx)
	}()

	time.Sleep(1 * time.Second)
	cancel()

	wantMax := dto.InfoWithName{
		Name: "xxx.ru",
		Info: dto.Info{
			IsAlive:      true,
			ResponseTime: 7878,
		},
	}

	gotMax := cache.GetMax()
	if gotMax != wantMax {
		t.Errorf("GetMax() = %v, want %v", gotMax, wantMax)
	}

	wantMin := dto.InfoWithName{
		Name: "habr.ru",
		Info: dto.Info{
			IsAlive:      true,
			ResponseTime: 222,
		},
	}
	gotMin := cache.GetMin()

	if gotMin != wantMin {
		t.Errorf("GetMin() = %v, want %v", gotMin, wantMin)
	}

	wantHabr := dto.Info{
		IsAlive:      true,
		ResponseTime: 222,
	}
	gotHabr, _ := cache.GetURL("habr.ru")
	if gotHabr != wantHabr {
		t.Errorf("GetURL() = %v, want %v", gotHabr, wantHabr)
	}

	_, notExistErr := cache.GetURL("aaaaaaaa.ru")

	if !errors.Is(notExistErr, ErrSiteNotFound) {
		t.Errorf("err = %v, want %v", notExistErr, ErrSiteNotFound)
	}
}
