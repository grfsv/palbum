import com.android.build.api.dsl.LibraryExtension
import com.grfsv.palbum.buildLogic.convention.android
import com.grfsv.palbum.buildLogic.convention.configureAndroidCommon
import com.grfsv.palbum.buildLogic.convention.configureKotlinCommon
import com.grfsv.palbum.buildLogic.convention.configureLinter
import com.grfsv.palbum.buildLogic.convention.getPluginId
import com.grfsv.palbum.buildLogic.convention.kotlinAndroid
import com.grfsv.palbum.buildLogic.convention.libs

import org.gradle.api.Plugin
import org.gradle.api.Project

class AndroidLibraryPlugin : Plugin<Project> {
    override fun apply(target: Project) {
        with(target) {
            with(pluginManager) {
                apply(libs.getPluginId("android-library"))
                apply(libs.getPluginId("kotlin-android"))
            }

            kotlinAndroid {
                configureKotlinCommon(this)
            }

            android<LibraryExtension> {
                configureAndroidCommon()
            }
            configureLinter()
        }
    }
}