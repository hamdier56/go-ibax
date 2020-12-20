/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package model

type VDEScheMember struct {
	ID                   int64  `gorm:"primary_key; not null" json:"id"`
	VDEPubKey            string `gorm:"not null" json:"vde_pub_key"`
	VDEComment           string `gorm:"not null" json:"vde_comment"`
	VDEName              string `gorm:"not null" json:"vde_name"`
	VDEIp                string `gorm:"not null" json:"vde_ip"`
	VDEType              int64  `gorm:"not null" json:"vde_type"`
	ContractRunHttp      string `gorm:"not null" json:"contract_run_http"`
	ContractRunEcosystem string `gorm:"not null" json:"contract_run_ecosystem"`

	UpdateTime int64 `gorm:"not null" json:"update_time"`
	CreateTime int64 `gorm:"not null" json:"create_time"`
}

func (VDEScheMember) TableName() string {
	return "vde_sche_member"
}

func (m *VDEScheMember) Create() error {
	return DBConn.Create(&m).Error
}

func (m *VDEScheMember) Updates() error {
	return DBConn.Model(m).Updates(m).Error
}

func (m *VDEScheMember) Delete() error {
	return DBConn.Delete(m).Error
}

func (m *VDEScheMember) GetAll() ([]VDEScheMember, error) {
	var result []VDEScheMember
	err := DBConn.Find(&result).Error
	return result, err
}
func (m *VDEScheMember) GetOneByID() (*VDEScheMember, error) {
	err := DBConn.Where("id=?", m.ID).First(&m).Error
	return m, err
}

func (m *VDEScheMember) GetOneByPubKey(VDEPubKey string) (*VDEScheMember, error) {
	err := DBConn.Where("vde_pub_key=?", VDEPubKey).First(&m).Error
	return m, err
}
