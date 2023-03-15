package impl

//import (
//	"context"
//	"fmt"
//
//	"github.com/tqtcloud/workflow-backend/apps/task"
//
//	"github.com/infraboard/mcube/exception"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//)
//
//func (s *service) save(ctx context.Context, ins *task.Task) error {
//	if _, err := s.col.InsertOne(ctx, ins); err != nil {
//		return exception.NewInternalServerError("inserted task(%s) document error, %s",
//			ins.Data.Name, err)
//	}
//	return nil
//}
//
//func (s *service) get(ctx context.Context, id string) (*task.Task, error) {
//	filter := bson.M{"_id": id}
//
//	ins := task.NewDefaultTask()
//	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
//		if err == mongo.ErrNoDocuments {
//			return nil, exception.NewNotFound("task %s not found", id)
//		}
//
//		return nil, exception.NewInternalServerError("find task %s error, %s", id, err)
//	}
//
//	return ins, nil
//}
//
//func newQueryTaskRequest(r *task.QueryTaskRequest) *queryTaskRequest {
//	return &queryTaskRequest{
//		r,
//	}
//}
//
//type queryTaskRequest struct {
//	*task.QueryTaskRequest
//}
//
//func (r *queryTaskRequest) FindOptions() *options.FindOptions {
//	pageSize := int64(r.Page.PageSize)
//	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)
//
//	opt := &options.FindOptions{
//		Sort: bson.D{
//			{Key: "create_at", Value: -1},
//		},
//		Limit: &pageSize,
//		Skip:  &skip,
//	}
//
//	return opt
//}
//
//func (r *queryTaskRequest) FindFilter() bson.M {
//	filter := bson.M{}
//	if r.Keywords != "" {
//		filter["$or"] = bson.A{
//			bson.M{"data.name": bson.M{"$regex": r.Keywords, "$options": "im"}},
//			bson.M{"data.author": bson.M{"$regex": r.Keywords, "$options": "im"}},
//		}
//	}
//	return filter
//}
//
//func (s *service) query(ctx context.Context, req *queryTaskRequest) (*task.TaskSet, error) {
//	resp, err := s.col.Find(ctx, req.FindFilter(), req.FindOptions())
//
//	if err != nil {
//		return nil, exception.NewInternalServerError("find task error, error is %s", err)
//	}
//
//	set := task.NewTaskSet()
//	// 循环
//	for resp.Next(ctx) {
//		ins := task.NewDefaultTask()
//		if err := resp.Decode(ins); err != nil {
//			return nil, exception.NewInternalServerError("decode task error, error is %s", err)
//		}
//
//		set.Add(ins)
//	}
//
//	// count
//	count, err := s.col.CountDocuments(ctx, req.FindFilter())
//	if err != nil {
//		return nil, exception.NewInternalServerError("get task count error, error is %s", err)
//	}
//	set.Total = count
//
//	return set, nil
//}
//
//func (s *service) update(ctx context.Context, ins *task.Task) error {
//	if _, err := s.col.UpdateByID(ctx, ins.Id, ins); err != nil {
//		return exception.NewInternalServerError("inserted task(%s) document error, %s",
//			ins.Data.Name, err)
//	}
//
//	return nil
//}
//
//func (s *service) deleteTask(ctx context.Context, ins *task.Task) error {
//	if ins == nil || ins.Id == "" {
//		return fmt.Errorf("task is nil")
//	}
//
//	result, err := s.col.DeleteOne(ctx, bson.M{"_id": ins.Id})
//	if err != nil {
//		return exception.NewInternalServerError("delete task(%s) error, %s", ins.Id, err)
//	}
//
//	if result.DeletedCount == 0 {
//		return exception.NewNotFound("task %s not found", ins.Id)
//	}
//
//	return nil
//}
