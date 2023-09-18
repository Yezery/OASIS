package repositories

// db.Delete(&user) // 将 user 标记为软删除状态，即将 user.deleted_at 设为当前时间
// db.Unscoped().Delete(&user) // 强制执行该语句将物理删除被标记的行
// db.Unscoped().Where("deleted_at is not null").Delete(&User{}) // 利用 Where 子句物理删除所有已经被标记删除的行
