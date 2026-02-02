package com.grfsv.palbum.core.datastore.di

import android.content.Context
import androidx.datastore.core.DataStoreFactory
import com.google.crypto.tink.Aead
import com.google.crypto.tink.integration.android.AndroidKeystore
import com.grfsv.palbum.core.datastore.AuthLocalDataSource
import com.grfsv.palbum.core.datastore.RefreshTokenSerializer
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.android.qualifiers.ApplicationContext
import dagger.hilt.components.SingletonComponent
import javax.inject.Singleton

@Module
@InstallIn(SingletonComponent::class)
object DataStoreModule {

    private const val KEY_ALIAS = "refresh_token_key"

    @Provides
    @Singleton
    internal fun providesAead() : Aead{
        if (!AndroidKeystore.hasKey(KEY_ALIAS)) {
            AndroidKeystore.generateNewAes256GcmKey(KEY_ALIAS)
        }
        return AndroidKeystore.getAead(KEY_ALIAS)
    }
    
    @Provides
    @Singleton
    internal fun providesAuthLocalDataSource(
        @ApplicationContext context: Context,
        refreshTokenSerializer: RefreshTokenSerializer
    ): AuthLocalDataSource {
        val dataStore = DataStoreFactory.create(
            serializer = refreshTokenSerializer,
            produceFile = { context.filesDir.resolve("refresh_token.pb") },
        )
        return AuthLocalDataSource(dataStore)
    }
}