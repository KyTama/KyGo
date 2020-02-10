package product
//
//import (
//	"NewGo/utils"
//	"github.com/juju/mgosession"
//	"gopkg.in/mgo.v2/bson"
//	"os"
//)
//
//type mongoRepo struct {
//	pool *mgosession.Pool
//}
////
//func NewMongoRepo(p *mgosession.Pool) Repository {
//	return &mongoRepo{p}
//}
//
//func (r *mongoRepo) FindAll() (data []Products, err error) {
//
//	return
//}
//
//func (r *mongoRepo) FindByID(id utils.ID) (data Products, err error) {
//	result := Products{}
//	session := r.pool.Session(nil)
//	coll := session.DB(os.Getenv("MONGODB_DATABASE")).C("product")
//	err = coll.Find(bson.M{"_id": id}).One(&result)
//	data = result
//	if err != nil {
//		return data, err
//	}
//
//	return
//}
//
//func (r *mongoRepo) InsertProduct(product Products) (data Products, err error) {
//
//	return
//}
