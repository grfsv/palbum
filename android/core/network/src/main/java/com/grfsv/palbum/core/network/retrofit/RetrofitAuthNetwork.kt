package com.grfsv.palbum.core.network.retrofit

import com.grfsv.palbum.core.model.LoginRequest
import com.grfsv.palbum.core.model.LogoutRequest
import com.grfsv.palbum.core.model.RefreshRequest
import com.grfsv.palbum.core.model.SignUpRequest
import com.grfsv.palbum.core.model.TokenPair
import com.grfsv.palbum.core.network.AuthNetworkDataSource
import com.grfsv.palbum.core.network.BuildConfig
import com.grfsv.palbum.core.network.model.NetworkResponse
import com.jakewharton.retrofit2.converter.kotlinx.serialization.asConverterFactory
import kotlinx.serialization.json.Json
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.OkHttpClient
import okhttp3.logging.HttpLoggingInterceptor
import retrofit2.Retrofit
import retrofit2.create
import retrofit2.http.Body
import retrofit2.http.GET
import retrofit2.http.POST
import javax.inject.Inject


private interface RetrofitAuthNetworkApi {
    @NoAuth
    @POST("auth/signup")
    suspend fun signUp(@Body request: SignUpRequest) : NetworkResponse<TokenPair>

    @NoAuth
    @POST("auth/login")
    suspend fun login(@Body request: LoginRequest) : NetworkResponse<TokenPair>

    @POST("auth/logout")
    suspend fun logout(@Body request: LogoutRequest) : NetworkResponse<Unit>

    @NoAuth
    @POST("auth/refresh")
    suspend fun refresh(@Body request: RefreshRequest) : NetworkResponse<TokenPair>
}

class RetrofitAuthNetwork @Inject constructor(
    networkJson: Json
): AuthNetworkDataSource {

    private val authApi: RetrofitAuthNetworkApi

    init {
        val clientBuilder = OkHttpClient.Builder()

        if (BuildConfig.DEBUG) {
            val loggingInterceptor = HttpLoggingInterceptor().apply {
                level = HttpLoggingInterceptor.Level.BODY
            }
            clientBuilder.addInterceptor(loggingInterceptor)
        }

        authApi = Retrofit.Builder()
            .baseUrl(BuildConfig.BASE_URL)
            .client(clientBuilder.build())
            .addConverterFactory(
                networkJson.asConverterFactory("application/json".toMediaType())
            )
            .build()
            .create()
    }

    override suspend fun signUp(request: SignUpRequest): TokenPair =
        authApi.signUp(request).data

    override suspend fun login(request: LoginRequest): TokenPair =
        authApi.login(request).data

    override suspend fun logout(request: LogoutRequest) =
        authApi.logout(request).data

    override suspend fun refresh(request: RefreshRequest): TokenPair =
        authApi.refresh(request).data
}