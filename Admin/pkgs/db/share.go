package db

import (
	"errors"
	"fmt"

	"gpc/db/mongodb/cst"

	"rcdd/project/VSTH/Admin/pkgs/utils"
)

//-----------------------------------------------------------------------------------------------------------//

type SharePackString struct{}

func (self *SharePackString) UserId(id int) string {
	return fmt.Sprintf("%s=%v", cst.Field_UserId, id)
}

func (self *SharePackString) InChiKey(key string) string {
	return fmt.Sprintf("%s=%v", cst.Field_InChiKey, key)
}

func (self *SharePackString) MolId(id int) string {
	return fmt.Sprintf("%s=%v", cst.Field_MolId, id)
}

//-----------------------------------------------------------------------------------------------------------//

type ShareChkField struct{}

func (self *ShareChkField) StdName(name string) error {
	if !utils.CheckStdName(name) {
		return errors.New(obj.pack.FieldInvalid(cst.Field_Name, "[a-z0-9A-Z]len=4,100"))
	}
	return nil
}

func (self *ShareChkField) User(user *User) (err error) {
	if user == nil || user.Id <= 0 {
		return errors.New(obj.pack.FieldInvalid(cst.Field_UserId))
	}
	return
}

func (self *ShareChkField) Mol(mol *ZincMol) (err error) {
	if mol == nil || mol.Id <= 0 {
		return errors.New(obj.pack.FieldInvalid(cst.Field_MolId))
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//

type ShareCollDoc struct{}

func (self *ShareCollDoc) CheckUser(user *User) (*User, error) {
	if err := obj.chk.User(user); err != nil {
		return nil, err
	}
	return self.CheckUserById(user.Id)
}

func (self *ShareCollDoc) CheckUserById(id int) (v *User, err error) {
	v = &User{Id: id}
	v.Init()
	if err = v.ReadById(); err == nil && !v.Status.IsOn() {
		return nil, obj.err.StatusOFF(obj.pack.Pack(CST_COLL_USER, v.Name, v.Id))
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//

func (self *ShareCollDoc) CheckMol(mol *ZincMol) (*ZincMol, error) {
	if err := obj.chk.Mol(mol); err != nil {
		return nil, err
	}
	return self.CheckMolById(mol.Id)
}

func (self *ShareCollDoc) CheckMolById(id int) (v *ZincMol, err error) {
	v = &ZincMol{Id: id}
	v.Init()
	if err = v.ReadById(); err == nil {
		return nil, obj.err.StatusOFF(obj.pack.Pack(CST_COLL_ZINCMOL, v.Name, v.Id))
	}
	return
}

//-----------------------------------------------------------------------------------------------------------//
