package com.grfsv.palbum.core.data

import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class SessionStateHolder @Inject constructor() : SessionExpiryListener {

    private val _sessionExpired = MutableStateFlow(false)
    val sessionExpired: Flow<Boolean> = _sessionExpired.asStateFlow()

    override fun onSessionExpired() {
        _sessionExpired.value = true
    }

    fun acknowledge() {
        _sessionExpired.value = false
    }
}