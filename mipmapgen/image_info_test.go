package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadImageInfo_png_rgba(t *testing.T) {
	info, e := LoadImageInfo("examples/image/128x128-rgba.png")
	assert.NoError(t, e)
	assert.NotNil(t, info)

	assert.Equal(t, info.Height, 128)
	assert.Equal(t, info.Height, 128)
	assert.Equal(t, info.Channels, "srgba")
	assert.Equal(t, info.Format, "PNG")
}

func TestLoadImageInfo_png_rgb(t *testing.T) {
	info, e := LoadImageInfo("examples/image/128x128-rgb.png")
	assert.NoError(t, e)
	assert.NotNil(t, info)

	assert.Equal(t, info.Height, 128)
	assert.Equal(t, info.Height, 128)
	assert.Equal(t, info.Channels, "srgb")
	assert.Equal(t, info.Format, "PNG")
}

func TestLoadImageInfo_jpeg(t *testing.T) {
	info, e := LoadImageInfo("examples/image/128x128-rgb.jpg")
	assert.NoError(t, e)
	assert.NotNil(t, info)

	assert.Equal(t, info.Height, 128)
	assert.Equal(t, info.Height, 128)
	assert.Equal(t, info.Channels, "srgb")
	assert.Equal(t, info.Format, "JPEG")
}
