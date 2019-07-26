package filesystem

import (
	"encoding/csv"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//读取文件，返回一个字符串切片，每行作为切片的一个元素
func ReadFileLines(filePath string) ([]string, error) {
	//并不认可这种提前声明变量的风格，当提前return 时，有些声明的变量根本就没用上
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

//打开CSV文件，返回CSV内容；一个二维切片
func ReadCSV(filePath string) ([][]string, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	return csv.NewReader(f).ReadAll()
}

//向CSV文件追加一行
func AppendCSVRow(filePath string, row []string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	if err = w.Write(row); err != nil {
		return err
	}
	w.Flush()
	return nil
}

//向CSV文件写入多行
func WriteCSVRows(filePath string, rows [][]string, mode int) error {
	f, err := os.OpenFile(filePath, mode, 0644)
	defer f.Close()
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	if err = w.WriteAll(rows); err != nil {
		return err
	}
	w.Flush()
	return nil
}

//向CSV文件覆盖多行
func OverWriteCSVRows(filePath string, rows [][]string) error {
	return WriteCSVRows(filePath, rows, os.O_CREATE|os.O_TRUNC|os.O_WRONLY)
}

//向CSV文件追加多行
func AppendCSVRows(filePath string, rows [][]string) error {
	return WriteCSVRows(filePath, rows, os.O_CREATE|os.O_APPEND|os.O_WRONLY)
}

//合并文件
func MergeFiles(toFile string, fromFiles ...string) error {
	var (
		errs    = errors.New("some errors: ") //由于接收多个文件名，所以定义一个错误变量，用来记录多个错误
		hasErrs = false                       //用来标注是否发生过错误
		//用来处理每个错误，把每个错误都记录到errs; 并标注hasErrs
		handleErr = func(fromFile string, err error) bool {
			if err != nil {
				hasErrs = true
				errs = errors.New(fromFile + err.Error() + "; ")
				return true
			}
			return false
		}
	)

	to, err := os.OpenFile(toFile, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		to.Close()
		return err
	}
	for _, fromFile := range fromFiles {
		from, err := os.OpenFile(fromFile, os.O_RDONLY, 0644)
		if handleErr(fromFile, err) {
			from.Close()
			continue
		}
		_, err = io.CopyBuffer(to, from, make([]byte, 64)) //io.CopyBuffer比io.Copy更安全
		handleErr(fromFile, err)
		from.Close()
	}
	to.Close()
	if hasErrs {
		return errs
	}
	return nil
}
