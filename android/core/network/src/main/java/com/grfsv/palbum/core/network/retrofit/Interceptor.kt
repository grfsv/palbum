package com.grfsv.palbum.core.network.retrofit

import com.grfsv.palbum.core.network.TokenProvider
import okhttp3.Interceptor
import okhttp3.Response
import retrofit2.Invocation

import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class Interceptor @Inject constructor(
    private val tokenProvider: TokenProvider,
): Interceptor {
    override fun intercept(chain: Interceptor.Chain): Response {
        val request = chain.request()

        val isNoAuth = request.tag(Invocation::class.java)
            ?.method()
            ?.isAnnotationPresent(NoAuth::class.java) == true

        if (isNoAuth) return chain.proceed(request)

        val accessToken = tokenProvider.getAccessToken()

        val newRequest =  request.newBuilder()
            .addHeader("Authorization", "Bearer $accessToken")
            .build()
        return chain.proceed(newRequest)
    }
}