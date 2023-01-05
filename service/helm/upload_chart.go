package service

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/20gu00/aBais/common/config"

	"go.uber.org/zap"
)

// chart文件上传
// mime多用途网络传输协议 邮件协议

// http请求四种常见的POST提交数据方式 服务端通常是根据请求头（headers）中的 Content-Type 字段来获知请
//求中的消息主体是用何种方式编码，再对主体进行解析 也就是说， Content-Type 指定了消息主体中的编码方式 。因
//此，POST 提交数据方案，直接跟 Content-Type 和消息主体两部分有关。

//http请求常见的content-type分为4种：application/json、x-www-form-urlencoded、multipart/form-data、text/plain。
//1.x-www-form-urlencoded  浏览器原生表单 发送前编码所有字符 默认
//2.multipart/form-data 不对文件进行编码,在使用包含文件上传空间的表单时,必须使用这个值 我们使用表单上传文件时，必须将enctype设为multipart/form-data
//上面两种 POST 数据方式，都是浏览器原生支持的，而且现阶段原生 form 表单也只支持这两种方式
//3.application/json 消息主题是json字符串
//4.text/plain 空格转换为"+",但不对特殊字符编码
func (*helmStore) UploadChartFile(file multipart.File, header *multipart.FileHeader) error {
	filename := header.Filename
	t := strings.Split(filename, ".")
	if t[len(t)-1] != "tgz" {
		zap.L().Error("chart文件必须以.tgz结尾")
		return errors.New("chart文件必须以.tgz结尾")
	}
	filePath := config.Config.UploadPath + "/" + filename
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		zap.L().Error("chart文件已存在")
		return errors.New("chart文件已存在")
	}
	out, err := os.Create(filePath)
	if err != nil {
		zap.L().Error("创建chart文件失败", zap.Error(err))
		return errors.New("创建chart文件失败" + err.Error())
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		zap.L().Error("拷贝chart文件失败", zap.Error(err))
		return errors.New("拷贝chart文件失败" + err.Error())
	}
	return nil
}
