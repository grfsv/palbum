package com.grfsv.palbum.core.data

import com.grfsv.palbum.core.datastore.AuthLocalDataSource
import com.grfsv.palbum.core.model.RefreshRequest
import com.grfsv.palbum.core.network.AuthNetworkDataSource
import com.grfsv.palbum.core.network.RefreshResult
import com.grfsv.palbum.core.network.TokenProvider
import retrofit2.HttpException
import javax.inject.Inject

class TokenProviderImpl @Inject constructor(
    private val authLocalDataSource: AuthLocalDataSource,
    private val authNetworkDataSource: AuthNetworkDataSource,
    private val sessionExpiryListener: SessionExpiryListener,
) : TokenProvider {

    override fun getAccessToken(): String? = authLocalDataSource.accessToken

    override suspend fun refresh(): RefreshResult {
        val refreshToken = authLocalDataSource.getRefreshToken()
            ?: run {
                sessionExpiryListener.onSessionExpired()
                authLocalDataSource.clear()
                return RefreshResult.SessionExpired
            }

        return try {
            val newTokenPair = authNetworkDataSource.refresh(RefreshRequest(refreshToken))
            authLocalDataSource.saveToken(newTokenPair)
            RefreshResult.Success
        } catch (e: HttpException) {
            if (e.code() in 400..499) {
                sessionExpiryListener.onSessionExpired()
                authLocalDataSource.clear()
                RefreshResult.SessionExpired
            } else {
                throw e
            }
        }
    }
}