package repo

import (
	"fmt"
	"gorm.io/gorm"
	"watchAlert/internal/models"
)

type (
	LdapRepo struct {
		entryRepo
	}

	InterLdapRepo interface {
		Create(r models.Ldap) error
		Update(r models.Ldap) error
		Get() (*models.Ldap, error)
	}
)

func newLdapInterface(db *gorm.DB, g InterGormDBCli) InterLdapRepo {
	return &LdapRepo{
		entryRepo{
			g:  g,
			db: db,
		},
	}
}

func (ur LdapRepo) Create(r models.Ldap) error {
	data, _ := ur.Get()
	if data != nil {
		return fmt.Errorf("ldap is already exist")
	}
	err := ur.g.Create(&models.Ldap{}, r)
	if err != nil {
		return err
	}
	return nil
}

func (ur LdapRepo) Update(r models.Ldap) error {
	u := Updates{
		Table: models.Ldap{},
		Where: map[string]interface{}{
			"id = ?": r.ID,
		},
		Updates: r,
	}

	err := ur.g.Updates(u)
	if err != nil {
		return err
	}
	return nil
}

func (ur LdapRepo) Get() (*models.Ldap, error) {
	var data *models.Ldap
	var db = ur.db.Model(models.Ldap{})
	err := db.Where("status = ?", 1).First(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
