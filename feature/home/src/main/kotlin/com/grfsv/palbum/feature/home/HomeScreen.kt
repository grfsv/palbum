package com.grfsv.palbum.feature.home

import androidx.compose.foundation.gestures.snapping.SnapPosition
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.hilt.lifecycle.viewmodel.compose.hiltViewModel


@Composable
internal fun HomeScreen(
    onMessageClick: (String) -> Unit,
    modifier: Modifier = Modifier,
    viewModel: HomeViewModel = hiltViewModel()
) {
    val uiState by viewModel.uiState.collectAsState()

    HomeScreen(
        onMessageClick = onMessageClick,
        uiState = uiState,
        modifier = modifier,
    )
}

@Composable
internal fun HomeScreen(
    onMessageClick: (String) -> Unit,
    uiState: HomeUiState,
    modifier: Modifier = Modifier,
) {

    when(uiState) {
        is HomeUiState.Idle -> {}
        is HomeUiState.Loading -> {
            Box(
                modifier = Modifier.fillMaxSize(),
            ) {
                CircularProgressIndicator(modifier = Modifier.align(Alignment.Center))
            }
        }
        is HomeUiState.Error -> TODO()

        is HomeUiState.Success -> TODO()
    }
}