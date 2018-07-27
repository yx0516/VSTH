package gsession

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/mgo.v2"

	mod "gpc/util/sync"
)

var (
	globalSession *mgo.Session          // 全局 Session
	globalDBName  string                // 全局 数据库名
	globalMod     = mod.NewSyncMod(100) // 控制 Session 的 Clone 和 Copy 切换

)

//-----------------------------------------------------------------------------------------------------------//

// 模块使用前(登录认证后)必须调用该方法初始化 session
func InitSession(session *mgo.Session, dataBaseName string) error {
	if session == nil {
		return errors.New("session is nil.")
	}
	if err := session.Ping(); err != nil {
		return fmt.Errorf("session.ping error:%s", err.Error())
	}

	globalDBName = strings.TrimSpace(dataBaseName)
	globalSession = session

	return nil
}

// 设置 多少次 Session.Clone 后切换一次 Session.Copy【默认每100次就Copy一个】
func SetSessionCopy(i int) {
	globalMod.SetMod(i)
}

// 关闭 mgo 的 全局 session
func CloseSession() {
	if globalSession != nil {
		globalSession.LogoutAll()
		globalSession.Close()
	}
}

//-----------------------------------------------------------------------------------------------------------//

// 克隆/复制 Session 使用后不要时，记得关闭：Session.Close()
func NewSession() *mgo.Session {
	if globalMod.Do() {
		return globalSession.Copy()
	} else {
		session := globalSession.Clone()
		if err := session.Ping(); err != nil {
			session = globalSession.Copy()
		}
		return session
	}
}

// 克隆/复制 Session 的包装器
func WithSession(fun func(session *mgo.Session)) {
	session := NewSession()
	defer session.Close()
	fun(session)
}

// 克隆/复制 Session 后定位到 数据库，使用后不要时，记得关闭： Database.Session.Close()
func NewDB() *mgo.Database {
	return NewSession().DB(globalDBName)
}

// 克隆/复制 Session.DB 的包装器
func WithDB(fun func(db *mgo.Database)) {
	db := NewDB()
	defer db.Session.Close()
	fun(db)
}

//-----------------------------------------------------------------------------------------------------------//
