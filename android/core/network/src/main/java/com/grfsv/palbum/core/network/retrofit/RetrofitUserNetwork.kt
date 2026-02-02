package com.grfsv.palbum.core.network.retrofit

import com.grfsv.palbum.core.network.UserNetworkDataSource
import retrofit2.Retrofit
import retrofit2.create
import javax.inject.Inject

class RetrofitUserNetwork @Inject constructor(
    private val retrofit: Retrofit
) : UserNetworkDataSource{
    private val userApi = retrofit.create<UserNetworkDataSource>()

    override suspend fun logout() {
        TODO("Not yet implemented")
    }
}