package com.grfsv.palbum.core.datastore

import androidx.datastore.core.DataStore
import com.grfsv.palbum.core.model.TokenPair
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.distinctUntilChanged
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.flow.map
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class AuthLocalDataSource @Inject constructor(
    private val refreshTokenDataStore: DataStore<ByteArray>
) {

    // ------------------------------------------------
    // 1. 通信 (Interceptor) 向け
    // ------------------------------------------------
    // 待機なし・同期アクセス。単なる変数なので爆速です。
    @Volatile // 別スレッド(Interceptor)から見た時に最新値が見えるようにする
    var accessToken: String? = null
        private set // 外部からは書き換え禁止（読み取り専用）

    // リフレッシュトークンは「更新時」にしか使わないので suspend (DataStore読み込み) でOK
    suspend fun getRefreshToken(): String? {
        val bytes = refreshTokenDataStore.data.first()
        return if (bytes.isEmpty()) null else String(bytes, Charsets.UTF_8)
    }

    // ------------------------------------------------
    // 2. UI (ViewModel) 向け
    // ------------------------------------------------
    // UIが知りたいのは「ログインしてるか？」だけ。
    // アクセストークンが変わっても通知は飛びません（無駄な更新を防ぐ）。
    val hasToken: Flow<Boolean> = refreshTokenDataStore.data
        .map { it.isNotEmpty() }
        .distinctUntilChanged()

    // ------------------------------------------------
    // 3. 保存・削除処理
    // ------------------------------------------------
    suspend fun saveToken(token: TokenPair) {
        // メモリ（変数）を即更新
        accessToken = token.accessToken

        // ディスク（DataStore）を保存
        refreshTokenDataStore.updateData {
            token.refreshToken.toByteArray(Charsets.UTF_8)
        }
    }

    suspend fun clear() {
        accessToken = null
        refreshTokenDataStore.updateData { ByteArray(0) }
    }
}