package com.grfsv.palbum.core.data.util

import com.grfsv.palbum.core.model.AppResult
import com.grfsv.palbum.core.model.InfraError
import retrofit2.HttpException
import java.io.IOException

suspend fun <T, E> safeApiCall(
    mapToDomainError: (code: Int, message: String?) -> E?,
    block: suspend () -> T
): AppResult<T, E> {
    return try {
        AppResult.Success(block())
    } catch (e: HttpException) {
        val code = e.code()
        when {
            code in 400..499 -> {
                val domainError = mapToDomainError(code, e.message())
                if (domainError != null) {
                    AppResult.DomainError(domainError)
                } else {
                    AppResult.InfraError(InfraError.Unknown(e))
                }
            }
            else -> AppResult.InfraError(InfraError.Server(code))
        }
    } catch (e: IOException) {
        AppResult.InfraError(InfraError.Network)
    } catch (e: Exception) {
        AppResult.InfraError(InfraError.Unknown(e))
    }
}