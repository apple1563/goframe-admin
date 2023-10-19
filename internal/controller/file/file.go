package file

import (
	"context"
	"goframe-starter/api/vfile"
	"goframe-starter/internal/service/fileService"
)

type File struct{}

var Ctrl = new(File)

func (u *File) AddFile(ctx context.Context, req *vfile.AddFileReq) (res *vfile.AddFileRes, err error) {
	return fileService.AddFile(ctx, req)
}

func (u *File) DeleteFile(ctx context.Context, req *vfile.DeleteFileReq) (res *vfile.DeleteFileRes, err error) {
	return fileService.DeleteFile(ctx, req)
}

func (u *File) UpdateFile(ctx context.Context, req *vfile.UpdateFileReq) (res *vfile.UpdateFileRes, err error) {
	return fileService.UpdateFile(ctx, req)
}

func (u *File) ListFile(ctx context.Context, req *vfile.ListFileReq) (res *vfile.ListFileRes, err error) {
	return fileService.ListFile(ctx, req)
}

func (u *File) OneFile(ctx context.Context, req *vfile.OneFileReq) (res *vfile.OneFileRes, err error) {
	return fileService.OneFile(ctx, req)
}

func (u *File) UploadChunk(ctx context.Context, req *vfile.UploadChunkReq) (res *vfile.UploadChunkRes, err error) {
	return fileService.UploadChunk(ctx, req)
}
