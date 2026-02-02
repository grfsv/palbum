package com.grfsv.palbum.core.data.di

import com.grfsv.palbum.core.data.TokenProviderImpl
import com.grfsv.palbum.core.network.TokenProvider
import dagger.Binds
import dagger.Module
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent

@Module
@InstallIn(SingletonComponent::class)
abstract class ImplementModule {

    @Binds
    abstract fun bindsTokenProvider(
        tokenProvider: TokenProviderImpl
    ): TokenProvider
}