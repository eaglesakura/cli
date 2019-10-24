package main

import (
	"strings"
)

/*
 Android OSのDPI設定管理
*/
type DotPerInch struct {
	Name string // dpi名称, xxxdpi等
}

/*
 mdpiを1.0とした標準倍率を取得する
*/
func (it *DotPerInch) GetSizeMultiply() float64 {
	switch it.Name {
	case "ldpi":
		return 0.75
	case "mdpi":
		return 1.0
	case "hdpi":
		return 1.5
	}

	// xの数だけ大きくなる
	result := 1.0
	dpiCheck := "xh"
	for strings.Index(it.Name, dpiCheck) >= 0 {
		result += 1.0
		dpiCheck = "x" + dpiCheck
	}
	return result
}

/*
 リサイズ対象のサイズ一覧を取得する
 リサイズの必要が無い場合、0以下を返却する
*/
func (it *DotPerInch) GetResizePixels(srcPixels int, targetDpi *DotPerInch) int {
	selfMult := it.GetSizeMultiply()
	targetMult := targetDpi.GetSizeMultiply()

	if selfMult < targetMult {
		// 自分のほうが倍率が低いため、リサイズの必要がない
		return 0
	}

	// 対象倍率をかけて返却
	return int(float64(srcPixels) * (targetMult / selfMult))
}
