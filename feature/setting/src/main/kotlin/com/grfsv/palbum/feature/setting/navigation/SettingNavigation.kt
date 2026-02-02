package com.grfsv.palbum.feature.setting.navigation

import androidx.navigation.NavController
import androidx.navigation.NavGraphBuilder
import androidx.navigation.NavOptions
import androidx.navigation.compose.composable
import androidx.navigation.navigation
import com.grfsv.palbum.feature.setting.SettingRoute
import kotlinx.serialization.Serializable


@Serializable data object SettingRoute

fun NavController.navigateToSetting(navOptions: NavOptions? = null) = navigate(route = SettingRoute, navOptions)

fun NavGraphBuilder.settingScreen() {
    composable<SettingRoute> {
        SettingRoute()
    }
}