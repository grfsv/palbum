package com.grfsv.palbum.core.data.di

import com.grfsv.palbum.core.data.SessionExpiryListener
import com.grfsv.palbum.core.data.SessionStateHolder
import com.grfsv.palbum.core.data.repository.AuthRepository
import com.grfsv.palbum.core.data.repository.OnlineAuthRepository
import dagger.Binds
import dagger.Module
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent

@Module
@InstallIn(SingletonComponent::class)
internal interface DataModule {

    @Binds
    fun bindsAuthRepository(
        authRepository: OnlineAuthRepository
    ): AuthRepository

    @Binds
    fun bindsSessionExpiryListener(
        sessionStateHolder: SessionStateHolder
    ): SessionExpiryListener
}