package com.grfsv.palbum.core.network.model

import kotlinx.serialization.Serializable

@Serializable
internal data class NetworkResponse<T>(
    val data: T,
)

