package com.grfsv.palbum.core.database

import androidx.room.Database
import androidx.room.RoomDatabase
import androidx.room.TypeConverters
import com.grfsv.palbum.core.database.dao.DtkDao
import com.grfsv.palbum.core.database.entity.Dtk
import com.grfsv.palbum.core.database.utils.LocalDateConvertor

@Database(
    entities = [Dtk::class],
    version = 1,
)
@TypeConverters(LocalDateConvertor::class)
abstract class Database : RoomDatabase() {
    abstract fun dtkDao() : DtkDao


}