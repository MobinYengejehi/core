package main

import (
	"github.com/MobinYengejehi/core/base/errors"
	"github.com/MobinYengejehi/core/core"
	"github.com/MobinYengejehi/core/events"
	"github.com/MobinYengejehi/core/styles"
	"github.com/MobinYengejehi/core/video"
)

func main() {
	b := core.NewBody("Video Example")
	v := video.NewVideo(b)
	// v.Rotation = -90
	v.Styler(func(s *styles.Style) {
		s.Grow.Set(1, 1)
	})
	errors.Log(v.Open("deer.mp4"))
	v.OnShow(func(e events.Event) {
		v.Play(0, 0)
	})
	b.RunMainWindow()
}
