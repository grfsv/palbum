package com.grfsv.palbum.core.database.di

import com.grfsv.palbum.core.database.Database
import com.grfsv.palbum.core.database.dao.DtkDao
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent

@Module
@InstallIn(SingletonComponent::class)
internal object DaoModule {

    @Provides
    fun providesDtkDao(database: Database): DtkDao {
        return database.dtkDao()
    }
}