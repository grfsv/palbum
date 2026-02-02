package com.grfsv.palbum.feature.auth

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.grfsv.palbum.core.data.repository.AuthRepository
import com.grfsv.palbum.core.model.AppResult
import com.grfsv.palbum.core.model.LoginRequest
import com.grfsv.palbum.core.model.toMessage
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class LoginViewModel @Inject constructor(
    private val authRepository: AuthRepository
) : ViewModel() {

    private val _uiState = MutableStateFlow<AuthUiState>(AuthUiState.Idle)
    val uiState: StateFlow<AuthUiState> = _uiState.asStateFlow()

    fun login(email: String, password: String) {
        viewModelScope.launch {
            _uiState.value = AuthUiState.Loading

            val result = authRepository.login(
                LoginRequest(
                    mail = email,
                    password = password
                )
            )

            _uiState.value = when (result) {
                is AppResult.Success -> AuthUiState.Success
                is AppResult.DomainError -> AuthUiState.FieldError(result.error.toMessage())
                is AppResult.InfraError -> AuthUiState.TransientError(result.error.toMessage())
            }
        }
    }

    fun transientErrorShown() {
        if (_uiState.value is AuthUiState.TransientError) {
            _uiState.value = AuthUiState.Idle
        }
    }
}
