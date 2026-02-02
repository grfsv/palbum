package com.grfsv.palbum.navigation

import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.navigation.compose.NavHost
import com.grfsv.palbum.feature.auth.navigation.authSection
import com.grfsv.palbum.feature.auth.navigation.navigateToLogin
import com.grfsv.palbum.ui.PalbumState
import com.grfsv.palbum.feature.detail.navigation.detailScreen
import com.grfsv.palbum.feature.detail.navigation.navigateToDetail
import com.grfsv.palbum.feature.auth.navigation.navigateToSignUp
import com.grfsv.palbum.feature.home.navigation.homeSection
import com.grfsv.palbum.feature.setting.navigation.settingScreen


@Composable
fun PalbumNavHost(
    appState: PalbumState,
    startDestination: Any,
    modifier: Modifier = Modifier,
) {
    val navController = appState.navController

    NavHost(
        navController = navController,
        startDestination = startDestination,
        modifier = modifier
    ) {
        // 認証セクション
        authSection(
            onShowSnackbar = appState::showSnackbar,
            onNavigateToSignUp = navController::navigateToSignUp,
            onNavigateToLogin = navController::navigateToLogin,
        )

        // ホームセクション
        homeSection(
            onMessageClick = navController::navigateToDetail,
        ) {
            detailScreen()
        }

        // 設定画面
        settingScreen()
    }
}