package handler

import (
	"context"
	"course/course-srv/proto/dto"
	"course/proto/basic"
	"course/public"
	"course/public/util"
	"github.com/micro/go-micro/v2/errors"
)

func (c *CourseServiceHandler) ListChapter(ctx context.Context, in *dto.ChapterPageDto, out *dto.ChapterPageDto) error {
	list, exception := chapterDao.List(in)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.CourseServiceName, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Rows = list
	return nil
}

func (c *CourseServiceHandler) SaveChapter(ctx context.Context, in *dto.ChapterDto, out *dto.ChapterDto) error {
	chapterDto, exception := chapterDao.Save(in)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.CourseServiceName, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, chapterDto)
	return nil
}

func (c *CourseServiceHandler) DeleteChapter(ctx context.Context, in *basic.String, out *basic.String) error {
	exception := chapterDao.Delete(in.Str)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.CourseServiceName, exception.Error(), exception.Code())
	}
	return nil
}
