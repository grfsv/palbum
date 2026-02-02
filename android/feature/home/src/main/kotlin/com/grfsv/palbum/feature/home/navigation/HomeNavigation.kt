package com.grfsv.palbum.feature.home.navigation

import androidx.navigation.NavController
import androidx.navigation.NavGraphBuilder
import androidx.navigation.NavOptions
import kotlinx.serialization.Serializable
import androidx.navigation.compose.composable
import androidx.navigation.navigation
import com.grfsv.palbum.feature.home.HomeScreen


@Serializable data object HomeRoute

@Serializable data object HomeBaseRoute

fun NavController.navigateToHome(navOptions: NavOptions? = null) = navigate(route = HomeRoute, navOptions = navOptions)

fun NavGraphBuilder.homeSection(
    onMessageClick: (String) -> Unit,
    detailDestination: NavGraphBuilder.() -> Unit,
) {
    navigation<HomeBaseRoute>(startDestination = HomeRoute) {
        composable<HomeRoute> {
            HomeScreen(onMessageClick)
        }
        detailDestination()
    }
}