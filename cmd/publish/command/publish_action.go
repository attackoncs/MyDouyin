/*
 * 发布视频 操作业务逻辑
 */

package command

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strings"

	"MyDouyin/kitex_gen/publish"
	"MyDouyin/pkg/minio"
	"MyDouyin/pkg/ttviper"

	"MyDouyin/dal/db"

	"github.com/gofrs/uuid"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type PublishActionService struct {
	ctx context.Context
}

// NewPublishActionService new PublishActionService
func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

// PublishAction publish video.
func (s *PublishActionService) PublishAction(req *publish.DouyinPublishActionRequest, uid int, cfg *ttviper.Config) (err error) {
	MinioVideoBucketName := minio.MinioVideoBucketName
	videoData := []byte(req.Data)

	// // 获取后缀
	// filetype := http.DetectContentType(videoData)

	// byte[] -> reader
	reader := bytes.NewReader(videoData)
	u2, err := uuid.NewV4()
	if err != nil {
		return err
	}
	fileName := u2.String() + "." + "mp4"
	// 上传视频
	err = minio.UploadFile(MinioVideoBucketName, fileName, reader, int64(len(videoData)))
	if err != nil {
		return err
	}
	// 获取视频链接
	url, err := minio.GetFileUrl(MinioVideoBucketName, fileName, 0)
	playUrl := strings.Split(url.String(), "?")[0]
	if err != nil {
		return err
	}

	u3, err := uuid.NewV4()
	if err != nil {
		return err
	}

	// 获取封面
	coverPath := u3.String() + "." + "jpg"
	coverData, err := readFrameAsJpeg(playUrl)
	if err != nil {
		return err
	}

	// 上传封面
	coverReader := bytes.NewReader(coverData)
	err = minio.UploadFile(MinioVideoBucketName, coverPath, coverReader, int64(len(coverData)))
	if err != nil {
		return err
	}

	// 获取封面链接
	coverUrl, err := minio.GetFileUrl(MinioVideoBucketName, coverPath, 0)
	if err != nil {
		return err
	}

	CoverUrl := strings.Split(coverUrl.String(), "?")[0]

	// 封装video
	videoModel := &db.Video{
		AuthorID:      uid,
		PlayUrl:       playUrl,
		CoverUrl:      CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.Title,
	}
	return db.CreateVideo(s.ctx, videoModel)
}

// ReadFrameAsJpeg
// 从视频流中截取一帧并返回 需要在本地环境中安装ffmpeg并将bin添加到环境变量
func readFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), err
}
