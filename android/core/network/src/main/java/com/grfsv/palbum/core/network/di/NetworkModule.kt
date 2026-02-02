package com.grfsv.palbum.core.network.di

import com.grfsv.palbum.core.network.AuthNetworkDataSource
import com.grfsv.palbum.core.network.BuildConfig
import com.grfsv.palbum.core.network.UserNetworkDataSource
import com.grfsv.palbum.core.network.retrofit.Interceptor
import com.grfsv.palbum.core.network.retrofit.Authenticator
import com.grfsv.palbum.core.network.retrofit.RetrofitAuthNetwork
import com.grfsv.palbum.core.network.retrofit.RetrofitUserNetwork
import com.jakewharton.retrofit2.converter.kotlinx.serialization.asConverterFactory
import dagger.Binds
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent
import kotlinx.serialization.json.Json
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.OkHttpClient
import okhttp3.logging.HttpLoggingInterceptor
import retrofit2.Retrofit
import java.util.concurrent.TimeUnit
import javax.inject.Singleton


@Module
@InstallIn(SingletonComponent::class)
internal object NetworkModule {

    @Provides
    @Singleton
    fun providesNetworkJson(): Json = Json {
        ignoreUnknownKeys = true
    }

    @Provides
    @Singleton
    internal fun providesOkClient(
        interceptor: Interceptor,
        authenticator: Authenticator
    ): OkHttpClient {
        val builder = OkHttpClient.Builder()
            .connectTimeout(30, TimeUnit.SECONDS)
            .readTimeout(30, TimeUnit.SECONDS)
            .writeTimeout(30, TimeUnit.SECONDS)
            .addInterceptor(interceptor)
            .authenticator(authenticator)

        // デバッグビルドのみログ出力
        if (BuildConfig.DEBUG) {
            val loggingInterceptor = HttpLoggingInterceptor().apply {
                level = HttpLoggingInterceptor.Level.BODY
            }
            builder.addInterceptor(loggingInterceptor)
        }

        return builder.build()
    }


    @Provides
    @Singleton
    internal fun providesRetrofit(
        networkJson: Json,
    okHttpClient: OkHttpClient
    ): Retrofit =
        Retrofit.Builder()
            .baseUrl(BuildConfig.BASE_URL)
            .client(okHttpClient)
            .addConverterFactory(
                networkJson.asConverterFactory("application/json".toMediaType())
            )
            .build()
}

@Module
@InstallIn(SingletonComponent::class)
internal interface NetworkDataSourceModule {

    @Binds
    fun bindsUserNetworkDataSource(
        userNetworkDataSource: RetrofitUserNetwork
    ): UserNetworkDataSource

    @Binds
    fun bindsAuthNetworkDataSource(
        authNetworkDataSource: RetrofitAuthNetwork
    ): AuthNetworkDataSource
}