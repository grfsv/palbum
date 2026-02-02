package com.grfsv.palbum.core.data.repository

import com.grfsv.palbum.core.model.AppResult
import com.grfsv.palbum.core.model.AuthError
import com.grfsv.palbum.core.model.LoginRequest
import com.grfsv.palbum.core.model.SignUpRequest
import kotlinx.coroutines.flow.Flow


interface AuthRepository {
    val isLoggedIn: Flow<Boolean>
    val sessionExpired: Flow<Boolean>

    fun acknowledgeSessionExpired()

    suspend fun signUp(request: SignUpRequest): AppResult<Unit, AuthError>

    suspend fun login(loginRequest: LoginRequest): AppResult<Unit, AuthError>

    suspend fun logout(): AppResult<Unit, AuthError>

    suspend fun refresh()
}