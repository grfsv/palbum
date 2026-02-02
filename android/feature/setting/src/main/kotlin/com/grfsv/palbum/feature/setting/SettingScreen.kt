package com.grfsv.palbum.feature.setting

import androidx.compose.material3.Text
import androidx.compose.material3.TextButton
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.hilt.lifecycle.viewmodel.compose.hiltViewModel

@Composable
internal fun SettingRoute(
    modifier: Modifier = Modifier,
    viewModel: SettingViewModel = hiltViewModel()
) {
    viewModel.uiState
    settingScreen(
        modifier = modifier,
        onLogoutClick = viewModel::logout
    )
}

@Composable
internal fun settingScreen(
    modifier: Modifier = Modifier,
    onLogoutClick: () -> Unit
) {
    Text("Setting Screen")
    TextButton(
        onClick = onLogoutClick
    ) {
        Text("logout")
    }
}