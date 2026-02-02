package com.grfsv.palbum.core.network

import retrofit2.http.POST

interface UserNetworkDataSource {
    @POST("logout")
    suspend fun logout()
}