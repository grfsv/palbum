package com.grfsv.palbum.core.database.dao

import androidx.room.Dao
import androidx.room.Insert
import androidx.room.Query
import com.grfsv.palbum.core.database.entity.Dtk
import kotlinx.coroutines.flow.Flow
import kotlinx.datetime.LocalDate

@Dao
interface DtkDao {
    @Insert
    fun insertAll(vararg Dtks: Dtk)

    @Query("SELECT * FROM dtk WHERE date = :date")
    fun listByDate(date: LocalDate): Flow<List<Dtk>>

    @Query("SELECT * FROM dtk WHERE friendUuid = :friendUUID")
    fun listByFriendUUID(friendUUID: String) : Flow<List<Dtk>>
}