package com.grfsv.palbum.buildLogic.convention

import org.gradle.api.Project
import org.gradle.kotlin.dsl.configure
import org.jlleitschuh.gradle.ktlint.KtlintExtension

fun Project.configureLinter() {
    pluginManager.apply(libs.getPluginId("ktlint"))

    extensions.configure<KtlintExtension> {
        android.set(true)
    }
}