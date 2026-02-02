package com.grfsv.palbum.core.model

import kotlinx.serialization.Serializable


@Serializable
data class SignUpRequest(
    val mail: String,
    val password: String,
    val name: String
)

