package main

import (
	"errors"
	"fmt"
	"github.com/eaglesakura/cli/commons/shell"
	"github.com/eaglesakura/cli/commons/utils"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var dpiTables = []*DotPerInch{
	{Name: "ldpi"},
	{Name: "mdpi"},
	{Name: "hdpi"},
	{Name: "xhdpi"},
	{Name: "xxhdpi"},
	{Name: "xxxhdpi"},
}

type Request struct {
	Path        string   `yaml:"path"`         // ファイル一覧へのパス
	Platform    string   `yaml:"platform"`     // 対象プラットフォーム
	Type        string   `yaml:"type"`         // 出力タイプ
	Format      string   `yaml:"format"`       // 出力ファイルフォーマット
	ConvertArgs []string `yaml:"convert_args"` // convertコマンドへ渡される引数
}

// mipmap出力設定ファイル
type Configure struct {
	Requests []Request
}

func parseConfigure(inputDirectory string) (*Configure, error) {
	configFile := filepath.Join(inputDirectory, "config.yaml")
	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed config file open, %v", configFile)
	}

	var result Configure
	if err = yaml.Unmarshal(buf, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
 処理対象のdpi一覧を取得する
*/
func (it *Request) getDpiList(path string) []*DotPerInch {
	var result []*DotPerInch
	for _, dir := range utils.ListDirectories(path) {
		result = append(result, &DotPerInch{
			Name: dir.Name(),
		})
	}
	return result
}

/*
 1ファイル単位でmipmapを生成する
*/
func (it *Request) generateAndroidMipmap(inputDirectory string, srcDpi *DotPerInch, src os.FileInfo, outputPath string) error {

	// 出力ファイル名を決定する
	dstFileName := src.Name()
	if len(it.Format) > 0 {
		// フォーマット変換が必要
		dstFileName = src.Name()[0:strings.LastIndex(src.Name(), ".")] + "." + strings.ToLower(it.Format)
	}

	srcFilePath := filepath.Join(inputDirectory, it.Path, srcDpi.Name, src.Name())

	// 画像情報
	info, err := LoadImageInfo(srcFilePath)
	if err != nil {
		return err
	}

	fmt.Printf("convert %v[%vx%v:%v] %v\n", srcDpi.Name, info.Width, info.Height, info.Format, srcFilePath)
	for _, dstDpi := range dpiTables {

		dstWidth := srcDpi.GetResizePixels(info.Width, dstDpi)
		dstHeight := srcDpi.GetResizePixels(info.Height, dstDpi)

		if dstWidth <= 0 || dstHeight <= 0 {
			continue
		}

		// convert経由で出力する
		dstFileDir := filepath.Join(outputPath, fmt.Sprintf("%v-%v", it.Type, dstDpi.Name))
		dstFilePath := filepath.Join(dstFileDir, dstFileName)
		_ = os.MkdirAll(dstFileDir, os.ModePerm)

		// 出力ファイルが存在したらskip
		if _, err := os.Stat(dstFilePath); err == nil {
			// ファイルが存在するので、出力しない
			fmt.Printf("  - %v[%vx%v] Skip \n", dstDpi.Name, dstWidth, dstHeight)
			continue
		}

		cmd := &shell.Shell{
			Commands: []string{
				"convert", srcFilePath,
			},
		}
		// リサイズの必要があるなら設定
		if dstWidth != info.Width || dstHeight != info.Height {
			cmd.Commands = append(cmd.Commands, "-resize", fmt.Sprintf("%vx%v", dstWidth, dstHeight))
		}

		// 引数を追加する
		for _, arg := range it.ConvertArgs {
			cmd.Commands = append(cmd.Commands, arg)
		}
		// 出力ファイルパス
		cmd.Commands = append(cmd.Commands, dstFilePath)

		fmt.Printf("  * %v[%vx%v] %v -> %v\n", dstDpi.Name, dstWidth, dstHeight, dstFilePath, cmd.Commands)

		// execute!
		_, stdErr, err := cmd.RunStdout()
		if err != nil {
			return errors.New(fmt.Sprintf("%v %v", err, stdErr))
		}
	}

	return nil
}

/*
 1リクエストを処理する
*/
func (it *Request) generateAndroid(inputPath string, outputPath string) error {
	// 出力先ディレクトリを作成する
	if err := os.MkdirAll(outputPath, os.ModePerm); err != nil {
		return err
	}

	// dpi一覧を取得する
	for _, dpi := range it.getDpiList(filepath.Join(inputPath, it.Path)) {
		// dpi内部のファイルを列挙する
		path := filepath.Join(inputPath, it.Path, dpi.Name)
		for _, srcFile := range utils.ListFiles(path) {
			// 1ファイルの生成を行う
			if err := it.generateAndroidMipmap(inputPath, dpi, srcFile, outputPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func (it *Request) Generate(inputPath string, outputPath string) error {
	switch it.Platform {
	case "android":
		return it.generateAndroid(inputPath, outputPath)
	default:
		return fmt.Errorf("invalid platform(%v)", it.Platform)
	}
}
