package repository

import (
	"context"
	"log/slog"
	"time"
	"vs-so-sanh/infrastructures/configuration"
	"vs-so-sanh/internal/device"
	"vs-so-sanh/internal/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoDbDeviceRepository struct {
	cfg    configuration.Config
	client *mongo.Client
}

func NewMongoDbDeviceRepository(cfg configuration.Config, client *mongo.Client) device.Repository {
	return &MongoDbDeviceRepository{
		cfg:    cfg,
		client: client,
	}
}

func (m *MongoDbDeviceRepository) Save(c context.Context, device model.Device) (*model.Device, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MongoDbDeviceRepository) FindByName(c context.Context, deviceName string) (*model.Device, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.client.Database(m.cfg.Database.DBName).Collection("brands")

	pipeline := mongo.Pipeline{
		// 1. Tìm document Brand chứa điện thoại này (để lọc bớt dữ liệu rác)
		{{"$match", bson.M{"devices.model_name": deviceName}}},

		// 2. Tách mảng devices ra thành các document riêng lẻ
		{{"$unwind", "$devices"}},

		// 3. Lọc lại một lần nữa để lấy đúng điện thoại cần tìm
		{{"$match", bson.M{"devices.model_name": deviceName}}},

		// 4. (Quan trọng) Đưa object device lên làm root document
		// Lúc này kết quả trả về sẽ khớp trực tiếp với struct model.Device
		{{"$replaceRoot", bson.M{"newRoot": "$devices"}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Kiểm tra xem có kết quả không
	if cursor.Next(ctx) {
		var device model.Device
		// Decode trực tiếp vào model.Device, không cần qua struct trung gian
		if err := cursor.Decode(&device); err != nil {
			return nil, err
		}
		return &device, nil
	}

	return nil, nil // Không tìm thấy
}

func (m *MongoDbDeviceRepository) FindTop20(c context.Context) ([]model.Device, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := m.client.Database(m.cfg.Database.DBName).Collection("brands")

	pipeline := mongo.Pipeline{
		// 1. Tách mảng devices ra thành từng dòng riêng biệt
		{{"$unwind", "$devices"}},

		// 2. Bỏ qua trường specifications (Project)
		{{"$project", bson.D{
			{"model_name", "$devices.model_name"},
			{"imageUrl", "$devices.imageUrl"},
			{"_id", 0}, // Không lấy ID của Brand
		}}},

		// 3. Giới hạn 20 cái đầu tiên
		{{"$limit", 20}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var devices []model.Device
	if err := cursor.All(ctx, &devices); err != nil {
		return nil, err
	}

	slog.Info("devices: ", devices)

	return devices, nil
}

func (m *MongoDbDeviceRepository) Update(c context.Context, device model.Device) (*model.Device, error) {
	//TODO implement me
	panic("implement me")
}
