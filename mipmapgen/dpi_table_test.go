package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDotPerInch_GetResizePixels(t *testing.T) {
	ldpi := &DotPerInch{
		Name: "ldpi",
	}
	mdpi := &DotPerInch{
		Name: "mdpi",
	}
	hdpi := &DotPerInch{
		Name: "hdpi",
	}
	xhdpi := &DotPerInch{
		Name: "xhdpi",
	}
	xxhdpi := &DotPerInch{
		Name: "xxxhdpi",
	}
	xxxhdpi := &DotPerInch{
		Name: "xxxhdpi",
	}

	// 入力はxxxhdpiとする
	{
		assert.Equal(t, xxxhdpi.GetResizePixels(100, xxxhdpi), 100)
		assert.Equal(t, xxxhdpi.GetResizePixels(100, xxhdpi), 75)
		assert.Equal(t, xxxhdpi.GetResizePixels(100, xhdpi), 50)
		assert.Equal(t, xxxhdpi.GetResizePixels(100, mdpi), 25)
	}
	// 上方変換は不要
	{
		assert.Equal(t, ldpi.GetResizePixels(100, xxxhdpi), 0)
		assert.Equal(t, ldpi.GetResizePixels(100, mdpi), 0)
		assert.Equal(t, ldpi.GetResizePixels(100, hdpi), 0)
		assert.Equal(t, ldpi.GetResizePixels(100, xhdpi), 0)
		assert.Equal(t, ldpi.GetResizePixels(100, xxhdpi), 0)
		assert.Equal(t, ldpi.GetResizePixels(100, xxxhdpi), 0)
	}
}

func TestDotPerInch_GetSizeMultiply(t *testing.T) {
	{
		inch := &DotPerInch{
			Name: "ldpi",
		}
		assert.Equal(t, inch.GetSizeMultiply(), 0.75)
	}
	{
		inch := &DotPerInch{
			Name: "mdpi",
		}
		assert.Equal(t, inch.GetSizeMultiply(), 1.0)
	}
	{
		inch := &DotPerInch{
			Name: "hdpi",
		}
		assert.Equal(t, inch.GetSizeMultiply(), 1.5)
	}
	{
		inch := &DotPerInch{
			Name: "xhdpi",
		}
		assert.Equal(t, inch.GetSizeMultiply(), 2.0)
	}
	{
		inch := &DotPerInch{
			Name: "xxxhdpi",
		}
		assert.Equal(t, inch.GetSizeMultiply(), 3.0)
	}
	{
		inch := &DotPerInch{
			Name: "xxxhdpi",
		}
		assert.Equal(t, inch.GetSizeMultiply(), 4.0)
	}
	{
		inch := &DotPerInch{
			Name: "xxxxhdpi",
		}
		assert.Equal(t, inch.GetSizeMultiply(), 5.0)
	}
}
