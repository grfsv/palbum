package com.grfsv.palbum.core.model

sealed interface AppResult<out T, out E> {
    data class Success<T>(val data: T) : AppResult<T, Nothing>
    data class DomainError<E>(val error: E) : AppResult<Nothing, E>
    data class InfraError(val error: com.grfsv.palbum.core.model.InfraError) : AppResult<Nothing, Nothing>
}

inline fun <T, E> AppResult<T, E>.onSuccess(action: (T) -> Unit): AppResult<T, E> {
    if (this is AppResult.Success) action(data)
    return this
}

inline fun <T, E> AppResult<T, E>.onDomainError(action: (E) -> Unit): AppResult<T, E> {
    if (this is AppResult.DomainError) action(error)
    return this
}

inline fun <T, E> AppResult<T, E>.onInfraError(action: (InfraError) -> Unit): AppResult<T, E> {
    if (this is AppResult.InfraError) action(error)
    return this
}

fun <T, E> AppResult<T, E>.getOrNull(): T? = when (this) {
    is AppResult.Success -> data
    else -> null
}