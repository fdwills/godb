// auto generated file
// DO NOT EDIT!

package server

import (
	"../core"
	"../model"
	"../pb"
	"database/sql"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

type UserServer struct{}

func (s *UserServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.UserObject, error) {
	db := core.GetMasterDB(core.READ_MODE)
	var user model.User
	db.Find(&user, in.Uid)
	return &pb.UserObject{
		Uid:      user.Uid,
		UpdateAt: user.UpdateAt.String,
		CreateAt: user.CreateAt.String,
		Phone:    user.Phone.String,
		Password: user.Password.String,
		Name:     user.Name.String,
	}, nil
}

func (s *UserServer) CreateUser(ctx context.Context, in *pb.UserObject) (*pb.CreateUserResponse, error) {
	db := core.GetMasterDB(core.READ_MODE)

	r := db.Create(
		&model.User{
			Uid:      in.Uid,
			UpdateAt: in.UpdateAt,
			CreateAt: in.CreateAt,
			Phone:    sql.NullString{in.Phone, true},
			Password: sql.NullString{in.Password, true},
			Name:     sql.NullString{in.Name, true},
		})
	return &pb.CreateUserResponse{
		AffectNum: r.RowsAffected,
	}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	db := core.GetMasterDB(core.READ_MODE)
	var user model.User
	db.Find(&user, in.Uid)
	update_user := in.User

	row_affected := db.Model(&model.User{Uid: in.Uid}).Updates(
		model.User{
			Uid:      update_user.Uid,
			UpdateAt: update_user.UpdateAt,
			CreateAt: update_user.CreateAt,
			Phone:    sql.NullString{update_user.Phone, true},
			Password: sql.NullString{update_user.Password, true},
			Name:     sql.NullString{update_user.Name, true},
		}).RowsAffected
	return &pb.UpdateUserResponse{
		AffectNum: row_affected,
	}, nil
}

func (s *UserServer) QueryUser(ctx context.Context, in *pb.UserObject) (*pb.QueryUserResponse, error) {
	var user_objects []*pb.UserObject
	db := core.GetMasterDB(core.READ_MODE)
	var users []model.User

	db.Where(&model.User{Uid: in.Uid}).Find(&users)
	for _, user := range users {
		user_objects = append(user_objects,
			&pb.UserObject{
				Uid:      user.Uid,
				UpdateAt: user.UpdateAt.String,
				CreateAt: user.CreateAt.String,
				Phone:    user.Phone.String,
				Password: user.Password.String,
				Name:     user.Name.String,
			},
		)
	}

	return &pb.QueryUserResponse{
		Users: user_objects,
	}, nil
}

func (s *UserServer) QueryOneUser(ctx context.Context, in *pb.UserObject) (*pb.UserObject, error) {
	db := core.GetMasterDB(core.READ_MODE)
	var user model.User

	db.Where(&model.User{Uid: in.Uid}).Find(&user)
	if user.Uid != 0 {
		return &pb.UserObject{
			Uid:      user.Uid,
			UpdateAt: user.UpdateAt.String,
			CreateAt: user.CreateAt.String,
			Phone:    user.Phone.String,
			Password: user.Password.String,
			Name:     user.Name.String,
		}, nil
	}

	return &pb.UserObject{}, nil
}
