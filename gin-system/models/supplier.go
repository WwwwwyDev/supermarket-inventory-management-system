package models

import "time"

type Supplier struct {
	SupplierId         int       `gorm:"column:supplier_id"`
	SupplierName       string    `gorm:"column:supplier_name"`
	SupplierShortname  string    `gorm:"column:supplier_shortname"`
	SupplierAddress    string    `gorm:"column:supplier_address"`
	SupplierCtelephone string    `gorm:"column:supplier_ctelephone"`
	SupplierEmail      string    `gorm:"column:supplier_email"`
	SupplierLiaisonman string    `gorm:"column:supplier_liaisonman"`
	SupplierLtelephone string    `gorm:"column:supplier_ltelephone"`
	SupplierRemarks    string    `gorm:"column:supplier_remarks"`
	IsDel              bool      `gorm:"column:is_del"`
	UpdateTime         time.Time `gorm:"update_time"`
}
