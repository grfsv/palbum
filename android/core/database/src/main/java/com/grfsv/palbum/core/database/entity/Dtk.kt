package com.grfsv.palbum.core.database.entity

import androidx.room.Entity
import kotlinx.datetime.LocalDate

@Entity(primaryKeys = ["friendUuid", "date"])
data class Dtk (
    val friendUuid: String,
    val date: LocalDate,
    val dtk: String
)