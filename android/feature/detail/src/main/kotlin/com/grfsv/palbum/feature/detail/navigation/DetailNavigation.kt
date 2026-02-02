package com.grfsv.palbum.feature.detail.navigation

import androidx.navigation.NavController
import androidx.navigation.NavGraphBuilder
import androidx.navigation.NavOptions
import kotlinx.serialization.Serializable

@Serializable data class DetailRoute(val messageId: String)

fun NavController.navigateToDetail(messageId: String, navOptions: NavOptions? = null) {
    navigate(route = DetailRoute(messageId), navOptions = navOptions)
}
fun NavGraphBuilder.detailScreen() {

}
