package com.grfsv.palbum.core.network


import com.grfsv.palbum.core.model.LoginRequest
import com.grfsv.palbum.core.model.LogoutRequest
import com.grfsv.palbum.core.model.RefreshRequest
import com.grfsv.palbum.core.model.SignUpRequest
import com.grfsv.palbum.core.model.TokenPair


interface AuthNetworkDataSource {
    suspend fun signUp(request: SignUpRequest) : TokenPair
    suspend fun login(request: LoginRequest) : TokenPair
    suspend fun logout(request: LogoutRequest)
    suspend fun refresh(request: RefreshRequest) : TokenPair
}

