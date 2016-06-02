// auto generated file
// DO NOT EDIT!

package server

import (
	"../core"
	"../model"
	"../pb"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

type UserExtraServer struct{}

func (s *UserExtraServer) GetUserExtra(ctx context.Context, in *pb.GetUserExtraRequest) (*pb.UserExtraObject, error) {
	db := core.GetMasterDB(core.READ_MODE)
	var user_extra model.UserExtra
	db.Find(&user_extra, in.Uid)
	return &pb.UserExtraObject{
		Key2:     user_extra.Key2,
		Key1:     user_extra.Key1,
		UpdateAt: user_extra.UpdateAt.String,
		Uid:      user_extra.Uid,
		CreateAt: user_extra.CreateAt.String,
	}, nil
}

func (s *UserExtraServer) CreateUserExtra(ctx context.Context, in *pb.UserExtraObject) (*pb.CreateUserExtraResponse, error) {
	db := core.GetMasterDB(core.READ_MODE)

	r := db.Create(
		&model.UserExtra{
			Key2:     in.Key2,
			Key1:     in.Key1,
			UpdateAt: in.UpdateAt,
			Uid:      in.Uid,
			CreateAt: in.CreateAt,
		})
	return &pb.CreateUserExtraResponse{
		AffectNum: r.RowsAffected,
	}, nil
}

func (s *UserExtraServer) UpdateUserExtra(ctx context.Context, in *pb.UpdateUserExtraRequest) (*pb.UpdateUserExtraResponse, error) {
	db := core.GetMasterDB(core.READ_MODE)
	var user_extra model.UserExtra
	db.Find(&user_extra, in.Uid)
	update_user_extra := in.UserExtra

	row_affected := db.Model(&model.UserExtra{Uid: in.Uid}).Updates(
		model.UserExtra{
			Key2:     update_user_extra.Key2,
			Key1:     update_user_extra.Key1,
			UpdateAt: update_user_extra.UpdateAt,
			Uid:      update_user_extra.Uid,
			CreateAt: update_user_extra.CreateAt,
		}).RowsAffected
	return &pb.UpdateUserExtraResponse{
		AffectNum: row_affected,
	}, nil
}

func (s *UserExtraServer) QueryUserExtra(ctx context.Context, in *pb.UserExtraObject) (*pb.QueryUserExtraResponse, error) {
	var user_extra_objects []*pb.UserExtraObject
	db := core.GetMasterDB(core.READ_MODE)
	var user_extras []model.UserExtra

	db.Where(&model.UserExtra{Uid: in.Uid}).Find(&user_extras)
	for _, user_extra := range user_extras {
		user_extra_objects = append(user_extra_objects,
			&pb.UserExtraObject{
				Key2:     user_extra.Key2,
				Key1:     user_extra.Key1,
				UpdateAt: user_extra.UpdateAt.String,
				Uid:      user_extra.Uid,
				CreateAt: user_extra.CreateAt.String,
			},
		)
	}

	return &pb.QueryUserExtraResponse{
		UserExtras: user_extra_objects,
	}, nil
}

func (s *UserExtraServer) QueryOneUserExtra(ctx context.Context, in *pb.UserExtraObject) (*pb.UserExtraObject, error) {
	db := core.GetMasterDB(core.READ_MODE)
	var user_extra model.UserExtra

	db.Where(&model.UserExtra{Uid: in.Uid}).Find(&user_extra)
	if user_extra.Uid != 0 {
		return &pb.UserExtraObject{
			Key2:     user_extra.Key2,
			Key1:     user_extra.Key1,
			UpdateAt: user_extra.UpdateAt.String,
			Uid:      user_extra.Uid,
			CreateAt: user_extra.CreateAt.String,
		}, nil
	}

	return &pb.UserExtraObject{}, nil
}
