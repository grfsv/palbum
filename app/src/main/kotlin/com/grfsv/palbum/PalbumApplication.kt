package com.grfsv.palbum

import android.app.Application
import dagger.hilt.android.HiltAndroidApp

@HiltAndroidApp
class PalbumApplication: Application() {
    override fun onCreate() {
        super.onCreate()
    }
}