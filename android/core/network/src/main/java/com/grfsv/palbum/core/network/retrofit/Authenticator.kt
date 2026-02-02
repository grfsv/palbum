package com.grfsv.palbum.core.network.retrofit

import com.grfsv.palbum.core.network.RefreshResult
import com.grfsv.palbum.core.network.TokenProvider
import kotlinx.coroutines.runBlocking
import kotlinx.coroutines.sync.Mutex
import kotlinx.coroutines.sync.withLock
import okhttp3.Authenticator
import okhttp3.Request
import okhttp3.Response
import okhttp3.Route
import retrofit2.Invocation
import javax.inject.Inject
import javax.inject.Singleton


@Singleton
class Authenticator @Inject constructor(
    private val tokenProvider: TokenProvider,
) : Authenticator {

    private val mutex = Mutex()

    private val Response.retryCount: Int
        get() {
            var currentResponse = priorResponse
            var result = 0
            while (currentResponse != null) {
                result++
                currentResponse = currentResponse.priorResponse
            }
            return result
        }

    override fun authenticate(route: Route?, response: Response): Request? {
        val isNoAuth = response.request.tag(Invocation::class.java)
            ?.method()
            ?.isAnnotationPresent(NoAuth::class.java) == true
        if (isNoAuth) return null

        when {
            response.retryCount > 2 -> return null
            else -> {
                val result = try {
                    runBlocking {
                        mutex.withLock { tokenProvider.refresh() }
                    }
                } catch (e: Exception) {
                    return null
                }

                return when (result) {
                    RefreshResult.Success -> {
                        tokenProvider.getAccessToken()?.let { token ->
                            response.request.newBuilder()
                                .header("Authorization", "Bearer $token")
                                .build()
                        }
                    }
                    RefreshResult.SessionExpired -> null
                }
            }
        }
    }
}