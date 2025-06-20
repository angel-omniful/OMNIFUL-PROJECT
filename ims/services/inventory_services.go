package services

import (
	"context"
	"time"
	"log"
	"github.com/angel-omniful/ims/models"
	"github.com/angel-omniful/ims/myDb"
	"github.com/omniful/go_commons/redis"
)

var cache *redis.Client=myDb.GetCache()


func CreateInventory(ctx context.Context, inv *models.Inventory) error {
	err:=db.Create(inv).Error
	if err != nil {
		return err
	}
    key:=inv.SKUID + ":" + inv.HubID
	_,err1:=cache.Set(ctx,key,"valid",1*time.Hour)
	if err1 != nil {
		log.Println("Error setting cache for inventory:", err1)
	}
	return nil
}
func GetAllInventory(ctx context.Context) ([]models.Inventory, error) {
	var inventory []models.Inventory
	err := db.Find(&inventory).Error
	return inventory, err
}

func GetInventoryByID(ctx context.Context, id string) (models.Inventory, error) {
	var inv models.Inventory
	
	err := db.First(&inv, "id = ?", id).Error

	key:=inv.SKUID + ":" + inv.HubID
	_,err1:=cache.Set(ctx,key,"valid",1*time.Hour)
	if err1 != nil {
		log.Println("Error setting cache for inventory:", err1)
	}
	return inv, err
}

func UpdateInventory(ctx context.Context, id string, updated *models.Inventory) error {
	err:=db.Model(&models.Inventory{}).Where("id = ?", id).Updates(updated).Error
	if err != nil {
		return err	
	}
	key := updated.SKUID + ":" + updated.HubID
	_,err1:=cache.Set(ctx,key,"valid",1*time.Hour)
	if err1 != nil {
		log.Println("Error updating cache for inventory:", err1)
	}	
	return nil	
}

func DeleteInventory(ctx context.Context, id string) error {
	inv,err:=GetInventoryByID(ctx, id)

	if err != nil {
		log.Println("Error fetching inventory for deletion:", err)
		return err
	}

	key:=inv.SKUID + ":" + inv.HubID

	_,err1:=cache.Del(ctx, key)
	if err1 != nil {
		log.Println("Error deleting cache for inventory:", err1)
	}

	err2:=db.Delete(&models.Inventory{}, "id = ?", id).Error
	return err2
}

func ValidateInventory(ctx context.Context, skuID, hubID string) (bool, error) {
	key := skuID + ":" + hubID
	val, err := cache.Get(ctx, key)
	if err != nil {
		return false, err
	}
	return val == "valid", nil
}
