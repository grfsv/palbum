package com.grfsv.palbum.core.model

sealed interface InfraError {
    data object Network : InfraError
    data class Server(val code: Int) : InfraError
    data class Unknown(val throwable: Throwable) : InfraError
}

fun InfraError.toMessage(): String = when (this) {
    InfraError.Network -> "通信エラーが発生しました。接続を確認してください"
    is InfraError.Server -> "サーバーエラーが発生しました"
    is InfraError.Unknown -> "予期しないエラーが発生しました"
}