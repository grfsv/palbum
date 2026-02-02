package com.grfsv.palbum.feature.auth.navigation

import androidx.navigation.NavController
import androidx.navigation.NavGraphBuilder
import androidx.navigation.NavOptions
import androidx.navigation.compose.composable
import androidx.navigation.navigation
import com.grfsv.palbum.feature.auth.LoginScreen
import com.grfsv.palbum.feature.auth.SignUpScreen
import kotlinx.serialization.Serializable

@Serializable data object AuthBaseRoute

@Serializable data object LoginRoute

@Serializable data object SignUpRoute

fun NavController.navigateToLogin(navOptions: NavOptions? = null) =
    navigate(route = LoginRoute, navOptions)

fun NavController.navigateToSignUp(navOptions: NavOptions? = null) =
    navigate(route = SignUpRoute, navOptions)

fun NavGraphBuilder.authSection(
    onShowSnackbar: (String) -> Unit,
    onNavigateToSignUp: () -> Unit,
    onNavigateToLogin: () -> Unit,
) {
    navigation<AuthBaseRoute>(startDestination = LoginRoute) {
        composable<LoginRoute> {
            LoginScreen(
                onShowSnackbar = onShowSnackbar,
                onNavigateToSignUp = onNavigateToSignUp
            )
        }
        composable<SignUpRoute> {
            SignUpScreen(
                onShowSnackbar = onShowSnackbar,
                onNavigateToLogin = onNavigateToLogin
            )
        }
    }
}
