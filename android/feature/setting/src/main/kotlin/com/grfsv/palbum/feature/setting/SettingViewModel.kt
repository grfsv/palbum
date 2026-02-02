package com.grfsv.palbum.feature.setting

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.grfsv.palbum.core.data.repository.AuthRepository
import com.grfsv.palbum.core.model.AppResult
import com.grfsv.palbum.core.model.InfraError
import com.grfsv.palbum.core.model.toMessage
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class SettingViewModel @Inject constructor(
    private val authRepository: AuthRepository
): ViewModel() {
    private val _uiState = MutableStateFlow<SettingUiState>(SettingUiState.Idle)
    val uiState: StateFlow<SettingUiState> = _uiState

    fun logout() {
        viewModelScope.launch {
            val result = authRepository.logout()

            _uiState.value =  when(result) {
                is AppResult.Success -> SettingUiState.Logout
                is AppResult.DomainError -> SettingUiState.Logout
                is AppResult.InfraError -> SettingUiState.Error(result.error.toMessage())
            }
        }
    }
}