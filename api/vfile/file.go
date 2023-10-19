package vfile

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe-starter/api/vcommon"
	"goframe-starter/internal/model/entity"
)

type AddFileReq struct {
	g.Meta `tags:"文件" method:"post" path:"/file" mime:"multipart/form-data" sm:"添加" dc:"添加文件" `
	Drive  string              `json:"drive"      description:"上传驱动：本地，ucloud"`
	File   []*ghttp.UploadFile `json:"file" type:"file" `
}
type AddFileRes struct {
	List []string `json:"list"`
}

type DeleteFileReq struct {
	g.Meta `tags:"文件" method:"delete" path:"/file"  sm:"添加" dc:"删除文件" `
	Id     uint `json:"id" v:"required" description:""`
}
type DeleteFileRes struct{}

type UpdateFileReq struct {
	g.Meta `tags:"文件" method:"put" path:"/file"  sm:"修改" dc:"更新文件" `
	*entity.File
}
type UpdateFileRes struct{}

type ListFileReq struct {
	g.Meta `tags:"文件" method:"get" path:"/file/list"  sm:"列表" dc:"文件列表" `
	*entity.File
	*vcommon.CommonPageReq
}
type ListFileRes struct {
	List []*entity.File `json:"list"`
	*vcommon.CommonPageRes
}

type OneFileReq struct {
	g.Meta `tags:"文件" method:"get" path:"/file"  sm:"单个" dc:"文件" `
	Id     uint `json:"id" v:"required" description:""`
}
type OneFileRes *entity.File

type UploadChunkReq struct {
	g.Meta      `tags:"文件" method:"post" path:"/file/chunk" mime:"multipart/form-data" sm:"添加" dc:"分片上传" `
	Drive       string            `json:"drive"      description:"上传驱动：本地，ucloud"`
	Chunk       *ghttp.UploadFile `json:"chunk" type:"file" `
	ChunkIndex  int               `json:"chunkIndex"`
	TotalChunks int               `json:"totalChunks"`
	Temp        string            `json:"temp"`
	OriginName  string            `json:"originName"`
	Ext         string            `json:"ext"`
	Size        uint64            `json:"size"`
}
type UploadChunkRes struct {
	FileUrl string `json:"fileUrl"`
}
