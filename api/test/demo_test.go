package test

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func arrayContains(target string, arr []string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}

func TestCopyGoFile(t *testing.T) {
	ginPath := `D:\Develop\buildx\public\gin`
	_ = os.RemoveAll(ginPath)
	_ = os.Mkdir(ginPath, 0777)
	_ = filepath.WalkDir("../../", func(path string, d fs.DirEntry, err error) error {
		newPath := filepath.Join(ginPath, strings.ReplaceAll(path, "..\\", ""))
		if d.IsDir() {
			if d.Name() == ".." {
				return nil
			}
			if arrayContains(d.Name(), []string{"node_modules", "dist", "vendor", "build", ".idea", "demo", "data", "tmp", "logs", ".git", "test"}) {
				return fs.SkipDir
			}
			_ = os.MkdirAll(newPath, os.ModePerm)
		} else {
			if arrayContains(d.Name(), []string{".air.toml", "buildx", "gen_gin_tpl.ini", "test.db"}) {
				return nil
			}
			newPath = fmt.Sprintf("%s.tpl", newPath)
			file, _ := os.ReadFile(path)
			data := strings.ReplaceAll(string(file), "gen_gin_tpl", "[//{ .ProjectName }//]")
			if d.Name() == "variable.go" {
				strings.ReplaceAll(data, "\"知识库\"", "[//{ .ProjectName }//]")
			}
			_ = os.WriteFile(newPath, []byte(data), os.ModePerm)
			fmt.Printf("✅ 创建文件：%s\n", newPath)
		}

		return nil
	})
}

func TestDemo(t *testing.T) {
	filepath.WalkDir(`D:\Develop\MsgBox\src\assets\plugins\MsgBox`, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if ext != ".js" {
			return nil
		}
		file, _ := os.ReadFile(path)
		fmt.Println(filepath.ToSlash(path[3:]))
		fmt.Println("\n```javaScript")
		fmt.Print(string(file))
		fmt.Printf("```\n")
		return nil
	})
}
