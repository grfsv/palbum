package com.grfsv.palbum.feature.setting

sealed class SettingUiState {
    object Idle : SettingUiState()
    data object Logout : SettingUiState()
    data class Error(val message: String) : SettingUiState()
}