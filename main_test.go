package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type BoolTest struct {
	Active bool `gorm:"type:tinyint;default:0"`
}

func TestGORM(t *testing.T) {

	if err := DB.AutoMigrate(&BoolTest{}); err != nil {
		t.Errorf("AutoMigrate failed, got error: %v", err)
	}
}
