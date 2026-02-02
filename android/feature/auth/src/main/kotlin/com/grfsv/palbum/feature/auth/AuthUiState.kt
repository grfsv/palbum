package com.grfsv.palbum.feature.auth

/**
 * Login/SignUp画面で共通のUiState
 */
sealed interface AuthUiState {
    data object Idle : AuthUiState
    data object Loading : AuthUiState
    data object Success : AuthUiState

    /** ドメインエラー（入力エラーなど）- インライン表示用 */
    data class FieldError(val message: String) : AuthUiState

    /** インフラエラー（通信エラーなど）- Snackbar表示用 */
    data class TransientError(val message: String) : AuthUiState
}
