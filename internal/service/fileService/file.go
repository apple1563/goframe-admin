package fileService

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"goframe-starter/api/vcommon"
	"goframe-starter/api/vfile"
	"goframe-starter/internal/consts"
	"goframe-starter/internal/dao"
	"goframe-starter/internal/model/entity"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var fileCols = dao.File.Columns()

func AddFile(ctx context.Context, req *vfile.AddFileReq) (res *vfile.AddFileRes, err error) {
	if req.File == nil {
		return nil, consts.ErrNoFileInput
	}
	var wg sync.WaitGroup
	errCh := make(chan error, len(req.File))
	res = &vfile.AddFileRes{
		List: make([]string, 0),
	}
	for _, file := range req.File {
		wg.Add(1)
		if file == nil {
			continue // 跳过空文件
		}
		go func(f *ghttp.UploadFile) {
			defer wg.Done()
			var folder = time.Now().Format("2006-01-02")
			var dirPath = g.Cfg().MustGet(ctx, "server.uploadRoot").String() + "/" + folder
			filename, err := f.Save(dirPath, true)
			if err != nil {
				errCh <- err
				return
			}
			var slice = gstr.SplitAndTrim(f.Filename, ".")
			_, err = dao.File.Ctx(ctx).Data(g.Map{
				fileCols.FileUrl:    "upload/" + folder + "/" + filename,
				fileCols.Size:       f.Size,
				fileCols.Drive:      req.Drive,
				fileCols.OriginName: f.Filename,
				fileCols.Ext:        slice[len(slice)-1],
				fileCols.Uid:        ctx.Value("uid"),
			}).Insert()
			if err != nil {
				_ = gfile.Remove(filepath.Join(gfile.MainPkgPath() + dirPath + filename))
				return
			}
			res.List = append(res.List, filename)
		}(file)

	}
	go func() {
		wg.Wait()
		close(errCh) // 关闭通道，表示不会再有错误传递
	}()
	// 在这里可以从 errCh 通道读取错误
	for err := range errCh {
		// 处理错误，可以将错误记录到日志或其他操作
		if err != nil {
			return nil, err
		}
	}
	return
}

func UpdateFile(ctx context.Context, req *vfile.UpdateFileReq) (res *vfile.UpdateFileRes, err error) {
	var data = g.Map{}
	data[fileCols.Remark] = req.Remark
	_, err = dao.File.Ctx(ctx).Where(fileCols.Id, req.Id).Data(data).Update()
	if err != nil {
		return nil, err
	}
	return
}

func DeleteFile(ctx context.Context, req *vfile.DeleteFileReq) (res *vfile.DeleteFileRes, err error) {
	value, err := dao.File.Ctx(ctx).WherePri(req.Id).Fields(fileCols.FileUrl).Value()
	if err != nil {
		return nil, err
	}
	_, err = dao.File.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	_ = gfile.Remove(gfile.MainPkgPath() + "/" + g.Cfg().MustGet(ctx, "server.uploadRoot").String() + "/" + value.String())
	return
}

func ListFile(ctx context.Context, req *vfile.ListFileReq) (res *vfile.ListFileRes, err error) {
	res = &vfile.ListFileRes{
		List:          make([]*entity.File, 0),
		CommonPageRes: &vcommon.CommonPageRes{},
	}
	res.PageIndex = req.PageIndex
	res.PageSize = req.PageSize
	var model = dao.File.Ctx(ctx).OrderDesc(fileCols.UpdatedAt)
	if req.Drive != "" {
		model = model.Where(fileCols.Drive, req.Drive)
	}
	if req.OriginName != "" {
		model = model.WhereLike(fileCols.OriginName, "%"+req.OriginName+"%")
	}
	if req.Ext != "" {
		model = model.WhereLike(fileCols.Ext, "%"+req.Ext+"%")
	}
	err = model.Page(req.PageIndex, req.PageSize).ScanAndCount(&res.List, &res.Total, false)
	if err != nil {
		return nil, err
	}
	return
}

func OneFile(ctx context.Context, req *vfile.OneFileReq) (res *vfile.OneFileRes, err error) {
	err = dao.File.Ctx(ctx).Where(fileCols.Id, req.Id).Scan(&res)
	if err != nil {
		return nil, err
	}
	return
}

func UploadChunk(ctx context.Context, req *vfile.UploadChunkReq) (res *vfile.UploadChunkRes, err error) {
	res = &vfile.UploadChunkRes{}
	if req.Chunk == nil {
		return nil, consts.ErrNoFileInput
	}
	//  保存到临时目录
	var folder = req.Temp
	var uploadPath = g.Cfg().MustGet(ctx, "server.uploadRoot").String() + "/"
	var tempPath = uploadPath + folder
	_, err = req.Chunk.Save(tempPath)
	if err != nil {
		return nil, err
	}
	// 分片上完，合并分片
	if req.ChunkIndex+1 == req.TotalChunks {
		var folder = time.Now().Format("2006-01-02")
		var dirPath = uploadPath + folder
		_ = os.Mkdir(dirPath, 0777)
		var filename = guid.S() + "." + req.OriginName
		var outputFilePath = dirPath + "/" + filename
		err = combineChunks(req.TotalChunks, outputFilePath, tempPath)
		if err != nil {
			return nil, err
		}
		//  插入数据库
		var fileUrl = folder + "/" + filename
		_, err = dao.File.Ctx(ctx).Data(g.Map{
			fileCols.FileUrl:    "upload/" + fileUrl,
			fileCols.Size:       req.Size,
			fileCols.Drive:      req.Drive,
			fileCols.OriginName: req.OriginName,
			fileCols.Ext:        req.Ext,
			fileCols.Uid:        ctx.Value("uid"),
		}).Insert()
		res.FileUrl = fileUrl
		if err != nil {
			_ = gfile.Remove(filepath.Join(gfile.MainPkgPath() + dirPath + filename))
			return nil, err
		}
	}
	return
}
func combineChunks(totalChunks int, outputFilePath string, tempDir string) error {
	// 创建最终文件
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}

	for i := 0; i < totalChunks; i++ {
		chunkFilePath := filepath.Join(gfile.MainPkgPath(), tempDir, gconv.String(i))

		// 打开分片文件
		chunkFile, err := os.Open(chunkFilePath)
		if err != nil {
			return err
		}
		// 从分片文件中读取数据并写入最终文件
		_, err = io.Copy(outputFile, chunkFile)
		if err != nil {
			err2 := chunkFile.Close()
			if err2 != nil {
				return err2
			}
			return err
		}
		err = chunkFile.Close()
		if err != nil {
			return err
		}
	}
	err = outputFile.Close()
	if err != nil {
		return err
	}
	_ = os.RemoveAll(tempDir)

	return nil
}
