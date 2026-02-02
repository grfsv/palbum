package com.grfsv.palbum.core.network

interface TokenProvider {
    fun getAccessToken(): String?
    suspend fun refresh(): RefreshResult
}