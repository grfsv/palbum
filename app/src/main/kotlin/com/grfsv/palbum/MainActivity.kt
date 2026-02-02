package com.grfsv.palbum

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.activity.viewModels
import com.grfsv.palbum.core.data.repository.AuthRepository
import com.grfsv.palbum.core.design.theme.AppTheme
import com.grfsv.palbum.ui.PalbumApp
import com.grfsv.palbum.ui.rememberPalbumState
import dagger.hilt.android.AndroidEntryPoint
import javax.inject.Inject


@AndroidEntryPoint
class MainActivity : ComponentActivity() {

    @Inject
    lateinit var authRepository: AuthRepository

    override fun onCreate(savedInstanceState: Bundle?) {

        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContent {
            val appState = rememberPalbumState(
                authRepository = authRepository
            )

            AppTheme {
                PalbumApp(
                    appState,
                )
            }
        }

    }
}
