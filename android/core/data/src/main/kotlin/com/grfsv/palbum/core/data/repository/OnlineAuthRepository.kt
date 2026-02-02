package com.grfsv.palbum.core.data.repository

import com.grfsv.palbum.core.data.SessionStateHolder
import com.grfsv.palbum.core.data.util.safeApiCall
import com.grfsv.palbum.core.datastore.AuthLocalDataSource
import com.grfsv.palbum.core.model.AppResult
import com.grfsv.palbum.core.model.AuthError
import com.grfsv.palbum.core.model.LoginRequest
import com.grfsv.palbum.core.model.LogoutRequest
import com.grfsv.palbum.core.model.RefreshRequest
import com.grfsv.palbum.core.model.SignUpRequest
import com.grfsv.palbum.core.network.AuthNetworkDataSource
import kotlinx.coroutines.flow.Flow
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class OnlineAuthRepository @Inject constructor(
    private val authLocalDataSource: AuthLocalDataSource,
    private val authNetworkDataSource: AuthNetworkDataSource,
    private val sessionStateHolder: SessionStateHolder,
) : AuthRepository {

    override val isLoggedIn: Flow<Boolean> = authLocalDataSource.hasToken

    override val sessionExpired: Flow<Boolean> = sessionStateHolder.sessionExpired

    override fun acknowledgeSessionExpired() {
        sessionStateHolder.acknowledge()
    }

    override suspend fun signUp(request: SignUpRequest): AppResult<Unit, AuthError> =
        safeApiCall(
            mapToDomainError = { code, _ ->
                when (code) {
                    400 -> AuthError.InvalidEmail
                    else -> null
                }
            }
        ) {
            val newTokenPair = authNetworkDataSource.signUp(request)
            authLocalDataSource.saveToken(newTokenPair)
        }


    override suspend fun login(loginRequest: LoginRequest): AppResult<Unit, AuthError> =
        safeApiCall(
            mapToDomainError = { code, _ ->
                when (code) {
                    400 -> AuthError.InvalidCredentials
                    401 -> AuthError.InvalidCredentials
                    403 -> AuthError.EmailNotVerified
                    else -> null
                }
            }
        ) {
            val newTokenPair = authNetworkDataSource.login(loginRequest)
            authLocalDataSource.saveToken(newTokenPair)
        }

    override suspend fun logout() : AppResult<Unit, AuthError> =
        safeApiCall(
            mapToDomainError = { _, _ ->
                null
            }
        ) {
            val refreshToken = authLocalDataSource.getRefreshToken()
            authLocalDataSource.clear()

            refreshToken?.let {
                authNetworkDataSource.logout(LogoutRequest(refreshToken))
            }
        }


    override suspend fun refresh() {
        val refreshToken = authLocalDataSource.getRefreshToken()

        if (refreshToken == null) {
            authLocalDataSource.clear()
        } else {
            val newTokenPair = authNetworkDataSource.refresh(RefreshRequest(refreshToken))
            authLocalDataSource.saveToken(newTokenPair)
        }
    }
}