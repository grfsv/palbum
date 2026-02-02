package com.grfsv.palbum.core.model

import kotlinx.serialization.Serializable

@Serializable
data class RefreshRequest (
    val refreshToken: String
)