package com.grfsv.palbum.core.network

sealed interface RefreshResult {
    data object Success : RefreshResult
    data object SessionExpired : RefreshResult
}