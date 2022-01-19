package model

import "GoHttpServerBestPractice/service/core"

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/18 17:16
*/
type Student struct {
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty" db:"name" validate:""`
	Age       int64  `json:"age,omitempty" db:"age" validate:"required"`
	Class     string `json:"class,omitempty" db:"class" validate:"required"`
	CreatedAt int64  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt int64  `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt int64  `json:"deleted_at,omitempty" db:"deleted_at"`
}

var Stu = core.Model{
	M:     nil, // M: new(Student) 传入模型的结构体指针
	Table: "student",
	CreateField: core.CreateField{
		CreatedFields:        nil,
		CreatedIgnoreFields:  []string{"deleted_at"},
		CreatedSetTimeFields: []string{"created_at", "updated_at"},
	},
	SoftDeleteField: core.SoftDeleteField{
		DeletedFields: "deleted_at",
	},
	UpdateField: core.UpdateField{
		UpdateFields:        nil,
		UpdateIgnoreFields:  []string{"created_at", "deleted_at"},
		UpdateSetTimeFields: []string{"updated_at"},
	},
	SelectField: core.SelectField{
		SelectFields:       nil,
		SelectIgnoreFields: []string{"deleted_at"},
	},
	SelectFieldList: core.SelectFieldList{
		Search:  []string{"name", "age"},
		Filter:  nil,
		Sort:    []string{"id"},
		PageMax: 100,
		PageMin: 17,
	},
}
