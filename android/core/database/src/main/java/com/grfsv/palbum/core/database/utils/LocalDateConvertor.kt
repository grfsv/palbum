package com.grfsv.palbum.core.database.utils

import androidx.room.TypeConverter
import kotlinx.datetime.LocalDate

class LocalDateConvertor {
    @TypeConverter
    fun localDateToLong(date: LocalDate?): Long? =
        date?.toEpochDays()

    @TypeConverter
    fun longToLocalDate(epochDay: Long?): LocalDate? =
        epochDay?.let { LocalDate.fromEpochDays(it) }
}