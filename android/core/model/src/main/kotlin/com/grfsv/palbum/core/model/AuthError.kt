package com.grfsv.palbum.core.model

sealed interface AuthError {
    data object InvalidCredentials : AuthError
    data object InvalidEmail : AuthError
    data object AccountLocked : AuthError
    data object EmailNotVerified : AuthError
}

fun AuthError.toMessage(): String = when (this) {
    AuthError.InvalidEmail -> "メールアドレスがすでに利用されています"
    AuthError.InvalidCredentials -> "メールアドレスまたはパスワードが間違っています"
    AuthError.AccountLocked -> "アカウントがロックされています"
    AuthError.EmailNotVerified -> "メールアドレスが確認されていません"
}