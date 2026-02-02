package com.grfsv.palbum.navigation

import androidx.annotation.StringRes
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Home
import androidx.compose.material.icons.filled.Settings
import androidx.compose.material.icons.outlined.Home
import androidx.compose.material.icons.outlined.Settings
import androidx.compose.ui.graphics.vector.ImageVector
import com.grfsv.palbum.feature.home.navigation.HomeBaseRoute
import com.grfsv.palbum.feature.home.navigation.HomeRoute
import com.grfsv.palbum.feature.setting.navigation.SettingRoute
import kotlin.reflect.KClass
import com.grfsv.palbum.feature.home.R as homeRes
import com.grfsv.palbum.feature.setting.R as settingRes


enum class TopLevelDestination(
    val selectedIcon: ImageVector,
    val unselectedIcon: ImageVector,
    @field:StringRes val iconTextId: Int,
    @field:StringRes val titleTextId: Int,
    val route: KClass<*>,
    val baseRoute: KClass<*> = route,
) {
   HOME(
       selectedIcon = Icons.Default.Home,
       unselectedIcon = Icons.Outlined.Home,
       iconTextId = homeRes.string.future_home_title,
       titleTextId = homeRes.string.future_home_title,
       route = HomeRoute::class,
       baseRoute = HomeBaseRoute::class
   ),

    SETTING(
        selectedIcon = Icons.Default.Settings,
        unselectedIcon = Icons.Outlined.Settings,
        iconTextId = settingRes.string.future_setting_title,
        titleTextId = settingRes.string.future_setting_title,
        route = SettingRoute::class,
    ),
}

