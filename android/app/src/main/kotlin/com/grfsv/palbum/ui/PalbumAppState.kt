package com.grfsv.palbum.ui

import androidx.compose.material3.SnackbarHostState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.Stable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.navigation.NavDestination
import androidx.navigation.NavDestination.Companion.hasRoute
import androidx.navigation.NavGraph.Companion.findStartDestination
import androidx.navigation.NavHostController
import androidx.navigation.compose.rememberNavController
import androidx.navigation.navOptions
import com.grfsv.palbum.core.data.repository.AuthRepository
import com.grfsv.palbum.feature.home.navigation.navigateToHome
import com.grfsv.palbum.feature.setting.navigation.navigateToSetting
import com.grfsv.palbum.navigation.TopLevelDestination
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.coroutineScope
import kotlinx.coroutines.flow.SharingStarted
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.map
import kotlinx.coroutines.flow.stateIn
import kotlinx.coroutines.launch


@Composable
fun rememberPalbumState(
    authRepository: AuthRepository,
    coroutineScope: CoroutineScope = rememberCoroutineScope(),
    navController: NavHostController = rememberNavController(),
): PalbumState {
    return remember(
        authRepository,
        coroutineScope,
        navController
    ) {
       PalbumState(
           authRepository = authRepository,
           coroutineScope = coroutineScope,
           navController = navController
       )
    }
}
@Stable
class PalbumState(
    private val authRepository: AuthRepository,
    val coroutineScope: CoroutineScope,
    val navController: NavHostController
) {
    private val previousDestination = mutableStateOf<NavDestination?>(null)

    val snackbarHostState = SnackbarHostState()

    fun showSnackbar(message: String) {
        coroutineScope.launch {
            snackbarHostState.showSnackbar(message)
        }
    }
    val currentDestination: NavDestination?
        @Composable get() {
            // Collect the currentBackStackEntryFlow as a state
            val currentEntry = navController.currentBackStackEntryFlow
                .collectAsState(initial = null)

            // Fallback to previousDestination if currentEntry is null
            return currentEntry.value?.destination.also { destination ->
                if (destination != null) {
                    previousDestination.value = destination
                }
            } ?: previousDestination.value
        }

    // 現在地がトップレベルのどれに該当するか判定
    val currentTopLevelDestination: TopLevelDestination?
        @Composable get() {
            return TopLevelDestination.entries.firstOrNull { topLevelDestination ->
                currentDestination?.hasRoute(route = topLevelDestination.route) == true
            }
        }
    val topLevelDestinations: List<TopLevelDestination> = TopLevelDestination.entries

    val authState: StateFlow<AuthState> = authRepository.isLoggedIn
        .map { isLoggedIn ->
            if (isLoggedIn) AuthState.LoggedIn else AuthState.LoggedOut
        }
        .stateIn(
            scope = coroutineScope,
            started = SharingStarted.WhileSubscribed(5_000),
            initialValue = AuthState.Loading,
        )

    val sessionExpired: StateFlow<Boolean> = authRepository.sessionExpired
        .stateIn(
            scope = coroutineScope,
            started = SharingStarted.WhileSubscribed(5_000),
            initialValue = false,
        )

    fun acknowledgeSessionExpired() {
        authRepository.acknowledgeSessionExpired()
    }

    // ボトムバー選択時の遷移ロジック
    fun navigateToTopLevelDestination(topLevelDestination: TopLevelDestination) {
        val topLevelNavOptions = navOptions {
            // 1. バックスタックの管理: グラフの開始地点(Home)まで戻す（状態は保存）
            popUpTo(navController.graph.findStartDestination().id) {
                saveState = true
            }
            // 2. 多重起動防止
            launchSingleTop = true
            // 3. 状態の復元
            restoreState = true
        }
        when (topLevelDestination) {
            TopLevelDestination.HOME -> navController.navigateToHome(topLevelNavOptions)
            TopLevelDestination.SETTING -> navController.navigateToSetting(topLevelNavOptions)
        }
    }
}

sealed interface AuthState {
    data object Loading : AuthState
    data object LoggedIn : AuthState
    data object LoggedOut : AuthState
}