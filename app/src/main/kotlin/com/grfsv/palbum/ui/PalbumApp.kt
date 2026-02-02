package com.grfsv.palbum.ui

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.AlertDialog
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.Icon
import androidx.compose.material3.NavigationBar
import androidx.compose.material3.NavigationBarItem
import androidx.compose.material3.Scaffold
import androidx.compose.material3.SnackbarHost
import androidx.compose.material3.SnackbarHostState
import androidx.compose.material3.Text
import androidx.compose.material3.TextButton
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import androidx.navigation.NavDestination
import androidx.navigation.NavDestination.Companion.hasRoute
import androidx.navigation.NavDestination.Companion.hierarchy
import com.grfsv.palbum.feature.auth.navigation.AuthBaseRoute
import com.grfsv.palbum.feature.home.navigation.HomeBaseRoute
import com.grfsv.palbum.navigation.PalbumNavHost
import com.grfsv.palbum.navigation.TopLevelDestination
import kotlin.reflect.KClass

@Composable
fun PalbumApp(
    appState: PalbumState,
) {
    val authState by appState.authState.collectAsStateWithLifecycle()
    val sessionExpired by appState.sessionExpired.collectAsStateWithLifecycle()

    if (sessionExpired) {
        AlertDialog(
            onDismissRequest = { appState.acknowledgeSessionExpired() },
            title = { Text("セッション期限切れ") },
            text = { Text("ログインセッションが期限切れになりました。再度ログインしてください。") },
            confirmButton = {
                TextButton(onClick = { appState.acknowledgeSessionExpired() }) {
                    Text("OK")
                }
            }
        )
    }

    when (authState) {
        AuthState.Loading -> {
            Box(
                modifier = Modifier.fillMaxSize(),
                contentAlignment = Alignment.Center
            ) {
                CircularProgressIndicator()
            }
        }
        AuthState.LoggedIn -> {
            MainContent(
                appState = appState,
                startDestination = HomeBaseRoute,
                showBottomBar = true
            )
        }
        AuthState.LoggedOut -> {
            MainContent(
                appState = appState,
                startDestination = AuthBaseRoute,
                showBottomBar = false
            )
        }
    }
}

@Composable
private fun MainContent(
    appState: PalbumState,
    startDestination: Any,
    showBottomBar: Boolean
) {
    val currentDest = appState.currentTopLevelDestination

    Scaffold(
        snackbarHost = { SnackbarHost(appState.snackbarHostState) },
        bottomBar = {
            if (showBottomBar && currentDest != null) {
                PalbumBottomBar(
                    destinations = appState.topLevelDestinations,
                    onNavigateToDestination = appState::navigateToTopLevelDestination,
                    currentDestination = appState.currentDestination
                )
            }
        }
    ) { innerPadding ->
        PalbumNavHost(
            appState = appState,
            startDestination = startDestination,
            modifier = Modifier.padding(innerPadding)
        )
    }
}

@Composable
private fun PalbumBottomBar(
    destinations: List<TopLevelDestination>,
    onNavigateToDestination: (TopLevelDestination) -> Unit,
    currentDestination: NavDestination?
) {
    NavigationBar {
        destinations.forEach { destination ->
            val selected = currentDestination
                .isRouteInHierarchy(destination.baseRoute)

            NavigationBarItem(
                selected = selected,
                onClick = { onNavigateToDestination(destination) },
                icon = {
                    Icon(
                        imageVector = if (selected) destination.selectedIcon else destination.unselectedIcon,
                        contentDescription = stringResource(id = destination.iconTextId)
                    )
                },
                label = { Text(stringResource(id = destination.titleTextId)) }
            )
        }
    }
}

private fun NavDestination?.isRouteInHierarchy(route: KClass<*>) =
    this?.hierarchy?.any {
        it.hasRoute(route)
    } ?: false
