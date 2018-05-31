package dbbootstrap

import enm "filmworldservice/common/enums"
import sch "filmworldservice/dataservice/sqlservice/dataschema"

type  InfDbInvoker interface {
	InvokeDatabaseInitialize(op enm.DbOperation)(bool)
}
type  BoosTapper  bool
func(b *BoosTapper) InvokeDatabaseInitialize(op enm.DbOperation)(bool){
	scheme:= new(sch.GeneratedSchema)
	switch op{
	   case  enm.SqlConnection:
		return scheme.GenerateSchema()
	   case enm.MockConnection:
		return false
	   default:
		return false
	}
	return false
}